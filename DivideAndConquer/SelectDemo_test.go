package DivideAndConquer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	data1w := ReadDataFromFile("1w_data.dat")
	data10w := ReadDataFromFile("10w_data.dat")
	data100w := ReadDataFromFile("100w_data.dat")
	data1000w := ReadDataFromFile("1000w_data.dat")

	type args struct {
		a []int
		p int
		r int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1w_data",
			args: args{
				a: data1w,
				p: 0,
				r: 9999,
			},
		},
		{
			name: "10w_data",
			args: args{
				a: data10w,
				p: 0,
				r: 99999,
			},
		},
		{
			name: "100w_data",
			args: args{
				a: data100w,
				p: 0,
				r: 999999,
			},
		},
		{
			name: "1000w_data",
			args: args{
				a: data1000w,
				p: 0,
				r: 9999999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var totalTime time.Duration
			for i := 0; i < 1000; i++{
				// 线性时间选择
				x := rand.Intn(tt.args.r-tt.args.p+1) + tt.args.p
				start := time.Now()
				Select(tt.args.a, tt.args.p, tt.args.r, x)
				elapsed := time.Since(start)
				totalTime += elapsed
			}

			averageTime := totalTime / 1000

			fmt.Println("数据规模：",tt.name)
			fmt.Println("运行1000次线性时间选择的平均时间为：",averageTime)
		})
	}
}
