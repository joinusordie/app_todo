package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) DeleteUser(userId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	DeleteItemsQuery := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
							WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id=$1`,
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err = r.db.Exec(DeleteItemsQuery, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	DeleteListsQuery := fmt.Sprintf(`DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1`,
		todoListTable, usersListsTable)
	_, err = r.db.Exec(DeleteListsQuery, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	DeleteUsersQuery := fmt.Sprintf("DELETE FROM %s u WHERE u.id = $1", userTable)
	_, err = r.db.Exec(DeleteUsersQuery, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
