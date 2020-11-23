package DivideAndConquer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	data1w := ReadDataFromFile("ordered_1w_data.dat")
	data10w := ReadDataFromFile("ordered_10w_data.dat")
	data100w := ReadDataFromFile("ordered_100w_data.dat")
	data1000w := ReadDataFromFile("ordered_1000w_data.dat")

	type args struct {
		a []int
		x int
		n int
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
				n: 10000,
			},
		},
		{
			name: "10w_data",
			args: args{
				a: data10w,
				n: 100000,
			},
		},
		{
			name: "100w_data",
			args: args{
				a: data100w,
				n: 1000000,
			},
		},
		{
			name: "1000w_data",
			args: args{
				a: data1000w,
				n: 10000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var totalTime time.Duration
			for i := 0; i < 1000; i++{
				// 二分搜索
				x := rand.Intn(tt.args.n)
				start := time.Now()
				BinarySearch(tt.args.a, tt.args.a[x], tt.args.n)
				elapsed := time.Since(start)
				totalTime += elapsed
			}

			averageTime := totalTime / 1000

			fmt.Println("数据规模：",tt.name)
			fmt.Println("运行1000次二分搜索的平均时间为：",averageTime)
		})
	}
}
