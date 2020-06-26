package consul

import (
	"fmt"
	"testing"
)

func TestDockerCreateServiceConsul(t *testing.T) {
	r ,err := CreateServiceConsul("127.0.0.2", 9000,"test","ping", "http")
	if err != nil {
		t.Error(err)
	}
	err = r.ServiceRegister(nil)
	if err != nil {
		t.Error(err)
	}
}

func TestLocalIP(t *testing.T) {
	ip := LocalIP()
	fmt.Println(ip)
}

func TestExtranetIP(t *testing.T) {
	ip := ExtranetIP()
	fmt.Println(ip)
}