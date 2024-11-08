import os

from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType
import pyspark.sql.functions as F


def getenv(env):
    env = os.getenv(env)
    if env == "":
        raise Exception(f"{env} is not defined")
    return env


HADOOP_URI = getenv("HADOOP_URI")
HADOOP_STAGE_EVENTS_DIR = getenv("HADOOP_STAGE_EVENTS_DIR")
HADOOP_FINAL_EVENTS_DIR = getenv("HADOOP_FINAL_EVENTS_DIR")

# Must respect the table at "/docs/services/clickstream/collector/index.md#events"
eventSchema = StructType(
    [
        StructField("pid", StringType(), True),
        StructField("productId", StringType(), False),
        StructField("channel", StringType(), True),
        StructField("event", StringType(), True),
        StructField("interaction", StringType(), True),
        StructField("createdAt", StringType(), True),
    ]
)

channels = ["mobile", "web"]

spark = SparkSession.builder.master("local").appName("process_raw_event").getOrCreate()
df = spark.read.json(f"{HADOOP_URI}/{HADOOP_STAGE_EVENTS_DIR}", schema=eventSchema)

for channel in channels:
    channelDf = df.where(F.expr("channel") == channel)
    channelDf.write.save(
        f"{HADOOP_URI}/{HADOOP_FINAL_EVENTS_DIR}/{channel}",
        format="parquet",
        mode="append",
    )
