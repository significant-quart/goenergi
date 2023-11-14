package main

import (
	"fmt"
	"goenergi/cmd/config"
	"goenergi/cmd/fetcher"
	"strconv"
	"time"
)

func main() {
	if len(config.SERIAL_NO) == 0 {
		panic("SERIAL_NO not set in .env")
	}

	if len(config.API_KEY) == 0 {
		panic("API_KEY not set in .env")
	}

	if len(config.DIRECTOR_URL) == 0 {
		panic("DIRECTOR_URL not set in .env")
	}

	if len(config.API_FREQUENCY) == 0 {
		config.API_FREQUENCY = "5"

		// TODO: Handle Error
	}

	i, err := strconv.Atoi(config.API_FREQUENCY)
	if err != nil {
		panic("API_FREQUENCY")
	}

	deviceData := fetcher.FetchInit(i)

	for {
		time.Sleep(1 * time.Second)

		fmt.Print("\033[u\033[K")
		fmt.Printf("Power from grid (Eddie) %d W\n", (*deviceData)[0].Eddi[0].Grd)
		fmt.Printf("CT-1 (Eddi) from grid %d W\n", (*deviceData)[0].Eddi[0].Ectp1)
		fmt.Printf("CT-2 (Eddi) from grid %d W\n", (*deviceData)[0].Eddi[0].Ectp2)
		fmt.Printf("CT-3 (Eddi) from grid %d W\n", (*deviceData)[0].Eddi[0].Ectp3)
		fmt.Printf("CT-1 (Harvi) from solar %d W\n", (*deviceData)[2].Harvi[0].Ectp1)
		fmt.Printf("CT-2 (Harvi) from solar %d W\n", (*deviceData)[2].Harvi[0].Ectp2)
		fmt.Printf("CT-3 (Harvi) from solar %d W\n", (*deviceData)[2].Harvi[0].Ectp3)
		fmt.Printf("CT-1 (Harvi) from solar %d W\n", (*deviceData)[2].Harvi[0].Ect1P)
		fmt.Printf("CT-2 (Harvi) from solar %d W\n", (*deviceData)[2].Harvi[0].Ect2P)
		fmt.Printf("CT-3 (Harvi) from solar %d W\n", (*deviceData)[2].Harvi[0].Ect3P)
		fmt.Print("\033[A")
	}
}
