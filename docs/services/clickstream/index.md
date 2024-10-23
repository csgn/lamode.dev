<p>
    <a href="/docs/index.md">Home</a> /
    <a href="/docs/services/index.md">Services</a> /
    <span>Clickstream</span>
</p>

# Overview
The **Clickstream** service will collect user events and store the raw data on the 
```Apache Hadoop``` cluster, and **internal consumers** can access the cluster via
```Apache Spark```, ```Presto``` etc. to read processed data in order to use
it in their business logic.

# Purpose of Service
We are collecting user events for the reasons listed below:

1. Recommend better **product** or **content** to users based on their 
historical behaviour.
2. Creating **advertising campaigns** based on users' behaviour.
3. Understand how **different page layouts** or **contents** affect users'
behaviour.

# Architecture
```mermaid
flowchart TD
    %% Event Collector Components
        EXTERNAL_USER(("User"))
        EC_GO("Go")
        EC_KAFKA(["Kafka"])
        EC_VALIDATOR("Validator")

        %% Style
            style EC_GO fill:#00ADD8,stroke:#00728f,color:#fff
            style EC_KAFKA fill:#444,stroke:#333,color:#fff
            style EC_VALIDATOR fill:#DF3D3B,stroke:#DF3D3B,color:#fff
    %% %

    %% Storage Components
        HADOOP[("Hadoop")]

        %% Style
            style HADOOP fill:#f9e64a,stroke:#f79625,color:#000
    %% %

    %% Data Processing Components
        DP_SPARK("Spark")
        DP_HIVE[("Hive")]

        %% Style
            style DP_SPARK fill:#e25a1c,stroke:#b13700,color:#fff
            style DP_HIVE fill:#fdee21,stroke:#bdb002,color:#000
    %% %

    %% Internal Consumers Components
        IC_SPARK("Spark")
        IC_PRESTO("Presto")

        %% Style
            style IC_SPARK fill:#e25a1c,stroke:#b13700,color:#fff
            style IC_PRESTO fill:#5d88d6,stroke:#000,color:#fff
    %% %

    %% Subgraphs
        subgraph "event_collector" ["Event Collector"]
            EC_GO
            EC_KAFKA
            EC_VALIDATOR 
        end


        subgraph "data_processing" ["Data Processing"]
            DP_SPARK
            DP_HIVE
        end

        subgraph "internal_consumers" ["Internal Consumers"]
            IC_SPARK
            IC_PRESTO
            ...((...))
        end
    %% %

    %% Relations
        EXTERNAL_USER -.-> event_collector
        
        EC_GO -- "raw event" --> EC_KAFKA
        EC_KAFKA -- "raw event" --> EC_VALIDATOR 
        EC_VALIDATOR -- "validated event" --> HADOOP

        HADOOP -- "read raw event" --> DP_SPARK

        DP_SPARK -- "write processed event" --> HADOOP
        DP_SPARK -- "write metadata" --> DP_HIVE

        HADOOP -. "read" .-> internal_consumers
    %% %
```

# Components
[🔗 Documentation of Clickstream Components](/services/clickstream/docs/index.md)
