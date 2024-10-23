package validator

import org.apache.spark.SparkContext
import org.apache.spark.sql.{SparkSession, DataFrame, Row}
import org.apache.spark.sql.streaming.DataStreamWriter

trait Task {
  protected def read()(implicit
      spark: SparkSession,
      sparkContext: SparkContext
  ): DataFrame

  protected def transform[T](
      input: DataFrame
  )(implicit
      spark: SparkSession,
      sparkContext: SparkContext
  ): T

  protected def write(result: DataFrame)(implicit
      spark: SparkSession,
      sparkContext: SparkContext
  ): DataStreamWriter[Row]

  def run()(implicit
      spark: SparkSession,
      sparkContext: SparkContext
  ): Unit
}
