package main

import "fmt"

func main() {
	// findIntersectionOf2SortedArrays()
	// mergeTwoSortedArrays()
	// findConcurrentEventsInCalendar()
	unionOfIntervals()
}

func unionOfIntervals() {
	events := [][2]int{
		[2]int{1, 5}, [2]int{2, 7}, [2]int{4, 5}, [2]int{6, 10}, [2]int{8, 9},
		[2]int{9, 17}, [2]int{11, 13}, [2]int{12, 15}, [2]int{14, 15},
	}
	unionEvents, _ := unionIntervals(events, [][2]string{})
	fmt.Println(events)
	fmt.Println("Union of events: ", unionEvents)
	fmt.Println()
	events = [][2]int{
		[2]int{0, 3}, [2]int{1, 1}, [2]int{2, 4}, [2]int{3, 4}, [2]int{5, 7}, [2]int{7, 8},
		[2]int{8, 11}, [2]int{9, 11}, [2]int{12, 14}, [2]int{12, 16}, [2]int{13, 15}, [2]int{16, 17},
	}
	eventsIntervalType := [][2]string{
		[2]string{"o", "o"}, [2]string{"c", "c"}, [2]string{"c", "c"}, [2]string{"c", "o"}, [2]string{"c", "o"}, [2]string{"c", "o"},
		[2]string{"c", "o"}, [2]string{"o", "c"}, [2]string{"c", "c"}, [2]string{"o", "c"}, [2]string{"o", "o"}, [2]string{"o", "o"},
	}
	unionEvents, eventTypes := unionIntervals(events, eventsIntervalType)
	fmt.Println(events)
	fmt.Println(eventsIntervalType)
	fmt.Println("Union of events: ", unionEvents)
	fmt.Println("Union of event types: ", eventTypes)
}

// func findConcurrentEventsInCalendar() {
// 	events := [][2]int{
// 		[2]int{1, 5}, [2]int{2, 7}, [2]int{4, 5}, [2]int{6, 10}, [2]int{8, 9},
// 		[2]int{9, 17}, [2]int{11, 13}, [2]int{12, 15}, [2]int{14, 15},
// 	}
// 	concurrentEvents := findConcurrentEvents(events)
// 	fmt.Println(events)
// 	fmt.Println("Concurrent events: ", concurrentEvents)
// 	fmt.Println()
// 	events = [][2]int{
// 		[2]int{0, 3}, [2]int{1, 1}, [2]int{2, 4}, [2]int{3, 4}, [2]int{5, 7}, [2]int{7, 8},
// 		[2]int{8, 11}, [2]int{9, 11}, [2]int{12, 14}, [2]int{12, 16}, [2]int{13, 15}, [2]int{16, 17},
// 	}
// 	concurrentEvents = findConcurrentEvents(events)
// 	fmt.Println(events)
// 	fmt.Println("Concurrent events: ", concurrentEvents)
// }

// func mergeTwoSortedArrays() {
// 	arr1 := []int{5, 13, 17}
// 	arr2 := []int{3, 7, 11, 19}
// 	arr := merge2SortedArrays(arr1, arr2)
// 	fmt.Printf("Array1:%v\nArray2:%v\nMerged:%v\n", arr1, arr2, arr)
// }

// func findIntersectionOf2SortedArrays() {
// 	arr1 := []int{2, 3, 3, 5, 5, 6, 7, 7, 8, 12}
// 	arr2 := []int{5, 5, 6, 8, 8, 9, 10, 10}
// 	arr := findIntersection(arr1, arr2)
// 	fmt.Printf("Array1:%v\nArray2:%v\nIntersection:%v\n", arr1, arr2, arr)
// 	fmt.Println()
// 	arr1 = []int{2, 3, 3, 5, 7, 11}
// 	arr2 = []int{3, 3, 7, 15, 31}
// 	arr = findIntersection(arr1, arr2)
// 	fmt.Printf("Array1:%v\nArray2:%v\nResult:%v\n", arr1, arr2, arr)
// }
