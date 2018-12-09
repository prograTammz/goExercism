package pythagorean

type Triplet [3]int

func Range(min int, max int) []Triplet {
	var result []Triplet
	const a, b, c int = 3, 4, 5
	value1, value2, value3 := a, b, c
	for value3 <= max {
		if value1 >= min && value3 <= max {
			temp := Triplet{value1, value2, value3}
			result = append(result, temp)

		}
		value1 += a
		value2 += b
		value3 += c
	}
	return result
}
func Sum(p int) []Triplet {
	var result []Triplet
	const a, b, c int = 3, 4, 5
	value1, value2, value3 := a, b, c

	for value3 <= p/2 {
		if value1+value2+value3 == p {
			temp := Triplet{value1, value2, value3}
			result = append(result, temp)

			value1 += a
			value2 += b
			value3 += c
		}
	}
	return result
}
