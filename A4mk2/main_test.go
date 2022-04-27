package main

import (
	"github.com/jacksongodsey/golearning/A4mk2/playgame"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < 4; i++ {
		playgame.Playgame()
	}
}
