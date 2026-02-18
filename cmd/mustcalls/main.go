package main

import (
	"github.com/nametake/mustcalls"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(mustcalls.Analyzer) }
