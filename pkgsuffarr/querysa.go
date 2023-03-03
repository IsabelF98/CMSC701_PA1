// This file containsfunctions for perfoming a query search

package pkgsuffarr

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func QuarySa(index string, queries string, querymode string, output string) {
	/*
		This function perfoms a query search based on a suffix array computed prior.
		The two possilbe search modes are neive or simple accelerant algorithum.

		Inputs:
			index(string): location to suffix array text file
			queries(string): location to queries FASTA file
			querymode(string): either "naive" or "simpaccel"
			output(string): output location as tab seperated list
		Outputs:
			None
	*/

	// open suffix array file
	readSaFile, err := os.Open(index)
	if err != nil {
		fmt.Println("Couldn't open file", err)
	}
	defer readSaFile.Close()

	var suffixarray MyIndex
	suffixarray.Read(readSaFile)

	text := string(suffixarray.Data) // original text as the string
	sa := suffixarray.Sa.Int32       // suffix array as an array

	suffixes := []string{} // empty array of suffixes

	// appends each suffix into the array of suffixes in order
	for i := 0; i < suffixarray.Sa.len(); i++ {
		idx := sa[i]
		suff := text[idx:]
		suffixes = append(suffixes, suff)
	}

	// open querries file
	readQueriesFile, err := os.Open(queries)
	if err != nil {
		fmt.Println("Could not open file", err)
	}

	// line by line scanner of queries
	fileScanner := bufio.NewScanner(readQueriesFile)
	fileScanner.Split(bufio.ScanLines)

	// create output file
	writeQueriesFile, err := os.Create(output)
	if err != nil {
		fmt.Println("Could not create file", err)
	}
	defer writeQueriesFile.Close()

	for fileScanner.Scan() {
		name := fileScanner.Text() // query name
		writeQueriesFile.WriteString(name[1:] + "\t")
		fileScanner.Scan()
		queryaux := fileScanner.Text() // query
		if querymode == "naive" {      // naive search
			query1 := string(queryaux[1 : len(queryaux)-2]) // query

			start := time.Now()
			indx1 := NaiveSearch(suffixes, query1)
			duration := time.Since(start)
			fmt.Println(query1+" CPU time:", duration.Seconds(), "sec") //print time

			query2 := NextOrderQuery(query1) //next order query
			indx2 := NaiveSearch(suffixes, query2)
			writeQueriesFile.WriteString(fmt.Sprint(indx2-indx1) + "\t")
			for i := indx1; i < indx2; i++ {
				writeQueriesFile.WriteString(fmt.Sprint(i) + "\t")
			}
			writeQueriesFile.WriteString("\n")
		} else if querymode == "simpaccel" { // simple accel search
			query1 := string(queryaux[1 : len(queryaux)-2]) // query

			start := time.Now() //start time
			indx1 := SimpAccel(suffixes, query1)
			duration := time.Since(start)
			fmt.Println(query1+" CPU time:", duration.Seconds(), "sec") //print time

			query2 := NextOrderQuery(query1) // next order query
			indx2 := SimpAccel(suffixes, query2)
			writeQueriesFile.WriteString(fmt.Sprint(indx2-indx1) + "\t")
			for i := indx1; i < indx2; i++ {
				writeQueriesFile.WriteString(fmt.Sprint(i) + "\t")
			}
			writeQueriesFile.WriteString("\n")
		}

	}
	readQueriesFile.Close()

}
