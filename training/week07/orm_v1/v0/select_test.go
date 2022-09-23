package v0

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelector_Build(t *testing.T) {
	tests := []struct {
		name    string
		s       QueryBuilder
		want    *Query
		wantErr error
	}{
		{
			name: "from",
			s:    NewSelector[TestModel]().From("test_model_tab"),
			want: &Query{
				SQL: "SELECT * FROM `test_model_tab`;",
			},
		},
		{
			name: "no from",
			s:    NewSelector[TestModel](),
			want: &Query{
				SQL: "SELECT * FROM `TestModel`;",
			},
		},
		{
			name: "from but empty",
			s:    NewSelector[TestModel]().From(""),
			want: &Query{
				SQL: "SELECT * FROM `TestModel`;",
			},
		},
		{
			name: "with db name",
			s:    NewSelector[TestModel]().From("test_db.test_model_tab"),
			want: &Query{
				SQL: "SELECT * FROM `test_db`.`test_model_tab`;",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Build()
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)

		})
	}
}

type TestModel struct {
	Id        int
	FirstName string
	Age       int8
	LastName  *sql.NullString
}
