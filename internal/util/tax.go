package util

import (
	"log/slog"
	"os"
	"strconv"
)

func getTaxThreshold() float32 {
	taxThresholdStr := os.Getenv("TAX_THRESHOLD")
	taxThreshold, err := strconv.ParseFloat(taxThresholdStr, 32)
	if err != nil {
		slog.Warn("Expected tax threshold to be number ", "threshold", taxThresholdStr)
		return 50
	}
	return float32(taxThreshold)

}
func getTaxRate() float32 {
	taxRateStr := os.Getenv("TAX_RATE")
	taxRate, err := strconv.ParseFloat(taxRateStr, 32)
	if err != nil {
		slog.Warn("Expected tax rate to be number", "rate", taxRateStr)
		return 0.05
	}
	return float32(taxRate)
}
