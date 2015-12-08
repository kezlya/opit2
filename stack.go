package main

type Stack struct {
	top        *Element
	size       int
	collection map[int]*Piece
}

type Element struct {
	key  int
	next *Element
}

func InitStack() *Stack {
	stack := Stack{
		collection: make(map[int]*Piece),
	}
	return &stack
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(piece *Piece) {
	if !s.Exist(piece.Key) {
		s.top = &Element{piece.Key, s.top}
		s.collection[piece.Key] = piece
		s.size++
	}
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() *Piece {
	if s.size > 0 {
		key := s.top.key
		s.top = s.top.next
		s.size--
		return s.Peek(key)
	}
	return nil
}

func (s *Stack) Exist(key int) bool {
	_, exist := s.collection[key]
	return exist
}

func (s *Stack) Peek(key int) *Piece {
	el, ok := s.collection[key]
	if ok {
		return el
	}
	return nil
}
