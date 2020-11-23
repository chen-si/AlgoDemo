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
	// 读取文件中的数据
	res := ReadDataFromFile("100w_data.dat")

	var totalTime time.Duration
	for i := 0; i < 1000; i++{
		temp := make([]int, 10000000)
		copy(temp, res)
		// 快速排序
		start := time.Now()
		RandomizedQuickSort(temp, 0, 999999)
		elapsed := time.Since(start)
		totalTime += elapsed
	}
	RandomizedQuickSort(res, 0, 999999)

	averageTime := totalTime / 1000

	fmt.Println("数据规模：100w")
	fmt.Println("运行1000次快速排序的平均时间为：",averageTime)
	// 输出到文件
	WriteDataToFile("ordered_100w_data.dat",res)
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
 * 随机生成数据
 * input: fileName 生成的数据文件的名字 dataCap 生成的数据的数量
 */
func GeneratedRandomData(fileName string, dataCap int) {
	// 判断文件是否已经存在
	exist, err := exists(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 如果文件存在，则删除这个文件
	if exist {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// 创建一个新文件
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 得到一个Writer 采用buffer 提高I/O效率
	w := bufio.NewWriter(file)

	// 循环向文件中写入随机数据
	for i := 0; i < dataCap; i++ {
		// 随机数据
		r := rand.Int()
		// 写入文件
		_, err = w.WriteString(strconv.Itoa(r) + "\n")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	// 清空缓冲区
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
		return
	}
}

/**
 * 判断文件是否存在
 * input: path 文件的路径
 * return: true表示文件存在，false表示文件不存在
 */
func exists(path string) (bool, error) {
	// 得到文件转态
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	// 判断文件是否存在
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
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
func WriteDataToFile(fileName string, data []int) {
	// 打开文件
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 创建Writer
	w := bufio.NewWriter(file)
	// 遍历数据
	for _, v := range data {
		// 将数据写入文件
		_, err = w.WriteString(strconv.Itoa(v) + "\n")
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
