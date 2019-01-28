package core

import "encoding/json"

type Empleo struct {
	Title    string
	Company  string
	Location string
	Link     string
	Tags     []string
	Time     string
}

func (e *Empleo) Serialize() (string, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(b), err
}
