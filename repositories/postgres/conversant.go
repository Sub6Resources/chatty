package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/ryan-berger/chatty/repositories"
)

const updateOrCreateConversant = `
INSERT INTO conversant (id, display_name) VALUES (:id, :display_name) 
ON CONFLICT (id) DO 
  UPDATE SET display_name = :display_name
`

type ConversantRepository struct {
	db *sqlx.DB
}

func NewConversantRepository(db *sqlx.DB) *ConversantRepository {
	return &ConversantRepository{
		db: db,
	}
}

func (repo *ConversantRepository) UpdateOrCreate(conversant repositories.Conversant) (*repositories.Conversant, error) {
	_, err := repo.db.NamedExec(updateOrCreateConversant, &conversant)

	if err != nil {
		return nil, err
	}

	return &conversant, nil
}
