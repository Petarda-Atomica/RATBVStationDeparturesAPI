# STAGE 1: Build the binary
FROM golang:1.25.9-alpine AS builder

# Install git
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy and download dependencies first (improves build caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the app 
# CGO_ENABLED=0 ensures the binary is statically linked
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# STAGE 2: Run the binary
FROM alpine:latest  

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]