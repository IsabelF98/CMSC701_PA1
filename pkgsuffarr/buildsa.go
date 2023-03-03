package pkgsuffarr

import (
	"encoding/binary"
	"fmt"
	"os"
)

func BuildSa(refrence string, outputsa string, outputpreftab string, makepreftab bool, k int) {
	/*
		This function builds the suffix array and prefix table for a given gene. Note that go does
		not have a way of defining default or optional inputs so the outputpreftab and k value is
		requiered. If makepreftab=false the k value can be anything you want it will not effect
		the output of the suffix array.

		Inputs:
			refrence(string): location to refrence FASTA file
			outputsa(string): location to suffix array output file as .txt
			outputpreftab(string): location to prefix output file as .bin
			makepreftab(bool): "true" if you want to creat a prefix table "false" otherwise
			k(int): size of prefixes for prefix table
		Outputs:
			None
	*/

	gene := ReadGeneFASTA(refrence) // read FASTA file
	MakeSuffArr(gene, outputsa)     // make suffix array

	// if user wants to creat a prefix table
	if makepreftab {
		file1, err := os.Open(outputsa) // open suffix array file
		if err != nil {
			fmt.Println("Couldn't open file", err)
		}
		defer file1.Close()

		var suffixaray MyIndex
		suffixaray.Read(file1) // read suffix array file

		pref_table := MakePrefixTable(k, &suffixaray) // make dense prefix table

		file2, err := os.Create(outputpreftab) // create file for prefix table
		if err != nil {
			fmt.Println("Couldn't open file", err)
		}

		err = binary.Write(file2, binary.LittleEndian, pref_table) // write prefix table into pinary file
		if err != nil {
			fmt.Println("Write failed", err)
		}
	}

}
