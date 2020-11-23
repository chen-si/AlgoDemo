package DynamicProgramming

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GeneratedRandomString(fileName string, length1 int, length2 int){
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
	if err != nil {
		log.Fatal(err)
		return
	}
	w := bufio.NewWriter(file)

	str1 := GetRandomString(length1)
	str2 := GetRandomString(length2)

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

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
		return
	}
}

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

func ReadStringFromFile(fileName string) []string {
	var res []string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	r := bufio.NewReader(file)

	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
			return nil
		}
		if err == io.EOF {
			return res
		}

		res = append(res, line)
	}

	return res
}

func GeneratedKnapsackData(fileName string, n int, c int){
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
	if err != nil {
		log.Fatal(err)
		return
	}
	w := bufio.NewWriter(file)

	_, err = w.WriteString(strconv.Itoa(c) + "\n")
	if err != nil {
		log.Fatal(err)
		return
	}

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

	_ = w.Flush()
}

func ReadKnapsackDataFromFile(fileName string, n int)(c int, w []int, v []int){
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	r := bufio.NewReader(file)

	line1, _ := r.ReadString('\n')
	//fmt.Println(line1)
	line1 = strings.TrimSuffix(line1, "\n")
	c, _ = strconv.Atoi(line1)

	line2, _ := r.ReadString('\n')
	wStr := strings.Split(line2, " ")
	for i := 0; i < len(wStr) - 1; i++{
		wi, _ := strconv.Atoi(wStr[i])
		w = append(w, wi)
	}

	line3, _ := r.ReadString('\n')
	vStr := strings.Split(line3, " ")
	for i := 0; i < len(vStr) - 1; i++{
		vi, _ := strconv.Atoi(vStr[i])
		v = append(v, vi)
	}

	return
}

func WriteKnapsackResultToFile(fileName string, x []int, value int){
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
	if err != nil {
		log.Fatal(err)
		return
	}
	w := bufio.NewWriter(file)

	_, _ = w.WriteString("MaxValue:" + strconv.Itoa(value) + "\n")
	_, _ = w.WriteString("X:" + "\n")

	for i := 0; i < len(x); i++{
		_, err = w.WriteString(strconv.Itoa(x[i]) + " ")
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	_, _ = w.WriteString("\n")

	_ = w.Flush()
}