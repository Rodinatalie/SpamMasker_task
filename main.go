package main

import (
    "bufio"
    "fmt"
    
    "os"
    "strings"
)

func main() {
    
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Show your spam!")
    scanner.Scan()
    str := scanner.Text()
    hg := "http://"
    bt := []byte(str)

    var bidx int
    strnewb := make([]byte, len(bt), cap(bt))
    copy(strnewb, bt)
    //idx := strings.Index(g, hg)

    aidx := strings.LastIndex(str, hg) //индекс первого символа после http://
    if aidx > 0{
        aidx += len(hg)
    
        for i := aidx; str[i] != ' '; i++{
            bidx = i
        } //находим индекс последнего символа ссылки
   
        for i := aidx; i != bidx + 1; i++{
            strnewb[i] = '*'
        } //заменяем ссылку звездочками

        strnew := string(strnewb)
        fmt.Println(strnew)
    }else{
        return
    }
    
    
}