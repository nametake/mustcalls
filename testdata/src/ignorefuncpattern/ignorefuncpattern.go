package ignorefuncpattern

func mustCall() {}

func IgnoreFunc1(num int) {
}

func Func1(num int) { // want "Func1 is not calling mustCall"
}
