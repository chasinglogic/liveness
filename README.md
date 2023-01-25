# Liveness Checks Made Easy

When working in a Kubernetes environment you should be defining meaningful
[readinessProbes and livenessProbes.](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/)
However, most teams implement a simple `/healthz` endpoint which only confirms
that the server is ready to serve requests. This is not a great indicator of
"liveness" as there are often dependencies which need to be alive and reachable
as well such as a database. If you already include a DB connection attempt and
cache ping in your `/healthz` endpoint then this tool isn't for you. But if
you're looking for a quick way to add meaningful liveness to your kubernetes
pods then read on. 

`liveness` is a simple process which you can run either inside your Docker
container as a background process or as a sidecar container. It takes a list of
resources via command line flags which it will monitor for liveness. It runs an
HTTP server that whenever it receives a request to any URL it will check all the
provided resources to ensure they are still accepting connections. If any
resource has become unavailable it will respond with a 503 Server Unavailable
response otherwise, it will send a 200.

In this way you don't need to reinvent the wheel if you're just looking for
simple liveness. It also is very useful for long running services which aren't
HTTP server likes background workers (such as those provided by Celery) since it
can add that functionality to check all the relevant dependencies are still
alive.

## Installation

Download the latest binary from the [Github Releases](https://github.com/chasinglogic/liveness/releases) 
and copy this into your container.

Alternatively you can add this snippet to your Dockerfile to install as part of your image build:

```dockerfile
ARG LIVENESS_VERSION=0.1.0
# Download the liveness release tarball for our LIVENESS_VERSIOn and extract the
# binary into /usr/bin. The results of uname -m must be passed through sed since
# it returns aarch64 but the release tarballs use arm64.
RUN curl -o /tmp/liveness.tar.gz -L https://github.com/chasinglogic/liveness/releases/download/${LIVENESS_VERSION}/liveness_Linux_$(uname -m | sed s/aarch64/arm64/).tar.gz && \
    tar -C /usr/bin -xzvf /tmp/liveness.tar.gz liveness
```

> **Note:** You'll often need to install curl before this with a command like
> `RUN apt update && apt install curl` but this will depend on your base image so
> it is omitted here.

This layer will be cached based on the result of LIVENESS_VERSION. You can pass
in the version you want to install as a build arg like so:

```
docker build --build-arg LIVENESS_VERSION=0.1.0 .
```

## Usage

## Supported Check Types

### HTTP checks

### Postgres checks

### TCP checks

