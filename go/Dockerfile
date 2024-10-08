# ---
FROM --platform=$BUILDPLATFORM registry-1.docker.io/library/golang:1.23-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . /build/

ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "-s -w" -a -o /out/app ./main.go

# ---
FROM --platform=$BUILDPLATFORM registry-1.docker.io/library/alpine as certs
RUN apk update && apk add ca-certificates

# ---
FROM registry-1.docker.io/library/busybox:1.36.1

ARG TARGETOS TARGETARCH
ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static-$TARGETARCH /tini

RUN chmod +x /tini

RUN mkdir /app
RUN addgroup -S app && adduser -S -G app app
WORKDIR /app

COPY --from=builder /out/app .
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

RUN chown -R app:app .
USER app

EXPOSE 3000

ENTRYPOINT [ "/tini", "--", "./app" ]
