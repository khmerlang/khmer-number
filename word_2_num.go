package khmer_number

import (
  "fmt"
  "strings"
)

func text2Token(text string) ([]string, error) {
	cleanText := strings.ReplaceAll(text, " ", "")
	tokens := []string{}
	for {
		token := ""
		for _, word := range NUM_WORDS_LIST {
			wordLength := len(word)
			if wordLength <= len(cleanText) {
				if word == cleanText[0:wordLength] {
					token = word
					cleanText = cleanText[wordLength:]
					break
				}
			}
		}

		if token == "" && cleanText != "" {
			return nil, fmt.Errorf("cannot segment word")
		}
		tokens = append(tokens, token)
		if len(cleanText) <= 0 {
			break
		}
	}

	return tokens, nil
}

func text2Int(tokens []string) string {
	num := 0
  num_word := 0
	for _, token := range tokens {
    if n, isFound := KHMER_DIGIT_WORD_MAP[token]; isFound {
      num = n
    } else if n, isFound := KHMER_TENTH_MAP[token]; isFound {
      num = n
      num_word += num
    } else if token == "លាន" {
      num *= 1000000
      num_word += num
    } else if token == "សែន" {
      num *= 100000
      num_word += num
    } else if token == "ម៉ឺន" {
      num *= 10000
      num_word += num
    } else if token == "ពាន់" {
      num *= 1000
      num_word += num
    } else if token == "រយ" {
      num *= 100
      num_word += num
    }
	}

	if num < 10 {
		num_word += num
	}

	return fmt.Sprintf("%d", num_word)
}

func text2number(text string) (string, error) {
	tokens, err := text2Token(text)
	if err != nil {
		return "", err
	}

	posDot := IndexOf("ចុច", tokens)

	if posDot == -1 {
		return text2Int(tokens), nil
	}

	digit := tokens[0:posDot]
  precision := tokens[posDot + 1:]
  leadZero := ""
  count := 0
  for _, value := range precision {
    if value == "សូន្យ" {
      leadZero += "0"
      count += 1
    } else {
      break
    }
  }

  if count == len(precision) {
		return fmt.Sprintf("%s.%s", text2Int(digit), leadZero), nil
  } else {
		return fmt.Sprintf("%s.%s%s", text2Int(digit), leadZero, text2Int(precision)), nil
  }
}

func Word2NumEN(text string) (string, error) {
	num, err := text2number(text)
	if err != nil {
		return "", err
	}

	return Khmer2RomanNum(num), nil
}

func Word2NumKH(text string) (string, error) {
	num, err := text2number(text)
	if err != nil {
		return "", err
	}

	return Roman2KhmerNum(num), nil
}
