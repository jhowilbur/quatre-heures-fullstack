FROM golang:1.20-alpine as develop-stage
LABEL authors="wilbur"

WORKDIR /app
COPY . .

RUN go build -o myapp

EXPOSE 8081

CMD ["./myapp", "db", "migrate", "./myapp", "web"]