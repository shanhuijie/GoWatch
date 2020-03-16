package leetcode
import(
	"fmt"
)

type ListNode struct {
    Val 	int
  	Next 	*ListNode
}

// func hasCycle(head *ListNode) bool {
    
// }

func (this *ListNode)AddTailList(n int) {
	var temp ListNode
	
		for this.Next != nil{
			this = this.Next
		}
		temp.Val = n
		temp.Next = nil
		this.Next = &temp	
			
}

func (this *ListNode)AddHeadList(n int) {
	var temp ListNode
	temp.Val = n
	temp.Next = nil
	this.Next = &temp
	
	
}


func PrintList(list *ListNode) {
	p := list
	var i int
	for  {
		fmt.Printf("%d-%v-%p \n", p.Val, i, p.Next)
		p = p.Next
		if p.Next == nil{
			fmt.Printf("%d-%v-%p \n", p.Val, i, p.Next)
			break
		}
		
		i++
	}
	fmt.Println()
}