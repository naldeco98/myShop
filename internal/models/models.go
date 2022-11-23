package models

import (
	"context"
	"database/sql"
	"time"
)

// DBModel is the type for database connection values
type DBModel struct {
	DB *sql.DB
}

type Models struct {
	DB DBModel
}

// NewModels returns a new model type with database connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{db},
	}
}

// Widget type for all widgets
type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

func (m *DBModel) GetWidget(id int) (*Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var widget Widget

	row := m.DB.QueryRowContext(ctx, "SELECT id, name, description, inventory_level, price FROM widgets WHERE id = ?", id)
	if err := row.Scan(&widget.ID, &widget.Name, &widget.Description, &widget.InventoryLevel, &widget.Price); err != nil {
		return nil, err
	}

	return &widget, nil
}
