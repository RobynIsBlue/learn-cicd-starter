package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input        http.Header
		output       string
		err          error
		errSubString string
	}{
		"simple":    {input: http.Header{"Authorization": []string{"ApiKey heheheHI!!!"}}, output: "heheheHI!!!", err: nil},
		"no header": {input: http.Header{"haha": []string{"ApiKey", "111"}}, output: "", err: ErrNoAuthHeaderIncluded},
		"empty API": {input: http.Header{"Authorization": []string{"ApiKey"}}, output: "", errSubString: "malformed"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(test.input)
			assert.Empty(t, cmp.Diff(test.output, got))
			if test.errSubString == "" {
				require.True(t, errors.Is(test.err, err))
			} else {
				require.NotContains(t, test.errSubString, err)
			}
		})
	}
}
