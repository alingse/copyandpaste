package copyandpaste_test

import (
	"testing"

	"github.com/alingse/copyandpaste"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		settings copyandpaste.LinterSetting
	}{
		{
			desc:     "repeatoptions",
			settings: copyandpaste.LinterSetting{},
		},
		{
			desc:     "repeatargs",
			settings: copyandpaste.LinterSetting{},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			a, err := copyandpaste.NewAnalyzer(test.settings)
			if err != nil {
				t.Fatal(err)
			}

			analysistest.Run(t, analysistest.TestData(), a, test.desc)
		})
	}
}
