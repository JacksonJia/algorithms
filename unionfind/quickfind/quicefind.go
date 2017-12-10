package quickfind

type QuickFind struct {
	Items []int
}

func New(size int) *QuickFind {
	return &QuickFind{Items: make([]int, size)}
}

func (qf *QuickFind) Union(p, q int) {

}

func (qf *QuickFind) Connected(p, q int) bool {
	return false
}

func (qf *QuickFind) InitFromFile(path string) {

}
