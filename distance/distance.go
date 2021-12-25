package distance

import (
	"bufio"
	"fmt"
	"os"
)

func ReadArr(address string) []string {
	r, _ := os.Open(address)
	defer r.Close()
	s := bufio.NewScanner(r)
	arr := []string{}
	for s.Scan() { // 循环直到文件结束
		line := s.Text() // 这个 line 就是每一行的文本了，string 类型
		arr = append(arr, line)
		fmt.Println(line)
	}
	return arr
}

func ToIntArr(str []string) [][]int {

}
