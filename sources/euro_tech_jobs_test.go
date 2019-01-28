package main

import "testing"

func TestIrishJobsIe_Fetch(t *testing.T) {
	i := EuroTechJobs{}
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
