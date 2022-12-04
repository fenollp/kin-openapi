package openapi3

import (
	"testing"
	// fuzz "github.com/AdaLogics/go-fuzz-headers"
)

func FuzzLoadFromData(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		sl := NewLoader()

		doc, err := sl.LoadFromData(data)
		if err != nil {
			return // Try harder at generating YAML
		}

		_ = doc.Validate(sl.Context)
	})
}
