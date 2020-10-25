package proverb

// Proverb will output a relevant proverb based on a list of input words.
func Proverb(rhyme []string) []string {
	var proverb []string
	if len(rhyme) == 0 {
		return proverb
	}

	for i, word := range rhyme {
		if i < len(rhyme)-1 {
			proverb = append(proverb, "For want of a "+word+" the "+rhyme[i+1]+" was lost.")
		} else {
			proverb = append(proverb, "And all for the want of a "+rhyme[0]+".")
		}
	}
	return proverb
}
