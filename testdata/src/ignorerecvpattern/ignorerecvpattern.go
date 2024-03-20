package ignorefuncpattern

func mustCall() {}

type Struct struct{}

func (s *Struct) Method(num int) { // want "Method is not calling mustCall"
}

type IgnoreStruct struct{}

func (s *IgnoreStruct) Method(num int) {
}
