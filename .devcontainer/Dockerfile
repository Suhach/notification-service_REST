FROM golang:1.24-alpine

RUN apk add --no-cache \
    docker-cli\
    git \
    curl \
    bash \
    make \
    build-base 

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1 \
    GO111MODULE=on  \
    PATH="/go/bin:${PATH}"

CMD ["bash"]