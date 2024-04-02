package coins

import (
	"sort"
)

// MinCoins is a function that takes a value of change and a list of available coin values.
// It counts the minimum number of coins needed to make up the given value.
// This function works only under particular cases,
// such as when the list of coins is sorted or when the minimum amount contains the greatest value of coins.
func MinCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

// MinCoins2 does the same as MinCoins but handles errors.
// It returns an error for a negative value or an empty list of coins values.
// MinCoins2 is a recursive function that finds the best combination of coins.
func MinCoins2(val int, coins []int) []int {
	/*
		Variable tmp stores the temporary result for coin combinations to compare it with the current best result.
		Variable curSum stores the current best result for the minimum number of coins.
		Variable curRes stores the current best combination of coins.
	*/
	var tmp = make([]int, 0)
	var curSum = val
	var curRes = make([]int, 0)

	/*
		The loop iterates over all given coin values.
	*/
	for _, v := range coins {
		/*
			Check if the current coin value is greater than the given value, if so, skip to the next.
		*/
		if v > val {
			continue
		}

		/*
			Check if the difference between the given value and the coin value is negative,
			if so, this branch of recursion doesn't suit. Return the empty slice.
		*/
		if val-v < 0 {
			return []int{}
		}

		/*
			Check if we find one of the possible solutions.
		*/
		if v == val && val-v == 0 {
			return []int{v}
		}

		/*
			Recursion entry.
		*/
		tmp = append(MinCoins2(val-v, coins), v)

		/*
			If an empty slice was obtained, return to the beginning of the recursion.
			Else, check if the found solution is better than the current best solution.
		*/
		if len(tmp) == 0 {
			return []int{}
		} else if len(tmp) <= curSum {
			if len(tmp) == 1 && val-tmp[0] != 0 {
				continue
			}
			curSum = len(tmp)
			curRes = make([]int, curSum)
			copy(curRes, tmp)
		}

		/*
			Reset tmp variable.
		*/
		tmp = make([]int, 0)
	}

	/*
		Sort the final slice to satisfy the task.
	*/
	sort.Slice(curRes, func(i, j int) bool {
		return curRes[i] > curRes[j]
	})

	/*
		Return the best result.
	*/
	return curRes
}
