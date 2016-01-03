package tmpl_util

import (
	"fmt"
	"html/template"
)

func divf64(a, b int) float64 {
	return float64(a) / float64(b)
}

func divint(a, b int) int {
	return a / b
}

func max(slice []int) int {
	m := 0
	for _, e := range slice {
		if e > m {
			m = e
		}
	}
	return m
}

func mod(a, b int) int {
	return a % b
}

func sum(slice []int) int {
	sum := 0
	for _, e := range slice {
		sum += e
	}
	return sum
}

func toPercent(format string, a float64) string {
	return fmt.Sprintf(format, a*100)
}

var funcMap = template.FuncMap{
	"divf64":    divf64,
	"divint":    divint,
	"max":       max,
	"mod":       mod,
	"sum":       sum,
	"toPercent": toPercent,
}

func GetFuncMap() template.FuncMap {
	return funcMap
}
