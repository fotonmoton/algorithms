package unionfind

import "log"

func testCount(i implementation) {
	qf := i.create(10)

	if qf.Count() != 10 {
		log.Fatalf("%v: Before any union number of components should be equal to number of sites", i.name)
	}
}
