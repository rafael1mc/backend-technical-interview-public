package models

type Athlete struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	League string `json:"league" db:"league"`

	Stats Stats `json:"stats" db:"-"`
}

func (a Athlete) GetID() int {
	return a.ID
}
func (a Athlete) GetLeague() string {
	return a.League
}
func (a *Athlete) SetStats(s Stats) {
	a.Stats = s
}
