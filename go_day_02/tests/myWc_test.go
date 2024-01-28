package tests

import (
	"os/exec"
	"testing"
)

func TestW(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "test1",
			filePath: "files/file1.txt",
			expected: "118 files/file1.txt\n",
		},
		{
			name:     "test2",
			filePath: "files/file2.txt",
			expected: "120 files/file2.txt\n",
		},
		{
			name:     "test3",
			filePath: "files/file3.txt",
			expected: "90 files/file3.txt\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args = []string{"-w"}
			data, _ := exec.Command("../myWc/myWc", append(args[0:], tt.filePath)...).Output()
			if string(data) != tt.expected {
				t.Errorf("Got \"%v\", expected \"%v\"", string(data), tt.expected)
			}
		})
	}
}

func TestL(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "test1",
			filePath: "files/file1.txt",
			expected: "10 files/file1.txt\n",
		},
		{
			name:     "test2",
			filePath: "files/file2.txt",
			expected: "15 files/file2.txt\n",
		},
		{
			name:     "test3",
			filePath: "files/file3.txt",
			expected: "8 files/file3.txt\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args = []string{"-l"}
			data, _ := exec.Command("../myWc/myWc", append(args[0:], tt.filePath)...).Output()
			if string(data) != tt.expected {
				t.Errorf("Got \"%v\", expected \"%v\"", string(data), tt.expected)
			}
		})
	}
}

func TestM(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "test1",
			filePath: "files/file1.txt",
			expected: "832 files/file1.txt\n",
		},
		{
			name:     "test2",
			filePath: "files/file2.txt",
			expected: "841 files/file2.txt\n",
		},
		{
			name:     "test3",
			filePath: "files/file3.txt",
			expected: "644 files/file3.txt\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args = []string{"-m"}
			data, _ := exec.Command("../myWc/myWc", append(args[0:], tt.filePath)...).Output()
			if string(data) != tt.expected {
				t.Errorf("Got \"%v\", expected \"%v\"", string(data), tt.expected)
			}
		})
	}
}
