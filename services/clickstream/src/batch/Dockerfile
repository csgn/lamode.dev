FROM apache/spark-py

USER root

ENV CRON_LOG_PATH=/var/log/cron.log
ENV CRON_ERROR_LOG_PATH=/var/log/cron_error.log

RUN apt-get update && apt-get install -y cron
RUN touch ${CRON_LOG_PATH} ${CRON_ERROR_LOG_PATH}

WORKDIR /app
COPY . .

RUN crontab -u root /app/cron
CMD printenv > /etc/environment && cron -f && tail -f ${CRON_LOG_PATH}
