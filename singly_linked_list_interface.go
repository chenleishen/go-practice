package main

import "fmt"

// List represents a singly-linked list that holds values of any type.
type List struct {
	next *List
	val  interface{}
}

func (l *List) append(s interface{}) {
	if l == nil {
		*l = List{next: nil, val: s}
	}

	if l.val == nil {
		l.val = s
		return
	}

	cur := l
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &List{val: s}
}

func (l *List) print() {
	for l != nil {
		fmt.Println(l.val)
		l = l.next
	}
}

func (l *List) length() int {
	count := 0
	for l != nil {
		count += 1
		l = l.next
	}
	return count
}

func (l *List) prepend(s interface{}) {
	if l == nil {
		*l = List{next: nil, val: s}
	}

	if l.val == nil {
		l.val = s
		return
	}

	old_head := &List{next: l.next, val: l.val}
	*l = List{next: old_head, val: s}
}

func (l *List) contains(s interface{}) bool {
	if l == nil {
		return false
	}

	for l != nil {
		if l.val == s {
			return true
		}
		l = l.next
	}

	return false
}

func (l *List) delete(s interface{}) {
	if l == nil {
		return
	}

	if l.val == s {
		if l.next == nil {
			l.val = nil
			l.next = nil
		} else {
			l.val = l.next.val
			l.next = l.next.next
		}
		return
	}

	prev := l
	l = l.next

	for l != nil {
		if l.val == s {
			prev.next = l.next
			return
		}
		prev = l
		l = l.next
	}
}

func main() {
	var test List

	// append + print
	test.append("hello!")
	test.append("world!")
	test.print()

	// length
	fmt.Println(test.length())

	// prepend
	test.prepend("shelley says: ")
	test.print()

	// contains
	fmt.Println(test.contains("hello!"))
	fmt.Println(test.contains("shelley"))

	// delete
	test.delete("world!")
	test.print()
}

