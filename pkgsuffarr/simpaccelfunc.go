// This file contains the function for the simple accelerant algorithum

package pkgsuffarr

import (
	"math"
)

func lcp(str1 string, str2 string) int {
	/*
		This function computes the longest common prefix given two strings

		Inputs:
			str1(string): the first string
			str2(string): the second string

		Output:
			index(int): the index of the first instance where the characters no longer mathch
	*/
	n := int(math.Min(float64(len(str1)), float64(len(str2)))) // minimum legth of string 1 and 2
	for i := 0; i < n; i++ {
		if str1[i] == str2[i] { // if the characters match
			continue
		} else { // if the characters dont match
			return i
		}
	}
	return n + 1 // index if all characters of the smallest string match
}

func SimpAccel(suffixes []string, query string) int {
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
	} else if query > suffixes[center] {
		left = int(center) // change left index
		aux = float64(left+right) / 2.0
		center = int(math.Floor(aux)) //change cetnter index
	} else { // query is a direct match
		return center // done
	}

	lcpL := lcp(query, suffixes[left])  // LCP of query and left
	lcpR := lcp(query, suffixes[right]) // LCP of query and right

	for i := 0; i < len(suffixes); i++ {
		idx := int(math.Min(float64(lcpL), float64(lcpR))) // new starting index based on LCP
		if query[idx:] < suffixes[center][idx:] {          // look at top half
			if center == left+1 { // stopping criteria
				return center
			}
			right = int(center)
			lcpR = lcp(query, suffixes[right])
			aux = float64(left+right) / 2.0
			center = int(math.Floor(aux))
			continue
		} else if query[idx:] > suffixes[center][idx:] { // look at bottom half
			if center == right-1 { // stopping criteria
				return right
			}
			left = int(center)
			lcpL = lcp(query, suffixes[left])
			aux = float64(left+right) / 2.0
			center = int(math.Floor(aux))
			continue
		} else { // query is a direct match
			return center // done
		}
	}
	return -1 // return -1 if error occured (not a possible index)
}
