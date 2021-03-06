package tx
type SavePointStack struct {
	i    int
	data []string //方法队列
}

func (it SavePointStack) New() SavePointStack {
	return SavePointStack{
		data: []string{},
		i:    0,
	}
}

func (s *SavePointStack) Push(k string) {
	s.data = append(s.data, k)
	s.i++
}

func (s *SavePointStack) Pop() *string {
	if s.i == 0 {
		return nil
	}
	s.i--
	var ret = s.data[s.i]
	s.data = s.data[0:s.i]
	return &ret
}

func (s *SavePointStack) Len() int {
	return s.i
}
