from datetime import datetime, timedelta

from airflow.models.dag import DAG
from airflow.operators.dummy_operator import DummyOperator
from airflow.operators.bash_operator import BashOperator
from airflow.contrib.operators.spark_submit_operator import SparkSubmitOperator

with DAG(
    "dag",
    description="A simple DAG",
    schedule="@daily",
    start_date=datetime(2024, 11, 8),
    catchup=False,
) as dag:
    start = DummyOperator(
        task_id="start_task",
    )

    # if this task failed go end otherwise go t2
    t1 = BashOperator(
        task_id="move_raw_events_task",
        bash_command="/opt/airflow/jobs/scripts/move_raw_events_job.sh ",
    )

    # run this task when t1 is succeed
    t2 = SparkSubmitOperator(
        task_id="process_raw_event_task",
        application="/opt/airflow/jobs/python/process_raw_event.py",
        env_vars={
            "HADOOP_CONF_DIR": "/opt/hadoop",
            "YARN_CONF_DIR": "/opt/hadoop",
        },
    )

    t3 = BashOperator(
        task_id="move_stage_events_task",
        bash_command="/opt/airflow/jobs/scripts/move_stage_events_job.sh ",
    )

    end = DummyOperator(
        task_id="end_task",
        trigger_rule="one_failed",
    )

    start >> t1
    t1 >> t2
    t1 >> end

    t2 >> t3
    t2 >> end

    t3 >> end
