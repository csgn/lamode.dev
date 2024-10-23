FROM eclipse-temurin:11-alpine AS builder

ENV JOB_NAME=validator
ENV SPARK_VERSION=3.5.1
ENV SCALA_VERSION=2.12
ENV SBT_VERSION=1.10.2
ENV SPARK_TGZ_URL=https://archive.apache.org/dist/spark/spark-${SPARK_VERSION}/spark-${SPARK_VERSION}-bin-hadoop3.tgz
ENV SBT_TGZ_URL=https://github.com/sbt/sbt/releases/download/v${SBT_VERSION}/sbt-${SBT_VERSION}.tgz
ENV SPARK_HOME=/opt/spark
ENV SBT_HOME=/opt/sbt
ENV TARGET_JAR=/app/target/scala-2.12/${JOB_NAME}_2.12-${PROJECT_VERSION}.jar

RUN apk --update add wget tar bash
# REMOVE IN PROD
WORKDIR /spark-workdir
# REMOVE IN PROD
COPY ./dev_dependencies/spark.tgz .

RUN export SPARK_TMP="$(mktemp -d)"; \
    cd ${SPARK_TMP}; \
    # REMOVE IN PROD
    cp /spark-workdir/spark.tgz .; \
    # wget -nv -O spark.tgz "${SPARK_TGZ_URL}"; \
    tar -xf spark.tgz; \
    mv spark-${SPARK_VERSION}-bin-hadoop3 ${SPARK_HOME}; \
    # rm spark-${SPARK_VERSION}-bin-hadoop3.tgz; \
    ln -s ${SPARK_HOME}/bin/spark-submit /usr/bin/spark-submit

# REMOVE IN PROD
WORKDIR /sbt-workdir

# REMOVE IN PROD
COPY ./dev_dependencies/sbt.tgz .

RUN export SBT_TMP="$(mktemp -d)"; \
    cd ${SBT_TMP}; \
    # REMOVE IN PROD
    cp /sbt-workdir/sbt.tgz .; \
    #wget -nv -O sbt.tgz "${SBT_TGZ_URL}"; \
    tar -xf sbt.tgz; \
    mv sbt ${SBT_HOME};\
    # rm sbt.tgz;\
    ln -s ${SBT_HOME}/bin/sbt /usr/bin/sbt

WORKDIR /app
COPY . .

RUN sbt clean package

CMD ["sh", "-c", "spark-submit --packages org.apache.spark:spark-sql-kafka-0-10_${SCALA_VERSION}:${SPARK_VERSION} --deploy-mode client ${TARGET_JAR}"]
