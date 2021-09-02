## HTTP Monitor

Micro golang app for monitoring http endpoints

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
./http_monitor.bin
```

#### Docker Setup:

2 - Build the image:

```bash
docker build . --tag http-monitor
```

3 - Run it!:

```bash
docker run --rm http-monitor
```

You should see some logs matching: `[URL]: [Status Code], in [Request Time]`. Example:

```bash
[INFO] http://www.google.com: 200 OK, in 203.559672ms
[INFO] http://www.yahoo.com: 200 OK, in 1.142937649s
[INFO] http://www.yahoo.com: 200 OK, in 792.987539ms
[INFO] http://www.google.com: 200 OK, in 185.066743ms
[INFO] http://www.yahoo.com: 200 OK, in 816.863777ms
[INFO] http://www.yahoo.com: 200 OK, in 852.272979ms
```
