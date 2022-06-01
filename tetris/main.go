package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"strconv"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tetromino interface {
	color() color.Color
}

type TetrominoI struct{}
type TetrominoO struct{}

//type TetrominoT struct{}
//type TetrominoS struct{}
//type TetrominoZ struct{}
//type TetrominoJ struct{}
//type TetrominoL struct{}

func (t TetrominoI) color() color.Color {
	return color.RGBA{49, 199, 239, 255}
}

func (t TetrominoO) color() color.Color {
	return color.RGBA{247, 211, 8, 255}
}

var tetrominoes = []Tetromino{
	&TetrominoI{},
	&TetrominoO{},
	//&TetrominoT{},
	//&TetrominoS{},
	//&TetrominoZ{},
	//&TetrominoJ{},
	//&TetrominoL{},
}

var Matrix [][]color.Color

type Game struct{}

func (g *Game) Update() error {
	return nil
}

var CellSize = 10

var Counter = 0

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Size()
	screen.Fill(color.RGBA{255, 255, 255, 255})
	for i, row := range Matrix {
		for j, c := range row {
			cell := ebiten.NewImage(10, 10)
			cell.Fill(c)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(j*(CellSize+1)+5), float64(i*(CellSize+1)+3-(20*(CellSize+1))))
			screen.DrawImage(cell, op)

		}
	}
	Counter++
	ebitenutil.DebugPrint(screen, "Frame! "+strconv.Itoa(Counter))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 120, 227
}

func main() {
	Matrix = make([][]color.Color, 40)
	for i := range Matrix {
		Matrix[i] = make([]color.Color, 10)
		// TODO remove me
		for j := range Matrix[i] {
			Matrix[i][j] = color.RGBA{0, 100, 0, 255}
		}
	}

	ebiten.SetWindowSize(640, 900)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
