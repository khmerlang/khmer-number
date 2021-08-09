package khmer_number

import (
  "testing"
)

func TestingNum2Word(num string, word string, t *testing.T) {
  result := Num2Word(num, "")
  if result != word {
    t.Errorf("(%s) expect = (%s), got (%s)", num, word, result)
  }
}

func TestingNum2WordWithSpace(num string, word string, t *testing.T) {
  result := Num2Word(num, " ")
  if result != word {
    t.Errorf("(%s) expect = (%s), got (%s)", num, word, result)
  }
}

func TestConvertNumToWord(t *testing.T) {
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
  TestingNum2Word("12.0005", "ដប់ពីរចុចសូន្យសូន្យសូន្យប្រាំ", t)
  TestingNum2Word("12.0000", "ដប់ពីរចុចសូន្យសូន្យសូន្យសូន្យ", t)
}

func TestConvertNumToWordWithSpace(t *testing.T) {
  TestingNum2WordWithSpace("0", "សូន្យ", t)
  TestingNum2WordWithSpace("1", "មួយ", t)
  TestingNum2WordWithSpace("2", "ពីរ", t)
  TestingNum2WordWithSpace("3", "បី", t)
  TestingNum2WordWithSpace("4", "បួន", t)
  TestingNum2WordWithSpace("5", "ប្រាំ", t)
  TestingNum2WordWithSpace("6", "ប្រាំមួយ", t)
  TestingNum2WordWithSpace("7", "ប្រាំពីរ", t)
  TestingNum2WordWithSpace("8", "ប្រាំបី", t)
  TestingNum2WordWithSpace("9", "ប្រាំបួន", t)
  TestingNum2WordWithSpace("10", "ដប់ ", t)
  TestingNum2WordWithSpace("20", "ម្ភៃ ", t)
  TestingNum2WordWithSpace("30", "សាមសិប ", t)
  TestingNum2WordWithSpace("40", "សែសិប ", t)
  TestingNum2WordWithSpace("50", "ហាសិប ", t)
  TestingNum2WordWithSpace("60", "ហុកសិប ", t)
  TestingNum2WordWithSpace("70", "ចិតសិប ", t)
  TestingNum2WordWithSpace("80", "ប៉ែតសិប ", t)
  TestingNum2WordWithSpace("90", "កៅសិប ", t)
  TestingNum2WordWithSpace("100", "មួយរយ ", t)
  TestingNum2WordWithSpace("1000", "មួយពាន់ ", t)
  TestingNum2WordWithSpace("10000", "មួយម៉ឺន ", t)
  TestingNum2WordWithSpace("100000", "មួយសែន ", t)
  TestingNum2WordWithSpace("1000000", "មួយលាន ", t)
  TestingNum2WordWithSpace("1234", "មួយពាន់ ពីររយ សាមសិប បួន", t)
  TestingNum2WordWithSpace("1234.78", "មួយពាន់ ពីររយ សាមសិប បួន ចុច ចិតសិប ប្រាំបី", t)
  TestingNum2WordWithSpace("10005.55", "មួយម៉ឺន ប្រាំ ចុច ហាសិប ប្រាំ", t)
  TestingNum2WordWithSpace("12.05", "ដប់ ពីរ ចុច សូន្យ ប្រាំ", t)
  TestingNum2WordWithSpace("12.0005", "ដប់ ពីរ ចុច សូន្យ សូន្យ សូន្យ ប្រាំ", t)
  TestingNum2WordWithSpace("12.0000", "ដប់ ពីរ ចុច សូន្យ សូន្យ សូន្យ សូន្យ ", t)
}

func TestConvertKhmerNumToWord(t *testing.T) {
  TestingNum2Word("០", "សូន្យ", t)
  TestingNum2Word("១", "មួយ", t)
  TestingNum2Word("២", "ពីរ", t)
  TestingNum2Word("៣", "បី", t)
  TestingNum2Word("៤", "បួន", t)
  TestingNum2Word("៥", "ប្រាំ", t)
  TestingNum2Word("៦", "ប្រាំមួយ", t)
  TestingNum2Word("៧", "ប្រាំពីរ", t)
  TestingNum2Word("៨", "ប្រាំបី", t)
  TestingNum2Word("៩", "ប្រាំបួន", t)
  TestingNum2Word("១០", "ដប់", t)
  TestingNum2Word("២០", "ម្ភៃ", t)
  TestingNum2Word("៣០", "សាមសិប", t)
  TestingNum2Word("៤០", "សែសិប", t)
  TestingNum2Word("៥០", "ហាសិប", t)
  TestingNum2Word("៦០", "ហុកសិប", t)
  TestingNum2Word("៧០", "ចិតសិប", t)
  TestingNum2Word("៨០", "ប៉ែតសិប", t)
  TestingNum2Word("៩០", "កៅសិប", t)
  TestingNum2Word("១០០", "មួយរយ", t)
  TestingNum2Word("១០០០", "មួយពាន់", t)
  TestingNum2Word("១០០០០", "មួយម៉ឺន", t)
  TestingNum2Word("១០០០០០", "មួយសែន", t)
  TestingNum2Word("១០០០០០០", "មួយលាន", t)
  TestingNum2Word("១២៣៤", "មួយពាន់ពីររយសាមសិបបួន", t)
  TestingNum2Word("១២៣៤.៧៨", "មួយពាន់ពីររយសាមសិបបួនចុចចិតសិបប្រាំបី", t)
  TestingNum2Word("១០០០៥.៥៥", "មួយម៉ឺនប្រាំចុចហាសិបប្រាំ", t)
  TestingNum2Word("១២.០៥", "ដប់ពីរចុចសូន្យប្រាំ", t)
  TestingNum2Word("១២.០០០៥", "ដប់ពីរចុចសូន្យសូន្យសូន្យប្រាំ", t)
  TestingNum2Word("១២.០០០០", "ដប់ពីរចុចសូន្យសូន្យសូន្យសូន្យ", t)
}
