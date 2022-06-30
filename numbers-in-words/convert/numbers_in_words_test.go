package convert

import "testing"

func TestDoSomthing(t *testing.T) {
	want := true
	got := DoSomethingWrong()
	if got != want {
		t.Errorf("DoSomething() = %t, want %t", got, want)
	}
}

func TestConvert1DigitNumberToWord(t *testing.T) {
	want := "six dollars"
	input := 6
	got := ConvertNumberToWord(input)
	if got != want {
		t.Errorf("ConvertNumberToWord() = %s, want %s", got, want)
	}
}

func TestConvert2DigitNumberToWord(t *testing.T) {
	want := "sixteen dollars"
	input := 16
	got := ConvertNumberToWord(input)
	if got != want {
		t.Errorf("ConvertNumberToWord() = %s, want %s", got, want)
	}
}
