ARG GO_VERSION

FROM golang:${GO_VERSION}

WORKDIR /workdir
COPY . /workdir

# ENTRYPOINT ["go", "run", "main.go"]
# CMD ["sleep", "1d"]
ENTRYPOINT ["tail", "-f", "/dev/null"]
