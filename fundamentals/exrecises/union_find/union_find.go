package unionfind

type UnionFind interface {
	Find(site int) int               // returns "component" to which "site" belongs
	Union(aSite, bSite int)          // links two sites. After union a and b belongs to same component
	Connected(aSite, bSite int) bool // checks if two sites belongs to same component
	Count() int                      // returns number of
}
