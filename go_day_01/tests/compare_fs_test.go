package tests

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestCompareFS(t *testing.T) {
	tests := []struct {
		name    string
		FlagOld string
		FlagNew string
		Old     string
		New     string
		result  string
	}{
		{
			name:    "test1",
			FlagOld: "--old",
			FlagNew: "--new",
			Old:     "./50/database/snapshot1.txt",
			New:     "./50/database/snapshot2.txt",
			result:  "./50/database/compareFS.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args = []string{""}
			data, err := exec.Command("../app", append(args, tt.FlagOld, tt.Old, tt.FlagNew, tt.New)...).Output()
			if err == nil {
				file, err := os.Open(tt.result)
				if err == nil {
					log.Println(err)
					sc := bufio.NewScanner(file)
					if string(data) != sc.Text() {
						t.Errorf("Got \"%v\", expected \"%v\"", string(data), sc.Text())
					}
				}
			}
		})
	}
}
