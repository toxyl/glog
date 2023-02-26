package glog

import "bytes"

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

	progressBar.WriteString(Percentage(percent, LoggerConfig.AutoFloatPrecision))

	return progressBar.String()
}
