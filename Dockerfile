FROM golang:1.16.7-alpine3.14

WORKDIR /src

COPY . .

RUN go mod tidy

RUN go build -o /app/bin/main

RUN rm -R /src

ENTRYPOINT [ "/app/bin/main" ]