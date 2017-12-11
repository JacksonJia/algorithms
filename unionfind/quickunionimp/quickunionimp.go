package quickunionimp

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/JacksonJia/algorithms/errors"
	"github.com/JacksonJia/algorithms/utils"
)

type QuickUnionImp struct {
	Items []int
	Size  []int
}

func New(size int) *QuickUnionImp {
	items := make([]int, size)
	for i := range items {
		items[i] = i
	}

	s := make([]int, size)
	for i := range s {
		s[i] = 1
	}

	return &QuickUnionImp{Items: items, Size: s}
}

func (qui *QuickUnionImp) Union(p, q int) {
	pRoot := qui.getRoot(p)
	qRoot := qui.getRoot(q)

	pSize := qui.getSize(p)
	qSize := qui.getSize(q)
	if pSize <= qSize {
		qui.Items[pRoot] = qRoot
		return
	}
	qui.Items[qRoot] = pRoot
}

func (qui *QuickUnionImp) Connected(p, q int) bool {
	return qui.getRoot(p) == qui.getRoot(q)
}

func (qui *QuickUnionImp) getSize(p int) int {
	return qui.Size[p]
}

func (qui *QuickUnionImp) getRoot(p int) int {
	var parent int
	for {
		parent = qui.Items[p]
		if qui.Items[parent] == parent {
			return parent
		}
		p = parent
	}
}

func InitFromFile(path string) (*QuickUnionImp, error) {
	if !utils.IsFileExists(path) {
		return &QuickUnionImp{}, errors.ErrFileNotExists
	}

	file, err := os.Open(path)
	if err != nil {
		return &QuickUnionImp{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return &QuickUnionImp{}, errors.ErrFileNotExists
	}

	size, _ := strconv.Atoi(scanner.Text())
	qui := New(size)

	for scanner.Scan() {
		text := scanner.Text()
		pq := strings.Split(text, " ")
		p, _ := strconv.Atoi(pq[0])
		q, _ := strconv.Atoi(pq[1])
		qui.Union(p, q)
	}

	if err := scanner.Err(); err != nil {
		return &QuickUnionImp{}, err
	}

	return qui, err
}
