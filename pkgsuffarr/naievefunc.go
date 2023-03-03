// This file contains the function for the naive search algorithum

package pkgsuffarr

import (
	"math"
)

func NaiveSearch(suffixes []string, query string) int {
	/*
		This function finds the index in the suffix array where the first instance of the query occurs

		Inputs:
			suffixarray(MyIndex): the suffix array as it was created in the BuildSA function
			query(string): the query we are looking for in the sffix array

		Output:
			index(int): the index of the first instance of the query
	*/
	// starting index of the left, center, and right index
	left := 0
	right := len(suffixes) - 1
	aux := float64(left+right) / 2.0
	center := int(math.Floor(aux))

	// initial checks
	if query < suffixes[center] { // look at top half
		right = int(center) // change right index
		aux = float64(left+right) / 2.0
		center = int(math.Floor(aux)) //change cneter index
	}
	if query > suffixes[center] {
		left = int(center) // change left index
		aux = float64(left+right) / 2.0
		center = int(math.Floor(aux)) //change cetnter index
	}
	if query == suffixes[center] { // query is a direct match
		return center // done
	}

	for i := 0; i < len(suffixes); i++ {
		// stopping criterial
		if query < suffixes[center] && center == left+1 {
			return center
		}
		if query > suffixes[center] && center == right-1 {
			return right
		}
		if query == suffixes[center] { // query is a direct match
			return center
		}

		if query < suffixes[center] { // look at top half
			right = int(center)
			aux = float64(left+right) / 2.0
			center = int(math.Floor(aux))
			continue
		}
		if query > suffixes[center] { // look at bottom half
			left = int(center)
			aux = float64(left+right) / 2.0
			center = int(math.Floor(aux))
			continue
		}

	}
	return -1 // return -1 if error occured (not a possible index)
}

// }
// 	// starting index of the left and right most index
// 	left := 0
// 	right := len(suffixes) - 1

// 	for i := 0; i < len(suffixes); i++ {
// 		// compute center index
// 		aux := float64(left+right) / 2.0
// 		center := int(math.Floor(aux))

// 		// look at top half
// 		if query < suffixes[center] {
// 			if center == left+1 { //stopping criteria
// 				return center
// 			}
// 			right = int(center) // change right index
// 			continue
// 		}

// 		// look at bottom half
// 		if query > suffixes[center] {
// 			if center == right+1 { // stopping criteria
// 				return right
// 			}
// 			left = int(center) //change left index
// 			continue
// 		}

// 	}
// 	return -1 // return -1 if error occured (not a possible index)
// }
