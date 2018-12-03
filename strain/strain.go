package strain

type Ints []int
type Lists [][]int
type Strings []string

func (val Ints) Keep(boolVal func(int) bool) Ints {
	var result Ints
	for _, num := range val {
		if boolVal(num) {
			result = append(result, num)
		}
	}
	return result
}
