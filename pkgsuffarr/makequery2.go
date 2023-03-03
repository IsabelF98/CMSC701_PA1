// This file contains the function for finding the next highest order query for query seach

package pkgsuffarr

func replacechar(old string, char rune, i int) string {
	/*
		This fucntion replaces a character in a string with another character at index i
		based on genome alphabet A, C, G, T

		Inputs:
			old(string): the original string being modifies
			char(rune): the character being added
			i(int): the index at which the swap is happening

		Outputs:
			new(string): the modified string
	*/
	out := []rune(old) // change string to rune type (easier to work with)
	out[i] = char      // swap in the new character
	return string(out) // return new string
}

func NextOrderQuery(query string) string {
	/*
		This function creates the next highest order query based on the initial query

		Inputs:
			query(string): the initial query

		Outputs:
			query2(string): the next order query
	*/

	n := len(query)     // length of query
	char := query[n-1:] // the last character in the query

	// switch last character of the query with the next highest character
	if char == "A" {
		query2 := replacechar(query, 'C', n-1)
		return query2
	} else if char == "C" {
		query2 := replacechar(query, 'G', n-1)
		return query2
	} else if char == "G" {
		query2 := replacechar(query, 'T', n-1)
		return query2
	} else {
		// if the last character is T need to start at the begining of query
		for j := 0; j < len(query); j++ {
			char := string(query[j])
			if char == "A" {
				query2 := replacechar(query, 'C', j)
				return query2
			} else if char == "C" {
				query2 := replacechar(query, 'G', j)
				return query2
			} else if char == "G" {
				query2 := replacechar(query, 'T', j)
				return query2
			} else {
				continue
			}
		}
		// querry is all T's
		query2 := query + "A"
		return query2
	}
}
