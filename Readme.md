<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="http://acresoftware.com">
    <img src="https://global-uploads.webflow.com/5d2cef25b4a76dcc2236ace0/5d2e058599a69033b540bbfb_side-by-side-blue.svg" alt="Logo" width="240" height="82">
  </a>

  <h3 align="center">Transforming the way we buy property</h3>

  <p align="center">
      A monorepo for all Acre's backend services
    <br />
</p>

```
TL;DR. All roads lead to README.md.

Everything you need to get started is documented here. We try to keep it dry and consice as much as we can. In-depth articles or related resources can be placed on Notion but they must be introduced here.

Enjoy.
```

<!-- TABLE OF CONTENTS -->

## Table of Contents

- [Architecture](#architecture)
  - [Built With](#built-with)
- [Go Directories](#go-directories)

<!-- ABOUT THE PROJECT -->

## Architecture

[![Acre-backend System Design][acre-backend-system-design]]()

### Built With

- [Go](https://golang.org)
- [gRPC](https://grpc.io/)
- [gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

### Go Directories

The following structure highlights some of the main components of our project.

```
  acre-backend
  ├── api
  │   ├── proto
  │   │   └── v1
  │   │       └── protection-service.proto
  │   └── swagger
  │       └── v1
  │           └── protection-service.swagger.json
  ├── cmd
  │   └── server
  │       └── server.go
  ├── pkg
  │   ├── api
  │   │   └── v1
  │   │       ├── protection-service.pb.go
  │   │       └── protection-service.pb.gw.go
  │   ├── cmd
  │   ├── protocol
  │   │   ├── rest
  │   │   │   ├── middleware
  │   │   │   └── server.go
  │   │   └── grpc
  │   │       ├── middleware
  │   │       └── server.go
  │   └── service
  │       └── v1
  │           └── protection-service.go
  ├── vendor
  ├── ...
  ├── Makefile
  ├── Gopkg.lock
  ├── Gopkg.toml
  └── ...
```

#### `/cmd`

Main applications for this project.

#### `/vendor`

Application dependencies managed by dependency management tool [`dep`](https://github.com/golang/dep)

### Service Application Directories

#### `/api`

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[config-workspace]: config/index.ts
[acre-backend-system-design]: assets/acre-backend-system-design.png
