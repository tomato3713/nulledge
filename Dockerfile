FROM golang:1.21-bookworm AS builder
ENV WORKDIR /app
WORKDIR $WORKDIR

COPY . $WORKDIR

RUN go mod download
RUN go install github.com/cespare/reflex
RUN go build -o /nulledge ${WORKDIR}/server.go

FROM debian:bookworm-slim AS release
WORKDIR /

COPY --from=builder /nulledge /nulledge

EXPOSE 8080

CMD ["/nulledge"]
