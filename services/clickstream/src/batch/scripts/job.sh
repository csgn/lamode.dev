#!/bin/bash

set -e

${SPARK_HOME}/bin/spark-submit \
    --deploy-mode client \
    /app/main.py
