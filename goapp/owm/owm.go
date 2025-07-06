package owm

import (
	"encoding/json"
	"fmt"
	"github.com/briandowns/openweathermap"
	"os"
	"path/filepath"
	"time"
)

func tempDir() (dir string, err error) {
	if dir, err = os.Getwd(); err != nil {
		return "", err
	}
	tmpDir := filepath.Join(dir, "temp")

	if err = os.MkdirAll(tmpDir, 0755); err != nil {
		return "", err
	}
	return tmpDir, nil
}

func cachedFile(name string, data any) (info os.FileInfo, err error) {
	var tmpDir string
	if tmpDir, err = tempDir(); err != nil {
		return nil, err
	}
	path := filepath.Join(tmpDir, name)

	var stat os.FileInfo
	if stat, err = os.Stat(path); err != nil {
		return nil, err
	}

	var open *os.File
	if open, err = os.Open(path); err != nil {
		return nil, err
	}

	defer open.Close()
	return stat, json.NewDecoder(open).Decode(data)
}

func writeCache(name string, data any) (err error) {
	var tmpDir string
	if tmpDir, err = tempDir(); err != nil {
		return err
	}
	path := filepath.Join(tmpDir, name)

	var indent []byte
	if indent, err = json.MarshalIndent(data, "", "  "); err != nil {
		return err
	}
	return os.WriteFile(path, indent, 0644)
}

func Current() (data *openweathermap.CurrentWeatherData, err error) {
	var key = os.Getenv("OWM_KEY")
	var zip = os.Getenv("OWM_ZIP")

	var info os.FileInfo
	data = &openweathermap.CurrentWeatherData{}
	info, err = cachedFile("current.json", data)

	// return cache data
	if err == nil && (info != nil && time.Since(info.ModTime()) < time.Minute*30) {
		return data, nil
	}

	if key == "" || zip == "" {
		return data, nil
	}

	if data, err = openweathermap.NewCurrent("F", "en", key); err != nil {
		return nil, err
	}
	if err = data.CurrentByZipcode(zip, "US"); err != nil {
		return nil, err
	}

	if err = writeCache("current.json", data); err != nil {
		return nil, err
	}
	return data, nil
}

func Forecast() (data *openweathermap.Forecast5WeatherData, err error) {
	var key = os.Getenv("OWM_KEY")
	var zip = os.Getenv("OWM_ZIP")

	var info os.FileInfo
	data = &openweathermap.Forecast5WeatherData{}
	info, err = cachedFile("forecast.json", data)

	// return cache data
	if err == nil && (info != nil && time.Since(info.ModTime()) < time.Hour) {
		return data, nil
	}

	if key == "" || zip == "" {
		fmt.Printf("key=%q zip=%q\n", key, zip)
		return data, nil
	}

	// make the api call
	var fc *openweathermap.ForecastWeatherData
	if fc, err = openweathermap.NewForecast("5", "f", "en", key); err != nil {
		return nil, err
	}
	if err = fc.DailyByZipcode(zip, "US", 99); err != nil {
		return nil, err
	}

	data = fc.ForecastWeatherJson.(*openweathermap.Forecast5WeatherData)
	if err = writeCache("forecast.json", data); err != nil {
		return nil, err
	}
	return data, nil
}
