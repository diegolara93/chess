package main

import (
	"fmt"
)

type ChessPiece struct {
	x           int
	y           int
	color       string
	typeOfPiece string
}

func (c *ChessPiece) isValidRookMove(x int, y int, newX int, newY int) bool {
	return true
}
func (c *ChessPiece) isValidKingMove(x int, y int, newX int, newY int) bool {
	return true
}
func (c *ChessPiece) isValidQueenMove(x int, y int, newX int, newY int) bool {
	return true
}
func (c *ChessPiece) isValidPawnMove(x int, y int, newX int, newY int) bool {
	return true
}
func (c *ChessPiece) isValidBishopMove(x int, y int, newX int, newY int) bool {
	return true
}

func (c *ChessPiece) isValidMove(x int, y int, newX int, newY int) bool {
	switch c.typeOfPiece {
	case "rook":
		return c.isValidRookMove(x, y, newX, newY)
	case "king":
		return c.isValidKingMove(x, y, newX, newY)
	case "queen":
		return c.isValidQueenMove(x, y, newX, newY)
	case "pawn":
		return c.isValidPawnMove(x, y, newX, newY)
	case "bishop":
		return c.isValidBishopMove(x, y, newX, newY)
	default:
		return false
	}
}

func (c *ChessPiece) move(x int, y int) {
	if c.isValidMove(c.x, c.y, x, y) {
		c.x = x
		c.y = y
	} else {
		fmt.Println("Invalid move")
	}
}

type ChessBoard struct {
	pieces []ChessPiece
}

func (c *ChessBoard) printCurrentBoard() {
	for _, piece := range c.pieces {
		fmt.Printf("Piece: %s, Color: %s, X: %d, Y: %d\n", piece.typeOfPiece, piece.color, piece.x, piece.y)
	}
}

func (c *ChessBoard) checkIfWinningMove(x int, y int) bool {
	return true
}

func main() {
	whiteKing := ChessPiece{1, 5, "white", "king"}
	blackKing := ChessPiece{8, 5, "black", "king"}
	whiteQueen := ChessPiece{1, 4, "white", "queen"}
	blackQueen := ChessPiece{8, 4, "black", "queen"}
	whiteKnight1 := ChessPiece{1, 2, "white", "knight"}
	whiteKnight2 := ChessPiece{1, 7, "white", "knight"}
	blackKnight1 := ChessPiece{8, 2, "black", "knight"}
	blackKnight2 := ChessPiece{8, 7, "black", "knight"}
	whiteBishop1 := ChessPiece{1, 3, "white", "bishop"}
	whiteBishop2 := ChessPiece{1, 6, "white", "bishop"}
	blackBishop1 := ChessPiece{8, 3, "black", "bishop"}
	blackBishop2 := ChessPiece{8, 6, "black", "bishop"}
	whiteLeftRook := ChessPiece{1, 1, "white", "rook"}
	whiteRightRook := ChessPiece{1, 8, "white", "rook"}
	blackLeftRook := ChessPiece{8, 1, "black", "rook"}
	blackRightRook := ChessPiece{8, 8, "black", "rook"}
	whitePawn1 := ChessPiece{2, 1, "white", "pawn"}
	whitePawn2 := ChessPiece{2, 2, "white", "pawn"}
	whitePawn3 := ChessPiece{2, 3, "white", "pawn"}
	whitePawn4 := ChessPiece{2, 4, "white", "pawn"}
	whitePawn5 := ChessPiece{2, 5, "white", "pawn"}
	whitePawn6 := ChessPiece{2, 6, "white", "pawn"}
	whitePawn7 := ChessPiece{2, 7, "white", "pawn"}
	whitePawn8 := ChessPiece{2, 8, "white", "pawn"}
	blackPawn1 := ChessPiece{7, 1, "black", "pawn"}
	blackPawn2 := ChessPiece{7, 2, "black", "pawn"}
	blackPawn3 := ChessPiece{7, 3, "black", "pawn"}
	blackPawn4 := ChessPiece{7, 4, "black", "pawn"}
	blackPawn5 := ChessPiece{7, 5, "black", "pawn"}
	blackPawn6 := ChessPiece{7, 6, "black", "pawn"}
	blackPawn7 := ChessPiece{7, 7, "black", "pawn"}
	blackPawn8 := ChessPiece{7, 8, "black", "pawn"}

	chessBoard := ChessBoard{[]ChessPiece{
		whiteLeftRook, whiteRightRook, blackLeftRook, blackRightRook,
		whiteKnight1, whiteKnight2, blackKnight1, blackKnight2,
		whiteBishop1, whiteBishop2, blackBishop1, blackBishop2,
		whiteKing, blackKing, whiteQueen, blackQueen,
		whitePawn1, whitePawn2, whitePawn3, whitePawn4, whitePawn5, whitePawn6, whitePawn7, whitePawn8,
		blackPawn1, blackPawn2, blackPawn3, blackPawn4, blackPawn5, blackPawn6, blackPawn7, blackPawn8}}
	chessBoard.printCurrentBoard()
}
