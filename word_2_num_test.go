package khmer_number

import (
  "testing"
)

func TestingWord2Num(word string, num string, t *testing.T) {
  result, _ := Word2NumEN(word)
  if result != num {
    t.Errorf("(%s) expect = (%s), got (%s)", word, num, result)
  }
}

func TestingWord2NumKH(word string, num string, t *testing.T) {
  result, _ := Word2NumKH(word)
  if result != num {
    t.Errorf("(%s) expect = (%s), got (%s)", word, num, result)
  }
}

func TestConvertWordToNum(t *testing.T) {
  TestingWord2Num("សូន្យ", "0", t)
  TestingWord2Num("មួយ", "1", t)
  TestingWord2Num("ពីរ", "2", t)
  TestingWord2Num("បី", "3", t)
  TestingWord2Num("បួន", "4", t)
  TestingWord2Num("ប្រាំ", "5", t)
  TestingWord2Num("ប្រាំមួយ", "6", t)
  TestingWord2Num("ប្រាំពីរ", "7", t)
  TestingWord2Num("ប្រាំបី", "8", t)
  TestingWord2Num("ប្រាំបួន", "9", t)
  TestingWord2Num("ដប់", "10", t)
  TestingWord2Num("ម្ភៃ", "20", t)
  TestingWord2Num("សាមសិប", "30", t)
  TestingWord2Num("សែសិប", "40", t)
  TestingWord2Num("ហាសិប", "50", t)
  TestingWord2Num("ហុកសិប", "60", t)
  TestingWord2Num("ចិតសិប", "70", t)
  TestingWord2Num("ប៉ែតសិប", "80", t)
  TestingWord2Num("កៅសិប", "90", t)
  TestingWord2Num("មួយរយ", "100", t)
  TestingWord2Num("មួយពាន់", "1000", t)
  TestingWord2Num("មួយម៉ឺន", "10000", t)
  TestingWord2Num("មួយសែន", "100000", t)
  TestingWord2Num("មួយលាន", "1000000", t)
  TestingWord2Num("មួយពាន់ពីររយសាមសិបបួន", "1234", t)
  TestingWord2Num("មួយពាន់ពីររយសាមសិបបួនចុចចិតសិបប្រាំបី", "1234.78", t)
  TestingWord2Num("មួយម៉ឺនប្រាំចុចហាសិបប្រាំ", "10005.55", t)
  TestingWord2Num("ដប់ពីរចុចសូន្យប្រាំ", "12.05", t)
  TestingWord2Num("ដប់ពីរចុចសូន្យសូន្យសូន្យប្រាំ", "12.0005", t)
  TestingWord2Num("ដប់ពីរចុចសូន្យសូន្យសូន្យសូន្យ", "12.0000", t)
  TestingWord2Num("មួយរយដប់ប្រាំមួយលានប្រាំបីរយដប់ពាន់ពីររយពីរ", "116810202", t)
}

func TestConvertWordToKhmerNum(t *testing.T) {
  TestingWord2NumKH("សូន្យ", "០", t)
  TestingWord2NumKH("មួយ", "១", t)
  TestingWord2NumKH("ពីរ", "២", t)
  TestingWord2NumKH("បី", "៣", t)
  TestingWord2NumKH("បួន", "៤", t)
  TestingWord2NumKH("ប្រាំ", "៥", t)
  TestingWord2NumKH("ប្រាំមួយ", "៦", t)
  TestingWord2NumKH("ប្រាំពីរ", "៧", t)
  TestingWord2NumKH("ប្រាំបី", "៨", t)
  TestingWord2NumKH("ប្រាំបួន", "៩", t)
  TestingWord2NumKH("ដប់", "១០", t)
  TestingWord2NumKH("ម្ភៃ", "២០", t)
  TestingWord2NumKH("សាមសិប", "៣០", t)
  TestingWord2NumKH("សែសិប", "៤០", t)
  TestingWord2NumKH("ហាសិប", "៥០", t)
  TestingWord2NumKH("ហុកសិប", "៦០", t)
  TestingWord2NumKH("ចិតសិប", "៧០", t)
  TestingWord2NumKH("ប៉ែតសិប", "៨០", t)
  TestingWord2NumKH("កៅសិប", "៩០", t)
  TestingWord2NumKH("កៅសិបមួយ", "៩១", t)
  TestingWord2NumKH("មួយរយ", "១០០", t)
  TestingWord2NumKH("មួយរយមួយ", "១០១", t)
  TestingWord2NumKH("មួយពាន់", "១០០០", t)
  TestingWord2NumKH("មួយម៉ឺន", "១០០០០", t)
  TestingWord2NumKH("មួយសែន", "១០០០០០", t)
  TestingWord2NumKH("មួយលាន", "១០០០០០០", t)
  TestingWord2NumKH("មួយពាន់ពីររយសាមសិបបួន", "១២៣៤", t)
  TestingWord2NumKH("មួយពាន់ពីររយសាមសិបបួនចុចចិតសិបប្រាំបី", "១២៣៤.៧៨", t)
  TestingWord2NumKH("មួយម៉ឺនប្រាំចុចហាសិបប្រាំ", "១០០០៥.៥៥", t)
  TestingWord2NumKH("ដប់ពីរចុចសូន្យប្រាំ", "១២.០៥", t)
  TestingWord2NumKH("ដប់ពីរចុចសូន្យសូន្យសូន្យប្រាំ", "១២.០០០៥", t)
  TestingWord2NumKH("ដប់ពីរចុចសូន្យសូន្យសូន្យសូន្យ", "១២.០០០០", t)
}
