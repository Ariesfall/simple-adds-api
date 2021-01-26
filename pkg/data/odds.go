package data

import (
	"fmt"
	"strconv"
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

type MatchOdds struct {
	*Match
	Sites []*Odds `json:"sites"`
}

func ListMatch(db *sqlx.DB, in *Match) ([]*MatchOdds, error) {
	now := int(time.Now().Unix())
	q := fmt.Sprintf("SELECT sport_key, sport_nice, commence_time, home_team, team_a, team_b, sites_count FROM matchs "+
		"WHERE commence_time>%d", now)

	if in.SportKey != "" {
		q += fmt.Sprintf(" AND sport_key='%s'", in.SportKey)
	}

	matchs := []*Match{}
	err := db.Select(&matchs, q+" ORDER BY commence_time")
	if err != nil {
		return nil, err
	}

	res := []*MatchOdds{}
	for _, m := range matchs {
		matchKey := m.HomeTeam + strconv.Itoa(m.CommenceTime)
		q = fmt.Sprintf("SELECT match_key, site_key, site_nice, last_update, odds_team_a, odds_team_b, odds_draw FROM odds WHERE match_key='%s'", matchKey)

		sites := []*Odds{}

		err = db.Select(&sites, q)
		if err != nil {
			return nil, err
		}

		res = append(res, &MatchOdds{
			Match: m,
			Sites: sites,
		})
	}
	return res, nil
}

func GetMatch(db *sqlx.DB, in *Match) (*Match, error) {
	q := fmt.Sprintf("SELECT sport_key, sport_nice, commence_time, home_team, team_a, team_b, sites_count FROM matchs "+
		"WHERE SportKey='%s' AND commence_time=%d AND home_team='%s'", in.SportKey, in.CommenceTime, in.HomeTeam)

	res := &Match{}
	err := db.Get(res, q)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateMatch(db *sqlx.DB, in *Match) error {
	q := "INSERT INTO matchs (sport_key, sport_nice, commence_time, home_team, team_a, team_b, sites_count)" +
		" VALUES (:sport_key, :sport_nice, :commence_time, :home_team, :team_a, :team_b, :sites_count)"

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
	q := fmt.Sprintf("SELECT match_key, site_key, site_nice, last_update, odds_team_a, odds_team_b, odds_draw FROM odds"+
		"WHERE match_key='%s' AND site_key='%s'", in.MatchKey, in.SiteKey)

	res := &Odds{}
	err := db.Get(res, q)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func CreateOdds(db *sqlx.DB, in *Odds) error {
	q := "INSERT INTO odds (match_key, site_key, site_nice, last_update, odds_team_a, odds_team_b, odds_draw)" +
		" VALUES (:match_key, :site_key, :site_nice, :last_update, :odds_team_a, :odds_team_b, :odds_draw)"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOdds(db *sqlx.DB, in *Odds) error {
	q := "UPDATE odds SET last_update=:last_update, odds_team_a=:odds_team_a, odds_team_b=Lodds_team_b, odds_draw=:odds_draw" +
		"WHERE match_key=:match_key AND site_key=:site_key"

	_, err := db.NamedExec(q, in)
	if err != nil {
		return err
	}

	return nil
}
