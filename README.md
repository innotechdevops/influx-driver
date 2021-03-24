# influx-driver

## Install

```
$ go get github.com/innotechdevops/influx-driver
```

## How to use

- With env

```golang
driver := influxdriver.New(influxdriver.ConfigEnv())
```

- With config

```golang
driver := influxdriver.New(influxdriver.Config{
    URL:   os.Getenv("INFLUX_URL"),
    Token: os.Getenv("INFLUX_TOKEN"),
})
```