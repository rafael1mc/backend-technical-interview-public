package models

type Entity interface {
	GetID() int
	GetLeague() string
	GetEntityType() string
	SetStats(Stats)

	CalculateFpts() float64
}
