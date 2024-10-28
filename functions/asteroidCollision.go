package functions

// import(
// 	"fmt"
// )

// // Stack represents a stack data structure
// type Stack struct {
// 	items []int
// }

// // Push adds an element to the stack
// func (s *Stack) Push(item int) {
// 	s.items = append(s.items, item)
// }

// // Pop removes and returns the top element from the stack
// func (s *Stack) Pop() (int, error) {
// 	if len(s.items) == 0 {
// 		return 0, fmt.Errorf("stack is empty")
// 	}
// 	item := s.items[len(s.items)-1]
// 	s.items = s.items[:len(s.items)-1]
// 	return item, nil
// }

// // Peek returns the top element of the stack without removing it
// func (s *Stack) Peek() (int, error) {
// 	if len(s.items) == 0 {
// 		return 0, fmt.Errorf("stack is empty")
// 	}
// 	return s.items[len(s.items)-1], nil
// }

// // IsEmpty checks if the stack is empty
// func (s *Stack) IsEmpty() bool {
// 	return len(s.items) == 0
// }

// // Size returns the number of elements in the stack
// func (s *Stack) Size() int {
// 	return len(s.items)
// }

// func modulus (a int) int {
// 	if a > 0 { return a } else {return -1*a}
// }

// func asteroidCollision(asteroids []int) []int {
// 	s := Stack{}

// 	for i, a := range asteroids {
// 		if s.Size() == 0 {
// 			s.Push(a)
// 		} else {
// 			element, status := s.Peek()
// 			if status == nil {
// 				if element * a < 0 {
// 					if 
// 				} else {
// 					s.Push(a)
// 				}
// 			}

// 		}  
		
// 	}

// }