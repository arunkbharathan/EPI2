package main

func findConcurrentEvents(arr [][2]int) int {

	start := [24]int{}
	end := [24]int{}
	maxConcurrentEvents := 0
	for _, events := range arr {
		start[events[0]]++
		end[events[1]]++
	}
	concurrentEvents := 0
	for i := 0; i < 24; i++ {
		concurrentEvents = start[i] - end[i] + concurrentEvents
		maxConcurrentEvents = max(concurrentEvents, maxConcurrentEvents)
	}
	return maxConcurrentEvents
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
