package data

import (
	"testing"

	"github.com/Ariesfall/simple-odds-api/pkg/conn"
	"github.com/jmoiron/sqlx"
)

var db, _ = conn.ConnectPgdb(conn.DefaultConn)

func TestListSports(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{db}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListSports(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSports() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v\n", got)
		})
	}
}
