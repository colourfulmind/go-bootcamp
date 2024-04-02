package tests

import (
	"bufio"
	"os"
	"os/exec"
	"testing"
)

func TestXargs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		xargs    []string
		filePath string
		expected string
	}{
		{
			name:     "test1",
			args:     []string{"-f", "-ext", "'txt'"},
			xargs:    []string{"../myWc", "-l"},
			filePath: "files",
			expected: "6 50/files/file1.txt\n" +
				"19 50/files/file2.txt\n" +
				"8 50/files/file3.txt\n",
		},
		{
			name:     "test2",
			filePath: "dirs",
			expected: "0 50/dirs/dir1/file1.txt\n" +
				"0 50/dirs/dir1/file2.txt\n" +
				"0 50/dirs/dir2/dir1/file1.txt\n" +
				"0 50/dirs/dir3/file1.txt\n" +
				"0 50/dirs/dir3/file2.txt\n",
		},
	}
	for _, tt := range tests {
		sc := bufio.NewScanner(os.Stdin)
		t.Run(tt.name, func(t *testing.T) {
			exec.Command("../myFind/myFind", append(tt.args, tt.filePath)...)

			for sc.Scan() {
				res, _ := exec.Command("../myXargs/myXargs", append(tt.xargs, sc.Text())...).Output()
				if string(res) != tt.expected {
					t.Errorf("Got \"%v\", expected \"%v\"", res, tt.expected)
				}
			}
		})
	}
}
