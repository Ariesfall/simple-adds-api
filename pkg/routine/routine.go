package routine

import (
	"log"

	"github.com/Ariesfall/simple-odds-api/pkg/conn"
	"github.com/Ariesfall/simple-odds-api/pkg/data"
	"github.com/robfig/cron"
)

func Jobs() *cron.Cron {
	c := cron.New()

	// per one mins
	c.AddFunc("0 */1 * * * *", inplayMatch)
	// per one hours
	c.AddFunc("0 0 * * * *", nonInPlayMatch)

	return c
}

// per one mins
func inplayMatch() {
	log.Println("sync inplay match")
	SyncOdds("upcoming")
}

// hourly
func nonInPlayMatch() {
	log.Println("sync sports match")

	db := conn.GetPgDB()

	sportList, err := data.ListSports(db)
	if err != nil {
		log.Fatalln("Error: list sports" + err.Error())
		return
	}

	for _, sport := range sportList {
		err = SyncOdds(sport.Key)
		if err != nil {
			log.Fatalf("Error: sync sport %s, msg %s\n", sport.Key, err.Error())
			return
		}
	}
}
