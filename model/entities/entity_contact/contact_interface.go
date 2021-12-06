package entity_contact

type impl interface {
	Save() (int, error)
	FindAll() ([]contact, error)
	FindById(id int) (contact, error)
	UpdateById(id int) (int64, error)
	DeleteById(id int) (int64, error)
}

type service struct {
	impl
}

func NewService(contact impl) service {
	return service{contact}
}
