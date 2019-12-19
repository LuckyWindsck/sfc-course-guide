ARG GO_VERSION

FROM golang:${GO_VERSION}

WORKDIR /workdir
COPY . /workdir

# ENTRYPOINT ["tail", "-f", "/dev/null"]
ENTRYPOINT ["go", "run", "cmd/server/main.go"]
