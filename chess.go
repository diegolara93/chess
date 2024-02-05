package main

import (
	"fmt"
)

type ChessPiece interface {
	move()
	isValidMove()
}

type King struct {
	x     int
	y     int
	color string
}

func (k *King) move() {
	fmt.Println("King moves")
}

func (k *King) isValidMove() {
	fmt.Println("King is valid move")
}

type Queen struct {
	x     int
	y     int
	color string
}

func (q *Queen) move() {
	fmt.Println("Queen moves")
}

func (q *Queen) isValidMove() {
	fmt.Println("Queen is valid move")
}

type Rook struct {
	x     int
	y     int
	color string
}

func (r *Rook) move() {
	fmt.Printf("%s rook moves\n", r.color)
}

func (r *Rook) isValidMove() {
	fmt.Println("Rook is valid move")
}

type Bishop struct {
	x     int
	y     int
	color string
}

func (b *Bishop) move() {
	fmt.Println("Bishop moves")
}

func (b *Bishop) isValidMove() {
	fmt.Println("Bishop is valid move")
}

type Knight struct {
	x     int
	y     int
	color string
}

func (k *Knight) move() {
	fmt.Println("Knight moves")
}

func (k *Knight) isValidMove() {
	fmt.Println("Knight is valid move")
}

type Pawn struct {
	x     int
	y     int
	color string
}

func (p *Pawn) move() {
	fmt.Println("Pawn moves")
}

func (p *Pawn) isValidMove() {
	fmt.Println("Pawn is valid move")
}

type ChessBoard struct {
	pieces []ChessPiece
}

func (c *ChessBoard) printCurrentBoard() {
	for _, piece := range c.pieces {
		fmt.Printf("%T, %+v\n", piece, piece)
	}
}

func main() {
	whiteKing := King{1, 5, "white"}
	blackKing := King{8, 5, "black"}
	whiteQueen := Queen{1, 4, "white"}
	blackQueen := Queen{8, 4, "black"}
	whiteKnight1 := Knight{1, 2, "white"}
	whiteKnight2 := Knight{1, 7, "white"}
	blackKnight1 := Knight{8, 2, "black"}
	blackKnight2 := Knight{8, 7, "black"}
	whiteBishop1 := Bishop{1, 3, "white"}
	whiteBishop2 := Bishop{1, 6, "white"}
	blackBishop1 := Bishop{8, 3, "black"}
	blackBishop2 := Bishop{8, 6, "black"}
	whiteLeftRook := Rook{1, 1, "white"}
	whiteRightRook := Rook{1, 8, "white"}
	blackLeftRook := Rook{8, 1, "black"}
	blackRightRook := Rook{8, 8, "black"}
	whitePawn1 := Pawn{2, 1, "white"}
	whitePawn2 := Pawn{2, 2, "white"}
	whitePawn3 := Pawn{2, 3, "white"}
	whitePawn4 := Pawn{2, 4, "white"}
	whitePawn5 := Pawn{2, 5, "white"}
	whitePawn6 := Pawn{2, 6, "white"}
	whitePawn7 := Pawn{2, 7, "white"}
	whitePawn8 := Pawn{2, 8, "white"}
	blackPawn1 := Pawn{7, 1, "black"}
	blackPawn2 := Pawn{7, 2, "black"}
	blackPawn3 := Pawn{7, 3, "black"}
	blackPawn4 := Pawn{7, 4, "black"}
	blackPawn5 := Pawn{7, 5, "black"}
	blackPawn6 := Pawn{7, 6, "black"}
	blackPawn7 := Pawn{7, 7, "black"}
	blackPawn8 := Pawn{7, 8, "black"}

	chessBoard := ChessBoard{[]ChessPiece{
		&whiteLeftRook, &whiteRightRook, &blackLeftRook, &blackRightRook,
		&whiteKnight1, &whiteKnight2, &blackKnight1, &blackKnight2,
		&whiteBishop1, &whiteBishop2, &blackBishop1, &blackBishop2,
		&whiteKing, &blackKing, &whiteQueen, &blackQueen,
		&whitePawn1, &whitePawn2, &whitePawn3, &whitePawn4, &whitePawn5, &whitePawn6, &whitePawn7, &whitePawn8,
		&blackPawn1, &blackPawn2, &blackPawn3, &blackPawn4, &blackPawn5, &blackPawn6, &blackPawn7, &blackPawn8}}
	chessBoard.printCurrentBoard()
}
