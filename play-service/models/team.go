package models

type Team struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	League string `json:"league" db:"league"`

	Stats Stats `json:"stats" db:"-"`
}

type NFLTeam struct {
	Team
}

func (nfl *NFLTeam) CalculateFpts() float64 {
	stats := nfl.Stats.(*NFLStats)

	fpts := stats.FieldGoalsMade * 3
	fpts += stats.Assists * 2
	fpts += stats.Fumbles

	return float64(fpts)
}

type NBATeam struct {
	Team
}

func (nba *NBATeam) CalculateFpts() float64 {
	stats := nba.Stats.(*NBAStats)

	fpts := stats.ThreePointsMade * 3
	fpts += stats.TwoPointsMade * 2
	fpts += stats.FreeThrowsMade

	return float64(fpts)
}
