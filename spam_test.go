package main

import (
	"testing"
)
//fhfhf
func TestSpam(t *testing.T) {
	//Arrange
	testTable := []struct {
		str      string
		expected string
	}{
		{
			str:      "http://sdfsfkfmsef.jj fsesefhgjhjhs",
			expected: "http://************** fsesefhgjhjhs",
		},
		{
			str:      "Перейди по ссылке и получи 100500 мильенов http://youarelose.com",
			expected: "Перейди по ссылке и получи 100500 мильенов http://**************",
		},
		{
			str:      "узнай кем ты был в прошлой жизни и кем будешь в следующей: http://youarelose",
			expected: "узнай кем ты был в прошлой жизни и кем будешь в следующей: http://**********",
		},
		{
			str:      "http://youarelose.jhjhjh - раздаем биткоины по 1к доллапров",
			expected: "http://***************** - раздаем биткоины по 1к доллапров",
		},
		{
			str:      "тут весь компромат на твоего соседа: http://sosedyam.net данные доступны в течении трех часов",
			expected: "тут весь компромат на твоего соседа: http://************ данные доступны в течении трех часов",
		},
		{
			str:      "ШОК! Этот отвар поможет избавиться от порчи раз и навсегда!: hTTp://magicotvar.ru ",
			expected: "ШОК! Этот отвар поможет избавиться от порчи раз и навсегда!: hTTp://magicotvar.ru ",
		},
	}
	//Act
	for _, testCase := range testTable {
		result := SpamMasker(testCase.str)

		//Assert
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}

}
