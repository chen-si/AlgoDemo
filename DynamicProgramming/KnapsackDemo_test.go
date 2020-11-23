package DynamicProgramming

import (
	"fmt"
	"testing"
	"time"
)

func TestKnapsack(t *testing.T) {
	GeneratedKnapsackData("knapsack_50.dat",50, 250)
	c, w, v := ReadKnapsackDataFromFile("knapsack_50.dat", 50)

	x := make([]int, len(w))
	var p [][]int
	p = make([][]int, 50000)
	for i := 0; i < 50000 ; i++{
		p[i] = make([]int, 2)
	}
	start := time.Now()
	value := Knapsack(len(w), c, v, w, p, x)
	elapsed := time.Since(start)

	fmt.Println("算法耗时：", elapsed)

	WriteKnapsackResultToFile("Knapsack_50_Result.dat",x, value)
}
