package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vicanso/go-charts/v2"
)

func writeFile(buf []byte, survey_id string, user_id string) error {
	tmpPath := "./radarchart/img"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, survey_id+"_"+user_id+".png")
	err = ioutil.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func GenRadarChart(survey_title string, indicators []string, indicators_weight []float64, indicator_values [][]float64,
	survey_id string, user_id string) {
	p, err := charts.RadarRender(
		indicator_values,
		charts.TitleTextOptionFunc(survey_title),
		charts.LegendLabelsOptionFunc(nil),
		charts.RadarIndicatorOptionFunc(indicators, indicators_weight),
	)
	if err != nil {
		panic(err)
	}

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf, survey_id, user_id)
	if err != nil {
		panic(err)
	}
}

func GenGetRadarChart(survey_title string, indicators []string, indicators_weight []float64, indicator_values [][]float64,
	survey_id string, user_id string) ([]byte, error) {
	p, err := charts.RadarRender(
		indicator_values,
		charts.TitleTextOptionFunc(survey_title),
		charts.LegendLabelsOptionFunc(nil),
		charts.RadarIndicatorOptionFunc(indicators, indicators_weight),
	)
	if err != nil {
		panic(err)
	}

	buf, err := p.Bytes()
	if err != nil {
		panic(err)
	}
	err = writeFile(buf, survey_id, user_id)
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile("./radarchart/img/" + survey_id + "_" + user_id + ".png")
	if err != nil {
		return nil, err
	}
	return file, nil
}
