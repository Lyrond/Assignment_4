package domain

type Group struct {
	Id   int
	Name string
}

var groupCounter int

func NewGroup(groupName string) *Group {
	if len(groupName) > 250 {
		return nil
	}

	groupCounter++

	return &Group{
		Id:   groupCounter,
		Name: groupName,
	}
}
