FROM golang:latest AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

WORKDIR /app
COPY ./ ./
RUN GCO_ENABLED=0 GOOS=linux go build -o thicccat

FROM ubuntu:mantic 
COPY --from=builder /app/thicccat /thicccat
EXPOSE 4917 9093
ENTRYPOINT ["./thicccat"]