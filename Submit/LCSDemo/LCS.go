package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 读取字符序列
	str := ReadStringFromFile("string2000_2500.dat")

	x := []byte(str[0])
	y := []byte(str[1])

	// 初始化操作
	m := len(x)
	n := len(y)

	var c [][]int
	var b [][]int

	c = make([][]int, m + 1)
	for i := 0; i <= m ; i ++{
		c[i] = make([]int, n + 1)
	}

	b = make([][]int, m + 1)
	for i := 0; i <= m ; i ++{
		b[i] = make([]int, n + 1)
	}

	start := time.Now()
	// 计算最优值
	LCSLength(m, n, x, y, c, b)
	period1 := time.Since(start)

	file, err := os.OpenFile("out.dat", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	w := bufio.NewWriter(file)

	start = time.Now()
	// 得到最长公共子序列，结果输出到w中
	LCS(m, n, x, b, w)
	period2 := time.Since(start)

	// 打印耗时
	fmt.Println(period1 + period2)

	// 清空缓存
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

/**
 * @Description: 计算最优值
 * @param m 字符序列x长度
 * @param n 字符序列y长度
 * @param x 字符序列x
 * @param y 字符序列y
 * @param c 最优值矩阵
 * @param b 最优解矩阵
 */
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

/**
 * @Description: 得到最优解
 * @param i 字符序列x当前扫描位置（在b中）
 * @param j 字符序列y当前扫描位置（在b中）
 * @param x 字符序列x
 * @param b 最优解矩阵
 * @param w 最长公共子序列输出位置
 */
func LCS(i int, j int, x []byte,b [][]int, w *bufio.Writer){
	if i == 0 || j == 0{
		return
	}
	if b[i][j] == 1{
		LCS(i - 1, j - 1, x, b, w)
		// 找到相同的字符，输入进w
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

/**
 * @Description: 自动生成随机字符序列
 * @param fileName 文件名
 * @param length1 序列1长度
 * @param length2 序列2长度
 */
func GeneratedRandomString(fileName string, length1 int, length2 int){
	// 判断文件是否存在
	exist, err := exists(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 如果存在 先删除原有文件
	if exist {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	// 创建新文件
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	w := bufio.NewWriter(file)

	// 得到两个随机字串
	str1 := GetRandomString(length1)
	str2 := GetRandomString(length2)

	// 写入文件
	_, err = w.WriteString(str1 + "\n")
	if err != nil{
		log.Fatal(err)
		return
	}

	_, err = w.WriteString(str2 + "\n")
	if err != nil{
		log.Fatal(err)
		return
	}

	// 清空缓存
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
		return
	}
}

/**
 * @Description: 得到指定长度的随机字符串
 * @param length 长度
 * @return string 随机字符串
 */
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}


/**
 * @Description: 判断文件是否存在
 * @param path 文件路径
 * @return bool 文件是否存在
 * @return error 错误
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
 * @Description: 从文件中读取字符串
 * @param fileName 文件名
 * @return []string 字符串切片
 */
func ReadStringFromFile(fileName string) []string {
	var res []string

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	r := bufio.NewReader(file)

	// 循环读取数据
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
			return nil
		}
		// 如果到文件尾
		if err == io.EOF {
			return res
		}

		res = append(res, line)
	}

	return res
}