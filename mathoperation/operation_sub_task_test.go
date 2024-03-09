package mathoperation_test

import (
	"testing"

	"github.com/pylinuxtech/go_lang_testing/mathoperation"
)

func TestSubTaskAdd(t *testing.T) {

	t.Run("Addition of two positive number ", func(t *testing.T) {
		// input
		a := 1
		b := 2
		expected := 3

		result := mathoperation.Add(a, b)

		if expected != result {
			t.Error("Test Fail")
		}

	})

	t.Run("Addition of two negative number ", func(t *testing.T) {

		a := -1
		b := -2
		expected := -3

		result := mathoperation.Add(a, b)

		if expected != result {
			t.Error("Test Fail")
		}

	})

}
