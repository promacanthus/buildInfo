# buildInfo

Automatically set git commit information into the buildInfo, so that you can use it in your code.

# Installation

`go get -u github.com/promacanthus/buildInfo`

# Quick Start

Import the package at main.go.

```go
import _ "github.com/promacanthus/buildInfo"

func main() {
    // Use the buildInfo anywhere.
    info.Infos()
    // ...
    // Your application logic here.
}
```
Build the application as an executable binary. Run the following command to build the application.

```shell
go build -ldflags "-s -w -X github.com/promacanthus/buildInfo/info.GitCommit=$(git log -1 --pretty=format:%h)" -a -o main main.go
```

Build the application as a docker image. Add the following command into your dockerfile.

```Dockerfile
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} \
    go build -ldflags "-s -w -X github.com/promacanthus/buildInfo/info.GitCommit=$(git log -1 --pretty=format:%h)" -a -o main main.go
```

Run the following command to test it the commit info inject successfully.

`go run -ldflags "-X github.com/promacanthus/buildInfo/info.GitCommit=a1b2c3d4" main.go`