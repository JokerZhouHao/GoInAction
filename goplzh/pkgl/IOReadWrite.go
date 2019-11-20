package main

/* 有这么几种读取文件方式，写类似
1、ioutil.ReadFile(path)
2、os.Open(path)		fi.Read
3、os.Open(path)		bufio.NewReader(fi)
3、os.Open(path)		ioutil.ReadAll(fi)
*/

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func read0(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

func read1(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024)
	buf := make([]byte, 8096)

	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:]...)
	}
	return string(chunks)
}

func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024)
	buf := make([]byte, 8096)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func test(fc func(string) string, path string) {
	start := time.Now()
	fc(path)
	fmt.Printf("Cost time %v\n", time.Now().Sub(start))
}

func main() {
	fmt.Println(filepath.Abs(filepath.Dir(".")))

	path := "./src/GoLearn/data/read_test.txt"

	test(read0, path)
	test(read1, path)
	test(read2, path)
	test(read3, path)
}
