<p>
    <a href="/docs/index.md">Home</a> /
    <a href="/docs/services/index.md">Services</a> /
    <a href="/docs/services/clickstream/index.md">Clickstream</a> /
    <span>Batch</span>
</p>

<a href="/services/clickstream/src/batch/README.md">README</a>

# Overview
There are few jobs in order to perform ETL process and we are using `Apache Airflow` for scheduling them.

# Stages
![clickstream-batch-stages-diagram.png](/docs/resources/diagrams/clickstream-batch-stages-diagram.png)

[Edit diagram](stages-diagram.mmd)


| Name                          | Description                                                                                                   |
| ---                           | --                                                                                                            |
| move_raw_events_task          | Move all `raw events` from `raw/` directory into `stage/` directory for prepare the data in order to process. |
| process_raw_events_task       | Reads the `staged events` from `stage/` directory and performs data transformations and then writes the `transformed events` into `final/` directory.               |
| move_stage_events_task        | After the processing, move all `staged events` from `stage/` directory into `archive/` directory. It performs backing up the data for any fault or reprocessing circumstances.                  |

# Changelogs
- [v0.1.0-alpha.1](/services/clickstream/src/batch/CHANGELOG.md#v010-alpha1)
