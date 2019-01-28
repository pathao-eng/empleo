package main

import (
	"testing"
)

func TestWeLoveGolang_Fetch(t *testing.T) {
	i := WeLoveGolang{}
	if err := i.Init(); err != nil {
		t.Error(err)
		return
	}

	for jobs, ok := i.Fetch(); ; {
		for _, j := range jobs {
			v, err := j.Serialize()
			if err != nil {
				t.Error(err)
				continue
			}
			t.Logf("%s", v)
		}
		if !ok {
			break
		}
	}
}
