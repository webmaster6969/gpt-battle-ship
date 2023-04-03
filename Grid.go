package main

type Grid struct {
	Cells []*Cell `json:"Cells"`
}

func (g *Grid) AddCell(Location Location, Status int) {
	g.Cells = append(g.Cells, &Cell{Location, Status})
}

func (g *Grid) FindCell(Location Location) *Cell {
	for _, cell := range g.Cells {
		if cell.Location.X == Location.X && cell.Location.Y == Location.Y {
			return cell
		}
	}

	return nil
}
