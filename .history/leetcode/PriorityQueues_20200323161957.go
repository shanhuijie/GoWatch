package leetcode

import(
	"fmt"
)

type Queue struct{
	Index		int
	Value		int
	Priority	int
}

func Initqueue() *Queue {
	 var q *Queue
	 q = new(Queue)
	 return q
}

func(queue *Queue) Insertqueue(data int){
	for k,v := range queue{
		fmt.Println(k,v)
	}
}
