package linetocmd

import (
	"testing"
)

func TestParse(t *testing.T) {
	const command = `ping "127.0.0.1" -c 1`
	cmd, err := Parse(command)
	if err != nil {
		t.Fatalf(err.Error())
	}
	data, err := cmd.Output()
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("%s", data)
}