package models

type Stats interface {
	GetID() string
	GetEntityID() int
	GetEntityType() string
}

type BaseStats struct {
	ID         string `json:"id" db:"id"`
	EntityID   int    `json:"entityId" db:"entity_id"`
	EntityType string `json:"entityType" db:"entity_type"`
}

func (s BaseStats) GetID() string {
	return s.ID
}
func (s BaseStats) GetEntityID() int {
	return s.EntityID
}
func (s BaseStats) GetEntityType() string {
	return s.EntityType
}

type NFLStats struct {
	BaseStats
	FieldGoalsMade int `json:"fieldGoalsMade" db:"field_goals_made"`
	Assists        int `json:"assists" db:"assists"`
	Fumbles        int `json:"fumbles" db:"fumbles"`
}

type NBAStats struct {
	BaseStats
	ThreePointsMade int `json:"threePointsMade" db:"three_points_made"`
	TwoPointsMade   int `json:"twoPointsMade" db:"two_points_made"`
	FreeThrowsMade  int `json:"freeThrowsMade" db:"free_throws_made"`
}
