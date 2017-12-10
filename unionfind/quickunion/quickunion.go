package quickunion

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/JacksonJia/algorithms/errors"
	"github.com/JacksonJia/algorithms/utils"
)

type QuickUnion struct {
	Items []int
}

func New(size int) *QuickUnion {
	items := make([]int, size)
	for i := range items {
		items[i] = i
	}

	return &QuickUnion{Items: items}
}

func (qu *QuickUnion) Union(p, q int) {
	//get p, q's root
	pRoot := qu.getRoot(p)
	qRoot := qu.getRoot(q)
	if pRoot == qRoot {
		return
	}

	//set pRoot value qRoot also qRoot's value
	qu.Items[pRoot] = qRoot
}

func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.getRoot(p) == qu.getRoot(q)
}

func (qu *QuickUnion) getRoot(p int) int {
	var parent int

	for {
		parent = qu.Items[p]
		if qu.Items[parent] == parent {
			return parent
		}
		p = parent
	}
}

func InitFromFile(path string) (*QuickUnion, error) {
	if !utils.IsFileExists(path) {
		return &QuickUnion{}, errors.ErrFileNotExists
	}

	file, err := os.Open(path)
	if err != nil {
		return &QuickUnion{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return &QuickUnion{}, errors.ErrFileNotExists
	}

	size, _ := strconv.Atoi(scanner.Text())
	qu := New(size)

	for scanner.Scan() {
		text := scanner.Text()
		pq := strings.Split(text, " ")
		p, _ := strconv.Atoi(pq[0])
		q, _ := strconv.Atoi(pq[1])
		qu.Union(p, q)
	}

	if err := scanner.Err(); err != nil {
		return &QuickUnion{}, err
	}

	return qu, err
}
