package khmer_number

import (
  "strings"
)

var KHMER_DIGIT_WORD_MAP = map[string]int {
  "សូន្យ": 0,
  "មួយ": 1,
  "ពីរ": 2,
  "បី": 3,
  "បួន": 4,
  "ប្រាំ": 5,
  "ប្រាំមួយ": 6,
  "ប្រាំពីរ": 7,
  "ប្រាំបី": 8,
  "ប្រាំបួន": 9,
}

var KHMER_TENTH_MAP = map[string]int {
  "ដប់": 10,
  "ម្ភៃ": 20,
  "សាមសិប": 30,
  "សែសិប": 40,
  "ហាសិប": 50,
  "ហុកសិប": 60,
  "ចិតសិប": 70,
  "ប៉ែតសិប": 80,
  "កៅសិប": 90,
}

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

var NUM_WORDS_LIST = []string {
  "ប្រាំមួយ",
  "ប្រាំពីរ",
  "ប្រាំបួន",
  "ប៉ែតសិប",
  "ប្រាំបី",
  "សាមសិប",
  "ហុកសិប",
  "ចិតសិប",
  "សូន្យ",
  "សែសិប",
  "ហាសិប",
  "កៅសិប",
  "ប្រាំ",
  "ពាន់",
  "ម៉ឺន",
  "ម្ភៃ",
  "ចុច",
  "សែន",
  "លាន",
  "ដប់",
  "មួយ",
  "ពីរ",
  "បួន",
  "រយ",
  "បី",
}

func Khmer2RomanNum(num string) string {
  khNum := ""
  khNum = strings.ReplaceAll(num, "០", "0")
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

func Roman2KhmerNum(num string) string {
  enNum := ""
  enNum = strings.ReplaceAll(num, "0", "០")
  enNum = strings.ReplaceAll(enNum, "1", "១")
  enNum = strings.ReplaceAll(enNum, "2", "២")
  enNum = strings.ReplaceAll(enNum, "3", "៣")
  enNum = strings.ReplaceAll(enNum, "4", "៤")
  enNum = strings.ReplaceAll(enNum, "5", "៥")
  enNum = strings.ReplaceAll(enNum, "6", "៦")
  enNum = strings.ReplaceAll(enNum, "7", "៧")
  enNum = strings.ReplaceAll(enNum, "8", "៨")
  enNum = strings.ReplaceAll(enNum, "9", "៩")
  enNum = strings.ReplaceAll(enNum, ",", "")
  enNum = strings.ReplaceAll(enNum, ".", ".")
  return enNum
}

func IndexOf(word string, data []string) (int) {
  for k, v := range data {
    if word == v {
      return k
    }
  }

  return -1    //not found.
}
