package convert

import "testing"

const succeeded = "\u2713"
const failed = "\u2717"

func TestCelsiusToFahrenheit(t *testing.T) {
	tt := []struct {
		name           string
		input          Celsius
		expectedOutput Fahrenheit
	}{
		{
			name:           "test 1",
			input:          0,
			expectedOutput: 32,
		},
		{
			name:           "test 2",
			input:          31,
			expectedOutput: 87.8,
		},
	}

	t.Log("Test C -> F")
	{
		for i, tst := range tt {
			t.Logf("\tTest (%d) %s", i, tst.name)
			output := CelsiusToFahrenheit(tst.input)

			if output != tst.expectedOutput {
				t.Fatalf("\t%s\t Values dont match: exp[%v] act[%v]", failed, tst.expectedOutput, output)
			}
			t.Logf("\t%s\tTest passed", succeeded)

		}
	}
}
