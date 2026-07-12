FROM golang:1.26.1-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/api/exe /app/cmd/api

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/cmd/api/exe /app
COPY --from=builder /app/frontend /app/frontend
ENV PROJECT_ROOT=/app
CMD ["/app/exe"]