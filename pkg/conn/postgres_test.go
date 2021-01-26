package conn

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestConnectPgdb(t *testing.T) {
	type args struct {
		p *Postgres
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{DefaultConn}, false},
		{"test2", args{&Postgres{
			Dsn:             MakeDsn("localhost", "5432", "postgres", "1qaz@WSX", "sbet"),
			ConnMaxLiftTime: 10,
			MaxOpenConns:    10,
			MaxIdleConns:    1,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectPgdb(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectPgdb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.Close()
		})
	}
}
