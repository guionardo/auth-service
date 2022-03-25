# auth-service
Auth and authorization generic microsservice

## Features

* Authentication by credentials
* Authorization by JWT token
  

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
    CBP -- no --> UE{is password valid?};
    UE -- yes --> GT(generates token) --> ST[(stores token and\n user data in cache)] --> R(token response 200);
    UE -- no --> IBP(register last bad password\nrequest timestamp) --> UR(unauthorized response 401);
    U -- no --> UR(unauthorized response 401);    
```

### Authorization by token

```curl
curl {HOST}/auth -H "Accept: application/json" -H "Authorization: Bearer {token}"
```

```mermaid
flowchart TD
    Z[Auth] --> Y(request with token);
    Y --> C[(cache)] --> QC{is token in cache?};
    QC -- yes --> QCV{is token still valid?};
    QCV -- yes --> R(user data response 200);
    QCV -- no --> UR(expired token response 401);
```