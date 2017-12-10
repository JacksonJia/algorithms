package main

import (
	"fmt"

	"github.com/JacksonJia/algorithms/unionfind/quickfind"
)

func main() {
	qf, _ := quickfind.InitFromFile("data/unionfind/tinyUF.txt")

	fmt.Println(qf.Connected(1, 3))
}
