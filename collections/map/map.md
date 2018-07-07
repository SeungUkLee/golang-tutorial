# Map

Map은 키(key)에 대응하는 값(value)을 신속히 찾는 hash table 을 구현한 자료구조이다.
Go 언어는 Map 타입을 내장하고 있는데 `map[Key_type]Value_type` 으로 선언할 수 있다.
즉 다음과 같이 선언 할 수 있다. 

~~~ go
var idMap map[int]string
~~~


이 때 map은 reference 타입이므로 idMap은 nil 값을 갖는다.


map을 초기화하는 방법으로 `make()`함수를 사용할 수도 있다.

~~~ go
idMap = make(map[int]string)
~~~

`make()` 함수는 hash table 를 메모리에 생성하고 그 포인터를 리턴한다. 
따라서 idMap 은 hash table 을 가르키는 포인터를 갖는다.

 
`make()` 함수를 사용하지 않고 literal 을 사용해 초기화 할 수 있다.

~~~ go
mapVariable := map[string]string {
    "1": "test",
    "2": "test2",
    "3": "test3"
}
~~~

