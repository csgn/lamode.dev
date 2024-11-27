package ingest

import org.apache.spark.sql.types._

// Must respect the schema at "/docs/services/clickstream/collector/index.md#events"
object EventSchema {
  def apply() = {
    StructType(
      Array(
        StructField("sid", StringType, nullable = false),
        StructField("channel", StringType, nullable = false),
        StructField("timestamp", StringType, nullable = false),
        StructField("culture", StringType, nullable = false),
        StructField("action", StringType, nullable = false),
        StructField("type", StringType, nullable = false),
        StructField("ipAddress", StringType, nullable = false),
        StructField("latitude", FloatType, nullable = false),
        StructField("longitude", FloatType, nullable = false),
        StructField("productId", StringType, nullable = true),
        StructField("searchQuery", StringType, nullable = true)
      )
    )
  }
}
