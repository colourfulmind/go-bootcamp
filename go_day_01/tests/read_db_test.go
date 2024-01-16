package tests

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestReadDB(t *testing.T) {
	tests := []struct {
		name     string
		flag     string
		FileName string
		result   string
	}{
		{
			name:     "test1",
			flag:     "-f",
			FileName: "./tests/database/original_database.xml",
			result:   "./tests/database/original_database.json",
		},
		{
			name:     "test2",
			flag:     "-f",
			FileName: "./tests/database/stolen_database.json",
			result:   "./tests/database/stolen_database.xml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args = []string{""}
			data, err := exec.Command("../app", append(args, tt.flag, tt.FileName)...).Output()
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
