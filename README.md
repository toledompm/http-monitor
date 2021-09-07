## HTTP Monitor

Micro golang app for monitoring http endpoints.

http_monitor takes a path to a config file as the first argument, it proceeds to fetch the provided URLs, logging the response status code, as well as a breakdown of the request time.

### Async processing

Go routines allows for the monitor keep track of every provided url in the provided intervals, without being influenced by bad endpoints taking too long to respond.

- `time.AfterFunc`: AfterFunc creates it's own go-routine after the provided time has passed, without holding up any processing, or beeing held up itself.

### Setup:

1 - Create a config file:
`configs/config.json`

```json
[
  {
    "url": "http://www.google.com",
    "interval": 20
  },
  {
    "url": "http://www.yahoo.com",
    "interval": 10
  }
]
```

_"interval" is measured in seconds, the minimum allowed value is 1_

#### Local Setup

2 - Build the binary:

```bash
go build -o http_monitor.bin
```

3 - Run it!:

```bash
./http_monitor.bin ./configs/config.json
```

#### Docker Setup:

2 - Build the image:

```bash
docker build . --tag http-monitor
```

3 - Run it!:

```bash
docker run --rm -v /absolute/path/to/config.json:/app/config.json http-monitor
```

_you need to provide a configFile to the container at:/app/config.json_

#### Output

You should see some logs matching: `[URL]: [Status Code] - [Request Time]`. Followed by a break down of the total request time. Example:

```bash
[INFO] https://www.yahoo.com: 200 OK - 637.766229ms
DNS Lookup: 37.76337ms
Connection: 165.026824ms
TLS Handshake: 168.809246ms
Time to first byte: 637.608928ms
```
