<p>
    <a href="/docs/index.md">Home</a> /
    <a href="/docs/services/index.md">Services</a> /
    <a href="/docs/services/clickstream/index.md">Clickstream</a> /
    <a href="/services/clickstream/docs/index.md">Components</a> /
    <span>Validator</span>
</p>

# Overview
The `Validator` is a Spark streaming job in order to validate events.

# Architecture
```mermaid
flowchart TD
    %% Event Collector Components
        EC_KAFKA(["Kafka"])
        EC_VALIDATOR("Validator")

        %% Style
            style EC_KAFKA fill:#444,stroke:#333,color:#fff
            style EC_VALIDATOR fill:#DF3D3B,stroke:#DF3D3B,color:#fff
    %% %

    %% Storage Components
        HADOOP[("Hadoop")]

        %% Style
            style HADOOP fill:#f9e64a,stroke:#f79625,color:#000
    %% %

    %% Subgraphs
        subgraph "event_collector" ["Event Collector"]
            EC_KAFKA
            EC_VALIDATOR
        end
    %% %

    %% Relations
        EC_KAFKA -- "raw event" --> EC_VALIDATOR 
        EC_VALIDATOR -- "validated event" --> HADOOP
    %% %
```
> [Diagram Reference](/docs/services/clickstream/index.md#architecture)

# Changelogs
- [v0.1.0-alpha.1 - 23/10/2024](/services/clickstream/src/validator/CHANGELOG.md#v010-alpha1---23102024)
