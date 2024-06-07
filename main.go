package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HextoInt(hex string) string {
	number, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(number)
}

func BintoInt(bin string) string {
	number, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(number)
}

func Apostrophe(s string) string {
	count := 1
	r := []rune(s)
	for i := 0; i <= len(r)-1; i++ {
		if r[i] == '\'' && count%2 != 0 && i == len(r)-1 {
			// r = append(r[:i], r[i+1:]...)
			continue
		}
		if count%2 != 0 {
			if r[i] == '\'' {
				if r[i+1] == ' ' {
					r = append(r[:i+1], r[i+2:]...)
					count++
					continue
				} else {
					count++
					i++
				}
			}
		}
		if count%2 == 0 {
			if r[i] == '\'' && r[i-1] == ' ' {
				count++
				r = append(r[:i-1], r[i:]...)

			}
		}
	}
	return string(r)
}

func Punctuations(c rune) bool {
	if c == ',' || c == '.' || c == '!' || c == '?' || c == ':' || c == ';' {
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]

	// reading first file

	Text, _ := os.ReadFile(args[0])
	count := 0
	for i := 0; i < len(Text); i++ {
		if Text[len(Text)-1] == '\\' || Text[len(Text)-1] == '/' || Text[len(Text)-1] == ',' || Text[len(Text)-1] == ' ' || Text[len(Text)-1] == 'p' || Text[len(Text)-1] == 'w' {
			continue
		}
		if Text[i] == '(' {
			continue
		}
		if Text[i] == '(' {
			i++
			for Text[i+1] != ')' {
				if Text[i] != ')' {
					continue
				}
				if Text[i] == ' ' {
					count++
					Text = append(Text[:i], Text[i+1:]...)
					continue
				}
				i++
				if Text[i] == ' ' {
					count++
				}
			}
			for Text[i] == ' ' {
				Text = append(Text[:i], Text[i+1:]...)
			}
		}
	}

	for i := 0; i < len(Text); i++ {
		if Text[i] == ',' {
			Text = append(Text[:i+1], append([]byte{' '}, Text[i+1:]...)...)
		}
	}

	// array to push the words1 into
	word := strings.Split(string(Text), " ")
	var words1 []string
	// words := Apostrophe((word))
	for _, c := range word {
		if c == "" {
			continue
		} else {
			words1 = append(words1, c)
		}
	}
	TransA(words1)
	// Punc(words1)
	for i := 0; i < len(words1); i++ {
		if i == 0 && (words1[i] == "(cap)" || words1[i] == "(cap, " || words1[i] == "(Cap, " || words1[i] == "(CAP, " ||
			words1[i] == "(up, " || words1[i] == "(Up, " || words1[i] == "(UP, " || words1[i] == "(low, " || words1[i] == "(Low, " ||
			words1[i] == "(LOW, " || words1[i] == "(up)" || words1[i] == "(low)" || words1[i] == "(CAP)" || words1[i] == "(LOW)" ||
			words1[i] == "(UP)" || words1[i] == "(Cap)" || words1[i] == "(Low)" || words1[i] == "(Up)" || words1[i] == "(bin)" ||
			words1[i] == "(hex)" || words1[i] == "(BIN)" || words1[i] == "(Bin)" || words1[i] == "(HEX)" || words1[i] == "(Hex)" ||
			words1[i] == "." || words1[i] == "," || words1[i] == ";" || words1[i] == "!" || words1[i] == ":" || words1[i] == "?") {
			continue
		} else if words1[i] == "(up" || words1[i] == "(cap" || words1[i] == "(low" || (words1[i] == "(low," && i == len(words1)-1) || (words1[i] == "(cap," && i == len(words1)-1) || (words1[i] == "(up," && i == len(words1)-1) {
			continue
		}
		// if (words1[i] == "(hex)" && words1[i-1] != HextoInt(words1[i-1])) || (words1[i] == "(bin)" && words1[i-1] != BintoInt(words1[i-1])) {
		// 	words1 = words1[:i]
		// 	continue
		// }
		if words1[i] == "(cap)" && i == len(words1)-1 {
			words1[i-1] = strings.ToLower(words1[i-1])
			words1[i-1] = strings.Title(words1[i-1])
			words1 = words1[:i]
			continue
		} else if words1[i] == "(up)" && i == len(words1)-1 {
			words1[i-1] = strings.ToUpper(words1[i-1])
			words1 = words1[:i]
			continue
		} else if words1[i] == "(low)" && i == len(words1)-1 {
			words1[i-1] = strings.ToLower(words1[i-1])
			words1 = words1[:i]
			continue
		}
		if words1[i] == "(up)" && i > 0 {
			words1[i-1] = strings.ToUpper(words1[i-1])
			words1 = append(words1[:i], words1[i+1:]...)
			i--
			continue
		} else if words1[i] == "(low)" && i > 0 {
			words1[i-1] = strings.ToLower(words1[i-1])
			words1 = append(words1[:i], words1[i+1:]...)
			i--
			continue
		} else if words1[i] == "(cap)" && i > 0 {
			words1[i-1] = strings.ToLower(words1[i-1])
			words1[i-1] = strings.Title(words1[i-1])
			words1 = append(words1[:i], words1[i+1:]...)
			i--
			continue
		} else if words1[i] == "(hex)" {
			words1[i-1] = HextoInt(words1[i-1])
			words1 = append(words1[:i], words1[i+1:]...)
			i--
			continue
		} else if words1[i] == "(bin)" {
			words1[i-1] = BintoInt(words1[i-1])
			words1 = append(words1[:i], words1[i+1:]...)
			i--
			continue
			// upper with number
		} else if words1[i] == "(up," {
			b := strings.Trim(string(words1[i+1]), words1[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number && i-j >= 0; j++ {
				words1[i-j] = strings.ToUpper(words1[i-j])
			}
			words1 = append(words1[:i], words1[i+2:]...)
			i--
			continue
			// lower with number
		} else if words1[i] == "(low," {
			b := string(words1[i+1][:1])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number && i-j >= 0; j++ {
				words1[i-j] = strings.ToLower(words1[i-j])
			}
			words1 = append(words1[:i], words1[i+2:]...)
			i--
			continue
			// capitalize with num
		} else if words1[i] == "(cap," {
			b := strings.Trim(string(words1[i+1]), words1[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number && i-j >= 0; j++ {
				words1[i-j] = strings.Title(words1[i-j])
			}
			words1 = append(words1[:i], words1[i+2:]...)
			i--
			continue
		}
	}
	TransA(words1)

	for i := 0; i < len(words1); i++ {
		resPonc := ""
		resInit := ""
		for _, c := range words1[i] {
			if resInit == "" && Punctuations(c) {
				resPonc += string(c)
			} else {
				resInit += string(c)
			}
		}
		if resPonc != "" {
			index := i
			for index > 0 {
				if words1[index-1] != "" {
					words1[index-1] += resPonc
					break
				}
				index--
			}
			words1[i] = resInit
			resInit = ""
			resPonc = ""
		}
	}

	res := ""
	for i := 0; i < len(words1); i++ {
		if i == 0 && (words1[i] == "(cap)" || words1[i] == "(cap," || words1[i] == "(Cap," || words1[i] == "(CAP," ||
			words1[i] == "(up," || words1[i] == "(Up," || words1[i] == "(UP," || words1[i] == "(low," || words1[i] == "(Low," ||
			words1[i] == "(LOW," || words1[i] == "(up)" || words1[i] == "(low)" || words1[i] == "(CAP)" || words1[i] == "(LOW)" ||
			words1[i] == "(UP)" || words1[i] == "(Cap)" || words1[i] == "(Low)" || words1[i] == "(Up)" || words1[i] == "(bin)" ||
			words1[i] == "(hex)" || words1[i] == "(BIN)" || words1[i] == "(Bin)" || words1[i] == "(HEX)" || words1[i] == "(Hex)" ||
			words1[i] == "." || words1[i] == "," || words1[i] == ";" || words1[i] == "!" || words1[i] == ":" || words1[i] == "?") {
			continue
		}
		if words1[i] == "(cap)." || words1[i] == "(up)." || words1[i] == "(low)." {
			continue
		}

		if words1[i] != "" && words1[i] != " " && i < len(words1)-1 {
			res += words1[i] + " "
		} else {
			res += words1[i]
		}
	}

	resFinal := strings.Trim(res, " ")
	ress := Apostrophe(resFinal)

	re := regexp.MustCompile(`''`)
	ress = re.ReplaceAllString(ress, "' '")

	err := os.WriteFile(args[1], []byte(ress), 0o644)
	if err != nil {
		return
	}
}

func TransA(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

	for i := 0; i < len(s)-1; i++ {
		for _, letter := range vowels {
			if s[i] == "a" && strings.HasPrefix(strings.ToLower(s[i+1]), letter) {
				s[i] = "an"
			} else if s[i] == "A" && strings.HasPrefix(s[i+1], letter) {
				s[i] = "An"
			} else if s[i] == "'a" && strings.HasPrefix(strings.ToLower(s[i+1]), letter) {
				s[i] = "'an"
			} else if s[i] == "'A" && strings.HasPrefix(s[i+1], letter) {
				s[i] = "'An"
			}
		}
	}
	return s
}
