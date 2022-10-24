package main

func DNAtoRNA(dna string) string {
	var rnaString string
	for _, c := range dna {
		if c == 'T' {
			rnaString += "U"
		} else {
			rnaString += string(c)
		}
	}

	return rnaString
}
