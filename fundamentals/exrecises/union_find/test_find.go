package unionfind

import "log"

func testFind(i implementation) {
	qf := i.create(2)

	if qf.Find(0) != 0 || qf.Find(1) != 1 {
		log.Fatalf("%v Before union all sites belongs to component with same number", i.name)
	}
}
