package unionfind

import "log"

func testUnion(i implementation) {
	qf := i.create(4)

	qf.Union(0, 1)

	if !qf.Connected(0, 1) {
		log.Fatalf("%v: sites should be connected after union", i.name)
	}

	if qf.Find(0) != qf.Find(1) {
		log.Fatalf("%v after union sites should be in the same component", i.name)
	}

	qf.Union(2, 3)

	if !qf.Connected(2, 3) {
		log.Fatalf("%v: sites should be connected after union", i.name)
	}

	if qf.Find(2) != qf.Find(3) {
		log.Fatalf("%v after union sites should be in the same component", i.name)
	}

	qf.Union(1, 2)

	if qf.Count() != 1 {
		log.Fatalf("%v after union count should be decreased", i.name)
	}
}
