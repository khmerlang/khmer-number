package khmer_number

import (
  "fmt"
  "strconv"
  "strings"
)

var KHMER_DIGIT_WORD = []string {
  "សូន្យ",
  "មួយ",
  "ពីរ",
  "បី",
  "បួន",
  "ប្រាំ",
  "ប្រាំមួយ",
  "ប្រាំពីរ",
  "ប្រាំបី",
  "ប្រាំបួន",
}

var KHMER_TENTH = []string{
  "",
  "ដប់",
  "ម្ភៃ",
  "សាមសិប",
  "សែសិប",
  "ហាសិប",
  "ហុកសិប",
  "ចិតសិប",
  "ប៉ែតសិប",
  "កៅសិប",
}

func int2Khmer(num int64) string {
  if num >= 1000000 {
    if num % 1000000 != 0 {
      return fmt.Sprintf("%sលាន%s", int2Khmer(num / 1000000), int2Khmer(num % 1000000))
    } else {
      return fmt.Sprintf("%sលាន", int2Khmer(num / 1000000))
    }
  } else if num >= 100000 {
    if num % 100000 != 0 {
      return fmt.Sprintf("%sសែន%s", int2Khmer(num / 100000), int2Khmer(num % 100000))
    } else {
      return fmt.Sprintf("%sសែន", int2Khmer(num / 100000))
    }
  } else if num >= 10000 {
    if num % 10000 != 0 {
      return fmt.Sprintf("%sម៉ឺន%s", int2Khmer(num / 10000), int2Khmer(num % 10000))
    } else {
      return fmt.Sprintf("%sម៉ឺន", int2Khmer(num / 10000))
    }
  } else if num >= 1000 {
    if num % 1000 != 0 {
      return fmt.Sprintf("%sពាន់%s", int2Khmer(num / 1000), int2Khmer(num % 1000))
    } else {
      return fmt.Sprintf("%sពាន់", int2Khmer(num / 1000))
    }
  } else if num >= 100 {
    if num % 100 != 0 {
      return fmt.Sprintf("%sរយ%s", int2Khmer(num / 100), int2Khmer(num % 100))
    } else {
      return fmt.Sprintf("%sរយ", int2Khmer(num / 100))
    }
  } else if num >= 10 {
    if num % 10 != 0 {
      return fmt.Sprintf("%s%s", KHMER_TENTH[num / 10], KHMER_DIGIT_WORD[num % 10])
    } else {
      return fmt.Sprintf("%s", KHMER_TENTH[num / 10])
    }
  } else if num >= 0 {
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

func Num2Word(num string) string {
  nums := strings.Split(num, ".")
  if len(nums) == 1 {
    digit, _ := strconv.ParseInt(nums[0], 10, 64)
    return int2Khmer(digit)
  }

  digit, _ := strconv.ParseInt(nums[0], 10, 64)
  precision, _ := strconv.ParseInt(nums[1], 10, 64)
  lead_zero := ""
  for _, value := range nums[1] {
    if string(value) == "0" {
      lead_zero += "សូន្យ"
    } else {
      break
    }
   }
  return fmt.Sprintf("%sចុច%s%s", int2Khmer(digit), lead_zero, int2Khmer(precision))
}
