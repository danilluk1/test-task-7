package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ApiUrl = "https://api.openweathermap.org/data/2.5/weather"
)

type forecast struct {
	Name       string    `json:"name"`
	Dt         int       `json:"dt"`
	DtTxt      string    `json:"dt_txt"`
	Visibility int       `json:"visibility"`
	Main       main      `json:"main"`
	Weather    []weather `json:"weather"`
	Clouds     clouds    `json:"clouds"`
	Wind       wind      `json:"wind"`
}

type main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  float32 `json:"pressure"`
	SeaLevel  float32 `json:"sea_level"`
	Humidity  float32 `json:"humidity"`
	TempKf    float32 `json:"temp_kf"`
}

type weather struct {
	ID          int32  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type clouds struct {
	All int `json:"all"`
}

type wind struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
	Gust  float32 `json:"gust"`
}

func GetCurrentWeather(city string) (*forecast, error) {
	resp, err := http.Get(fmt.Sprintf("%s?q=%s&lang=ru&units=metric", ApiUrl, city))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var forecast forecast
	err = json.Unmarshal(body, &forecast)
	if err != nil {
		return nil, err
	}

	return &forecast, nil
}
