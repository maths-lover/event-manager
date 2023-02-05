package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
)

func TestQueries_CreateOrganizer(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		ctx context.Context
		arg CreateOrganizerParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Organizer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "First",
			fields: fields{db: testQuery.db},
			args: args{
				ctx: context.Background(),
				arg: CreateOrganizerParams{
					Name:  "Organization1",
					Email: "organization1@mail.com",
					Phone: sql.NullString{
						String: "9898989898",
						Valid:  true,
					},
					Company: sql.NullString{
						String: "Company1",
						Valid:  true,
					},
					Logo: []byte("https://cdn.logo.com/hotlink-ok/logo-social.png"),
					Address: sql.NullString{
						String: "This guy lives at home",
						Valid:  true,
					},
				},
			},
			want: Organizer{
				ID:    1,
				Name:  "Organization1",
				Email: "organization1@mail.com",
				Phone: sql.NullString{
					String: "9898989898",
					Valid:  true,
				},
				Company: sql.NullString{
					String: "Company1",
					Valid:  true,
				},
				Logo: []byte("https://cdn.logo.com/hotlink-ok/logo-social.png"),
				Address: sql.NullString{
					String: "This guy lives at home",
					Valid:  true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := testQuery
			got, err := q.CreateOrganizer(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CreateOrganizer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.CreateOrganizer() = %v, want %v", got, tt.want)
			}
		})
	}
}
