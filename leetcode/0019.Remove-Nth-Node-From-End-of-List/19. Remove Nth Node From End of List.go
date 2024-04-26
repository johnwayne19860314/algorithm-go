package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法一
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}

// 解法二
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	if n > len {
		return head
	}
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}
	current = head
	i := 0
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head
}

func removeNthNodeFromEnd0423(head *ListNode, n int) *ListNode {

	tmpNode := head
	// count the number of nodes 
	num := 0
	for tmpNode != nil {
		tmpNode = tmpNode.Next
		num++
	}
	tmpNode = head
	leftNodeNum := (num-n)
	
	if leftNodeNum == 0 {
		// case : remove the head node
		return tmpNode.Next
	}else if leftNodeNum < 0 {
		// case : remove the node before the head which is invalid
		return tmpNode
	}else if leftNodeNum == 1 {
		// case : remove the head's next node
		// need init the leftNode val for it will not be inited in the for loop after this
		removeNode := tmpNode.Next
		tmpNode.Next = removeNode.Next
		removeNode.Next = nil
		return tmpNode

	}
	
	// start from 1 instead of 0;
	iter := 1
	for iter < num {
		tmpNode = tmpNode.Next
		iter++
		if iter == leftNodeNum {
			removeNode := tmpNode.Next
			tmpNode.Next = removeNode.Next
			removeNode.Next = nil
			break
		}
		
	}
	return head

}
