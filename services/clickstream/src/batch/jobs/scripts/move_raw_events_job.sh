#!/bin/bash
set -e

RAW_EVENTS_DIR="${HADOOP_URI}/${HADOOP_RAW_EVENTS_DIR}"
STAGE_EVENTS_DIR="${HADOOP_URI}/${HADOOP_STAGE_EVENTS_DIR}"

isEmpty=$(hdfs dfs -count ${RAW_EVENTS_DIR} | awk '{print $2}')
if [[ isEmpty -eq 0 ]];then
    echo "There is no elements to move. Skipping..."
else
    hdfs dfs -mv ${RAW_EVENTS_DIR}/*.json ${STAGE_EVENTS_DIR}
fi
