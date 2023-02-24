package pkgsuffarr

import (
	"bufio"
	//"fmt"
	"log"
	"os"
)

// Function reads FASTA file and outputs file contents as a string
func readFASTA(refrence string) string {
	file, err := os.Open(refrence) // open fasta file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //close fasta file

	var gene string = ""

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip first linee
	for scanner.Scan() {
		gene += string(scanner.Text())
		// fmt.Println(string(scanner.Text()))
	}
	gene += "$"
	return gene

}
