FROM golang:1.21-bullseye as builder

RUN apt update && \
    apt-get install -y \
        build-essential \
        ca-certificates \
        curl

WORKDIR /build

# cache dependencies.
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build go install -v ./...

RUN make build

FROM debian:bullseye

RUN useradd -m w3t && \
    apt update && \
    apt install -y \
        ca-certificates \
        curl

USER w3t

COPY --from=builder /go/bin/w3t /bin

WORKDIR /apps

ADD config.yaml /etc/w3t_exporter/config.yaml

EXPOSE 9115
ENTRYPOINT  [ "/bin/w3t" ]
CMD         [ "server" ]