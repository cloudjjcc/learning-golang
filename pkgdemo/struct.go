package pkgdemo

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	data *Score
}

func (s *Student) GetName() string {
	return s.Name
}

type Score struct {
	Math    int `json:"math"`
	English int `json:"english"`
}

func (s *Score) GetMath() int {
	return s.Math
}
