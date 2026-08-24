FROM golang:1.22-alpine

WORKDIR /app

# 1. Copy dependency files first
COPY go.mod  ./

# 2. Download dependencies
RUN go mod download

# 3. Copy application source code
COPY . .

# 4. Build application
RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
