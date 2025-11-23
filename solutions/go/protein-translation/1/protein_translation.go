package protein

import (
    "errors"
    "strings"
)

var (
    ErrStop = errors.New("Stop sequence encountered.")
    ErrInvalidBase = errors.New("Invalid codon base.")
)

func FromRNA(rna string) ([]string, error) {
	var codon string
    ret := []string{}

    rna = strings.ToUpper(rna)
    
	for i := 0; i < len(rna); i += 3 {
        codon = rna[i : i + 3]
        prot, err := FromCodon(codon)

        if err != nil {
            if err == ErrInvalidBase {
            	return []string{}, err
            } else {
                return ret, nil
            }
        }

        ret = append(ret, prot)
    }

    return ret, nil
}

func FromCodon(codon string) (string, error) {
	codonProt := map[string]string{
        "AUG": "Methionine",
        "UUU": "Phenylalanine",
        "UUC": "Phenylalanine",
        "UUA": "Leucine",
        "UUG": "Leucine",
        "UCU": "Serine",
        "UCC": "Serine",
        "UCA": "Serine",
        "UCG": "Serine",
        "UAU": "Tyrosine",
        "UAC": "Tyrosine",
        "UGU": "Cysteine",
        "UGC": "Cysteine",
        "UGG": "Tryptophan",
        "UAG": "STOP",
        "UAA": "STOP",
        "UGA": "STOP",
    }

    prot, present := codonProt[codon]

    if !present {
        return "", ErrInvalidBase
    }

    if prot == "STOP" {
        return prot, ErrStop
    }

    return prot, nil
}
