package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Position struct {
	row int
	col int
}

type Block interface {
	Draw(offsetX, offsetY int)
	Move(rows, cols int)
	Rotate(dir int)
	GetCellPositions() []Position
	GetId() int
}

type BaseBlock struct {
	id            int
	cellSize      int
	rotationState int
	cells         map[int][]Position
	rowOffset     int
	colOffset     int
}

func (b *BaseBlock) Draw(offsetX, offsetY int) {
	positions := b.GetCellPositions()

	for _, position := range positions {
		x := (int32)(position.col*b.cellSize + offsetX)
		y := (int32)(position.row*b.cellSize + offsetY)
		width := (int32)(b.cellSize - 1)
		height := (int32)(b.cellSize - 1)

		rl.DrawRectangle(x, y, width, height, CellColors[b.id])
	}
}

func (b *BaseBlock) GetId() int {
	return b.id
}

func (b *BaseBlock) Move(rows, cols int) {
	b.rowOffset += rows
	b.colOffset += cols
}

func (b *BaseBlock) Rotate(dir int) {
	b.rotationState = (b.rotationState + dir + 4) % 4
}

func (b *BaseBlock) GetCellPositions() []Position {
	tiles := b.cells[b.rotationState]

	// fmt.Println(tiles)

	movesTiles := make([]Position, 0)

	for _, pos := range tiles {
		movesTiles = append(movesTiles, Position{pos.row + b.rowOffset, pos.col + b.colOffset})
	}

	// fmt.Println(movesTiles)

	return movesTiles
}

type LBlock struct {
	BaseBlock
}

type JBlock struct {
	BaseBlock
}

type IBlock struct {
	BaseBlock
}

type OBlock struct {
	BaseBlock
}

type SBlock struct {
	BaseBlock
}

type TBlock struct {
	BaseBlock
}

type ZBlock struct {
	BaseBlock
}

func NewLBlock() *LBlock {
	lBlock := &LBlock{}
	lBlock.id = 1
	lBlock.cellSize = 30
	lBlock.cells = getLBlockRotationStates()
	lBlock.rotationState = 0
	lBlock.Move(0, 3)

	return lBlock
}

func NewJBlock() *JBlock {
	jBlock := &JBlock{}
	jBlock.id = 2
	jBlock.cellSize = 30
	jBlock.cells = getJBlockRotationStates()
	jBlock.rotationState = 0
	jBlock.Move(0, 3)

	return jBlock
}

func NewIBlock() *IBlock {
	iBlock := &IBlock{}
	iBlock.id = 3
	iBlock.cellSize = 30
	iBlock.cells = getIBlockRotationStates()
	iBlock.rotationState = 0
	iBlock.Move(-1, 3)

	return iBlock
}

func NewSBlock() *SBlock {
	sBlock := &SBlock{}
	sBlock.id = 4
	sBlock.cellSize = 30
	sBlock.cells = getSBlockRotationStates()
	sBlock.rotationState = 0
	sBlock.Move(0, 3)

	return sBlock
}

func NewOBlock() *OBlock {
	oBlock := &OBlock{}
	oBlock.id = 5
	oBlock.cellSize = 30
	oBlock.cells = getOBlockRotationStates()
	oBlock.rotationState = 0
	oBlock.Move(0, 4)

	return oBlock
}

func NewTBlock() *TBlock {
	tBlock := &TBlock{}
	tBlock.id = 6
	tBlock.cellSize = 30
	tBlock.cells = getTBlockRotationStates()
	tBlock.rotationState = 0
	tBlock.Move(0, 3)

	return tBlock
}
func NewZBlock() *ZBlock {
	zBlock := &ZBlock{}
	zBlock.id = 7
	zBlock.cellSize = 30
	zBlock.cells = getZBlockRotationStates()
	zBlock.rotationState = 0
	zBlock.Move(0, 3)

	return zBlock
}

// func intializeBlock(b *Block, id int, getRotationState func() map[int]Position) {
// 	b.id = id
// 	b.cellSize = 30
// 	b.cells = getRotationState()
// 	b.rotationState = 0
// }

func getLBlockRotationStates() map[int][]Position {
	cells := make(map[int][]Position)

	cells[0] = []Position{{0, 2}, {1, 0}, {1, 1}, {1, 2}}
	cells[1] = []Position{{0, 1}, {1, 1}, {2, 1}, {2, 2}}
	cells[2] = []Position{{1, 0}, {1, 1}, {1, 2}, {2, 0}}
	cells[3] = []Position{{0, 0}, {0, 1}, {1, 1}, {2, 1}}

	return cells
}

func getJBlockRotationStates() map[int][]Position {
	cells := make(map[int][]Position)
	cells[0] = []Position{{0, 0}, {1, 0}, {1, 1}, {1, 2}}
	cells[1] = []Position{{0, 1}, {0, 2}, {1, 1}, {2, 1}}
	cells[2] = []Position{{1, 0}, {1, 1}, {1, 2}, {2, 2}}
	cells[3] = []Position{{0, 1}, {1, 1}, {2, 0}, {2, 1}}

	return cells
}

func getIBlockRotationStates() map[int][]Position {
	cells := make(map[int][]Position)

	cells[0] = []Position{{1, 0}, {1, 1}, {1, 2}, {1, 3}}
	cells[1] = []Position{{0, 2}, {1, 2}, {2, 2}, {3, 2}}
	cells[2] = []Position{{2, 0}, {2, 1}, {2, 2}, {2, 3}}
	cells[3] = []Position{{0, 1}, {1, 1}, {2, 1}, {3, 1}}

	return cells
}

func getOBlockRotationStates() map[int][]Position {
	cells := make(map[int][]Position)

	cells[0] = []Position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	cells[1] = []Position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	cells[2] = []Position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	cells[3] = []Position{{0, 0}, {0, 1}, {1, 0}, {1, 1}}

	return cells
}

func getSBlockRotationStates() map[int][]Position {
	cells := make(map[int][]Position)

	cells[0] = []Position{{0, 1}, {0, 2}, {1, 0}, {1, 1}}
	cells[1] = []Position{{0, 1}, {1, 1}, {1, 2}, {2, 2}}
	cells[2] = []Position{{1, 1}, {1, 2}, {2, 0}, {2, 1}}
	cells[3] = []Position{{0, 0}, {1, 0}, {1, 1}, {2, 1}}

	return cells
}

func getTBlockRotationStates() map[int][]Position {
	cells := make(map[int][]Position)

	cells[0] = []Position{{0, 1}, {1, 0}, {1, 1}, {1, 2}}
	cells[1] = []Position{{0, 1}, {1, 1}, {1, 2}, {2, 1}}
	cells[2] = []Position{{1, 0}, {1, 1}, {1, 2}, {2, 1}}
	cells[3] = []Position{{0, 1}, {1, 0}, {1, 1}, {2, 1}}

	return cells
}

func getZBlockRotationStates() map[int][]Position {
	cells := make(map[int][]Position)

	cells[0] = []Position{{0, 0}, {0, 1}, {1, 1}, {1, 2}}
	cells[1] = []Position{{0, 2}, {1, 1}, {1, 2}, {2, 1}}
	cells[2] = []Position{{1, 0}, {1, 1}, {2, 1}, {2, 2}}
	cells[3] = []Position{{0, 1}, {1, 0}, {1, 1}, {2, 0}}

	return cells
}
