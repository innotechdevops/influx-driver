package influxdriver

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"os"
)

// InfluxDBDriver is the interface
type InfluxDBDriver interface {
	Connect() influxdb2.Client
}

// Config is a model for connect InfluxDB
type Config struct {
	URL   string
	Token string
}

type influxDB struct {
	Conf Config
}

func (db *influxDB) Connect() influxdb2.Client {
	return influxdb2.NewClient(db.Conf.URL, db.Conf.Token)
}

// New for create influxdb driver
func New(config Config) InfluxDBDriver {
	return &influxDB{
		Conf: config,
	}
}

// ConfigEnv for create influxdb driver
func ConfigEnv() Config {
	return Config{
		URL:   os.Getenv("INFLUX_URL"),
		Token: os.Getenv("INFLUX_TOKEN"),
	}
}
