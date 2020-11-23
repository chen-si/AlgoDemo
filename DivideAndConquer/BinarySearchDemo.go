package DivideAndConquer

/**
 * 二分搜索
 * input: a 有序数列 x 待搜索的数字 n 数列的长度
 * return: p x在a中的位置，未找到则返回-1
 */
func BinarySearch(a []int, x int, n int) int {
	left := 0
	right := n - 1
	// 左边界小于右边界的时候 继续循环查找
	for left <= right {
		middle := (left + right) / 2
		if x == a[middle] {
			return middle
		}
		if x > a[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
