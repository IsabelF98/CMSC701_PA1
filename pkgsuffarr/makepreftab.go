// This file conatins the functions to create a dense prefix table
// This function takes a really long time to run and I am not sure how to make it faster
// Any help would be much apreciated

package pkgsuffarr

// set prefix table structure
type Perfidx struct {
	start int32
	end   int32
}

func findMinElement(arr []int) int {
	/*
		This function finds the minimum element in the array and returns the minimum number
	*/
	min_num := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min_num {
			min_num = arr[i]
		}
	}
	return min_num
}

func findMaxElement(arr []int) int {
	/*
		This function finds the maximum element in the array and returns the minimum number
	*/
	max_num := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max_num {
			max_num = arr[i]
		}
	}
	return max_num
}

func computeAllPrefix(alph []rune, k int) []string {
	/*
		This function computes all the prefix of leght k for a given alphabet

		Inputs:
			alph(rune): alphabet in sorted order
			k(int): length of prefix

		Outputs:
			prefs(string array): an array of all prefixs in sorted order
	*/
	if k <= 0 { // stop if at last character
		return []string{""}
	}
	prefs := make([]string, 0) // empty string of prefixs
	for _, c := range alph {   // for all letters in alphabet
		subprefs := computeAllPrefix(alph, k-1) // subprefix for given character using recursion
		for _, subComb := range subprefs {      // for all subprefix
			prefs = append(prefs, string(c)+subComb) // add subprefic to char
		}
	}
	return prefs
}

func MakePrefixTable(k int, suffixarray *MyIndex) []Perfidx {
	alph := []rune{'A', 'C', 'G', 'T'} // gene alphabet in order
	prefs := computeAllPrefix(alph, k) // array of all possible prefixes

	pref_table := []Perfidx{}  // empty prefx table
	sa := suffixarray.Sa.Int32 // suffix array

	for i := 0; i < len(prefs); i++ {
		offsets := suffixarray.Lookup([]byte(prefs[i]), -1) // find all instances of prefix in orignial text
		if offsets == nil {                                 // if prefix does not occur
			pref_table = append(pref_table, Perfidx{start: -1, end: -1}) // start and end index is -1
			continue
		}
		index_arr := []int{}                // empty array of suffix array indecies
		for j := 0; j < len(offsets); j++ { // looping through offset index array
			for k := 0; k < len(sa); k++ { // looping though suffix array
				if offsets[j] == int(sa[k]) { // if the idx is the same as the suffix array value
					index_arr = append(index_arr, k) // add suffix array index to list
				}
			}
		}
		// add start and end index to prefix talble
		pref_table = append(pref_table, Perfidx{start: int32(findMinElement(index_arr)), end: int32(findMaxElement(index_arr) + 1)})
	}

	return pref_table

}
