package DivideAndConquer

/**
 * 线性时间选择
 * input: a 待选择的数字序列 p 左边界 r 右边界 k 要寻找第k小的元素
 * return: a中第k小的元素的值
 */
func Select(a []int, p int, r int, k int) int {
	if r-p < 75 {
		RandomizedQuickSort(a, p, r)
		return a[p+k-1]
	}

	// 把每个组的中位数（第三个数）交换到区间[p,p+(r-p-4)/4]
	for i := 0; i <= (r-p-4)/5; i++ {
		s := p + 5*i
		t := s + 4
		// 冒泡排序，从后开始排，结果使得后三个数是排好顺序的（递增）
		for j := 0; j < 3; j++ {
			for n := s; n < t-j; n++ {
				if a[n] > a[n+1] {
					a[n], a[n+1] = a[n+1], a[n]
				}
			}
		}
		// 交换每组中的中位数到前面
		a[p+i], a[s+2] = a[s+2], a[p+i]
	}

	// 找出中位数的中位数
	x := Select(a, p, p+(r-p-4)/5, (r-p-4)/10)

	// 根据中位数的中位数进行划分
	i := PartitionForSelect(a, p, r, x)
	// 划分的位置
	j := i - p + 1

	if k <= j {
		return Select(a, p, i, k)
	} else {
		return Select(a, i+1, r, k-j)
	}
}

/**
 * 划分数字序列
 * input: a 待划分的序列 p 划分的左边界 r 划分的右边界 val 划分依据
 * return: 划分完成位置的下标
 */
func PartitionForSelect(a []int, p int, r int, val int) int {
	pos := 0
	for q := p; q <= r; q++ {
		if a[q] == val {
			pos = q
			break
		}
	}
	a[p], a[pos] = a[pos], a[p]

	return Partition(a, p, r)
}

