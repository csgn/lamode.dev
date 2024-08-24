# Authorization
The Authorization service validates the given users' credentials, stores
session information on `Session` ***(in memory cache)*** if credentials are 
valid, and then returns a `token` to the users. Otherwise, returns error message.

# Used Technologies
| Tecnology             | Purpose                                                |
| --------------------- | ------------------------------------------------------ |
| ```Scala```           | It serves users' requests and performs business logic. |
| ```CouchDB```         | It stores users' login informations.                   |
| ```Redis```           | It stores users' session informations on memory.       | 

# Architecture
```mermaid
flowchart
    %% Event Collector Components
        EXTERNAL_USER(("User"))
        API("Scala")
        DB[("CouchDB")]
        SESS[("Redis")]

        %% Style
            style API fill:#DD3735,stroke:#7F0C1D,color:#fff
            style DB fill:#22525E,stroke:#001E26;,color:#fff
            style SESS fill:#ff4438,stroke:#001E26;,color:#fff
    %% %

    %% Blocks
        subgraph "authorization_service" ["Authorization Service"]
            API
            DB
            SESS
        end
    %% %

    %% Relations
        EXTERNAL_USER --request--> API
        API --response--> EXTERNAL_USER

        API <-.-> DB
        API <-.-> SESS
    %% %
```

```mermaid
sequenceDiagram
    actor       User
    participant API
    participant DB
    participant Cache

    autonumber
    User ->>+ API: Login credentials
    API -->>- User: Validate credentials

    alt Invalid credentials
        API -->> User: Error message
    else Valid credentials
        API ->>+ DB: Connect to DB
        DB -->>- API: Check if user exists

        alt User does not exist, not authorized
            API -->> User: Error Message
        else User authorized
            Cache ->> API: get or update session
            API -->> User: Session token
        end
    end
```
