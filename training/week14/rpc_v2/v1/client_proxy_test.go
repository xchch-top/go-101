package v1

import "testing"

func TestClient_Invoke(t *testing.T) {
	client := &Client{}
	InitClientProxy(&UserServiceClient{}, client)
}
