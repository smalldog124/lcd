package lcd

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Lcd_Input_1_Should_Be_one(t *testing.T) {
	expected := `
|
|`

	actual := Lcd("1")

	t.Log(actual)
	assert.Equal(t, expected, actual)
}

func Test_Lcd_Input_2_Should_Be_Tow(t *testing.T) {
	expected := ` _ 
 _|
|_`

	actual := Lcd("2")

	assert.Equal(t, expected, actual)
}

func Test_Lcd_Input_74_Should_Be_SeventFourZero(t *testing.T) {
	expected := `_    
 ||_|
 |  |`

	actual := Lcd("74")
	assert.Equal(t, expected, actual)
}
