package config

import (
	"fmt"
)

type Env string

const (
	Dev  Env = "DEV"
	Prod Env = "PROD"
	Test Env = "TEST"
)

func (env Env) IsValid() error {
	switch env {
	case Dev, Prod, Test:
		return nil
	}
	return fmt.Errorf("invalid GO_MODE env variable, please specify either:\n- GO_MODE=%s\n- GO_MODE=%s\n- GO_MODE=%s",
		Dev,
		Prod,
		Test)
}
