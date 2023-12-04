// Package run sets the flags and runs the application.
package run

import (
	"go_day_00/internal/app/calculations"
)

func Run(f calculations.Flags) {
	if !f.Mean && !f.Median && !f.Sd && !f.Mode {
		f = SetFlags()
	}
	f.Input()
}

func SetFlags() calculations.Flags {
	return calculations.Flags{
		Mean:   true,
		Median: true,
		Mode:   true,
		Sd:     true,
	}
}
