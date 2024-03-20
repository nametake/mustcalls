package funcpattern

func TargetFunc1(num int) { // want "TargetFunc1 is not calling mustCall"
}

func Func1(num int) {
}

func NoTargetFunc1(num int) {
}
