//go:build e2e

package v5

import "testing"

func TestServer_Start(t *testing.T) {
	s := &Server{}
	s.Start(":8081")
}
