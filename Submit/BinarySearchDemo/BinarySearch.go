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
	data := ReadDataFromFile("ordered_100w_data.dat")

	var totalTime time.Duration
	var target []int
	var result []int
	for i := 0; i < 1000; i++{
		// 二分搜索
		x := rand.Intn(1000000)
		// 计算二分搜索的执行时间
		start := time.Now()
		k :=  BinarySearch(data, data[x], 1000000)
		elapsed := time.Since(start)

		// 记录每一次二分搜索的结果
		target = append(target, data[x])
		result = append(result, k)
		totalTime += elapsed
	}
	// 计算平均执行时间
	averageTime := totalTime / 1000

	fmt.Println("数据规模：100w")
	fmt.Println("运行1000次二分搜索的平均时间为：",averageTime)

	// 结果写入文件
	WriteSelectDataToFile("100w_binary_search_res.dat", target, result)
}

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

