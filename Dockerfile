FROM golang:1.20 as builder

ARG TARGETOS
ARG TARGETARCH
ARG TARGET

RUN echo "building for $TARGETOS/$TARGETARCH"

WORKDIR /go/src/app
COPY ../ .
RUN make --file=/go/src/app/Makefile TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH $TARGET

FROM scratch

WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./kbot", "start"]