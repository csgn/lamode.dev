package validator

import org.apache.spark.SparkContext
import org.apache.spark.sql._
import org.apache.spark.sql.functions._
import org.apache.spark.sql.streaming._
import org.apache.spark.sql.types._

private[validator] case class Task1(
    kafkaProps: KafkaProperties,
    hadoopProps: HadoopProperties
) extends Task {

  // Rely upon the table at "services/clickstream/docs/collector/index.md#events"
  private val schema = StructType(
    Array(
      StructField("pid", StringType),
      StructField("sid", StringType),
      StructField("channel", StringType),
      StructField("event", StringType),
      StructField("eventGroup", StringType),
      StructField("senderMethod", StringType),
      StructField("ip", StringType)
    )
  )

  protected override def read()(implicit
      spark: SparkSession,
      sparkContext: SparkContext
  ): DataFrame = {
    spark.readStream
      .format("kafka")
      .option(
        "kafka.bootstrap.servers",
        kafkaProps.bootstrapServers
      )
      .option("subscribe", kafkaProps.topic)
      .load()

    // spark.readStream
    //   .format("socket")
    //   .option("host", "localhost")
    //   .option("port", 9876) // nc -l 9876
    //   .load()
  }

  protected override def transform[T](input: DataFrame)(implicit
      spark: SparkSession,
      sparkContext: SparkContext
  ): T = {
    import spark.implicits._

    input
      .selectExpr(
        "cast(value as string)"
      )
      .select(
        from_json($"value", schema).as("value")
      )
      .select("value.*")
      .asInstanceOf[T]
  }

  protected override def write(result: DataFrame)(implicit
      spark: SparkSession,
      sparkContext: SparkContext
  ): DataStreamWriter[Row] = {
    result.writeStream
      .format("json")
      .outputMode("append")
      .trigger(Trigger.ProcessingTime("1 second"))
      .option("path", s"${hadoopProps.uri}/${hadoopProps.dataFolder}")
      .option("checkpointLocation", s"${hadoopProps.uri}/checkpoints")
  }

  override def run()(implicit
      spark: SparkSession,
      sparkContext: SparkContext
  ): Unit = {
    // • Input Source //
    val inputDF = read()

    // • Data Transformation //
    val resultDF = transform[DataFrame](inputDF)

    // • Output Sink //
    val writer = write(resultDF)

    // • Start Stream //
    val streamingQuery = writer.start()
    streamingQuery.awaitTermination()

    // • Tidy up //
    streamingQuery.stop()
  }
}
