package strain

type Ints []int
type Lists [][]int
type Strings []string

//Keep returns all the approved given value
func (right Ints) Keep(filter func(int) bool) Ints {
	var output Ints
	for _, val := range right {
		if filter(val) {
			output = append(output, val)
		}
	}
	return output
}

//Discard return all refused values
func (right Ints) Discard(filter func(int) bool) Ints {
	var output Ints
	for _, val := range right {
		if !filter(val) {
			output = append(output, val)
		}
	}
	return output
}
func (right Lists) Keep(filter func([]int) bool) Lists {
	var output Lists
	for _, val := range right {
		if filter(val) {
			output = append(output, val)
		}
	}
	return output
}
func (right Strings) Keep(filter func(string) bool) Strings {
	var output Strings
	for _, val := range right {
		if filter(val) {
			output = append(output, val)
		}
	}
	return output
}
