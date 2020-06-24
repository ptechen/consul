package consul

import "testing"

func TestDockerCreateServiceConsul(t *testing.T) {
	r ,err := CreateServiceConsul("127.0.0.2", 9000,"test","ping", "http")
	if err != nil {
		t.Error(err)
	}
	r.ServiceRegister(nil)
}