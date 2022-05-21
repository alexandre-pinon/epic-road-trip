package config

import "log"

type Env int

const (
	Dev Env = iota
	Prod
	Test
)

func (env *Env) GetFileName() string {
	var envFile string

	switch *env {
	case Dev:
		envFile = ".env"
	case Prod:
		envFile = ".env.prod"
	case Test:
		envFile = ".env.test"
	default:
		log.Fatal("Invalid env")
	}

	return envFile
}
