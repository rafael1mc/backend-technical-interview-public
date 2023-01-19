package models

type Entity interface {
	GetID() int
	GetLeague() string
	SetStats(Stats)

	MarshalJSON() ([]byte, error)
}
