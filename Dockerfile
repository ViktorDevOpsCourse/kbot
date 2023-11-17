FROM --platform=linux/arm64 golang:1.20 as builder

ARG TARGETOS
ARG TARGETARCH
ARG MAKEFILE_RULE

RUN echo "building for $TARGETOS/$TARGETARCH" > /log.

WORKDIR /go/src/app
COPY ../ .
RUN make --file=/go/src/app/Makefile TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH $MAKEFILE_RULE

FROM --platform=$BUILDPLATFORM scratch

ARG BUILDPLATFORM

WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./kbot", "start"]