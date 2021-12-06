package entity_contact

func (data contact) Save() (int, error) {
	return data.saveRepository()
}

func (data contact) FindAll() ([]contact, error) {
	return data.findAllRepository()
}

func (data contact) FindById(id int) (contact, error) {
	return data.findByIdRepository(id)
}

func (data contact) UpdateById(id int) (int64, error) {
	return data.updateByIdRepository(id)
}

func (data contact) DeleteById(id int) (int64, error) {
	return data.deleteByIdRepository(id)
}
