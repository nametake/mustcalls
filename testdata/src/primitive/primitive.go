package primitive

func mustCall() {}

func f1() {
	mustCall()
}

func f2() { // want "f2 is not calling mustCall"
}
