package copyandpaste

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		desc     string
		settings LinterSetting
	}{}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a, err := NewAnalyzer(test.settings)
			if err != nil {
				t.Fatal(err)
			}

			analysistest.Run(t, analysistest.TestData(), a, test.desc)
		})
	}
}
