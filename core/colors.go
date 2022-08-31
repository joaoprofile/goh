package core

import (
	"github.com/labstack/gommon/color"
)

func Green(msg string) string {
	return color.Green(msg)
}

func Blue(msg string) string {
	return color.Blue(msg)
}

func Red(msg string) string {
	return color.Red(msg)
}

func Yellow(msg string) string {
	return color.Yellow(msg)
}
