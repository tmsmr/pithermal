package pithermal

import (
	"os"
	"strconv"
	"strings"
)

const CpuTempFile = "/sys/class/thermal/thermal_zone0/temp"

// GetCpuTemp returns the RPi's CPU temperature in °C
func GetCpuTemp() (float64, error) {
	data, err := os.ReadFile(CpuTempFile)
	if err != nil {
		return -1, err
	}
	value, err := strconv.ParseFloat(strings.TrimSpace(string(data)), 64)
	if err != nil {
		return -1, err
	}
	return value / 1000, nil
}
