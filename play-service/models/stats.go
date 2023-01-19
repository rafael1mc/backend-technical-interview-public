package models

type Stats interface {
	GetID() int
	GetEntityID() string
	GetEntityType() string
}

type BaseStats struct {
	ID         int    `json:"id" db:"id"`
	EntityID   string `json:"entityId" db:"entity_id"`
	EntityType string `json:"entityType" db:"entity_type"`
}

func (bs *BaseStats) GetID() int {
	return bs.ID
}

func (bs *BaseStats) GetEntityID() string {
	return bs.EntityID
}

func (bs *BaseStats) GetEntityType() string {
	return bs.EntityType
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
