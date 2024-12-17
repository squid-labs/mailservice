package main

import (
	"flag"

	"mailservice/pkg/consumer"
)

const (
	DefaultPath = "config/mailservice.yml"
	DefaultTag  = "mailservice"
)

var (
	config = flag.String("config", DefaultPath, "Path to mailservice config file")
	tag    = flag.String("tag", DefaultTag, "Tag for RabbitMQ consumer")
)

func main() {
	flag.Parse()

	consumer.Run(*config, *tag)
}
