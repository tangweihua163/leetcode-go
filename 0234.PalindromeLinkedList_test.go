package leetcode_go

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(&ListNode{1, &ListNode{2, nil}}) == false)
	fmt.Println(isPalindrome(&ListNode{1, &ListNode{1, nil}}) == true)

	fmt.Println(isPalindrome(&ListNode{1, &ListNode{2, &ListNode{1, nil}}}) == true)
	fmt.Println(isPalindrome(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{1, nil}}}}) == false)
}
