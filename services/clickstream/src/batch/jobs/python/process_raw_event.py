import os
from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType


def getenv(env):
    env = os.getenv(env)
    if env == "":
        raise Exception(f"{env} is not defined")
    return env


HADOOP_URI = getenv("HADOOP_URI")
HADOOP_STAGE_EVENTS_DIR = getenv("HADOOP_STAGE_EVENTS_DIR")
HADOOP_FINAL_EVENTS_DIR = getenv("HADOOP_FINAL_EVENTS_DIR")

# Must respect the table at "services/clickstream/docs/components/collector/index.md#events"
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

spark = SparkSession.builder.master("local").appName("process_raw_event").getOrCreate()
df = spark.read.json(f"{HADOOP_URI}/{HADOOP_STAGE_EVENTS_DIR}", schema=eventSchema)
df.write.save(
    f"{HADOOP_URI}/{HADOOP_FINAL_EVENTS_DIR}", format="parquet", mode="append"
)
