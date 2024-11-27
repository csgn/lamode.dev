package ingest

import scala.reflect.ClassTag

trait StreamPropertyKey extends Product with Serializable
case object Kafka extends StreamPropertyKey
case object Hadoop extends StreamPropertyKey

trait StreamProperties extends Product with Serializable
case class KafkaProperties(
    bootstrapServers: String,
    topic: String
) extends StreamProperties
case class HadoopProperties(
    uri: String,
    dataFolder: String
) extends StreamProperties
