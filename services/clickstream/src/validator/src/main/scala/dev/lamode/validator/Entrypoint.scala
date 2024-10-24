package validator

import scala.util.Properties.envOrElse

import org.apache.spark.SparkContext
import org.apache.spark.sql.{SparkSession}

private[validator] object Entrypoint extends App {
  implicit val spark = SparkSession.builder
    .appName("Entrypoint")
    .master("local[3]")
    .getOrCreate()

  implicit val sc = spark.sparkContext

  implicit val kafkaProps = KafkaProperties(
    bootstrapServers = envOrElse("KAFKA_ADDR", "0.0.0.0:9092"),
    topic = envOrElse("KAFKA_TOPIC", "clickstream-dev-topic")
  )

  implicit val hadoopProps = HadoopProperties(
    uri = envOrElse("HADOOP_URI", "hdfs://localhost:9000"),
    dataFolder = envOrElse("HADOOP_DATA_FOLDER", "data")
  )

  println()
  println("============= SPARK JOB OUTPUT =============")
  println()
  Task1(kafkaProps, hadoopProps).run()
  println()
  println("============================================")
  println()

  spark.stop()
}
