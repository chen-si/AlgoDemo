package GreedyAlgo

import (
	"fmt"
	"testing"
	"time"
)

func TestGreedySelector(t *testing.T) {
	GeneratedRandomActivity("Activity_10000000.dat", 10000000)
	s, f := ReadActivityFromFile("Activity_10000000.dat")

	A := make([]bool, len(s))

	start := time.Now()
	GreedySelector(len(s), s, f, A)
	elapsed := time.Since(start)

	fmt.Println("计算时间:", elapsed)
	WriteSelectorToFile("Activity_10000000_Selector.dat", A)
}
