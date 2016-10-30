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
	ARC_RUNES = map[int]string{
		0: "z", 1: "a", 2: "æ", 3: "b", 4: "c",
		5: "d", 6: "ð", 7: "e", 8: "f",
		9: "ᵹ", 10: "h", 11: "i", 12: "l",
		13: "m", 14: "n", 15: "o", 16: "p",
		17: "r", 18: "ſ", 19: "t", 20: "þ",
		21: "u", 22: "ƿ", 23: "x", 24: "y",
		25: " ", 26: "\n", 27: "A", 28: "Æ",
		29: "B", 30: "C", 31: "D", 32: "Ð",
		33: "E", 34: "F", 35: "Ᵹ", 36: "H",
		37: "I", 38: "L", 39: "M", 40: "N",
		41: "O", 42: "P", 43: "R", 44: "S",
		45: "T", 46: "Þ", 47: "U", 48: "Ƿ",
		49: "X", 50: "Y", 51: "⁊", 52: "·",
		53: "˙", 54: ".", 55: "&", 56: "†",
		57: "‡", 58: "♀", 59: "☉", 60: "-"}
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

func Run(script string) (string, error) {
	var (
		output    string
		newOutput string
		err       error
	)

	tokens, tokenErr := Tokenize(script)

	if tokenErr != nil {
		return "", tokenErr
	}

	registers := new(Registers)

	for i := 0; i < len(tokens); {
		i, newOutput, err = whichRun(tokens, registers, i)
		if err != nil {
			return "", err
		}
		output += newOutput
	}

	return output, nil
}

func whichRun(tokens []int, registers *Registers, index int) (int, string, error) {
	var (
		activeRegister int
	)

	switch tokens[index] {
	case T_AUGHT, T_AN, T_TWEGEN, T_THRIE, T_FEOWER:
		bytes := []byte{uint8(tokens[index])}
		writeToRegister(registers, activeRegister, bytes)
		index++

	case T_ATIMBRAN:
		if len(tokens)-index-1 < 1 {
			return index, "", &Error{"insufficient tokens for operation", index, ARC_WORDS[tokens[index]]}
		}
		if isNumericToken(tokens[index+1]) {
			activeRegister = tokens[index+1]
		}
		index += 2

	case T_AEFTER:
		if activeRegister == 3 {
			activeRegister = 0
			break
		}
		activeRegister++
		index++

	case T_AEGTHER, T_CEOSAN, T_ONGEAN, T_HEAH, T_EBBA:
		result, err := runOperator(tokens, index)
		if err != nil {
			return index, "", err
		}
		writeToRegister(registers, activeRegister, []byte{result})
		index += 3

	case T_AFTERSONA:
		if index < 1 {
			return index, "", &Error{"previous operation does not exist", index, ARC_WORDS[tokens[index]]}
		}
		// TODO: make this more robust
		tokens[index] = -1
		index++

	case -1:
		index++

	case T_ACWETHAN:
		index++
		return index, outputAllRegisters(registers), nil

	default:
		return index, "", &Error{"invalid token found", index, ARC_WORDS[tokens[index]]}
	}

	return index, "", nil
}

// Hoh boy, is this function squirrely
// I know I'm going to look at this function tomorrow and say audibly "what the actual fuck"
func runOperator(tokens []int, index int) (byte, error) {
	var result byte

	if tokens[index] != T_ONGEAN && len(tokens)-index-1 < 2 {
		result = uint8(tokens[index+1]) & uint8(tokens[index+2])
		return 0, &Error{"insufficient tokens for operation", index, ARC_WORDS[tokens[index]]}
	}

	// Validate the next tokens to make sure they're valid for an operator
	switch tokens[index+1] {
	case T_AEGTHER, T_CEOSAN, T_ONGEAN, T_HEAH, T_EBBA:
		// Recursive!!
		result, err := runOperator(tokens, index+1)
		if err != nil {
			return result, err
		}

		if tokens[index] != T_ONGEAN {
			return result, &Error{"got in a bad state", index, ARC_WORDS[tokens[index]]}
		}
		// Only operator that only takes a single argument is NOT (T_ONGEAN)
		return ^uint8(result), nil

	// TODO: find a cleaner way to do this
	case T_AUGHT, T_AN, T_TWEGEN, T_THRIE, T_FEOWER:
		if tokens[index] != T_ONGEAN {
			switch tokens[index+2] {
			case T_AUGHT, T_AN, T_TWEGEN, T_THRIE, T_FEOWER:
				break
			default:
				return result, &Error{"invalid tokens supplied to operator", index, ARC_WORDS[tokens[index]]}
			}
		}

	default:
		return result, &Error{"invalid tokens supplied to operator", index, ARC_WORDS[tokens[index]]}
	}

	switch tokens[index] {
	case T_ONGEAN:
		result = ^uint8(tokens[index+1])
	case T_CEOSAN:
		result = uint8(tokens[index+1]) | uint8(tokens[index+2])
	case T_AEGTHER:
		result = uint8(tokens[index+1]) ^ uint8(tokens[index+2])
	// TODO: Implement all the operators
	case T_HEAH:
	case T_EBBA:
	}

	return result, nil
}

func writeToRegister(registers *Registers, active int, bytes []byte) {
	regEnd := len(registers[active]) - 1
	for i := 0; i < len(bytes); i++ {
		for j := regEnd; j > 0; j-- {
			registers[active][j-1] = registers[active][j]
		}
		registers[active][regEnd] = bytes[i]
	}
}

func outputAllRegisters(registers *Registers) string {
	var (
		output = ""
	)

	for i, readBytes, currReg := 0, 0, 0; readBytes < 12; readBytes++ {
		output += ARC_RUNES[int(registers[currReg][i])]
		registers[currReg][i] = 0
		if currReg == 3 {
			currReg = 0
			i++
		} else {
			currReg++
		}
	}

	return output
}

func isNumericToken(token int) bool {
	switch token {
	case T_AUGHT, T_AN, T_TWEGEN, T_THRIE, T_FEOWER:
		return true
	}
	return false
}

func main() {
	output, err := Run("Atimbran an þrie aught twegen ongean ceosan feower twegen æftersona aught ceosan an feower an ongean þrie æfter aught ceosan þrie feower ongean aught æftersona ongean twegen aught an aught feower æfter an twegen aught feower ongean feower aught ongean an feower æftersona æfter an ongean ceosan feower twegen þrie aught ceosan an feower aught ongean ceosan feower twegen an feower acweþan")
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Println(output)
}
