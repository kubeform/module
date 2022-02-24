# Build the manager binary
FROM golang:1.16 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM ubuntu:20.04
WORKDIR /
COPY --from=builder /workspace/manager .

RUN set -x \
  && apt-get update \
  && apt-get install -y --no-install-recommends apt-transport-https ca-certificates curl unzip git openssh-server

RUN set -x \
  && curl -O -fsSL https://releases.hashicorp.com/terraform/1.1.5/terraform_1.1.5_linux_amd64.zip \
  && unzip terraform_1.1.5_linux_amd64.zip \
  && chmod 755 terraform \
  && mv terraform /bin/terraform

ENTRYPOINT ["/manager"]
