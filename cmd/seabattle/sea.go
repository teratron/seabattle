package main

type Game struct {
	*Sea
	owner *Sea
	enemy *Sea

	player map[uint16]*Sea
}

type Sea struct {
	pool   [][]*cell
	width  uint8
	height uint8
	area   uint8
}

type cell struct {
	state Ship
}

type Fleet struct {
	Ship
}

type Ship struct {
	deck        uint8
	position    [][2]uint8
	orientation string
}

/*func main() {
	ship := ship{
		deck: 4,
		position: [][2]uint8{
			{2, 1},
			{2, 2},
			{2, 3},
			{2, 4},
		},
	}
	fmt.Println(ship)
}*/
