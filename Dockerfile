# Build Stage
FROM golang:1.18-alpine AS builder

# Define the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copies the source code into the container
COPY ./cmd ./cmd
COPY ./configs ./configs
COPY ./pkg ./pkg
COPY ./scripts ./scripts

# Compile the application to a binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o clock ./cmd/clock/main.go

# Final Stage
FROM alpine:latest  

WORKDIR /root/

# Copiar o binário compilado do estágio de build para o estágio final
COPY --from=builder /app/clock .
COPY --from=builder /app/configs/clock_config.json .
COPY --from=builder /app/scripts/wait-for-it.sh .


# Comando para executar o aplicativo
CMD ["./clock", "clock_config.json"]