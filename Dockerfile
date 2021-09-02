FROM golang:1.16-alpine as BUILD

WORKDIR /app

COPY . /app

RUN go build -o build-dir/http_monitor.bin

FROM golang:1.16-alpine as RUN

WORKDIR /app

COPY --from=BUILD /app/build-dir/http_monitor.bin ./http_monitor.bin
COPY --from=BUILD /app/configs/config.json ./configs/config.json

CMD [ "./http_monitor.bin" ]
