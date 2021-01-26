package data

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestListMatch(t *testing.T) {
	type args struct {
		db *sqlx.DB
		in *Match
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{db, &Match{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListMatch(tt.args.db, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v\n", got)
		})
	}
}
