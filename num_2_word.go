package khmer_number

import (
  "fmt"
  "strconv"
  "strings"
)

func int2Word(num int64, space string) string {
  if num == 0 {
    return "សូន្យ"
  } else {
    return int2Khmer(num, space)
  }
}

func int2Khmer(num int64, space string) string {
  if num >= 1000000 {
    return fmt.Sprintf("%sលាន%s%s", int2Khmer(num / 1000000, space), space, int2Khmer(num % 1000000, space))
  } else if num >= 100000 {
    return fmt.Sprintf("%sសែន%s%s", int2Khmer(num / 100000, space), space, int2Khmer(num % 100000, space))
  } else if num >= 10000 {
    return fmt.Sprintf("%sម៉ឺន%s%s", int2Khmer(num / 10000, space), space, int2Khmer(num % 10000, space))
  } else if num >= 1000 {
    return fmt.Sprintf("%sពាន់%s%s", int2Khmer(num / 1000, space), space,  int2Khmer(num % 1000, space))
  } else if num >= 100 {
    return fmt.Sprintf("%sរយ%s%s", int2Khmer(num / 100, space), space,  int2Khmer(num % 100, space))
  } else if num >= 10 {
    return fmt.Sprintf("%s%s%s", KHMER_TENTH[num / 10], space, KHMER_DIGIT_WORD[num % 10])
  } else if num > 0 {
    return KHMER_DIGIT_WORD[num];
  }

  return ""
}

func countLeading(str string, item rune) int {
  counter := 0
  for _, value := range str {
    if value == item {
      counter++
    } else {
      break
    }
  }
  return counter
}

func Num2Word(num string, space string) string {
  nums := strings.Split(Khmer2RomanNum(num), ".")
  if len(nums) == 1 {
    digit, _ := strconv.ParseInt(nums[0], 10, 64)
    return int2Word(digit, space)
  }

  digit, _ := strconv.ParseInt(nums[0], 10, 64)
  precision, _ := strconv.ParseInt(nums[1], 10, 64)
  lead_zero := ""
  for _, value := range nums[1] {
    if string(value) == "0" {
      lead_zero += space
      lead_zero += "សូន្យ"
    } else {
      break
    }
  }
  return fmt.Sprintf("%s%sចុច%s%s%s", int2Word(digit, space), space, lead_zero, space, int2Khmer(precision, space))
}
