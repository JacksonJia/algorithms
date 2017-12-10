package quickfind

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/JacksonJia/algorithms/errors"
	"github.com/JacksonJia/algorithms/utils"
)

type QuickFind struct {
	Items []int
}

func New(size int) *QuickFind {
	items := make([]int, size)
	for i := range items {
		items[i] = i
	}

	return &QuickFind{Items: items}
}

func (qf *QuickFind) Union(p, q int) {
	if qf.Items[q] == qf.Items[p] {
		return
	}

	pPreValue := qf.Items[p]
	for index, value := range qf.Items {
		if value == pPreValue {
			qf.Items[index] = qf.Items[q]
		}
	}
}

func (qf *QuickFind) Connected(p, q int) bool {
	return qf.Items[p] == qf.Items[q]
}

func InitFromFile(path string) (*QuickFind, error) {
	if !utils.IsFileExists(path) {
		return &QuickFind{}, errors.ErrFileNotExists
	}

	file, err := os.Open(path)
	if err != nil {
		return &QuickFind{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return &QuickFind{}, errors.ErrFileNotExists
	}

	size, _ := strconv.Atoi(scanner.Text())
	qf := New(size)

	for scanner.Scan() {
		text := scanner.Text()
		pq := strings.Split(text, " ")
		p, _ := strconv.Atoi(pq[0])
		q, _ := strconv.Atoi(pq[1])
		qf.Union(p, q)
	}

	if err := scanner.Err(); err != nil {
		return &QuickFind{}, err
	}

	return qf, err
}
