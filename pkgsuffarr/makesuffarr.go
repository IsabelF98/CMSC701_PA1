package pkgsuffarr

import (
	"index/suffixarray"
)

func MakeSuffArr(gene string) *suffixarray.Index {
	gene_byte := []byte(gene)
	sa := suffixarray.New(gene_byte)
	return sa
}
