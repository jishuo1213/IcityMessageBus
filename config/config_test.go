package config

import (
	"testing"
)

func Test_GetToUrl(t *testing.T) {
	err := InitConfig("../")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(Config.unExactUrls)
	}

	un, err := Config.getUnExactUrl("/beta/as/app/getYearBill?mobile=15969716233")

	if err == nil {
		t.Log(un, "bbbbbbbb")
	} else {
		t.Log(err)
	}

	if host, found := Config.urlTable[un]; found {
		t.Log(host)
	} else {
		t.Log("not found")
	}

	host, url, err := Config.GetToUrl("/beta/as/app/getYearBill?mobile=15969716233")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(host, url)
	}
}
