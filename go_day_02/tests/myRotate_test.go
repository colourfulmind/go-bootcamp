package tests

import (
	"os/exec"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name: "test1",
			args: []string{"-a", "./files",
				"../tests/files/file1.log",
				"../tests/files/file2.log",
				"../tests/files/file3.log"},
			expected: "",
		},
		{
			name: "test2",
			args: []string{"-a", "./files",
				"../tests/files/file1.txt",
				"../tests/files/file2.txt",
				"../tests/files/file3.txt"},
			expected: "incorrect file format\n",
		},
		{
			name: "test3",
			args: []string{"../tests/files/file1.log",
				"../tests/files/file2.log",
				"../tests/files/file3.log"},
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, _ := exec.Command("../myRotate/myRotate", append(tt.args)...).Output()
			if string(data) != tt.expected {
				t.Errorf("Got \"%v\", expected \"%v\"", string(data), tt.expected)
			}
		})
	}
}
