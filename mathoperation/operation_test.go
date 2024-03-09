package mathoperation_test

import (
	"testing"

	"github.com/PylinuxTech/go_lang_testing/mathoperation"
)

func TestAdd(t *testing.T) {
	t.Log("Additon of two number 1,2")
	// input 1 , 2
	// result = 3
	// expection = 3
	result := mathoperation.Add(1, 2)

	if result != 3 {
		t.Error("Addition of two number which 1 , 2 is not equal to 3")
		t.Errorf("Addition of tow number which %d %d is not equal to %d", 1, 2, 3)
	}

}

func TestSub(t *testing.T) {
	t.Log("Additon of two number 1,2")
	// input 1 , 2
	// result = 3
	// expection = 3
	result := mathoperation.Sub(1, 2)

	if result != -1 {
		t.Error("Substraction of two number which 1 , 2 is not equal to -1")
		t.Errorf("Substraction of tow number which %d %d is not equal to %d", 1, 2, -1)
	}

}
