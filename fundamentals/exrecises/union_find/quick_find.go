package unionfind

type quickFind struct {
	id    []int
	count int
}

func NewQuickFind(count int) UnionFind {
	id := make([]int, count)

	for i := range id {
		id[i] = i
	}

	return &quickFind{id, count}
}

func (qf *quickFind) Find(site int) int {
	return qf.id[site]
}

func (qf *quickFind) Union(aSite, bSite int) {
	aComponent := qf.Find(aSite)
	bComponent := qf.Find(bSite)

	if aComponent == bComponent {
		return
	}

	for site, component := range qf.id {
		if component == aComponent {
			qf.id[site] = bComponent
		}
	}
	qf.count--
}

func (qf *quickFind) Connected(aSite, bSite int) bool {
	return qf.Find(aSite) == qf.Find(bSite)
}

func (qf *quickFind) Count() int {
	return qf.count
}
