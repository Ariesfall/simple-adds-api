package data

import (
	"github.com/jmoiron/sqlx"
)

// Sports is the detail of sports data
type Sports struct {
	ID     int    `db:"id"`
	Key    string `json:"key" db:"sport_key"`
	Active bool   `json:"active" db:"active"`
	Group  string `json:"group" db:"sport_group"`
	Detail string `json:"detail" db:"detail"`
	Title  string `json:"title" db:"title"`
}

func GetSports(db *sqlx.DB, key string) (*Sports, error) {
	q := "SELECT sport_key, active, sport_group, detail, title FROM sports WHERE sport_key = ?"

	res := &Sports{}
	err := db.Get(res, q, key)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ListSports(db *sqlx.DB) ([]*Sports, error) {
	q := "SELECT sport_key, active, sport_group, detail, title FROM sports WHERE active=true"

	res := []*Sports{}
	err := db.Select(&res, q)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateSports(db *sqlx.DB, in *Sports) error {
	q := "INSERT INTO sports (sport_key, active, sport_group, detail, title) VALUE (:sport_key, :active, :sport_group, :detail, :title)"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSports(db *sqlx.DB, in *Sports) error {

	q := "UPDATE sports SET active=:active, sport_group=:sport_group, detail=:detail, title=:title WHERE sport_key=:sport_key"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}
