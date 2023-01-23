package models

type Team struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	League string `json:"league" db:"league"`

	Stats Stats `json:"stats" db:"-"`
}

func (t Team) GetID() int {
	return t.ID
}
func (t Team) GetLeague() string {
	return t.League
}
func (t *Team) SetStats(s Stats) {
	t.Stats = s
}
