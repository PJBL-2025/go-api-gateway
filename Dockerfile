# Base image Go versi 1.24.2 (sesuaikan dengan versi di go.mod)
FROM golang:1.24.2

# Set working directory di container
WORKDIR /app

# Copy file go.mod dan go.sum dulu agar layer cache bisa dipakai
COPY go.mod go.sum ./

# Download semua module dependencies
RUN go mod download

# Copy seluruh source code
COPY . .

# Build aplikasi, hasil binary bernama "api-gateway"
RUN go build -o api-gateway main.go

# Expose port yang akan digunakan aplikasi (misal 8000)
EXPOSE 8000

# Jalankan binary
CMD ["./api-gateway"]
