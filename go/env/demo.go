package main

import (
	"fmt"

	"github.com/caarlos0/env"
)

type config struct {
	Env    string   `env:"ENV" envDefault:"development"`
	Hosts  []string `env:"HOSTS" envSeparator:":"`
	Port   int      `env:"PORT" envDefault:"2181"`
	CHRoot string   `env:"CHROOT" envDefault:"binlog"`
}

func main() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", cfg)
}
