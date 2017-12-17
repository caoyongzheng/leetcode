package leetcode

import "sort"

/*
	Given a collection of intervals, merge all overlapping intervals.

	For example,
	Given [1,3],[2,6],[8,10],[15,18],
	return [1,6],[8,10],[15,18].
*/

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

// Interval ...
type Interval struct {
	Start int
	End   int
}

// ByStart ...
type ByStart []Interval

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

func merge(intervals []Interval) (result []Interval) {
	if len(intervals) == 0 {
		return
	}
	sort.Sort(ByStart(intervals))
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start > result[len(result)-1].End {
			result = append(result, intervals[i])
		} else if result[len(result)-1].End < intervals[i].End {
			result[len(result)-1].End = intervals[i].End
		}
	}
	return
}
