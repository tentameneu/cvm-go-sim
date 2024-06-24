package simulator

import (
	"fmt"

	"github.com/tentameneu/cvm-go"
	"github.com/tentameneu/cvm-go-sim/internal/config"
	"github.com/tentameneu/cvm-go-sim/internal/logging"
)

var log = logging.Logger

func Run(stream []int) {
	intComparator := func(x, y int) int { return x - y }
	intCVM := cvm.NewCVM(config.BufferSize(), intComparator)

	length := len(stream)
	log().Info("Starting CVM Algorithm...")
	for i, element := range stream {
		logging.LogDeep(fmt.Sprintf("Processing element %d/%d...", i+1, length))
		n := intCVM.Process(element)
		logging.LogDeep("Current estimation after insert.", "N", n)
	}

	precision := float64(intCVM.N()) / float64(config.Distinct())
	if precision > 1.0 {
		precision = 1 / precision
	}

	log().Info("Done estimating number of distinct elements.", "N", intCVM.N(), "precision", fmt.Sprintf("%.2f%%", precision*100))
}
