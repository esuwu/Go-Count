package main

import (
	"bufio"
	"fmt"
	"github.com/esuwu/Counting-Go/counter"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	const k = 5
	result, total := counter.CountGo(reader, k)
	fmt.Println(result.String())
	fmt.Println("Total: ", total)
}
