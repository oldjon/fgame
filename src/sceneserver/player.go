package main

type Player struct {
	id       uint64
	skin     uint32 //低8位字节颜色表示rgb
	pos      *Position
	power    uint32
	way      *Vector
	movetime int64
}

/*
func NewPlayer(pos Position) *Player {位置信息从db load
}
*/
