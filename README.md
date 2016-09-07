# drone-sonar-runner

Drone plugin for publishing sonar-runner execution reports to a central sonar server. For the
usage information and a listing of the available options please take a look at
[the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t ypcloud/sonar-runner .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/sonar-runner' not found or does not exist..
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e DRONE_REPO=octocat/hello-world \
  -e DRONE_REPO_BRANCH=master \
  -e DRONE_COMMIT_BRANCH=master \
  -e PLUGIN_MOUNT=node_modules \
  -e SONAR_HOST=1.2.3.4:22 \
  -e SONAR_LOGIN=username \
  -e SONAR_PASSWORD=username \
  ypcloud/sonar-runner
```
