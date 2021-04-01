package No435

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] == left && intervals[i][1] >= right {
			count++
		} else if intervals[i][0] == left && intervals[i][1] < right {
			count++
			right = intervals[i][1]
		} else if intervals[i][0] >= right {
			right = intervals[i][1]
		} else if intervals[i][0] > left && intervals[i][1] >= right {
			count++
		} else if intervals[i][0] > left && intervals[i][1] < right {
			count++
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	return count
}
