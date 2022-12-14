package v7

import "testing"

func TestClient_Invoke(t *testing.T) {
	client := &Client{}
	InitClientProxy(&UserServiceClient{}, client)
}
