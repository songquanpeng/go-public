<p align="right">
   <a href="README.md">中文</a> | <strong>English</strong>
</p>

<div align="center">

# Go Public

_✨ Easily forward your local port to the public network. ✨_

</div>

<p align="center">
  <a href="https://raw.githubusercontent.com/songquanpeng/go-public/master/LICENSE">
    <img src="https://img.shields.io/github/license/songquanpeng/go-public?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/songquanpeng/go-public/releases/latest">
    <img src="https://img.shields.io/github/v/release/songquanpeng/go-public?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://github.com/songquanpeng/go-public/releases/latest">
    <img src="https://img.shields.io/github/downloads/songquanpeng/go-public/total?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://hub.docker.com/repository/docker/justsong/go-public">
    <img src="https://img.shields.io/docker/pulls/justsong/go-public?color=brightgreen" alt="docker pull">
  </a>
  <a href="https://goreportcard.com/report/github.com/songquanpeng/go-public">
  <img src="https://goreportcard.com/badge/github.com/songquanpeng/go-public" alt="GoReportCard">
  </a>
</p>

## Features
+ [x] Very easy to use
+ [x] Support TCP
+ [ ] Support UDP
+ [x] Support IP whitelist

## Usages

### Server Side

```bash
# Init config file
./go-public init server
# Save the generated token, will be used in the client side for authentication
cat go-public-server.yaml
# Start the server
./go-public
```

Or you can use docker to run the server:
```bash
docker run -d --restart always --name go-public -p 6871:6871 -p 8080:8080 -v /home/ubuntu/data/go-public:/app justsong/go-public
```

IP whitelist configuration example:
```yaml
# go-public-server.yaml
whitelist:
  - 123.213.241.5
  - 123.213.242.9
  - 125.216.243.1
```

```bash

### Client Side

```bash
# Initialize config file
./go-public init client
# Modify the config file with your saved token
vim go-public-client.yaml
# Start the client
# Please be aware that the remote port is not the port that the server listens on 
# as specified in the configuration file, but rather the port on which you want 
# to map the local port.
./go-public <local_port> <remote_port>
# For example:
./go-public 3000 8080  # Forward local port 3000 to remote port 8080
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