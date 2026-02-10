# ========================
# Stage 1: Build
# ========================
FROM golang:1.24-alpine AS builder

WORKDIR /app

# نصب ماژول‌ها
COPY go.mod go.sum ./
RUN go mod download

# کپی کل پروژه
COPY . .

# ساخت باینری
RUN go build -o server cmd/server/main.go

# ========================
# Stage 2: Run
# ========================
FROM alpine:latest

WORKDIR /app

# نصب curl برای تست Swagger (اختیاری)
RUN apk add --no-cache curl

# کپی باینری از مرحله build
COPY --from=builder /app/server .

# کپی docs برای Swagger
COPY --from=builder /app/docs ./docs

# پورت
EXPOSE 8080

# اجرای برنامه
CMD ["./server"]
