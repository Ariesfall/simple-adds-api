package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ariesfall/simple-odds-api/pkg/conn"
	"github.com/Ariesfall/simple-odds-api/pkg/odds"
	"github.com/Ariesfall/simple-odds-api/pkg/routine"
	"github.com/Ariesfall/simple-odds-api/pkg/serv"
	"github.com/spf13/cobra"
)

var (
	// cmd input license key
	apiLicense string

	pgAddr string = "127.0.0.1"
	pgPort string = "5432"
	pgUser string = "postgres"
	pgPass string

	rootCmd = &cobra.Command{
		Use:   "sbet",
		Short: "sbet collect the simple sport betting data from odds api",
		RunE:  OddsSrv,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&apiLicense, "license", "l", "", "key of license for the odds api")
	rootCmd.PersistentFlags().StringVarP(&pgAddr, "addr", "a", "127.0.0.1", "ip adds for the database")
	rootCmd.PersistentFlags().StringVarP(&pgPort, "port", "p", "5432", "ip port for the database")
	rootCmd.PersistentFlags().StringVarP(&pgUser, "user", "u", "postgres", "user for the database access")
	rootCmd.PersistentFlags().StringVarP(&pgPass, "pass", "P", "", "user password for the database access")

	rootCmd.MarkPersistentFlagRequired("license")

}

func initService() {
	odds.ApiKey = apiLicense
	log.Println("[StartUp] init odds api key: " + odds.ApiKey)
}

func OddsSrv(cmd *cobra.Command, args []string) error {
	initService()

	// connect pgsql
	dsn := conn.MakeDsn(pgAddr, pgPort, pgUser, pgPass, "sbet")
	db, err := conn.ConnectPgdb(&conn.Postgres{Dsn: dsn, ConnMaxLiftTime: 10, MaxOpenConns: 10, MaxIdleConns: 1})
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("[StartUp] connected to postgres database sbet")

	// store sports and upcoming data on start-up
	err = routine.InitSync()
	if err != nil {
		return err
	}

	// routine service in new thread
	job := routine.Jobs()
	go job.Start()
	log.Println("[StartUp] start rouine job")

	defer job.Stop()

	// http service
	r := serv.Service(db)

	// port listen
	log.Println("[StartUp] start service on port :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}

	return nil
}
