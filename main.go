package main

import (
	"fmt"

	"github.com/JacksonJia/algorithms/unionfind/quickfind"
	"github.com/JacksonJia/algorithms/unionfind/quickunion"
	"github.com/JacksonJia/algorithms/unionfind/quickunionimp"
)

func main() {
	qf, _ := quickfind.InitFromFile("data/unionfind/tinyUF.txt")
	fmt.Println(qf.Connected(1, 4))

	qu, _ := quickunion.InitFromFile("data/unionfind/tinyUF.txt")
	fmt.Println(qu.Connected(1, 4))

	qui, _ := quickunionimp.InitFromFile("data/unionfind/tinyUF.txt")
	fmt.Println(qui.Connected(1, 4))
}
