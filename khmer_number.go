package khmer_number

import (
  "fmt"
  "strconv"
  "strings"
)

var KHMER_DIGIT_WORD = []string {
  "",
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

func khmer2RomanNum(roman string) string {
  khNum := ""
  khNum = strings.ReplaceAll(roman, "០", "0")
  khNum = strings.ReplaceAll(khNum, "១", "1")
  khNum = strings.ReplaceAll(khNum, "២", "2")
  khNum = strings.ReplaceAll(khNum, "៣", "3")
  khNum = strings.ReplaceAll(khNum, "៤", "4")
  khNum = strings.ReplaceAll(khNum, "៥", "5")
  khNum = strings.ReplaceAll(khNum, "៦", "6")
  khNum = strings.ReplaceAll(khNum, "៧", "7")
  khNum = strings.ReplaceAll(khNum, "៨", "8")
  khNum = strings.ReplaceAll(khNum, "៩", "9")
  khNum = strings.ReplaceAll(khNum, ",", "")
  khNum = strings.ReplaceAll(khNum, ".", ".")
  return khNum
}


func Num2Word(num string, space string) string {
  nums := strings.Split(khmer2RomanNum(num), ".")
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
