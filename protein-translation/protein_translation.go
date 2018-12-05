package protein

import "errors"

var (
	STOP           error = errors.New("STOP")
	ErrInvalidBase error = errors.New("Invalid codon")
)

//FromRNA takes a long string of RNA and convert it to string of values
func FromRNA(input string) ([]string, error) {
	var Protien []string
	var count int
	var tempString string
	for _, i := range input {
		count++
		tempString += string(i)
		if count%3 == 0 {
			val, err := FromCodon(tempString)
			if err == STOP {
				break
			}
			if err != nil {
				return Protien, err
			}
			Protien = append(Protien, val)
			tempString = ""
		}
	}
	return Protien, nil
}

//FromCodon Takes a single Cordon and gives it's value
func FromCodon(input string) (string, error) {
	switch input {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", STOP
	default:
		return "", ErrInvalidBase
	}
}
