<p>
    <a href="/docs/index.md">Home</a> /
    <a href="/docs/services/index.md">Services</a> /
    <a href="/docs/services/clickstream/index.md">Clickstream</a> /
    <a href="/services/clickstream/docs/index.md">Components</a> /
    <span>Collector</span>
</p>

<a href="/services/clickstream/src/collector/README.md">README</a>

# Overview
The `Collector`, collects the `raw event(s)` data from the source and sends it
into the specified `Kafka` topic. The Collector provides an HTTP server 
and the server provides [endpoints](#endpoints) in order to collects event(s).

# Architecture
```mermaid
flowchart TD
    %% Event Collector Components
        EC_GO("Go")
        EC_KAFKA(["Kafka"])
        EXTERNAL_USER(("User"))

        %% Style
            style EC_GO fill:#00ADD8,stroke:#00728f,color:#fff
            style EC_KAFKA fill:#444,stroke:#333,color:#fff
    %% %

    %% Subgraphs
        subgraph "event_collector" ["Event Collector"]
            EC_GO
            EC_KAFKA
        end
    %% %

    %% Relations
        EXTERNAL_USER -.-> event_collector
        
        EC_GO -- "raw event" --> EC_KAFKA
    %% %
```
> [Diagram Reference](/docs/services/clickstream/index.md#architecture)


# Endpoints
```http
POST /e HTTP/1.1
Accept: application/json
Content-Length: 48
Content-Type: application/json

{
    "tv001": "111",
    "tv002": "123",
    "tv003": "web"
}


HTTP/1.1 200 OK
Content-Length: 0
Date: Thu, 03 Oct 2024 13:56:10 GMT
```
---

```http
GET /pixel?tv001=111&tv002=123&tv003=mobile HTTP/1.1
Accept: */*

HTTP/1.1 200 OK
Content-Length: 42
Content-Type: image/gif
Date: Thu, 03 Oct 2024 13:56:42 GMT
```

## Notes
In principle, events are sent by means of `Beacon API` but some 
clients, however, may not support. Therefore, we should use 
the `/pixel` endpoint in that case. The `/pixel` sends the event data within
its query and returns a `1px` image as a response.

Both endpoint responses are returns `200 OK` because we won't show an error
on the client side or resend the event. We always assume that the events which we sent,
are always arrived to the Collector.

As follows, there are some examples of how we can use these endpoints:

```js
// With Beacon API
navigator.sendBeacon(<COLLECTOR_ORIGIN>, <EVENT_DATA>);

// Without Beacon API
var img = document.createElement("img");
img.src = <COLLECTOR_ORIGIN_WITH_QUERY>
img.style.display = "none";
document.body.appendChild(img);
img.onload = function() {
    img.remove();
};
```

Eventually, it's up to you that how to use these endpoints in your project.

# Events
As follows, the table shows us the event types:

| Code              | Value                                 | Description           |
| -----             | ---                                   | -----                 |
| `pid`             | `string`                              | `pid`                 |
| `sid`             | `string`                              | `sid`                 |
| `channel`         | `web` \| `mobile`                     | `channel`             |
| `event`           | `click` \| `impression`               | `event action`        |
| `eventGroup`      | `search` \| `ad` \| `component`       | `event group`         |
| `senderMethod`    | `beacon` \| `other`                   | `sender method`       |
| `ip`              | `string`                              | `ip address`          |

# Changelogs
- [v0.1.0-alpha.1](/services/clickstream/src/collector/CHANGELOG.md#v010-alpha1)
