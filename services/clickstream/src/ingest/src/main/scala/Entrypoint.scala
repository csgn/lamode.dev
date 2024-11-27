package ingest

import scala.util.Properties.envOrNone

import org.apache.spark.SparkContext
import org.apache.spark.sql.{SparkSession}

object Entrypoint extends App {
  implicit val spark = SparkSession.builder
    .appName("Entrypoint")
    .master("local[3]")
    .getOrCreate()

  implicit val sc = spark.sparkContext

  val kafkaProps = KafkaProperties(
    bootstrapServers = envOrNone("KAFKA_ADDR").get,
    topic = envOrNone("KAFKA_TOPIC").get
  )

  val hadoopProps = HadoopProperties(
    uri = envOrNone("HADOOP_URI").get,
    dataFolder = envOrNone("HADOOP_RAW_EVENTS_DIR").get
  )

  println()
  println("============= SPARK JOB OUTPUT =============")
  println()
  IngestionStreamJob(kafkaProps, hadoopProps).run()
  println()
  println("============================================")
  println()

  spark.stop()
}
