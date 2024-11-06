FROM apache/airflow:2.10.3 AS airflow-base
ENV DEBIAN_FRONTEND noninteractive
ENV TERM linux

USER root
RUN apt-get update \
  && apt-get install -y --no-install-recommends \
         openjdk-17-jre-headless procps nano wget \
  && apt-get autoremove -yqq --purge \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64

FROM airflow-base AS spark-setup
ENV SPARK_VERSION=3.5.3
ENV SPARK_HOME=/opt/spark
RUN export SPARK_TMP="$(mktemp -d)" && \
    cd ${SPARK_TMP} && \
    wget -nv -O spark.tgz "https://dlcdn.apache.org/spark/spark-${SPARK_VERSION}/spark-${SPARK_VERSION}-bin-hadoop3.tgz" && \
    tar -xvzf spark.tgz && \
    mkdir -p "${SPARK_HOME}/bin" && \
    mkdir -p "${SPARK_HOME}/assembly/target/scala-2.12/jars" && \
    cp -a "spark-${SPARK_VERSION}-bin-hadoop3/bin/." "${SPARK_HOME}/bin/" && \
    cp -a "spark-${SPARK_VERSION}-bin-hadoop3/jars/." "${SPARK_HOME}/assembly/target/scala-2.12/jars/" && \
    rm -rf "spark-${SPARK_VERSION}-bin-hadoop3" "spark.tgz"
ENV PATH=$SPARK_HOME/bin:$PATH

FROM spark-setup AS hadoop-setup
ENV HADOOP_VERSION=3.4.1
ENV HADOOP_HOME=/opt/hadoop
RUN export HADOOP_TMP="$(mktemp -d)" && \
    cd ${HADOOP_TMP} && \
    wget -nv -O hadoop.tar.gz "https://dlcdn.apache.org/hadoop/common/hadoop-${HADOOP_VERSION}/hadoop-${HADOOP_VERSION}.tar.gz"; \
    tar -xvzf hadoop.tar.gz; \
    mkdir "${HADOOP_HOME}"; \
    cp -a "hadoop-${HADOOP_VERSION}/." "${HADOOP_HOME}"; \
    rm -rf "hadoop-${HADOOP_VERSION}" hadoop.tar.gz
ENV PATH=${HADOOP_HOME}/bin:$PATH

FROM hadoop-setup AS final-setup
USER airflow
COPY requirements.txt .
RUN pip install -r requirements.txt

FROM final-setup AS runner
COPY . .
#ENV AIRFLOW__CORE__DAGS_ARE_PAUSED_AT_CREATION=False

ENTRYPOINT [ "airflow", "standalone" ]
