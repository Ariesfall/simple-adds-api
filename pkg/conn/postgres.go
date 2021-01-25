package conn

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Postgres conn detail
type Postgres struct {
	Dsn             string
	ConnMaxLiftTime int
	MaxOpenConns    int
	MaxIdleConns    int
}

var (
	// postgres default conn detail
	DefaultConn = &Postgres{
		Dsn:             "postgres:password@localhost:5432/sbet",
		ConnMaxLiftTime: 10,
		MaxOpenConns:    10,
		MaxIdleConns:    1,
	}
)

// ConnectPgdb connect pg sql
func ConnectPgdb(p *Postgres) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", p.Dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(p.ConnMaxLiftTime))
	db.SetMaxOpenConns(p.MaxOpenConns)
	db.SetMaxIdleConns(p.MaxIdleConns)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MakeDsn(host, port, user, password, dbName string) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	return dsn
}
