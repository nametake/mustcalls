package primitive

func mustCall() {}

func f1() { // want "f1 is not calling mustCall."
	mustCall()
}
