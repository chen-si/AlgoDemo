package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//GeneratedKnapsackData("knapsack_1000.dat", 1000, 10000)
	// 从文件中读取数据
	c, w, v := ReadKnapsackDataFromFile("knapsack_1000.dat")

	// 初始化
	x := make([]int, 1000)
	var p [][]int
	p = make([][]int, 50000000)
	for i := 0; i < 50000000 ; i++{
		p[i] = make([]int, 2)
	}

	// 记录时间
	start := time.Now()
	value := Knapsack(1000, c, v, w, p, x)
	elapsed := time.Since(start)

	fmt.Println("算法耗时：", elapsed)

	// 将结果写入文件
	WriteKnapsackResultToFile("Knapsack_1000_Result.dat",x, value)
}

/**
 * 动态规划算法(跳跃点)求解0-1背包问题
 * @param n 物品数量
 * @param c 背包容量
 * @param v 物品价值
 * @param w 物品重量
 * @param p 跳跃点集 其中p[][0]代表重量 p[][1]代表价值
 * @param x 输出的0-1序列
 * @return 最大总价值
 */
func Knapsack(n int, c int, v []int, w []int, p [][]int, x []int) int{
	// 初始化一些变量
	head := make([]int, n + 2)
	head[n + 1] = 0
	p[0][0] = 0
	p[0][1] = 0
	left, right, next := 0, 0, 1
	head[n] = 1

	// 从最后一个物品往前追溯
	for i := n - 1; i >= 0; i --{
		k := left
		//fmt.Println(left, right)
		for j := left; j <= right ;j++{
			//fmt.Println(i, k, next)
			// 如果重量0超过背包容量
			if p[j][0] + w[i] > c{
				break
			}

			// y,m 用于指代跳跃点集 q
			y, m := p[j][0] + w[i], p[j][1] + v[i]
			// fmt.Println(i, "+", y, m)

			// 重量更小的跳跃点集-保留
			for k <= right && p[k][0] < y{
				p[next][0] = p[k][0]
				p[next][1] = p[k][1]
				next ++
				k ++
			}
			// 如果总重量相同但是价值更少，覆盖
			if k <= right && p[k][0] == y {
				if m < p[k][1]{
					m = p[k][1]
				}
				k ++
			}
			// 如果价值比之前的最大的还大 加到最后面
			if m >= p[next - 1][1]{
				p[next][0] = y
				p[next][1] = m
				next ++
			}
			// 清除上一个跳跃点集的无效部分（重量更大但是价值更少）
			for k <= right && p[k][1] <= p[next - 1][1]{
				k++
			}
		}
		// 填写相同的部分
		for k <= right {
			p[next][0] = p[k][0]
			p[next][1] = p[k][1]
			next ++
			k ++
		}

		left = right + 1
		right = next - 1
		// 记录起始点
		head[i] = next
	}
	Traceback(n, w, v, p, head, x)
	return p[next -1][1]
}

/**
 * 动态规划算法(跳跃点)求解0-1背包问题
 * @param n 物品数量
 * @param w 物品重量
 * @param v 物品价值
 * @param p 跳跃点集 其中p[][0]代表重量 p[][1]代表价值
 * @param head 记录跳跃点开始地点
 * @param x 输出的0-1序列
 */
func Traceback(n int, w []int, v []int, p [][]int, head []int, x []int){
	j, m := p[head[0] - 1][0], p[head[0] - 1][1]
	for i := 1; i <= n; i++{
		x[i - 1] = 0
		for k := head[i + 1]; k <= head[i] - 1;k ++{
			if p[k][0] + w[i - 1] == j && p[k][1] + v[i - 1] == m{
				// 将第i个物品放入背包中
				x[i - 1] = 1
				j = p[k][0]
				m = p[k][1]
				break
			}
		}
	}
}

/**
 * @Description: 随机生成0-1背包问题的数据
 * @param fileName 文件名
 * @param n 数据量
 * @param c 0-1 背包问题中背包容量
 */
func GeneratedKnapsackData(fileName string, n int, c int){
	// 判断文件是否已经存在
	exist, err := exists(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 如果存在 则删除
	if exist {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	// 新建文件
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 获取Writer
	w := bufio.NewWriter(file)

	// 写入背包容量
	_, err = w.WriteString(strconv.Itoa(c) + "\n")
	if err != nil {
		log.Fatal(err)
		return
	}

	// 写入物品重量
	for i := 0; i < n; i++{
		weight := rand.Intn(100) + 1
		_, err = w.WriteString(strconv.Itoa(weight) + " ")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	_, err = w.WriteString("\n")
	if err != nil {
		log.Fatal(err)
		return
	}

	// 写入物品价值
	for i := 0; i < n; i++{
		value := rand.Intn(100)
		_, err = w.WriteString(strconv.Itoa(value) + " ")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	_, err = w.WriteString("\n")
	if err != nil {
		log.Fatal(err)
		return
	}

	// 清空缓存
	_ = w.Flush()
}

/**
 * @Description: 从文件中读取数据
 * @param fileName 文件名
 * @return c 背包容量
 * @return w 物品重量切片
 * @return v 物品价值切片
 */
func ReadKnapsackDataFromFile(fileName string)(c int, w []int, v []int){
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 得到Reader
	r := bufio.NewReader(file)

	// 读取第一行（背包容量）
	line1, _ := r.ReadString('\n')
	//fmt.Println(line1)
	line1 = strings.TrimSuffix(line1, "\n")
	c, _ = strconv.Atoi(line1)

	// 读取第二行（物品重量）
	line2, _ := r.ReadString('\n')
	wStr := strings.Split(line2, " ")
	for i := 0; i < len(wStr) - 1; i++{
		wi, _ := strconv.Atoi(wStr[i])
		w = append(w, wi)
	}

	// 读取第三行（物品价值）
	line3, _ := r.ReadString('\n')
	vStr := strings.Split(line3, " ")
	for i := 0; i < len(vStr) - 1; i++{
		vi, _ := strconv.Atoi(vStr[i])
		v = append(v, vi)
	}

	return
}

/**
 * @Description: 将问题结果写入文件中
 * @param fileName 文件名
 * @param x 判断是否装入
 * @param value 最大价值
 */
func WriteKnapsackResultToFile(fileName string, x []int, value int){
	// 判断文件是否存在
	exist, err := exists(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 如果存在 则删除
	if exist {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	// 新建文件
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 得到Writer
	w := bufio.NewWriter(file)

	// 写入第一行： 最大总价值
	_, _ = w.WriteString("MaxValue:" + strconv.Itoa(value) + "\n")
	_, _ = w.WriteString("X:" + "\n")

	// 写入问题的解
	for i := 0; i < len(x); i++{
		_, err = w.WriteString(strconv.Itoa(x[i]) + " ")
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	_, _ = w.WriteString("\n")

	// 清空缓存
	_ = w.Flush()
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