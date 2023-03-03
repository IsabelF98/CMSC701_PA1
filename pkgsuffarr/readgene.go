// This file contains the functions for reading fasta files

package pkgsuffarr

import (
	"bufio"
	"log"
	"os"
)

func ReadGeneFASTA(refrence string) string {
	/*
		Function reads FASTA file and outputs file contents as a string

		Inputs:
			refrence(string): location to refrence FASTA file

		Outputs:
			gene(strin): the gene characters as a string with $ at the end
	*/
	file, err := os.Open(refrence) // open fasta file
	if err != nil {
		log.Fatal("Open file error", err)
	}
	defer file.Close() //close fasta file

	var gene string = "" // initialize gene string

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip first linee
	for scanner.Scan() {
		gene += string(scanner.Text()) // append line to gene string
	}
	gene += "$" // add $ at the end
	return gene

}
