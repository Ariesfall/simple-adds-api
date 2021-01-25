package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ariesfall/simple-odds-api/pkg/conn"
	"github.com/Ariesfall/simple-odds-api/pkg/serv"
	"github.com/spf13/cobra"
)

var (
	// cmd input license key
	userLicense string

	pgAddr string = "127.0.0.1"
	pgPort string = "5432"
	pgUser string = "postgres"
	pgPass string

	rootCmd = &cobra.Command{
		Use:   "sbet",
		Short: "sbet collect the simple sport betting data from odds api",
		RunE:  oddsSrv,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "key of license for the odds api")
	rootCmd.PersistentFlags().StringVarP(&pgAddr, "addr", "a", "127.0.0.1", "ip adds for the database")
	rootCmd.PersistentFlags().StringVarP(&pgPort, "port", "p", "5432", "ip port for the database")
	rootCmd.PersistentFlags().StringVarP(&pgUser, "user", "u", "postgres", "user for the database access")
	rootCmd.PersistentFlags().StringVarP(&pgPass, "pass", "P", "", "user password for the database access")

	rootCmd.MarkPersistentFlagRequired("license")
}

func oddsSrv(cmd *cobra.Command, args []string) error {

	// connect pgsql
	dsn := conn.MakeDsn(pgAddr, pgPort, pgUser, pgPass, "sbet")
	db, err := conn.ConnectPgdb(&conn.Postgres{Dsn: dsn, ConnMaxLiftTime: 10, MaxOpenConns: 10, MaxIdleConns: 1})
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("Connected to postgres database sbet")

	// http service
	r := serv.Service(db)

	// port listen
	log.Println("start service on port :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}

	return nil
}
