package config

import "testing"

func TestNewConfig(t *testing.T) {
	t.Run("testa se trata 'environment' correto", func(t *testing.T) {
		err := NewConfig(configValueDev)

		if err != nil {
			t.Errorf("\nExpected= %v | Result= %v", nil, err)
		}
	})

	t.Run("testa se trata 'environment' incorreto", func(t *testing.T) {
		err := NewConfig("xx")

		expected := "invalid environment"
		if err.Error() != expected {
			t.Errorf("\nExpected= %v | Result= %v", expected, err)
		}
	})
}
