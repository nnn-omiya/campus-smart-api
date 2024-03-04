#BUILD
FROM golang:latest AS build

WORKDIR /app

COPY ./api .

RUN go mod download

RUN go build -o main .

#DEPLOY

FROM gcr.io/distroless/base-debian12:latest

WORKDIR /

COPY --from=build /app/main /main

EXPOSE 8080

CMD ["/main"]
