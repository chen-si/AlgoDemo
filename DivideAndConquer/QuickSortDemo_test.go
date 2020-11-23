package DivideAndConquer

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomizedQuickSort(t *testing.T) {
	data1w := ReadDataFromFile("1w_data.dat")
	data10w := ReadDataFromFile("10w_data.dat")
	data100w := ReadDataFromFile("100w_data.dat")
	data1000w := ReadDataFromFile("1000w_data.dat")

	type args struct {
		a []int
		p int
		r int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1w_data",
			args:args{
				a: data1w,
				p: 0,
				r: 9999,
			},
		},
		{
			name: "10w_data",
			args:args{
				a: data10w,
				p: 0,
				r: 99999,
			},
		},
		{
			name: "100w_data",
			args:args{
				a: data100w,
				p: 0,
				r: 999999,
			},
		},
		{
			name: "1000w_data",
			args:args{
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
				temp := make([]int, 10000000)
				copy(temp, tt.args.a)
				// 快速排序
				start := time.Now()
				RandomizedQuickSort(temp, tt.args.p, tt.args.r)
				elapsed := time.Since(start)
				totalTime += elapsed
			}


			averageTime := totalTime / 1000

			fmt.Println("数据规模：",tt.name)
			fmt.Println("运行1000次快速排序的平均时间为：",averageTime)
		})
	}
}
