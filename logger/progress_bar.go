package logger

import (
	"bytes"

	"github.com/toxyl/glog/colorizers"
	"github.com/toxyl/glog/config"
)

func ProgressBar(percent float64, barLength int) string {
	numBars := int(percent * float64(barLength))

	var progressBar bytes.Buffer
	for i := 0; i < barLength; i++ {
		if i < numBars {
			progressBar.WriteString("■")
		} else {
			progressBar.WriteString("▫")
		}
	}
	progressBar.WriteString(" ")

	progressBar.WriteString(colorizers.Percentage(percent, config.LoggerConfig.AutoFloatPrecision))

	return progressBar.String()
}
