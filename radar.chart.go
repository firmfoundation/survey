package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vicanso/go-charts/v2"
)

func writeFile(buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, "radar-chart.png")
	err = ioutil.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}

func GenRadarChart(survey_title string, indicators []string, indicators_weight []float64, indicator_values [][]float64) {
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
	err = writeFile(buf)
	if err != nil {
		panic(err)
	}
}
