FROM golang:1.20-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 8448

CMD [ "/app/main" ]
