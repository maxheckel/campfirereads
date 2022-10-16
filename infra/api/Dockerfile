FROM golang:1.19-alpine

WORKDIR /app
COPY . .
ENV SERVER_PORT 8080
EXPOSE 8080
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]