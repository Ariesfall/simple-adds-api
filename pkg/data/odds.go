package data

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Match is the detail of match odds data
type Match struct {
	SportKey     string `json:"sport_key" db:"sport_key"`
	SportNice    string `json:"sport_nice" db:"sport_nice"`
	CommenceTime int    `json:"commence_time" db:"commence_time"`
	HomeTeam     string `json:"home_team" db:"home_team"`
	TeamA        string `json:"team_a" db:"team_a"`
	TeamB        string `json:"team_b" db:"team_b"`
	SitesCnt     int    `json:"sites_count" db:"sites_count"`
}

// OddsData is the detail of odds data, match_key is a key combine from `HomeTeam` and `CommenceTime`
type Odds struct {
	MatchKey   string  `json:"match_key" db:"match_key"`
	SiteKey    string  `json:"site_key" db:"site_key"`
	SiteNice   string  `json:"site_nice" db:"site_nice"`
	LastUpdate int     `json:"last_update" db:"last_update"`
	OddTeamA   float32 `json:"odds_team_a" db:"odds_team_a"`
	OddTeamB   float32 `json:"odds_team_b" db:"odds_team_b"`
	OddDraw    float32 `json:"odds_draw" db:"odds_draw"`
}

func GetMatch(db *sqlx.DB, in *Match) (*Match, error) {
	q := "SELECT sport_key, sport_nice, commence_time, home_team, team_a, team_b, sites_count FROM matchs " +
		"WHERE SportKey = ? AND  commence_time=? AND home_team=?"

	res := &Match{}
	err := db.Get(res, q, in.SportKey, in.CommenceTime, in.HomeTeam)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ListMatch(db *sqlx.DB, in *Match) ([]*Match, error) {
	now := time.Now().Unix()
	q := "SELECT sport_key, sport_nice, commence_time, home_team, team_a, team_b, sites_count FROM matchs " +
		"WHERE commence_time>?"

	res := []*Match{}
	err := db.Get(&res, q, in.SportKey, now)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateMatch(db *sqlx.DB, in *Match) error {
	q := "INSERT INTO matchs (sport_key, sport_nice, commence_time, home_team, team_a, team_b, sites_count)" +
		" VALUE (:sport_key, :sport_nice, :commence_time, :home_team, :team_a, :team_b, :sites_count)"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

func UpdateMatch(db *sqlx.DB, in *Match) error {

	q := "UPDATE matchs SET sites_count=:sites_count WHERE " +
		"sport_key=:sport_key AND commence_time=:commence_time AND home_team=:home_team"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

func GetOdds(db *sqlx.DB, in *Odds) (*Odds, error) {
	q := "SELECT match_key, site_key, site_nice, odds_team_a, odds_team_b, odds_draw FROM odds WHERE match_key=? AND site_key=?"

	res := &Odds{}
	err := db.Get(res, q, in.MatchKey, in.SiteKey)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateOdds(db *sqlx.DB, in *Odds) error {
	q := "INSERT INTO odds (match_key, site_key, site_nice, odds_team_a, odds_team_b, odds_draw)" +
		" VALUE (:match_key, :site_key, :site_nice, :odds_team_a, :odds_team_b, :odds_draw)"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOdds(db *sqlx.DB, in *Odds) error {
	q := "UPDATE odds SET odds_team_a=:odds_team_a, odds_team_b=Lodds_team_b, odds_draw=:odds_draw" +
		"WHERE match_key=:match_key AND site_key=:site_key"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}
