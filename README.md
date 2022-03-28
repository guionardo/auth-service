# auth-service
Auth and authorization generic microsservice

[![CodeQL Golang](https://github.com/guionardo/auth-service/actions/workflows/codeql-analysis-golang.yml/badge.svg)](https://github.com/guionardo/auth-service/actions/workflows/codeql-analysis-golang.yml)
[![CodeQL Python](https://github.com/guionardo/auth-service/actions/workflows/codeql-analysis-python.yml/badge.svg)](https://github.com/guionardo/auth-service/actions/workflows/codeql-analysis-python.yml)

## Features

* Authentication by credentials
* Authorization by JWT token

### User data feed upsert

```curl
curl -X POST {HOST}/user -d '{"userid":"","password_hash":"","payload":{}}' -H 'Content-Type: application/json' -H 'FEED-API-KEY: {feed API key}'

202 ACCEPTED
401 UNAUTHORIZED
```

```mermaid
flowchart TD
    Z[Upsert user data] --> Y(POST request with\nuser data);
    Y --> VK{Is valid FEED-API-KEY?};
    VK -- yes --> DB[(saves user data)] --> R(202 ACCEPTED);
    VK -- no --> RN(401 UNAUTHORIZED);

```

### User data feed delete

```curl
curl -X DELETE {HOST}/user/{userid} -H 'FEED-API-KEY: {feed API key}'

202 ACCEPTED
401 UNAUTHORIZED
```

```mermaid
flowchart TD
    Z[Delete user data] --> Y(DELETE request with\nuser id);
    Y --> VK{Is valid FEED-API-KEY?};
    VK -- yes --> DUD{user id is in database?};
    VK -- no --> RN(401 UNAUTHORIZED);
    DUD -- yes --> DB[(delete user data)] --> C[(remove all cached\nitems from user)] --> R(202 ACCEPTED);
    DUD -- no --> R(202 ACCEPTED)
```

### Authentication by credentials

```curl
curl -X POST {HOST}/auth -d '{"username":"","password":""} -H 'Content-Type: application/json'

200 { "data":"token" }
401 { "data":"UNAUTHORIZED" }
425 { "data":"WAIT BEFORE NEW LOGIN ATTEMPT" }
```

```mermaid
flowchart TD
    Z[Auth] --> Y(request with credentials) --> Q[(Database)];
    Q --> U{User exists?};
    U -- yes --> CBP{is last bad password\nrequest timestamp less\nthan 30 seconds?};
    CBP -- yes --> URR(unauthorized response 425);
    CBP -- no --> UE{is\npassword\nvalid?};
    UE -- yes --> GT(generates token) --> ST[(stores token and\n user data in cache)];
    ST --> DB[(register token by user id)] --> R(token response 200);
    UE -- no --> IBP(register last bad password\nrequest timestamp) --> UR(unauthorized\nresponse 401);
    U -- no --> UR(unauthorized\nresponse 401);    
```

### Authorization by token

```curl
curl {HOST}/auth -H "Accept: application/json" -H "Authorization: Bearer {token}"

200 { "data": { ... user data ...}}
401 { "data": "UNAUTHORIZED" }
404 { "data": "NOT FOUND" }
```

```mermaid
flowchart TD
    Z[Auth] --> Y(request\nwith token);
    Y --> C[(cache)] --> QC{is token\nin cache?};
    QC -- yes --> QCV{is token\nstill valid?};
    QC -- no --> RE(user not found 404)
    QCV -- yes --> R(user data\nresponse 200);
    QCV -- no --> UR(expired token\nresponse 401);
```

## Get refresh token

From a current valid token, get a special token, just used to refresh the current before it get expired.

```curl
curl {HOST}/refresh -H "Accept: application/json" -H "Authorization: Bearer {token}"

200 { "data": "new token" }
401 { "data": "UNAUTHORIZED" }
```

```mermaid
flowchart TD
    Z[Refresh] --> Y(request\nwith token);
    Y --> C[(cache)] --> QC{is token\nin cache?};
    QC -- yes --> RT[(remove old token\nfrom cache)] --> GT(generates token);
    QC -- no --> GT(generates token);
    GT --> ST[(stores token and\n user data in cache)] ;

    ST --> UR[(unregister old\ntoken by user id)];
    
    UR --> R(token response 200);
```
