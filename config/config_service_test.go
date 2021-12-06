package config

import (
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	t.Run("testa se retorna configuração 'common'", func(t *testing.T) {
		NewConfig(configValueDev)

		const expected = "common"
		result := os.Getenv("TEST_UNIT_COMMON")

		if expected != result {
			t.Errorf("\nExpected= %v | Result= %v", expected, result)
		}
	})

	t.Run("testa se retorna configuração 'dev'", func(t *testing.T) {
		NewConfig(configValueDev)

		const expected = configValueDev
		result := os.Getenv("TEST_UNIT_DEV")

		if expected != result {
			t.Errorf("\nExpected= %v | Result= %v", expected, result)
		}
	})

	t.Run("testa se retorna configuração 'test'", func(t *testing.T) {
		NewConfig(configValueTest)

		const expected = configValueTest
		result := os.Getenv("TEST_UNIT_TEST")

		if expected != result {
			t.Errorf("\nExpected= %v | Result= %v", expected, result)
		}
	})

	t.Run("testa se retorna configuração 'docker'", func(t *testing.T) {
		NewConfig(configValueDocker)

		const expected = configValueDocker
		result := os.Getenv("TEST_UNIT_DOCKER")

		if expected != result {
			t.Errorf("\nExpected= %v | Result= %v", expected, result)
		}
	})

	t.Run("testa se retorna configuração 'prod'", func(t *testing.T) {
		NewConfig(configValueProd)

		const expected = configValueProd
		result := os.Getenv("TEST_UNIT_PROD")

		if expected != result {
			t.Errorf("\nExpected= %v | Result= %v", expected, result)
		}
	})
}
