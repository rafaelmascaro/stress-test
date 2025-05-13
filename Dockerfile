FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o tester ./cmd/stresstest/main.go

FROM scratch
COPY --from=builder /app/tester /tester
ENTRYPOINT ["/tester"]