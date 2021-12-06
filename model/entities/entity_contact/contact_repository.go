package entity_contact

import "github.com/jairhdev/go-api-contact/external/database"

// ****** SAVE database
// ******
func (data contact) saveRepository() (int, error) {
	var db = database.NewService(database.NewDatabase())

	// START TRANSACTION
	tx, err := db.GetConn().Begin(db.GetCtx())
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(db.GetCtx())

	// RUN
	const query string = "INSERT INTO tb_contact (name, nickname, notes) VALUES ($1, $2, $3) RETURNING id;"

	rows, err := tx.Query(db.GetCtx(), query, data.Name, data.NickName, data.Notes)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	// END TRANSACTION
	if err := tx.Commit(db.GetCtx()); err != nil {
		return 0, err
	}
	return id, err
}

// ****** FIND ALL database
// ******
func (data contact) findAllRepository() ([]contact, error) {
	var db = database.NewService(database.NewDatabase())

	const query string = "SELECT id, name, nickname, notes FROM tb_contact;"

	rows, err := db.GetConn().Query(db.GetCtx(), query)
	if err != nil {
		return []contact{}, err
	}
	defer rows.Close()

	var result contact
	var results []contact

	for rows.Next() {
		if err := rows.Scan(
			&result.Id,
			&result.Name,
			&result.NickName,
			&result.Notes,
		); err != nil {
			return []contact{}, err
		}
		results = append(results, result)
	}
	return results, nil
}

// ****** FIND BY ID database
// ******
func (data contact) findByIdRepository(id int) (contact, error) {
	var db = database.NewService(database.NewDatabase())

	const query string = "SELECT id, name, nickname, notes FROM tb_contact WHERE id=$1;"

	rows, err := db.GetConn().Query(db.GetCtx(), query, id)
	if err != nil {
		return contact{}, err
	}
	defer rows.Close()

	var result contact
	for rows.Next() {
		if err := rows.Scan(
			&result.Id,
			&result.Name,
			&result.NickName,
			&result.Notes,
		); err != nil {
			return contact{}, err
		}
	}
	return result, nil
}

// ****** UPDATE BY ID database
// ******
func (data contact) updateByIdRepository(id int) (int64, error) {
	var db = database.NewService(database.NewDatabase())

	// START TRANSACTION
	tx, err := db.GetConn().Begin(db.GetCtx())
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(db.GetCtx())

	// RUN
	const query string = "UPDATE tb_contact SET name=$1, nickname=$2, notes=$3 WHERE id=$4;"

	rows, err := tx.Exec(db.GetCtx(), query, data.Name, data.NickName, data.Notes, id)
	if err != nil {
		return 0, err
	}

	// END TRANSACTION
	if err := tx.Commit(db.GetCtx()); err != nil {
		return 0, err
	}
	return rows.RowsAffected(), err
}

// ****** DELETE BY ID database
// ******
func (data contact) deleteByIdRepository(id int) (int64, error) {
	var db = database.NewService(database.NewDatabase())

	// START TRANSACTION
	tx, err := db.GetConn().Begin(db.GetCtx())
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(db.GetCtx())

	// RUN
	const query string = "DELETE FROM tb_contact WHERE id=$1;"

	rows, err := tx.Exec(db.GetCtx(), query, id)
	if err != nil {
		return 0, err
	}

	// END TRANSACTION
	if err := tx.Commit(db.GetCtx()); err != nil {
		return 0, err
	}
	return rows.RowsAffected(), err
}
