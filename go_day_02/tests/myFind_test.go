package tests

import (
	"os/exec"
	"testing"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "test1",
			filePath: "./dirs",
			expected: "./dirs\n" +
				"dirs/dir1\n" +
				"dirs/dir1/file1.txt\n" +
				"dirs/dir1/file2.txt\n" +
				"dirs/dir2\n" +
				"dirs/dir2/dir1\n" +
				"dirs/dir2/dir1/file1.txt\n" +
				"dirs/dir3\n" +
				"dirs/dir3/file1.txt\n" +
				"dirs/dir3/file2.txt\n" +
				"dirs/dir3/symlink1 -> dirs/dir1/file1.txt\n" +
				"dirs/dir3/symlink2 -> [broken]\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args = []string{""}
			data, err := exec.Command("../myFind/myFind", append(args, tt.filePath)...).Output()
			if string(data) != tt.expected && err != nil {
				t.Errorf("Got \"%v\", expected \"%v\"", string(data), tt.expected)
			}
		})
	}
}

func TestD(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "test1",
			filePath: "dirs",
			expected: "dirs\n" +
				"dirs/dir1\n" +
				"dirs/dir2\n" +
				"dirs/dir2/dir1\n" +
				"dirs/dir3\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args = []string{"-d"}
			data, err := exec.Command("../myFind/myFind", append(args, tt.filePath)...).Output()
			if string(data) != tt.expected && err != nil {
				t.Errorf("Got \"%v\", expected \"%v\"", string(data), tt.expected)
			}
		})
	}
}

func TestF(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		filePath string
		expected string
	}{
		{
			name:     "test1",
			args:     []string{"-f"},
			filePath: "dirs",
			expected: "dirs/dir1/file1.txt\n" +
				"dirs/dir1/file2.txt\n" +
				"dirs/dir2/dir1/file1.txt\n" +
				"dirs/dir3/file1.txt\n" +
				"dirs/dir3/file2.txt",
		},
		{
			name:     "test2",
			args:     []string{"-f", "-ext", "'txt'"},
			filePath: "dirs",
			expected: "dirs/dir1/file1.txt\n" +
				"dirs/dir1/file2.txt\n" +
				"dirs/dir2/dir1/file1.txt\n" +
				"dirs/dir3/file1.txt\n" +
				"dirs/dir3/file2.txt\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := exec.Command("../myFind/myFind", append(tt.args, tt.filePath)...).Output()
			if string(data) != tt.expected && err != nil {
				t.Errorf("Got \"%v\", expected \"%v\"", string(data), tt.expected)
			}
		})
	}
}

func TestSL(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected string
	}{
		{
			name:     "test1",
			filePath: "dirs",
			expected: "dirs/dir3/symlink1 -> dirs/dir1/file1.txt\n" +
				"dirs/dir3/symlink2 -> [broken]\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var args = []string{"-sl"}
			data, err := exec.Command("../myFind/myFind", append(args[0:], tt.filePath)...).Output()
			if string(data) != tt.expected && err != nil {
				t.Errorf("Got \"%v\", expected \"%v\"", string(data), tt.expected)
			}
		})
	}
}
