package tests

import (
	"bytes"
	"io"
	"main/internal/plants"
	"os"
	"testing"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func TestPlantsSuccess(t *testing.T) {
	tests := []struct {
		name  string
		plant interface{}
		want  string
	}{
		{
			name: "test1",
			plant: UnknownPlant{
				FlowerType: "rosa",
				LeafType:   "green",
				Color:      123,
			},
			want: "FlowerType:rosa\nLeafType:green\nColor(color_scheme=rgb):123\n",
		},
		{
			name: "test2",
			plant: AnotherUnknownPlant{
				FlowerColor: 10,
				LeafType:    "lanceolate",
				Height:      15,
			},
			want: "FlowerColor:10\nLeafType:lanceolate\nHeight(unit=inches):15\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := plants.DescribePlant(tt.plant)

			outC := make(chan string)

			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, r)
				outC <- buf.String()
			}()

			w.Close()
			os.Stdout = old
			text := <-outC

			if text != tt.want || err != nil {
				t.Errorf("expected: %v, got: %v", tt.want, text)
			}
		})
	}
}
