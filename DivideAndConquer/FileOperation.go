package DivideAndConquer

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func GeneratedRandomData(fileName string, dataCap int) {
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

	for i := 0; i < dataCap; i++ {
		r := rand.Int()
		_, err = w.WriteString(strconv.Itoa(r) + "\n")
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

func ReadDataFromFile(fileName string) []int {
	//fmt.Println(1)
	var res []int
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
		//fmt.Println("!!!!!!!         ", line)
		num, err := strconv.Atoi(strings.TrimSuffix(line, "\n"))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		res = append(res, num)
	}
	return res
}

func WriteDataToFile(fileName string, data []int) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	w := bufio.NewWriter(file)
	for _, v := range data {
		_, err = w.WriteString(strconv.Itoa(v) + "\n")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
