package DivideAndConquer

import "math/rand"

/**
 * 快速排序（随机化）
 * input: a 待排序的序列 p 排序的左边界 r 排序的右边界
 */
func RandomizedQuickSort(a []int, p int, r int) {
	if p < r {
		// 进行随机划分
		q := RandomizedPartition(a, p, r)
		// 对左半段排序
		RandomizedQuickSort(a, p, q-1)
		// 对右半段排序
		RandomizedQuickSort(a, q+1, r)
	}
}

/**
 * 随机选择划分标准进行划分数字序列
 * input: a 待划分的序列 p 划分的基准下标 r 划分的右边界
 * return: 划分完成位置的下标
 */
func RandomizedPartition(a []int, p int, r int) int {
	// 随机选择一个范围内的数据进行划分
	i := rand.Intn(r-p+1) + p
	a[i], a[p] = a[p], a[i]

	//划分操作
	return Partition(a, p, r)
}

/**
 * 划分数字序列
 * input: a 待划分的序列 p 划分的基准下标 r 划分的右边界
 * return: 划分完成位置的下标
 */
func Partition(a []int, p int, r int) int {
	i, j := p, r+1
	x := a[p]
	//将小于x的元素交换到左边区域，将大于x的元素交换到右边区域
	for {
		i++
		//如果左边的元素小于x，跳过
		for a[i] < x && i < r {
			i++
		}
		j--
		// 如果右边的元素大于x，跳过
		for a[j] > x {
			j--
		}

		// 循环停止条件
		if i >= j {
			break
		}

		// 交换
		a[i], a[j] = a[j], a[i]
	}

	a[p] = a[j]
	a[j] = x

	// j 是对数据随机划分的位置
	return j
}
