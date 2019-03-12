package main

type cell struct {
	T        terrain
	Explored bool
}

type terrain int

const (
	WallCell terrain = iota
	GroundCell
	DoorCell
	FungusCell
	BarrelCell
	StairCell
	StoneCell
)

func (c cell) IsFree() bool {
	switch c.T {
	case WallCell, BarrelCell:
		return false
	default:
		return true
	}
}

func (c cell) Flammable() bool {
	switch c.T {
	case FungusCell, DoorCell, BarrelCell:
		return true
	default:
		return false
	}
}

func (c cell) IsGround() bool {
	switch c.T {
	case GroundCell:
		return true
	default:
		return false
	}
}

func (c cell) IsNotable() bool {
	switch c.T {
	case StairCell, StoneCell, BarrelCell:
		return true
	default:
		return false
	}
}

func (c cell) ShortDesc(g *game, pos position) (desc string) {
	switch c.T {
	case WallCell:
		desc = "a wall"
	case GroundCell:
		desc = "the ground"
	case DoorCell:
		desc = "a door"
	case FungusCell:
		desc = "foliage"
	case BarrelCell:
		desc = "a barrel"
	case StairCell:
		desc = g.Objects.Stairs[pos].ShortDesc(g)
	}
	return desc
}

func (c cell) Desc(g *game, pos position) (desc string) {
	switch c.T {
	case WallCell:
		desc = "A wall is an impassable pile of rocks. It can be destructed by using some items."
	case GroundCell:
		desc = "This is just plain ground."
	case DoorCell:
		desc = "A closed door blocks your line of sight. Doors open automatically when you or a monster stand on them. Doors are flammable."
	case FungusCell:
		desc = "Blue dense foliage grows in the Underground. It is difficult to see through, and is flammable."
	case BarrelCell:
		desc = "A barrel. You can hide yourself inside it when no monsters see you. It is a safe place for resting and recovering."
	case StairCell:
		desc = g.Objects.Stairs[pos].Desc(g)
	}
	return desc
}

func (c cell) Style(g *game, pos position) (r rune, fg uicolor) {
	switch c.T {
	case WallCell:
		r, fg = '#', ColorFgLOS
	case GroundCell:
		r, fg = '.', ColorFgLOS
	case DoorCell:
		r, fg = '+', ColorFgPlace
	case FungusCell:
		r, fg = '"', ColorFgLOS
	case BarrelCell:
		r, fg = '_', ColorFgCollectable // TODO: change letter and color
	case StairCell:
		st := g.Objects.Stairs[pos]
		r, fg = st.Style(g)
	}
	return r, fg
}