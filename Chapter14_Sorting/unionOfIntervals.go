package main

func unionIntervals(arr [][2]int, eventType [][2]string) ([][2]int, [][2]string) {

	start := [24]int{}
	end := [24]int{}
	unionEvents := [][2]int{}
	unionEventTypes := [][2]string{}
	for _, events := range arr {
		start[events[0]]++
		end[events[1]]++
	}
	concurrentEvents := 0
	prevConcurrentEvents := 0
	startEvent, endEvent := -1, -1
	for i := 0; i < 24; i++ {
		concurrentEvents = start[i] - end[i] + concurrentEvents
		if prevConcurrentEvents == 0 && concurrentEvents > 0 {
			startEvent = i
		}
		if prevConcurrentEvents > 0 && concurrentEvents == 0 {
			endEvent = i
			unionEvents = append(unionEvents, [2]int{startEvent, endEvent})
		}
		prevConcurrentEvents = concurrentEvents
	}

	if len(eventType) > 0 {
		eventTypesOrdered := [24]string{}

		for ind, events := range arr {
			correspondingEventType := eventType[ind]
			if eventTypesOrdered[events[0]] != "c" {
				eventTypesOrdered[events[0]] = correspondingEventType[0]
			}
			if eventTypesOrdered[events[1]] != "c" {
				eventTypesOrdered[events[1]] = correspondingEventType[1]
			}
		}

		for _, events := range unionEvents {
			unionEventTypes = append(unionEventTypes, [2]string{eventTypesOrdered[events[0]], eventTypesOrdered[events[1]]})
		}

	}
	return unionEvents, unionEventTypes
}
