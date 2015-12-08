package main

type Stack struct {
	top        *Element
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

func (s *Stack) Push(piece *Piece) {
	if !s.Exist(piece.Key) {
		s.top = &Element{piece.Key, s.top}
		s.collection[piece.Key] = piece
	}
}

func (s *Stack) Pop() *Piece {
	if s.top != nil {
		key := s.top.key
		s.top = s.top.next
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
