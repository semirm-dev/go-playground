package ex

import (
	"github.com/sirupsen/logrus"
)

func Arrr() {
	arr := make([]int, 2, 6) // [0, 0], len 2, cap 6
	arr2 := make([]int, 0)   // empty array
	arr3 := make([]int, 2)   // [0, 0], len 2, cap 2

	logrus.Info(arr)
	logrus.Info(arr2)
	logrus.Info(arr3)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	logrus.Info(arr)

	logrus.Info(arr[1:5])

	// copy rest of the array
	copy(arr[2:], arr[2+1:]) // dst: 3, 4..., start: 4, 5...
	//arr[len(arr)-1] = nil
	arr = arr[:len(arr)-1]

	logrus.Info(arr)

	// replace with last element
	arr[3] = arr[len(arr)-1]
	//arr[len(arr)-1] = nil
	arr = arr[:len(arr)-1]

	logrus.Info(arr)
}

