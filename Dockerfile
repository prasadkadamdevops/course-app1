# ---------- Stage 1: Build ----------
FROM golang:1.23 AS builder

WORKDIR /app
COPY src/ ./

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o course-app main.go

# ---------- Stage 2: Distroless ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/course-app .

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app/course-app"]

