package models

type Athlete struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	League string `json:"league" db:"league"`

	Stats Stats `json:"stats" db:"-"`
}

type NFLAthlete struct {
	Athlete
}

func (a Athlete) GetID() int {
	return a.ID
}
func (a Athlete) GetLeague() string {
	return a.League
}
func (a Athlete) GetEntityType() string {
	return a.Stats.GetEntityType()
}

// func (a Athlete) SetStats(Stats) {
// 	return a.
// }

func (nfl *NFLAthlete) CalculateFpts() float64 {
	stats := nfl.Stats.(*NFLStats)

	fpts := stats.FieldGoalsMade * 3
	fpts += stats.Assists * 2
	fpts += stats.Fumbles

	return float64(fpts)
}

type NBAAthlete struct {
	Athlete
}

func (nba *NBAAthlete) CalculateFpts() float64 {
	stats := nba.Stats.(*NBAStats)

	fpts := stats.ThreePointsMade * 3
	fpts += stats.TwoPointsMade * 2
	fpts += stats.FreeThrowsMade

	return float64(fpts)
}
