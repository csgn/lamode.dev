<p>
    <a href="/docs/index.md">Home</a> /
    <a href="/docs/services/index.md">Services</a> /
    <span>Clickstream</span>
</p>

<a href="/services/clickstream/README.md">README</a>

# Overview
The **Clickstream** service will collect user events and store the raw data on the 
```Apache Hadoop``` cluster, and **internal consumers** can access the cluster via
```Apache Spark```, ```Presto``` etc. to read processed data in order to use
it in their own business purposes.

# Purpose of Service
We are collecting user events for the reasons listed below:

1. Recommend better **product** or **content** to users based on their 
historical behaviour.
2. Creating **advertising campaigns** based on users' behaviour.
3. Understand how **different page layouts** or **contents** affect users'
behaviour.

# Architecture
![clickstream_diagram](/docs/resources/diagrams/clickstream-diagram.png)

[Edit diagram](diagram.mmd)

# Components
1. [Collector](collector/index.md)
2. [Kafka](kafka/index.md)
3. [Validator](validator/index.md)
4. [Hadoop](hadoop/index.md)
4. [Batch](batch/index.md)
