package main

import "testing"

func TestReadEnvironment(t *testing.T) {
	t.Run("testa se arquivo existe e se consegue ler (arq existente)", func(t *testing.T) {
		const fileEnv string = "environment.properties"
		_, err := readEnvironment(fileEnv)

		if err != nil {
			t.Errorf("\nExpected= %v | Result= %v", "nil", err)
		}
	})

	t.Run("testa se arquivo existe (arq NÃO existente)", func(t *testing.T) {
		const fileEnv string = "XXXenvironment.properties"
		_, err := readEnvironment(fileEnv)

		expected := "stat XXXenvironment.properties: no such file or directory"

		if err.Error() != expected {
			t.Errorf("\nExpected= %v\nResult= %v", expected, err)
		}
	})

	t.Run("testa leitura ok do arquivo", func(t *testing.T) {
		const fileEnv string = "environment.properties"
		result, _ := readEnvironment(fileEnv)

		if result == "" {
			t.Errorf("\nExpected= %v | Result= %v", "dev, test, prod", result)
		}
	})
}
