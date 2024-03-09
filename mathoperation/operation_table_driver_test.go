package mathoperation_test

import (
	"testing"

	"github.com/pylinuxtech/go_lang_testing/mathoperation"
)

func TestTableDriven(t *testing.T) {

	type args struct {
		a int
		b int
	}

	listInput := []struct {
		testName string
		argument args
		expected int
	}{
		{
			testName: "Addition of two positve number",
			argument: args{1, 2},
			expected: 3,
		},
		{
			testName: "Addition of two negavtive number",
			argument: args{-1, -2},
			expected: -3,
		},
		{
			testName: "Addition of one negavtive and one +ve number",
			argument: args{-1, 2},
			expected: 1,
		},
	}

	for _, input := range listInput {
		t.Run(input.testName, func(t *testing.T) {
			result := mathoperation.Add(input.argument.a, input.argument.b)
			if result != input.expected {
				t.Error("Not able compare ")
			}
		})
	}

}
