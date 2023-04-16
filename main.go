package main

import (
	"encoding/json"
	"get_API/api/climacell"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	os.Setenv("CLIMACELL_API_KEY", "GBznxyP89LiJRv4o39Hqa6M3962DKu9X")

	req, err := http.NewRequest(http.MethodGet, "https://api.tomorrow.io/v4/timelines?location=40.75872069597532,-73.98529171943665&fields=temperature&timesteps=1h&units=metric&apikey=GBznxyP89LiJRv4o39Hqa6M3962DKu9X", nil)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("apikey", os.Getenv("CLIMACELL_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	log.Println("We got the response:", string(responseBytes))

	var weatherSamples []climacell.Weather
	if err := json.Unmarshal(responseBytes, &weatherSamples); err != nil {
		log.Fatalf("Ошибка десериализации данных погоды")
	}

	for _, w := range weatherSamples {
		if w.Temp != nil && w.Temp.Value != nil {
			log.Printf("Температура при %s равна %f градусов %s\n",
				w.ObservationTime.Value, *w.Temp.Value, w.Temp.Units)
		} else {
			log.Printf("Нет доступных температурных данных %s\n", w.ObservationTime.Value)
		}
	}

}
