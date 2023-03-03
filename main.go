package main

import (
	"fmt"
	"time"

	"github.com/IsabelF98/CMSC701_PA1/pkgsuffarr"
)

func main() {
	refrence := "ecoli.fa"
	outputsa := "ecoli.txt"
	outputpreftab := "ecoli_preftab.bin"
	queries := "queries.fa"
	queriesoutput := "ecoli_queries_hit.naive.txt"

	start := time.Now()
	pkgsuffarr.BuildSa(refrence, outputsa, outputpreftab, false, 2)
	duration := time.Since(start)
	fmt.Println("CPU time:", duration.Seconds(), "sec")

	pkgsuffarr.QuarySa(outputsa, queries, "naive", queriesoutput)

}
