# --- Stage 1: Build ---
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary with optimizations (static linking)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# --- Stage 2: Final Run ---
FROM alpine:3.19

# Set working directory
WORKDIR /app

# Install timezone data & certificates
RUN apk --no-cache add ca-certificates tzdata

# Copy binary from builder
COPY --from=builder /app/main .

# Copy public folder for static assets (if any)
COPY --from=builder /app/public ./public

# Expose port
EXPOSE 8080

# Run the app
CMD ["./main"]
