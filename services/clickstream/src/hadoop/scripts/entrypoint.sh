#!/bin/bash
set -e

sudo service ssh start
# sudo service cron start

if [ ! -d "/tmp/hadoop/dfs/name" ]; then
    $HADOOP_HOME/bin/hdfs namenode -format && echo "OK : HDFS namenode format operation finished successfully !"
fi

$HADOOP_HOME/sbin/start-dfs.sh

echo "YARNSTART = $YARNSTART"
if [[ -z $YARNSTART || $YARNSTART -ne 0 ]]; then
    echo "running start-yarn.sh"
    $HADOOP_HOME/sbin/start-yarn.sh
fi

$HADOOP_HOME/bin/hdfs dfs -mkdir /data
$HADOOP_HOME/bin/hdfs dfs -mkdir /data/raw
$HADOOP_HOME/bin/hdfs dfs -mkdir /data/stage
$HADOOP_HOME/bin/hdfs dfs -mkdir /data/archive
$HADOOP_HOME/bin/hdfs dfs -mkdir /data/final
$HADOOP_HOME/bin/hdfs dfs -mkdir /tmp
$HADOOP_HOME/bin/hdfs dfs -mkdir /users
$HADOOP_HOME/bin/hdfs dfs -mkdir /checkpoints
$HADOOP_HOME/bin/hdfs dfs -chmod 777 /data
$HADOOP_HOME/bin/hdfs dfs -chmod 777 /tmp
$HADOOP_HOME/bin/hdfs dfs -chmod 777 /users
$HADOOP_HOME/bin/hdfs dfs -chmod 777 /checkpoints
$HADOOP_HOME/bin/hdfs dfsadmin -safemode leave

# crontab -u root /app/cron 

# keep the container running indefinitely
tail -f $HADOOP_HOME/logs/hadoop-*-namenode-*.log
