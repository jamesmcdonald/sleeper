FROM golang as builder
WORKDIR /build
COPY sleeper.go /build
RUN go build sleeper.go

FROM debian
WORKDIR /app
COPY --from=builder /build/sleeper /app
ENTRYPOINT ./sleeper
