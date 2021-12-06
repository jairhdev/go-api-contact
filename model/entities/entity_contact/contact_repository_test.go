package entity_contact

import (
	"testing"

	"github.com/jairhdev/go-api-contact/config"
	"github.com/jairhdev/go-api-contact/external/database"
)

/*
** SEQUÊNCIA DE TESTES CRUD
- insere 1 registro
- pesquisa todos os registros. Valida se o registro inserido existe
- pesquisa registro inserido por id
- altera registro inserido. Valida com pesquisa ao registro e comparação se dados foram alterados
- apaga registro inserido. Valida com pesquisa de todos os registros e qtde encontrada (igual 0)
*/

func TestRepository(t *testing.T) {

	//****************************
	// Start config TEST
	const env string = "test"
	if err := config.NewConfig(env); err != nil {
		panic(err)
	}

	// Start database
	var db = database.NewService(database.NewDatabase())
	if err := db.NewConnect(); err != nil {
		panic(err)
	}
	//****************************

	// registro para inserção
	testData := contact{
		Id:       0,
		Name:     "Full Name A",
		NickName: "NickName A",
		Notes:    "Notes A",
	}

	t.Run("INSERT - valida id > 0", func(t *testing.T) {
		id, err := testData.saveRepository()
		if err != nil {
			t.Errorf("\nExpectedErr= %v\nResultErr= %v", "nil", err)
		}
		testData.Id = id

		if testData.Id == 0 {
			t.Errorf("\nExpectedId= %v\nResultId= %v", "valor maior que zero", testData.Id)
		}
	})

	t.Run("FIND ALL - valida registro correto e qtde de registro >= 1", func(t *testing.T) {
		results, err := testData.findAllRepository()
		if err != nil {
			t.Errorf("\nExpectedErr= %v\nResultErr= %v", "nil", err)
		}

		if len(results) < 1 {
			t.Errorf("\nExpected= %v\nResult= %v", 1, len(results))
		}

		for _, result := range results {
			ok := (result.Name == testData.Name &&
				result.NickName == testData.NickName &&
				result.Notes == testData.Notes)
			if !ok {
				t.Errorf("\nExpected= %v\nResult= %v", testData, result)
			}
		}
	})

	t.Run("FIND BY ID - valida registro correto", func(t *testing.T) {
		result, err := testData.findByIdRepository(testData.Id)
		if err != nil {
			t.Errorf("\nExpectedErr= %v\nResultErr= %v", "nil", err)
		}

		ok := (result.Id == testData.Id &&
			result.Name == testData.Name &&
			result.NickName == testData.NickName &&
			result.Notes == testData.Notes)
		if !ok {
			t.Errorf("\nExpected= %v\nResult= %v", testData, result)
		}
	})

	// registro para alteração
	testDataUpdate := contact{
		Id:       testData.Id,
		Name:     "XFull Name A",
		NickName: "XNickName A",
		Notes:    "XNotes A",
	}

	t.Run("UPDATE BY ID - altera registro e valida update em rows com findById", func(t *testing.T) {
		rows, err := testDataUpdate.updateByIdRepository(testData.Id)
		if err != nil {
			t.Errorf("\nExpectedErr= %v\nResultErr= %v", "nil", err)
		}

		if rows != 1 {
			t.Errorf("\nExpected= %v\nResult= %v", 1, rows)
		}

		resultUpdate, _ := testData.findByIdRepository(testData.Id)

		ok := (resultUpdate.Id == testDataUpdate.Id &&
			resultUpdate.Name == testDataUpdate.Name &&
			resultUpdate.NickName == testDataUpdate.NickName &&
			resultUpdate.Notes == testDataUpdate.Notes)
		if !ok {
			t.Errorf("\nExpected= %v\nResult= %v", testDataUpdate, resultUpdate)
		}
	})

	t.Run("DELETE BY ID - deleta registro e valida com qtde de findAll", func(t *testing.T) {
		rows, err := testData.deleteByIdRepository(testData.Id)
		if err != nil {
			t.Errorf("\nExpectedErr= %v\nResultErr= %v", "nil", err)
		}

		if rows != 1 {
			t.Errorf("\nExpected= %v\nResult= %v", 1, rows)
		}

		resultAfterDelete, _ := testData.findAllRepository()

		ok := len(resultAfterDelete) == 0
		if !ok {
			t.Errorf("\nExpected= %v\nResult= %v", 0, len(resultAfterDelete))
		}
	})
}
