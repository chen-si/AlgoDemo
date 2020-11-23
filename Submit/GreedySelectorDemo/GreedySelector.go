package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	//GeneratedRandomActivity("Activity_1000000.dat", 1000000)
	// 从文件中读取数据
	s, f := ReadActivityFromFile("Activity_1000000.dat")

	A := make([]bool, len(s))
	// 记录时间
	start := time.Now()
	GreedySelector(len(s), s, f, A)
	elapsed := time.Since(start)

	fmt.Println("算法耗时：", elapsed)
	// 输出结果到文件
	WriteSelectorToFile("Activity_1000000_Selector.dat", A)
}
/**
 * @Description: 活动安排问题的贪心算法
 * @param n 活动数量
 * @param s 活动的开始时间
 * @param f 活动的结束时间， 默认升序排列
 * @param A 是否安排此活动
 */
func GreedySelector(n int, s []int, f []int, A []bool){
	A[0] = true
	j := 0
	// 贪心策略：优先满足结束最早的
	for i := 1; i < n; i++{
		if s[i] >= f[j]{
			A[i] = true
			j ++
		}else{
			A[i] = false
		}
	}
}

/**
 * @Description: 随机生成活动安排问题的数据
 * @param fileName 文件名
 * @param n 数据量
 */
func GeneratedRandomActivity(fileName string, n int){
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

	f := make([]int,n)
	for i := 0; i < n; i++{
		// 结束时间的最大值设置为3*n
		f[i] = rand.Intn(3 * n) + 1
	}
	// 非降序排列
	sort.Ints(f)

	s := make([]int, n)
	for i := 0; i < n; i++{
		// 开始时间小于结束时间
		s[i] = rand.Intn(f[i])
	}

	// 写入文件
	for i := 0; i < n; i++{
		_, err = w.WriteString(strconv.Itoa(s[i]) + " ")
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

	for i := 0; i < n; i++{
		_, err = w.WriteString(strconv.Itoa(f[i]) + " ")
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
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
		return
	}
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
 * @Description: 从文件中读取数据
 * @param fileName 文件名
 * @return s 活动开始时间
 * @return f 活动结束时间
 */
func ReadActivityFromFile(fileName string) (s []int, f []int) {
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil,nil
	}
	r := bufio.NewReader(file)

	// 读取第一行 活动开始时间
	line1 , err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
		return nil,nil
	}

	// 读取第二行 活动结束时间
	line2 , err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
		return nil,nil
	}
	sString := strings.Split(line1, " ")
	fString := strings.Split(line2, " ")

	for i := 0; i < len(sString) - 1; i++{
		si, _ := strconv.Atoi(sString[i])
		s = append(s, si)
	}

	for i := 0; i < len(fString) - 1; i++{
		fi, _ := strconv.Atoi(fString[i])
		f = append(f, fi)
	}

	return
}

/**
 * @Description: 将结果写入文件
 * @param fileName 文件名
 * @param A 能否安排次活动的序列
 */
func WriteSelectorToFile(fileName string, A []bool){
	// 打开文件
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 获取Writer
	w := bufio.NewWriter(file)

	// 如果A[i] 为真，则输出1表示该活动可以安排，否则输出0
	for i := 0; i < len(A); i++{
		if A[i]{
			_, err = w.WriteString("1 ")
		}else{
			_, err = w.WriteString("0 ")
		}
	}

	// 清空缓存
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
