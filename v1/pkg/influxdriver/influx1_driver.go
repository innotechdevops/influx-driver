package influxdriver

import (
	"log"
	"os"

	influxdb1 "github.com/influxdata/influxdb1-client/v2"
)

// InfluxDB1Driver is the interface
type InfluxDB1Driver interface {
	Connect() influxdb1.Client
}

// Config is a model for connect InfluxDB
type Config struct {
	URL                string
	Username           string
	Password           string
	InsecureSkipVerify bool
}

type influxDB1 struct {
	Conf Config
}

func (db *influxDB1) Connect() influxdb1.Client {
	client, err := influxdb1.NewHTTPClient(influxdb1.HTTPConfig{
		Addr:               db.Conf.URL,
		Username:           db.Conf.Username,
		Password:           db.Conf.Password,
		InsecureSkipVerify: db.Conf.InsecureSkipVerify,
	})
	if err != nil {
		log.Fatalln("Cannot connect influxdb1", err)
	}
	return client
}

// New for create influxdb driver
func New(config Config) InfluxDB1Driver {
	return &influxDB1{
		Conf: config,
	}
}

// ConfigEnv for create influxdb driver
func ConfigEnv() Config {
	return Config{
		URL:                os.Getenv("INFLUX_URL"),
		Username:           os.Getenv("INFLUX_USER"),
		Password:           os.Getenv("INFLUX_PASS"),
		InsecureSkipVerify: false,
	}
}
