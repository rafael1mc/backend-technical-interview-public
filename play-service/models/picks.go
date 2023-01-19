package models

type Roster []Pick

type Pick struct {
	ID         int     `json:"id" db:"id"`
	UserID     int     `json:"userId" db:"user_id"`
	EntityID   int     `json:"-" db:"entity_id"`
	EntityType string  `json:"entityType" db:"entity_type"`
	Fpts       float64 `json:"fpts" db:"-"`

	Entity Entity `json:"entity" db:"-"`
}
