package main;

import "fmt";

const (
      boardSize int = 8
)

func getBoard(n int) string {
     // '+ n ' means '\n' at every row end
     var board = make([]rune, n*n+n);

     var x int = 0;
     for i := range board {
     	 if x != 0 && x % n == 0 {
	     board[i] = '\n';
	     x = 0;
	     continue;
	 }
	 x++;

	 if i % 2 == 1 {
	     board[i] = '#';
	 } else {
	     board[i] = ' ';
	 }
     }

     return string(board);
}

func main() {
     fmt.Println(getBoard(boardSize))
}
