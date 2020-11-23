package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 读取文件中的数据 这里使用了快速排序的生成数据
	data := ReadDataFromFile("100w_data.dat")

	var totalTime time.Duration
	var target []int
	var result []int
	for i := 0; i < 1000; i++{
		// 线性时间选择
		k := rand.Intn(1000000)
		// 计算线性时间选择的执行时间
		start := time.Now()
		x :=  Select(data, 0, 999999, k)
		elapsed := time.Since(start)

		// 记录每一次线性时间选择的结果
		target = append(target, k)
		result = append(result, x)
		totalTime += elapsed
	}
	// 计算平均执行时间
	averageTime := totalTime / 1000

	fmt.Println("数据规模：100w")
	fmt.Println("运行1000次线性时间选择的平均时间为：",averageTime)

	// 结果写入文件
	WriteSelectDataToFile("100w_select_res.dat", target, result)
}

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

/**
 * 从文件中读取数据
 * input: fileName 待读取数据的文件名
 * return: 文件中数据的切片
 */
func ReadDataFromFile(fileName string) []int {
	// 返回的切片
	var res []int

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// 创造一个Reader
	r := bufio.NewReader(file)
	// 循环读取数据
	for {
		// 按行读取数据
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
			return nil
		}
		if err == io.EOF {
			return res
		}
		//string to int
		num, err := strconv.Atoi(strings.TrimSuffix(line, "\n"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		// 添加到数据切片中
		res = append(res, num)
	}
	return res
}

/**
 * 将数据写入文件
 * input: fileName 存储数据的文件 target 查找目标切片 res 查找结果切片
 */
func WriteSelectDataToFile(fileName string, target []int, res []int ) {
	// 打开文件
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 创建Writer
	w := bufio.NewWriter(file)

	_, err = w.WriteString(fmt.Sprintf("target\t result\n"))
	// 遍历切片
	for i := 0; i < len(res); i++{
		// 写入数据
		_, err = w.WriteString(fmt.Sprintf("%d\t %d\n", target[i], res [i]))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	// 清空缓存
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}