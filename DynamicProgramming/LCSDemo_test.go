package DynamicProgramming

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestLCS(t *testing.T) {
	//GeneratedRandomString("string5000_7500.dat", 5000, 7500)
	str := ReadStringFromFile("string50_75.dat")
	//fmt.Println(str)

	x := []byte(str[0])
	y := []byte(str[1])

	m := len(x)
	n := len(y)

	var c [][]int
	var b [][]int

	c = make([][]int, m + 1)
	for i := 0; i <= m ; i ++{
		c[i] = make([]int, n + 1)
	}

	b = make([][]int, m + 1)
	for i := 0; i <= m ; i ++{
		b[i] = make([]int, n + 1)
	}

	start := time.Now()
	LCSLength(m, n, x, y, c, b)
	period1 := time.Since(start)

	file, err := os.OpenFile("out.dat", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	w := bufio.NewWriter(file)

	start = time.Now()
	LCS(m, n, x, b, w)
	period2 := time.Since(start)

	fmt.Println(period1 + period2)

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}