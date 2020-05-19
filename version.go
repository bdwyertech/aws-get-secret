package main

import (
	"fmt"
	"runtime"
)

var GitCommit, ReleaseVer, ReleaseDate string

func showVersion() {
	if GitCommit == "" {
		GitCommit = "DEVELOPMENT"
	}
	if ReleaseVer == "" {
		ReleaseVer = "DEVELOPMENT"
	}
	fmt.Println("version:", ReleaseVer)
	fmt.Println("commit:", GitCommit)
	fmt.Println("date:", ReleaseDate)
	fmt.Println("runtime:", runtime.Version())
}
