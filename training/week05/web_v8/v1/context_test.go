package v1

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestContext_BindJSON(t *testing.T) {
	testCases := []struct {
		name string

		req   *http.Request
		input any

		wantErr error
		wantVal any
	}{
		{
			name: "happy case",
			req: func() *http.Request {
				mockData := &bytes.Buffer{}
				mockData.WriteString(`{"name": "TOM"}`)
				req, err := http.NewRequest(http.MethodPost, "/user", mockData)
				if err != nil {
					t.Fatal(err)
				}
				return req
			}(),
			input:   &User{},
			wantVal: &User{Name: "TOM"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := &Context{
				Req: tc.req,
			}
			err := ctx.BindJSON(tc.input)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, tc.input)
		})
	}
}
