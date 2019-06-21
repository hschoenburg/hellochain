package types

type QueryResGreetings map[string][]Greeting

func (q QueryResGreetings) String() string {
	return fmt.Sprintf("%v", q)
}
