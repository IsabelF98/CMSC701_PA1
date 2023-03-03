// This file contains the functions for crating the suffix array for a gene
package pkgsuffarr

import (
	"bytes"
	"fmt"
	"index/suffixarray"
	"os"
)

func MakeSuffArr(gene string, output string) {
	/*
		This function computes the suffic array of a gene string and writes it to a file.
		Thisfunction utilizes the index/suffixarray package in GO.

		Inputs:
			gene(string): string of refrence gene
			output(string): location to suffix array output file as .txt
	*/

	gene_byte := []byte(gene)        // turn gene string into bytes
	sa := suffixarray.New(gene_byte) // create suffix array

	var bbuffer bytes.Buffer // create bytebuffer variable

	file, err := os.Create(output) // create file for suffix array
	if err != nil {
		fmt.Println("Couldn't create file")
	}
	defer file.Close()

	if err := sa.Write(&bbuffer); err != nil { // write suffix array to byte buffer
		panic(err)
	}

	_, err = bbuffer.WriteTo(file) // write byte buffer to output file
	if err != nil {
		fmt.Println("Couldn't write to file")
	}

}
