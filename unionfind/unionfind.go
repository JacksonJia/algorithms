package unionfind

type UF interface {
	Union(p, q int)
	Connected(p, q int) bool
}
