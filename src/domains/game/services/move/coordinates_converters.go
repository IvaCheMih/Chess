package move

func IndexToCoordinates(index int) string {
	y := int('8') - (index / 8)
	x := (index % 8) + int('A')

	return string(byte(x)) + string(byte(y))
}

func (g *Game) IndexToCoordinates(index int) string {
	y := int('8') - (index / g.N)
	x := (index % g.N) + int('A')

	return string(byte(x)) + string(byte(y))
}

func (g *Game) CoordinatesToIndex(coordinates string) int {
	x := int(coordinates[0]) - int('A')
	y := int('8') - int(coordinates[1])

	return (y * g.N) + x
}
