// Package calculations get the input.
// Then calculates the result and output it.
package calculations

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

// Statistics stores the final values
type Statistics struct {
	Average, Middle, RegularSD float64
	MostCommon                 int
}

// Mean calculates an average number of the array
func (s *Statistics) Mean(nums []int) {
	var SumNums int
	for _, v := range nums {
		SumNums += v
	}
	s.Average = float64(SumNums) / float64(len(nums))
}

// Median finds the middle number of the sorted array
func (s *Statistics) Median(nums []int) {
	slices.Sort(nums)
	if len(nums)%2 == 0 {
		s.Middle = float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2.0
	} else {
		s.Middle = float64(nums[len(nums)/2])
	}
}

// Mode finds the most common number in the array
func (s *Statistics) Mode(nums []int) {
	counts := make(map[int]int)
	var MaxValue int
	for _, v := range nums {
		counts[v] += 1
		if counts[v] > MaxValue {
			MaxValue = counts[v]
		}
	}
	var res []int
	for key, value := range counts {
		if value == MaxValue {
			res = append(res, key)
		}
	}
	slices.Sort(res)
	s.MostCommon = res[0]
}

// SD finds the regular standard deviation
func (s *Statistics) SD(nums []int) {
	s.Mean(nums)
	var SumValues float64
	average := s.Average
	for _, v := range nums {
		SumValues += (float64(v) - average) * (float64(v) - average)
	}
	s.RegularSD = math.Sqrt(SumValues / float64(len(nums)-1))
}

// Output outputs the result
func (s *Statistics) Output(f Flags, numbers []int) {
	if len(numbers) > 0 {
		if f.Mean {
			s.Mean(numbers)
			fmt.Printf("Mean:   %.2f\n", s.Average)
		}
		if f.Median {
			s.Median(numbers)
			fmt.Printf("Median: %.2f\n", s.Middle)
		}
		if f.Mode {
			s.Mode(numbers)
			fmt.Println("Mode:  ", s.MostCommon)
		}
		if f.Sd {
			s.SD(numbers)
			fmt.Printf("SD:     %.2f\n", s.RegularSD)
		}
	} else {
		err := errors.New("not enough data")
		fmt.Println(err)
	}
}
