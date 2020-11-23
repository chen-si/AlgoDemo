package GreedyAlgo

import (
	"fmt"
	"testing"
	"time"
)

func Test_buildTree(t *testing.T) {
	GeneratedRandomString("string_10000.dat", 10000)
	symFreq := ReadRuneAndFreqFromFile("string_10000.dat")

	var codes []HuffmanCode

	start := time.Now()
	tree := buildTree(symFreq)
	printCodes(tree, []byte{}, &codes)
	elapsed := time.Since(start)

	fmt.Println(elapsed)

	WriteHuffmanCodeToFile("string_10000_code.dat", codes)
}
