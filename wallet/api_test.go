package wallet

import (
	"encoding/json"
	"testing"
)

func TestJsonArray(t *testing.T) {
	ar := []string{"vite_15dac990004ae1cbf1af913091092f7c45205b88d905881d97",
		"vite_48c5a659e37a9a462b96ff49ef3d30f10137c600417ce05cda"}
	//as := Addresses{HexAddresses: ar}

	//ae :=[]string{}
	b, e := json.Marshal(ar)
	if e != nil {
		t.Fatal(e)
	}
	println(string(b))

	b, e = json.Marshal(nil)

	println(string(b))

}
func TestJsonMap(t *testing.T) {
	ma := make(map[string]string)
	ma["vite_15dac990004ae1cbf1af913091092f7c45205b88d905881d97"] = "Locked"
	ma["vite_48c5a659e37a9a462b96ff49ef3d30f10137c600417ce05cda"] = "Unlocked"

	b, e := json.Marshal(ma)
	if e != nil {
		t.Fatal(e)
	}
	println(string(b))

	b, e = json.Marshal(nil)

	println(string(b))

}
