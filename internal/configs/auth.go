package configs

import "time"

type AuthConfig struct {
	SignKey        []byte
	Alg            string
	ExpireDuration time.Duration
}
