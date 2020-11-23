package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 节点
type HuffmanTree interface {
	Freq() int
}

// 叶子节点（字符）
type HuffmanLeaf struct {
	freq  int
	value rune
}

// 中间节点
type HuffmanNode struct {
	freq        int
	left, right HuffmanTree
}

// 存储Huffman编码及其对应的频率以及字符
type HuffmanCode struct{
	value rune
	freq int
	code string
}


func (hLeaf HuffmanLeaf) Freq() int {
	return hLeaf.freq
}

func (hNode HuffmanNode) Freq() int {
	return hNode.freq
}

type treeHeap []HuffmanTree

// 构建堆
func (th treeHeap) Len() int { return len(th) }
func (th treeHeap) Less(i, j int) bool {
	return th[i].Freq() < th[j].Freq()
}
func (th *treeHeap) Push(ele interface{}) {
	*th = append(*th, ele.(HuffmanTree))
}
func (th *treeHeap) Pop() (popped interface{}) {
	popped = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return
}
func (th treeHeap) Swap(i, j int) { th[i], th[j] = th[j], th[i] }

func main() {
	//GeneratedRandomString("string_1000000.dat", 1000000)
	// 从文件中读取数据
	symFreq := ReadRuneAndFreqFromFile("string_100000.dat")

	var codes []HuffmanCode

	start := time.Now()
	tree := buildTree(symFreq)
	printCodes(tree, []byte{}, &codes)
	elapsed := time.Since(start)

	fmt.Println("算法耗时：", elapsed)
	// 结果写入文件
	WriteHuffmanCodeToFile("string_100000_code.dat", codes)
}

/**
 * @Description: 从字符频率构建huffman树
 * @param symFreq 字符及其频率的map
 * @return HuffmanTree Huffman树
 */
func buildTree(symFreq map[rune]int) HuffmanTree {
	var trees treeHeap

	for c, f := range symFreq {
		// 树集合，开始每个叶子节点都为一棵单独的树
		trees = append(trees, HuffmanLeaf{f, c})
	}
	// 初始化堆
	heap.Init(&trees)
	// 只要集合大于一个元素
	for trees.Len() > 1 {
		// 取出频率最小的两个集合
		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)

		// 将新树存入集合
		heap.Push(&trees, HuffmanNode{a.Freq() + b.Freq(), a, b})
	}
	// 堆顶部即为最后的Huffman树
	return heap.Pop(&trees).(HuffmanTree)
}

/**
 * @Description: 为Huffman树进行编码
 * @param tree Huffman树
 * @param prefix 当前编码前缀
 * @param codes 返回的编码集合
 */
func printCodes(tree HuffmanTree, prefix []byte, codes *[]HuffmanCode){
	switch i := tree.(type) {
	case HuffmanLeaf:
		// 如果是叶子节点 直接得到编码prefix
		c := HuffmanCode{
			value: i.value,
			freq:  i.freq,
			code:  string(prefix),
		}
		*codes = append(*codes, c)
	case HuffmanNode:
		// 左子树赋值0 递归
		prefix = append(prefix, '0')
		printCodes(i.left, prefix, codes)
		prefix = prefix[:len(prefix)-1]

		// 右子树赋值1 递归
		prefix = append(prefix, '1')
		printCodes(i.right, prefix, codes)
		prefix = prefix[:len(prefix)-1]
	}
}

/**
 * @Description:  随机生成字符串
 * @param fileName 文件名
 * @param n 数据量
 */
func GeneratedRandomString(fileName string, n int){
	// 判断文件是否存在
	exist, err := exists(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 如果存在 删除
	if exist {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	file, err := os.Create(fileName)
	defer file.Close()
	// 得到Writer
	w := bufio.NewWriter(file)

	// 得到随机字符串
	str := GetRandomString(n)

	// 写入文件中
	_, err = w.WriteString(str + "\n")
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
 * @Description: 得到随机字符串
 * @param length 长度
 * @return string 字符串
 */
func GetRandomString(length int) string {
	// 每个字符的出现频率不同，防止字符集中化
	str := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcddeeeeffggggggggggghijjkklklmnnopqrssssssssssssssssssssssssttuvwxxxxxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

/**
 * @Description: 读取字符及其在字符串中出现频率
 * @param fileName 文件名
 * @return map[rune]int 字符及其频率的map
 */
func ReadRuneAndFreqFromFile(fileName string) map[rune]int{
	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 得到Reader
	r := bufio.NewReader(file)

	str, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// 清除末尾\n
	strings.TrimSuffix(str, "\n")


	symFreq := make(map[rune]int)
	// 读取每一个字符并记录其频率
	for _, c := range str {
		if c == '\n'{
			break
		}
		symFreq[c]++
	}
	return symFreq
}

/**
 * @Description: 将编码结果写入文件内
 * @param fileName 文件名
 * @param codes Huffman编码
 */
func WriteHuffmanCodeToFile(fileName string,codes []HuffmanCode){
	// 打开文件
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 得到Writer
	w := bufio.NewWriter(file)

	// 写入标题栏
	_, err = w.WriteString("字符\t\t频率\t\t编码\t\t\n")

	//fmt.Println(codes)

	// 写入每一个编码
	for _, c := range codes{
		_, err = w.WriteString(fmt.Sprintf("%c\t\t%d\t\t%s\t\t\n", c.value, c.freq, c.code))
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// 清空缓存
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
		return
	}
}
