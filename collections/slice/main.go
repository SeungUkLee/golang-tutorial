package main

import "fmt"

func main() {
	/* Slice */
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println("slice : " , a)

	s := make([]int, 5, 10)
	println(len(s), cap(s))

	var s2 []int
	if s2 == nil {
		println("Nil Slice")
	}
	println(len(s2), cap(s2))


	/* Sub-Slice */
	s3 := []int{0, 1, 2, 3, 4, 5}
	s3 = s[2:5]
	fmt.Println(s3)

	/* Slice append and copy */
	s4 := []int{0, 1}
	s4 = append(s, 2) // 0, 1, 2
	s4 = append(s4, 3, 4, 5) // 0, 1, 2, 3, 4, 5
	fmt.Println(s4)

	sliceA := make([]int, 0, 3) // len = 0, cap = 3
	for i := 1; i <=15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA)

	/* Slice 병합 */
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	slice3 := append(slice1, slice2...)
	// ... 는 해당 슬라이스의 컬렉션을 표현하는 것. 두번째 슬라이스의 모든 요소들의 집합을 나타냄.
	// slice3 = append(slice1, 4, 5, 6)
	fmt.Println(slice3)

	/* Slice copy */
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source) * 2)
	copy(target, source) // source 슬라이스가 갖는 배열 데이터를 target 슬라이스가 갖는 배열로 복제.
	fmt.Println(target)
	fmt.Println(len(target), cap(target))
}
