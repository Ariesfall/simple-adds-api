package routine

import (
	"log"

	"github.com/robfig/cron"
)

func RoutineJob() *cron.Cron {
	c := cron.New()

	// per one mins
	c.AddFunc("0 */1 * * * *", inplayMatch)
	// hourly
	c.AddFunc("0 0 * * * *", nonInPlayMatch)

	return c
}

func Start(c *cron.Cron) {
	log.Printf("start run routine")
	c.Start()
}

func Stop(c *cron.Cron) {
	c.Stop()
}

// per one mins
func inplayMatch() {
}

// hourly
func nonInPlayMatch() {

}
