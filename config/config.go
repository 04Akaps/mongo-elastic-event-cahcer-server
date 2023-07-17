package config

import (
	"context"
	"mongo-event-cacher/types"
	"os"

	"github.com/naoina/toml"
)

type Collection struct {
	Name   string
	Struct struct{}
}

type Config struct {
	MongoDB struct {
		DataSource string
		DB         string
	}

	Elastic struct {
		Uri      string
		User     string
		Password string
	}

	Collections map[string]Collection

	CancelContext     context.Context
	CancelContextFunc context.CancelFunc
}

func NewConfig(file string) *Config {
	c := new(Config)

	if f, err := os.Open(file); err != nil {
		panic(err)
	} else {
		if err = toml.NewDecoder(f).Decode(c); err != nil {
			panic(err)
		} else {
			for key := range c.Collections {
				if err = types.VerifyStructMap(key); err != nil {
					panic(err)
				}
			}

			ctx := context.Background()

			c.CancelContext, c.CancelContextFunc = context.WithCancel(ctx)
			return c
		}
	}
}
