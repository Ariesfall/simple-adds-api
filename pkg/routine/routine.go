package routine

import (
	"log"

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
	log.Println("sync all match")
	SyncOdds("upcoming")
}
