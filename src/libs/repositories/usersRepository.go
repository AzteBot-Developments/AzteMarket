package repositories

import (
	"fmt"

	"github.com/RazvanBerbece/AzteMarket/src/libs/models/dax"
)

type DbUserRepository interface {
	GetUser(userId string) (*dax.User, error)
}

type UserRepository struct {
	DbContext AztebotDbContext
}

func NewUserRepository(connString string) UserRepository {
	repo := UserRepository{AztebotDbContext{
		ConnectionString: connString,
	}}
	repo.DbContext.Connect()
	return repo
}

func (r UserRepository) GetUser(userId string) (*dax.User, error) {

	query := "SELECT * FROM Users WHERE userId = ?"
	row := r.DbContext.SqlDb.QueryRow(query, userId)

	var user dax.User
	err := row.Scan(&user.Id,
		&user.DiscordTag,
		&user.UserId,
		&user.CurrentRoleIds,
		&user.CurrentCircle,
		&user.CurrentInnerOrder,
		&user.CurrentLevel,
		&user.CurrentExperience,
		&user.CreatedAt,
		&user.Gender)

	if err != nil {
		return nil, fmt.Errorf("an error ocurred while retrieving user with ID `%s`", userId)
	}

	return &user, nil

}
