FROM golang:1.17 as build

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY . /app/

RUN go build -a -o /app/main

# --------------------

FROM debian:stable-slim

WORKDIR /app

EXPOSE 8081

ENV APP_NAME=lambda_batch_processor
ENV APP_ADDRESS=:8081

ENV KAFKA_ADDRESS=172.17.0.1:29092,172.17.0.1:29093,172.17.0.1:29094
ENV KAFKA_CONSUMER_GROUP=audi.skripsi.lambda_batch_processor_group_tes
ENV KAFKA_IN_TOPIC=audi.skripsi.lambda_event_ingestion
ENV KAFKA_OUT_TOPIC=audi.skripsi.lambda_batch_event_identifier

ENV NAME_NODE_ADDRESS=172.17.0.1:9000

COPY --from=build /app/main /app/main

CMD ["./main"]