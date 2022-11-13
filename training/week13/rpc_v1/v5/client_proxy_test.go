package v5

import "testing"

func TestClient_Invoke(t *testing.T) {
	client := &Client{}
	InitClientProxy(&UserService{}, client)
}
