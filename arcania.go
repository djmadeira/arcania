package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	T_AUGHT = iota
	T_AN
	T_TWEGEN
	T_THRIE
	T_FEOWER
	T_AEGTHER
	T_CEOSAN
	T_ONGEAN
	T_HEAH
	T_EBBA
	T_ACWETHAN
	T_ATIMBRAN
	T_AFTERSONA
	T_AEFTER
)

var (
	ARC_RUNES = map[int]rune{
		1: 'a', 2: 'æ', 3: 'b', 4: 'c',
		5: 'd', 6: 'ð', 7: 'e', 8: 'f',
		9: 'ᵹ', 10: 'h', 11: 'i', 12: 'l',
		13: 'm', 14: 'n', 15: 'o', 16: 'p',
		17: 'r', 18: 'ſ', 19: 't', 20: 'þ',
		21: 'u', 22: 'ƿ', 23: 'x', 24: 'y',
		25: ' ', 26: '\n', 27: 'A', 28: 'Æ',
		29: 'B', 30: 'C', 31: 'D', 32: 'Ð',
		33: 'E', 34: 'F', 35: 'Ᵹ', 36: 'H',
		37: 'I', 38: 'L', 39: 'M', 40: 'N',
		41: 'O', 42: 'P', 43: 'R', 44: 'S',
		45: 'T', 46: 'Þ', 47: 'U', 48: 'Ƿ',
		49: 'X', 50: 'Y', 51: '⁊', 52: '·',
		53: '˙', 54: '.', 55: '&', 56: '†',
		57: '‡', 58: '♀', 59: '☉', 60: '-'}
	ARC_WORDS = map[int]string{
		T_AUGHT:     "aught",
		T_AN:        "an",
		T_TWEGEN:    "twegen",
		T_THRIE:     "þrie",
		T_FEOWER:    "feower",
		T_AEGTHER:   "ægþer",
		T_CEOSAN:    "ceosan",
		T_ONGEAN:    "ongean",
		T_HEAH:      "heah",
		T_EBBA:      "ebba",
		T_ACWETHAN:  "acweþan",
		T_ATIMBRAN:  "atimbran",
		T_AFTERSONA: "æftersona",
		T_AEFTER:    "æfter"}
)

type Registers [4][18]byte

type Error struct {
	errMessage    string
	errOffset     int
	errExpression string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s at word %d, expression: %s", e.errMessage, e.errOffset, e.errExpression)
}

func Tokenize(script string) (output []int, err *Error) {
	words := strings.Fields(script)

	output = make([]int, 0)

	// TODO: handle hanging punctuation, as in word.
	for i, word := range words {
		var w int
		switch strings.ToLower(word) {
		case "aught":
			w = T_AUGHT
		case "an":
			w = T_AN
		case "twegen":
			w = T_TWEGEN
		case "þrie":
			w = T_THRIE
		case "feower":
			w = T_FEOWER
		case "ægþer":
			w = T_AEGTHER
		case "ceosan":
			w = T_CEOSAN
		case "ongean":
			w = T_ONGEAN
		case "heah":
			w = T_HEAH
		case "ebba":
			w = T_EBBA
		case "acweþan":
			w = T_ACWETHAN
		case "atimbran":
			w = T_ATIMBRAN
		case "æftersona":
			w = T_AFTERSONA
		case "æfter":
			w = T_AEFTER
		case ".", "⁊", "·", "˙", "&", "†", "‡", "♀", "☉":
			continue
		default:
			return output, &Error{"invalid word found", i, word}
		}
		output = append(output, w)
	}

	return output, nil
}

func Run(script string) (output string, err error) {
	tokens, tokenErr := Tokenize(script)

	if tokenErr != nil {
		return "", tokenErr
	}

	registers := new(Registers)

	for i := 0; i < len(tokens); {
		i, err = runToken(&tokens, registers, i)
		if err != nil {
			return "", err
		}
	}

	return output, nil
}

func runToken(tokens *[]int, registers *Registers, index int) (int, error) {
	return index + 1, nil
}

func main() {
	output, err := Run("Atimbran an þrie aught twegen ongean ceosan feower twegen æftersona aught ceosan an feower an ongean þrie æfter aught ceosan þrie feower ongean aught æftersona ongean twegen aught an aught feower æfter an twegen aught feower ongean feower aught ongean an feower æftersona æfter an ongean ceosan feower twegen þrie aught ceosan an feower aught ongean ceosan feower twegen an feower")
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(output)
}
