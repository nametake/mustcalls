package main

import (
	"github.com/nametake/mustcalls"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(mustcalls.Analyzer) }
