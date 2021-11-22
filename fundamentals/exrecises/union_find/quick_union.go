package unionfind

type quickUnion struct {
	id    []int
	count int
}

func NewQuickUnion(count int) UnionFind {
	id := make([]int, count)

	for i := range id {
		id[i] = i
	}

	return &quickUnion{id, count}
}

func (qf *quickUnion) Find(site int) int {
	for site != qf.id[site] {
		site = qf.id[site]
	}

	return site
}

func (qf *quickUnion) Union(aSite, bSite int) {
	aComponent := qf.Find(aSite)
	bComponent := qf.Find(bSite)

	if aComponent == bComponent {
		return
	}

	qf.id[aComponent] = bComponent
	qf.count--
}

func (qf *quickUnion) Connected(aSite, bSite int) bool {
	return qf.Find(aSite) == qf.Find(bSite)
}

func (qf *quickUnion) Count() int {
	return qf.count
}
