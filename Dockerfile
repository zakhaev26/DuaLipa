# server containerization
FROM golang:1.17
WORKDIR /app
COPY . .
RUN go build main.go
EXPOSE 6969
CMD ["./main"]