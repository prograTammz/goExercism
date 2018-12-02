// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	//if the value returned is zero
	if len(rhyme) == 0 {
		return []string{}
	}
	count := len(rhyme) - 1
	var result []string
	for i := 0; i < count; i++ {
		result = append(result, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
	}
	result = append(result, "And all for the want of a "+rhyme[0]+".")
	return result

}
