package routine

import (
	"errors"
	"log"
	"strconv"

	"github.com/Ariesfall/simple-odds-api/pkg/conn"
	"github.com/Ariesfall/simple-odds-api/pkg/data"
	"github.com/Ariesfall/simple-odds-api/pkg/odds"
)

// InitSync store all available sports and all upcoming fixtures in db on Startup
func InitSync() (err error) {
	log.Println("[StartUp] store all available sports")
	err = SyncSport()
	if err != nil {
		return err
	}

	log.Println("[StartUp] store all upcoming fixtures")
	err = SyncOdds("upcoming")
	if err != nil {
		return err
	}

	return nil
}

// SyncSport make the request to odds, get result and update db
func SyncSport() error {
	res, err := odds.GetSports()
	if err != nil {
		log.Println("Error: request error" + err.Error())
		return err
	}

	err = UpsertSport(res)
	if err != nil {
		log.Println("Error: upsert error" + err.Error())
		return err
	}

	return nil
}

// SyncOdds make the request to odds, get result and update db
// sports is sport key or `upcoming`
func SyncOdds(sports string) error {
	res, err := odds.GetOdds(sports)
	if err != nil {
		log.Println("Error: request error" + err.Error())
		return err
	}

	err = UpsertMatch(res)
	if err != nil {
		log.Println("Error: upsert error" + err.Error())
		return err
	}

	return nil
}

// UpsertSport update or insert sport data
func UpsertSport(in *odds.SportsResp) error {
	if !in.Success {
		return errors.New("API request not success")
	}

	db := conn.GetPgDB()

	for _, s := range in.Data {
		// check exsit or not, exsit > update, not exsit > Create
		_, err := data.GetSports(db, s.Key)
		if err == nil {
			err = data.UpdateSports(db, s)
		} else {
			err = data.CreateSports(db, s)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// UpsertMatch update or insert match and odds data
func UpsertMatch(in *odds.OddsResp) error {
	if !in.Success {
		return errors.New("API request not success")
	}

	db := conn.GetPgDB()

	for _, m := range in.Data {
		if len(m.Teams) != 2 {
			log.Printf("Error: odds result:%+v\n", m)
			continue
		}

		in := &data.Match{
			SportKey:     m.SportKey,
			SportNice:    m.SportNice,
			CommenceTime: m.CommenceTime,
			HomeTeam:     m.HomeTeam,
			TeamA:        m.Teams[0],
			TeamB:        m.Teams[1],
			SitesCnt:     m.SitesCnt,
		}

		// check exsit or not, exsit > update, not exsit > Create
		_, err := data.GetMatch(db, in)
		if err == nil {
			err = data.UpdateMatch(db, in)
		} else {
			err = data.CreateMatch(db, in)
		}
		if err != nil {
			return err
		}

		// update or insert site odds data
		matchKey := m.HomeTeam + strconv.Itoa(m.CommenceTime)
		err = UpsertOdds(matchKey, m.Sites)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpsertOdds update or insert odds data
func UpsertOdds(matchkey string, sites []*odds.Site) error {
	db := conn.GetPgDB()
	for _, s := range sites {
		if len(s.Odds.H2h) != 3 {
			log.Printf("Error: odds h2h result:%+v\n", s)
			continue
		}

		in := &data.Odds{
			MatchKey:   matchkey,
			SiteKey:    s.SiteKey,
			SiteNice:   s.SiteNice,
			LastUpdate: s.LastUpdate,
			OddTeamA:   s.Odds.H2h[0],
			OddTeamB:   s.Odds.H2h[1],
			OddDraw:    s.Odds.H2h[2],
		}

		// check exsit or not, exsit and not last > update, not exsit > Create
		c, err := data.GetOdds(db, in)
		if err == nil && s.LastUpdate != c.LastUpdate {
			err = data.UpdateOdds(db, in)
		} else if err != nil {
			err = data.CreateOdds(db, in)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
