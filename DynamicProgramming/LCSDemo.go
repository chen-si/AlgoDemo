package DynamicProgramming

import (
	"bufio"
	"log"
)

// 计算最优值
func LCSLength(m int, n int, x []byte, y []byte, c [][]int, b [][]int){
	// if i or j = 0, c[i][j] = 0
	for i := 1; i <= m; i++{
		c[i][0] = 0
	}
	for i := 1; i <= n; i++{
		c[0][i] = 0
	}

	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			if x[i - 1] == y[j - 1]{
				// 如果这一位置的字符相等 则长度加一
				c[i][j] = c[i - 1][j - 1] + 1
				b[i][j] = 1
			}else if c[i - 1][j] >= c[i][j - 1] {
				// 判断应该往哪个方向延伸
				c[i][j] = c[i - 1][j]
				b[i][j] = 2
			}else{
				c[i][j] = c[i][j - 1]
				b[i][j] = 3
			}
		}
	}
}

func LCS(i int, j int, x []byte,b [][]int, w *bufio.Writer){
	if i == 0 || j == 0{
		return
	}
	if b[i][j] == 1{
		LCS(i - 1, j - 1, x, b, w)
		_, err := w.WriteString(string(x[i - 1]))
		if err != nil{
			log.Fatal(err)
			return
		}
	}else if b[i][j] == 2{
		LCS(i - 1, j, x, b, w)
	}else{
		LCS(i, j - 1, x, b, w)
	}
}
