package ingest

import org.apache.spark.SparkContext
import org.apache.spark.sql._
import org.apache.spark.sql.functions._
import org.apache.spark.sql.streaming._

case class IngestionStreamJob(
    kafkaProps: KafkaProperties,
    hadoopProps: HadoopProperties
) extends Job {

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
        from_json($"value", EventSchema()).as("value")
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
      .trigger(Trigger.ProcessingTime("5 minutes"))
      .option("maxRecordsPerFile", "100000")
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
