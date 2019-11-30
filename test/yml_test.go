package test

import "testing"
import "../../ymload"

func TestYmlLoad(t *testing.T) {
	d:= ymload.Format("./test.yml")
	t.Log(d)
	}
