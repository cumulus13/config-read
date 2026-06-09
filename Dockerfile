# Build stage
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git make build-base

WORKDIR /build

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w -X main.version=$(git describe --tags --always --dirty 2>/dev/null || echo 'dev') -X main.commit=$(git rev-parse --short HEAD 2>/dev/null || echo 'unknown') -X main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    -o /config-read ./cmd/config-read/

# Final stage
FROM alpine:3.19

RUN apk add --no-cache \
    ca-certificates \
    less \
    tzdata \
    # For extended character support
    musl-locales

# Create non-root user
RUN addgroup -g 1000 configread && \
    adduser -D -u 1000 -G configread configread

# Copy binary
COPY --from=builder /config-read /usr/local/bin/config-read

# Create config directory
RUN mkdir -p /home/configread/.config && \
    chown -R configread:configread /home/configread

# Set environment variables
ENV PAGER="less -R -F -X"
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

USER configread
WORKDIR /home/configread

ENTRYPOINT ["config-read"]
CMD ["--help"]
