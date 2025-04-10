#Building the base image
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o feedbackservice

#Building the final image AKA Lighter image
FROM alpine:latest

WORKDIR /root/ 

COPY --from=builder /app/feedbackservice . 

EXPOSE 8080

CMD ["./feedbackservice"]



