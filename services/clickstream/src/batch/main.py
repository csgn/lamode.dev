import os
import sys
import pyspark
from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType


def getenv(env):
    env = os.getenv(env)
    if env == "":
        raise Exception(f"{env} is not defined")
    return env


HADOOP_URI = getenv("HADOOP_URI")
HADOOP_DATA_FOLDER = getenv("HADOOP_DATA_FOLDER")

# Must respect the table at "services/clickstream/docs/collector/index.md#events"
eventSchema = StructType(
    [
        StructField("pid", StringType(), True),
        StructField("sid", StringType(), True),
        StructField("channel", StringType(), True),
        StructField("event", StringType(), True),
        StructField("eventGroup", StringType(), True),
        StructField("senderMethod", StringType(), True),
        StructField("ip", StringType(), True),
        StructField("createdAt", StringType(), True),
    ]
)

spark = SparkSession.builder.master("local").appName("batch").getOrCreate()
data = spark.read.json(f"{HADOOP_URI}/{HADOOP_DATA_FOLDER}", schema=eventSchema)
data.show(5)
