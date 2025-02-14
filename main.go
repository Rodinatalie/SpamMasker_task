package main

import (
	"bufio"
	"fmt"

	"os"
	"strings"
)
//коммит тест
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Show your spam!")
	scanner.Scan()
	str := scanner.Text()
	res := SpamMasker(str)
	fmt.Println(res)

}

func SpamMasker(str string) string {
	hg := "http://"
	bt := []byte(str)

	var bidx int
	strnewb := make([]byte, len(bt), cap(bt))
	copy(strnewb, bt)

	aidx := strings.Index(str, hg)
	var result string
	if aidx >= 0 {
		aidx += len(hg)

		for i := aidx; str[i] != ' '; i++ {
			bidx = i
			if i+1 == len(str) {
				break
			}

		} //находим индекс последнего символа ссылки

		for i := aidx; i != bidx+1; i++ {
			strnewb[i] = '*'
			if i+1 == len(str) {
				break
			}
		} //заменяем ссылку звездочками

		strnew := string(strnewb)
		result = strnew
	} else {
		result = str
	}
	return result
}
