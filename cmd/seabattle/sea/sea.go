package sea

import "fmt"

type game struct {
	*sea
	owner *sea
	enemy *sea

	player map[uint16]*sea
}

type sea struct {
	pool   [][]*cell
	width  uint8
	height uint8
	area   uint8
}

type cell struct {
	state ship
}

type fleet struct {
	ship
}

type ship struct {
	deck        uint8
	position    [][2]uint8
	orientation string
}

func main() {
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
}
