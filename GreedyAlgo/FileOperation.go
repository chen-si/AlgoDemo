package GreedyAlgo

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

// GreedySelector
func GeneratedRandomActivity(fileName string, n int){
	exist, err := exists(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	if exist {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	file, err := os.Create(fileName)
	defer file.Close()

	f := make([]int,n)
	for i := 0; i < n; i++{
		f[i] = rand.Intn(3 * n) + 1
	}
	sort.Ints(f)

	s := make([]int, n)
	for i := 0; i < n; i++{
		s[i] = rand.Intn(f[i])
	}

	w := bufio.NewWriter(file)

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

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

func ReadActivityFromFile(fileName string) (s []int, f []int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil,nil
	}
	r := bufio.NewReader(file)


	line1 , err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
		return nil,nil
	}

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

func WriteSelectorToFile(fileName string, A []bool){
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	w := bufio.NewWriter(file)

	for i := 0; i < len(A); i++{
		if A[i]{
			_, err = w.WriteString("1 ")
		}else{
			_, err = w.WriteString("0 ")
		}

	}


	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

// Huffman
func GeneratedRandomString(fileName string, n int){
	exist, err := exists(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	if exist {
		err = os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	file, err := os.Create(fileName)
	defer file.Close()

	str := GetRandomString(n)

	w := bufio.NewWriter(file)

	_, err = w.WriteString(str + "\n")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetRandomString(length int) string {
	str := "AAAAAAAABBBBBBBBBBBBBBBBBBBCDEFHHHHHHHHHHHHHHHGIIIIIIIIIJKLMMMMMMMMMMNOOOOOOOOOOOOOOOPQRRRRRSTUUUUUUUUUUUUUUUVWXYYYYYYYYYZ0000000000000111122222223444444445666666665789aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcddeeeeffggggggggggghijjkklklmnnopqrssssssssssssssssssssssssttuvwxxxxxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func ReadRuneAndFreqFromFile(fileName string) map[rune]int{
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	r := bufio.NewReader(file)

	str, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return nil
	}

	strings.TrimSuffix(str, "\n")

	symFreq := make(map[rune]int)
	// read each symbol and record the frequencies
	for _, c := range str {
		if c == '\n'{
			break
		}
		symFreq[c]++
	}
	return symFreq
}

func WriteHuffmanCodeToFile(fileName string,codes []HuffmanCode){
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	w := bufio.NewWriter(file)

	_, err = w.WriteString("字符\t\t频率\t\t编码\t\t\n")

	fmt.Println(codes)

	for _, c := range codes{
		_, err = w.WriteString(fmt.Sprintf("%c\t\t%d\t\t%s\t\t\n", c.value, c.freq, c.code))
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
		return
	}
}