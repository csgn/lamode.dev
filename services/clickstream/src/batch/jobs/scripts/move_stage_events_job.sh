#!/bin/bash
set -e

STAGE_EVENTS_DIR="${HADOOP_URI}/${HADOOP_STAGE_EVENTS_DIR}"
ARCHIVE_EVENTS_DIR="${HADOOP_URI}/${HADOOP_ARCHIVE_EVENTS_DIR}"

isEmpty=$(hdfs dfs -ls ${STAGE_EVENTS_DIR} | grep '\.json$' | wc -l)
if [[ isEmpty -eq 0 ]];then
    echo "There is no elements to move. Skipping..."
else
    hdfs dfs -mv ${STAGE_EVENTS_DIR}/*.json ${ARCHIVE_EVENTS_DIR}
fi
