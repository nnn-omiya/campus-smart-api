FROM golang:latest

WORKDIR /app

COPY ./api .

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

RUN go build -o main .

EXPOSE 8080

COPY ./setting/startup.sh ./startup.sh

RUN chmod 744 ./startup.sh
CMD ["./startup.sh"]
