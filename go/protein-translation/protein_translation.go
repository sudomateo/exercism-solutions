package protein

import (
	"errors"
)

var (
	// Custom errors.
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")

	// Translation map.
	translator = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine", "UUC": "Phenylalanine",
		"UUA": "Leucine", "UUG": "Leucine",
		"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
		"UAU": "Tyrosine", "UAC": "Tyrosine",
		"UGU": "Cysteine", "UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
	}
)

// FromCodon translates a condon to a protein.
func FromCodon(codon string) (string, error) {
	protein, ok := translator[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

// FromRNA translates an RNA to proteins.
func FromRNA(rna string) ([]string, error) {
	var ret []string
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			break
		}
		protein, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrStop {
				break
			}
			return ret, err
		}
		ret = append(ret, protein)
	}
	return ret, nil
}
