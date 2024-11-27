<p>
    <a href="/docs/index.md">Home</a> /
    <a href="/docs/services/index.md">Services</a> /
    <a href="/docs/services/clickstream/index.md">Clickstream</a> /
    <span>Collector</span>
</p>

<a href="/services/clickstream/src/collector/README.md">README</a>

# Overview
The `Collector`, collects the `raw event(s)` data from the source and sends it
into the specified `Kafka` topic. The Collector provides an HTTP server 
and the server provides [endpoints](#endpoints) in order to collects event(s).

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

Both endpoint responses are returns `200 OK` because we don't want to show an error
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
As follows, the JSON shows us the event types:
```json
{
    // Required
    "sid": "string",            // Logged in user's or guest's session id.
    "channel": "string",        // Which source does it come from? (e.g., mobile, web etc.)
    "timestamp": "string",      // ISO 8601 formatted timestamp (e.g., "2024-11-27T15:30:00Z")
    "culture": "string",        // Which language and country does it come from? (e.g., tr_TR, en_US, en_UK, etc.)
    "action": "string",         // How the event was occured. (e.g., click, impression)
    "type": "string",           // What is the event's type? (e.g., search, ad, component, etc.)
    "ipAddress": "string",      // IP address of the client
    "latitude": "float",        // Geolocation data
    "longitude": "float",       // Geolocation data

    // Optional
    "productId": "string",      // Which product was interacted with?
    "searchQuery": "string",    // What was searched?
}
```

# Changelogs
- [v0.1.0-alpha.1](/services/clickstream/src/collector/CHANGELOG.md#v010-alpha1)
