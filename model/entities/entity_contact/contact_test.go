package entity_contact

import (
	"fmt"
	"testing"
)

func TestNewContact(t *testing.T) {

	t.Run("condição ok com todos os dados", func(t *testing.T) {
		send := []byte(`{"name":"Full Name A","nickName":"nickName A","notes":"notes A"}`)
		expected1 := contact{Id: 0, Name: "Full Name A", NickName: "nickName A", Notes: "notes A"}
		var expected2 error

		result1, result2 := NewContact(send)

		if !(result1 == expected1 && result2 == expected2) {
			t.Errorf("\nExpected1= %v\nResult1= %v\n\nExpected2= %v\nResult2= %v \n",
				expected1, result1, expected2, result2)
		}
	})

	t.Run("condição ok sem todos os dados", func(t *testing.T) {
		send := []byte(`{"name":"Full Name A","notes":"notes A"}`)
		expected1 := contact{Id: 0, Name: "Full Name A", NickName: "", Notes: "notes A"}
		var expected2 error

		result1, result2 := NewContact(send)

		if !(result1 == expected1 && result2 == expected2) {
			t.Errorf("\nExpected1= %v\nResult1= %v\n\nExpected2= %v\nResult2= %v \n",
				expected1, result1, expected2, result2)
		}
	})

	t.Run("erro com JSON inválido", func(t *testing.T) {
		send := []byte(`xx{"name":"Full Name A","nickName":"nickName A","notes":"notes A"}`)
		expected1 := contact{Id: 0, Name: "", NickName: "", Notes: ""}
		expected2 := fmt.Errorf("invalid character 'x' looking for beginning of value")

		result1, result2 := NewContact(send)

		if !(result1 == expected1 && result2.Error() == expected2.Error()) {
			t.Errorf("\nExpected1= %v\nResult1= %v\n\nExpected2= %v\nResult2= %v \n",
				expected1, result1, expected2, result2)
		}
	})

}
