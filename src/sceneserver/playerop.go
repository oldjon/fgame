package main

type PlayerOp struct {
	optype uint32
	value  interface{}
}

type PlayerMove struct {
	UserId uint64
	Power  uint32
	Way    *Vector
}
