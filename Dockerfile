# syntax=docker/dockerfile:1

FROM golang:1.17.2-alpine

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /app/app

ENTRYPOINT /app/app

EXPOSE 3001

CMD [ "/app/app" ]
