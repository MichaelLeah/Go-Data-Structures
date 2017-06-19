// Package godatastructures provides basic implementation of various data structures such as Stacks and Queues.
package godatastructures

// Link is a singular node within the Linked list chain
type Link struct {
	name interface{} // Store any data
	next *Link       // The pointer to the next link address in the chain
}

// LinkedList is the container that holds the pointer to the first link in the chain
type LinkedList struct {
	firstLink *Link // The pointer to the first link in the chain
}

// Insert places the next link into the chain
func (list *LinkedList) Insert(link *Link) {
	link.next = list.firstLink
	list.firstLink = link
}

// Remove gets rid of the most recent link in the chain
func (list *LinkedList) Remove() {
	link := list.firstLink
	if list.firstLink != nil {
		list.firstLink = link.next
	}
}
