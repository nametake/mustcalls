package ignorefilepattern

func mustCall() {}

func f1(num int) { // want "f1 is not calling mustCall"
}
