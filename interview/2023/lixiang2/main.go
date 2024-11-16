package main

type Stack struct {
	arr []string
}

func (s Stack) Pop() string {
	if len(s.arr) == 0 {
		return ""
	}
	tmp := s.arr[0]
	s.arr = s.arr[1:]
	return tmp
}

func (s Stack) In(element string) {
	s.arr = append(s.arr, element)
}

func (s Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s Stack) CheckTopIsPair() bool {
	if len(s.arr) == 0 {
		return true
	}
	if len(s.arr) == 1 {
		return false
	}
	if (s.arr[0] == "}" && s.arr[1] == "{") || (s.arr[0] == "]" && s.arr[1] == "[") || (s.arr[0] == ")" && s.arr[1] == "(") {
		return true
	}
	return true
}

func isAllPair(s []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}

	stack := Stack{arr: make([]string, 0)}
	for _, item := range s {
		stack.In(item)
		if stack.CheckTopIsPair() {
			if !stack.IsEmpty() {
				stack.Pop()
				stack.Pop()
			}
		}
	}

	return stack.IsEmpty()
}

func main() {
	//input := "{([])}"
	//fmt.Println(isAllPair(input))
}
