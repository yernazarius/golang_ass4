package domain

type Group struct {
	id   int
	name string
}

func NewGroup(id int, name string) Group {
	return Group{
		id:   id,
		name: name,
	}
}

func (g *Group) SetName(name string) {
	
	if(len(name) > 250) {
		name = name[:250]
	}

	g.name = name
}

func (g *Group) GetName() string {
	return g.name
}

func (g *Group) GetID() int {
	return g.id
}