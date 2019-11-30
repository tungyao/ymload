package test

import "testing"
import "../../ymload"

func TestYmlLoad(t *testing.T) {
	ymload.Format("./test.yml")
}
