package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)

	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	fmt.Println(str)

	noData := m[999]
	fmt.Println(noData)

	fmt.Println(m)
	delete(m, 777)
	fmt.Println(m)

	/* Map 키 체크 */
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	// map변수[키] 를 읽을 때 2개의 리턴값을 리턴한다.
	// 첫번쨰 : 키에 상응하는 값
	// 두번째 : 존재유무
	val, exists := tickers["MSFT"]
	if !exists {
		fmt.Println("No MSFT ticker")
	}
	fmt.Println(val)
	fmt.Println(exists)

	/* Map 열거 */
	for key, val := range tickers {
		fmt.Println(key, val)
	}
}
