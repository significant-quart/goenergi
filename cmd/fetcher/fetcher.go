package fetcher

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"goenergi/cmd/config"
	"goenergi/cmd/schemas"

	"github.com/icholy/digest"

	"github.com/go-co-op/gocron"
)

func get(client *http.Client, endpoint string) (res *http.Response, err error) {
	client = &http.Client{
		Transport: &digest.Transport{
			Username: config.SERIAL_NO,
			Password: config.API_KEY,
		},
	}

	res, err = client.Get(endpoint)

	return res, err
}

func FetchInit(frequency int) *[]schemas.Devices {
	var client http.Client

	res, err := get(&client, config.DIRECTOR_URL)
	if err != nil {
		// TODO: Handle Error
	}

	var asn string
	for k, v := range res.Header {
		if k == "X_myenergi-Asn" {
			asn = "https://" + v[0]

			break
		}
	}
	res.Body.Close()
	if len(asn) == 0 {
		// TODO: Handle Error
	}

	var body []byte
	var deviceData []schemas.Devices

	s := gocron.NewScheduler(time.UTC)

	s.Every(frequency).Seconds().Do(func() {
		res, err = get(&client, asn+"/cgi-jstatus-*")
		if err != nil {
			// TODO: Handle Error
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			// TODO: Handle Error
		}

		if err := json.Unmarshal(body, &deviceData); err != nil {
			// TODO: Handle Error
		}
	})

	defer s.StartAsync()

	return &deviceData
}
