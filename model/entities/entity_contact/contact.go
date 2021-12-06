package entity_contact

import "encoding/json"

type contact struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Notes    string `json:"notes"`
}

// Cria novo objeto 'contact' tendo como origem body de uma requisição.
func NewContact(body []byte) (contact, error) {

	var data contact

	if err := json.Unmarshal(body, &data); err != nil {
		return contact{}, err
	}
	return data, nil
}

func NewContactEmpty() contact {
	return contact{}
}
