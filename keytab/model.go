package keytab

import (
	"encoding/json"
)

type Keytab struct {
	Path      string `json:"path,omitempty"`
	Principal string `json:"principal,omitempty"`
	Password  string `json:"password,omitempty"`
	Cipher    string `json:"cipher,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

func (k *Keytab) String() string {
	json, _ := json.Marshal(k)
	return string(json)
}
