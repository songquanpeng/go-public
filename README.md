# Go Public
>  Easily forward your local port to the public network.

<p>
  <a href="https://raw.githubusercontent.com/songquanpeng/go-public/main/LICENSE">
    <img src="https://img.shields.io/github/license/songquanpeng/go-public?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/songquanpeng/go-public/releases/latest">
    <img src="https://img.shields.io/github/v/release/songquanpeng/go-public?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://github.com/songquanpeng/go-public/releases/latest">
    <img src="https://img.shields.io/github/downloads/songquanpeng/go-public/total?color=brightgreen&include_prereleases" alt="release">
  </a>
</p>

## Usages

### Server Side
```bash
# init config file
./go-public init server
# check & save the generated token
cat go-public-server.yaml
# start the server
./go-public
```

### Client Side
```bash
# init config file
./go-public init client
# modify the config file with your saved token
vim go-public-client.yaml
# start the client
./go-public <local_port> <remote_port>
```

## Flowchart
```mermaid
sequenceDiagram
    participant Local Server
    participant Go Public Client
    participant Go Public Server
    participant User
    
    Go Public Client->>Go Public Server: Initialize connection
    User->>Go Public Server: Request
    Go Public Server->>Go Public Client: Assign connection uuid
    Go Public Client->>Go Public Server: Make a new connection with uuid
    Go Public Client->>Local Server: Make a new connection
    Go Public Server->>Go Public Client: Forward request
    Go Public Client->>Local Server: Forward request
    Local Server->>Go Public Client: Response
    Go Public Client->>Go Public Server: Forward response
    Go Public Server->>User: Forward Response
```