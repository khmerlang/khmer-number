package khmer_number

import (
  "testing"
)

func TestingNum2Word(num string, word string, t *testing.T) {
  result := Num2Word(num)
  if result != word {
    t.Errorf("(%s) expect = (%s), got (%s)", num, word, result)
  }
}

func TestConvertWordToRoman(t *testing.T) {
  TestingNum2Word("0", "សូន្យ", t)
  TestingNum2Word("1", "មួយ", t)
  TestingNum2Word("2", "ពីរ", t)
  TestingNum2Word("3", "បី", t)
  TestingNum2Word("4", "បួន", t)
  TestingNum2Word("5", "ប្រាំ", t)
  TestingNum2Word("6", "ប្រាំមួយ", t)
  TestingNum2Word("7", "ប្រាំពីរ", t)
  TestingNum2Word("8", "ប្រាំបី", t)
  TestingNum2Word("9", "ប្រាំបួន", t)
  TestingNum2Word("10", "ដប់", t)
  TestingNum2Word("20", "ម្ភៃ", t)
  TestingNum2Word("30", "សាមសិប", t)
  TestingNum2Word("40", "សែសិប", t)
  TestingNum2Word("50", "ហាសិប", t)
  TestingNum2Word("60", "ហុកសិប", t)
  TestingNum2Word("70", "ចិតសិប", t)
  TestingNum2Word("80", "ប៉ែតសិប", t)
  TestingNum2Word("90", "កៅសិប", t)
  TestingNum2Word("100", "មួយរយ", t)
  TestingNum2Word("1000", "មួយពាន់", t)
  TestingNum2Word("10000", "មួយម៉ឺន", t)
  TestingNum2Word("100000", "មួយសែន", t)
  TestingNum2Word("1000000", "មួយលាន", t)
  TestingNum2Word("1234", "មួយពាន់ពីររយសាមសិបបួន", t)
  TestingNum2Word("1234.78", "មួយពាន់ពីររយសាមសិបបួនចុចចិតសិបប្រាំបី", t)
  TestingNum2Word("10005.55", "មួយម៉ឺនប្រាំចុចហាសិបប្រាំ", t)
  TestingNum2Word("12.05", "ដប់ពីរចុចសូន្យប្រាំ", t)
}
