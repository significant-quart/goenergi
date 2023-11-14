package config

import "github.com/leapkit/core/envor"

var (
	SERIAL_NO     = envor.Get("SERIAL_NO", "")
	API_KEY       = envor.Get("API_KEY", "")
	DIRECTOR_URL  = envor.Get("DIRECTOR_URL", "")
	API_FREQUENCY = envor.Get("API_FREQUENCY", "")
	PORT          = envor.Get("PORT", "3000")
)
