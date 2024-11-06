FROM eclipse-temurin:11-jdk-focal

RUN apt-get update -y \
    && export DEBIAN_FRONTEND=noninteractive && apt-get install -y --no-install-recommends sudo ssh wget cron \
    && apt-get clean

COPY conf/ssh_config /etc/ssh/ssh_config

RUN ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa && cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys && chmod 0600 ~/.ssh/authorized_keys

ENV HADOOP_VERSION=3.4.0
ENV HADOOP_TGZ_URL=https://dlcdn.apache.org/hadoop/common/hadoop-${HADOOP_VERSION}/hadoop-${HADOOP_VERSION}.tar.gz 
ENV HADOOP_HOME=/hadoop

WORKDIR /hadoop
RUN wget -nv -O hadoop.tar.gz "${HADOOP_TGZ_URL}"; \
    tar -xf hadoop.tar.gz --strip-components 1; \
    rm hadoop.tar.gz; \
    rm -rf ${HADOOP_HOME}/share/doc

ENV HDFS_NAMENODE_USER=root
ENV HDFS_DATANODE_USER=root
ENV HDFS_SECONDARYNAMENODE_USER=root

ENV YARN_RESOURCEMANAGER_USER=root
ENV YARN_NODEMANAGER_USER=root

RUN echo "export JAVA_HOME=/opt/java/openjdk/" >> $HADOOP_HOME/etc/hadoop/hadoop-env.sh
COPY conf/core-site.xml $HADOOP_HOME/etc/hadoop/
COPY conf/hdfs-site.xml $HADOOP_HOME/etc/hadoop/
COPY conf/yarn-site.xml $HADOOP_HOME/etc/hadoop/

COPY scripts/entrypoint.sh $HADOOP_HOME/etc/hadoop/
ENV PATH=$PATH:$HADOOP_HOME/bin:$HADOOP_HOME/sbin

EXPOSE 50070 50075 50010 50020 50090 8020 9000 9864 9870 10020 19888 8088 8030 8031 8032 8033 8040 8042 22

WORKDIR /usr/local/bin
RUN sudo ln -s ${HADOOP_HOME}/etc/hadoop/entrypoint.sh .

ENV YARNSTART=0

WORKDIR /app

ENTRYPOINT [ "/usr/local/bin/entrypoint.sh" ]
