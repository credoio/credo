package credo

import (
	"fmt"
	"testing"
)

func TestParseLog(t *testing.T) {
	l, err := ParseLog("")
	if err == nil {
		t.Errorf("expected error")
	}
	if err != nil {
		fmt.Println(err)
	}
	l, err = ParseLog("abc")
	if err == nil {
		t.Errorf("expected error")
	}
	if err != nil {
		fmt.Println(err)
	}
	l, err = ParseLog("<aa>")
	if err == nil {
		t.Errorf("expected error")
	}
	if err != nil {
		fmt.Println(err)
	}
	l, err = ParseLog("<13>")
	if err == nil {
		t.Errorf("expected error")
	}
	l, err = ParseLog("<913>Sep  1 06:33:18 ubuntu dhclient[1009]: DHCPREQUEST of 172.16.206.128 on ens33 to 172.16.206.254 port 67 (xid=0x67389599)")
	if err == nil {
		t.Errorf("expected error")
	}
	if err != nil {
		fmt.Println(err)
	}
	l, err = ParseLog("<13>Sep  1 06:33:18 ubuntu dhclient[1009]: DHCPREQUEST of 172.16.206.128 on ens33 to 172.16.206.254 port 67 (xid=0x67389599)")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	fmt.Println(*l)
}
