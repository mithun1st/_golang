package repository

import (
	"fmt"
	"gin-basic/internal/database"
	model "gin-basic/internal/models"
	"gin-basic/internal/schema"
)

type AuthRepository struct {
	db *database.DB
}

func NewRepository(db *database.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindByEmail(email string) (*model.UserModel, error) {

	var sql string = fmt.Sprintf(`
select * from %s where
%s='%s'
limit 1
`, schema.SuperUserTable, schema.UserEmail, email)

	rows, err := r.db.SqlDB.Query(sql)
	if err != nil {
		return nil, err
	}

	var user model.UserModel
	for rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Password,
			&user.Email,
			&user.Phone,
		)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
