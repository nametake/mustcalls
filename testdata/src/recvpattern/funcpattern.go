package recvpattern

func mustCall() {}

type TargetStruct struct{}

func (s *TargetStruct) TargetMethod1(num int) { // want "TargetMethod1 is not calling mustCall"
}

func (TargetStruct) TargetMethod2(num int) { // want "TargetMethod2 is not calling mustCall"
}

type Struct struct{}

func (s *Struct) Method1(num int) {
}

type NoTargetStruct struct{}

func (s *NoTargetStruct) NoTargetMethod1(num int) {
}
