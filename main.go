package main

import (
	"fmt"

	"github.com/JacksonJia/algorithms/unionfind/quickfind"
	"github.com/JacksonJia/algorithms/unionfind/quickunion"
)

func main() {
	qf, _ := quickfind.InitFromFile("data/unionfind/tinyUF.txt")

	fmt.Println(qf.Connected(1, 2))

	qu, _ := quickunion.InitFromFile("data/unionfind/tinyUF.txt")

	fmt.Println(qu.Connected(1, 2))
}
