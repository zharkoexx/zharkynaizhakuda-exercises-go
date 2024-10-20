package gameplay

import "math"

func GetBestMove(board []string, token string) int {
	bestScore := math.Inf(-1)
	bestMove := -1

	for i := 0; i < 9; i++ {
		if board[i] == " " {
			board[i] = token
			score := minimax(board, 0, false, token, math.Inf(-1), math.Inf(1))
			board[i] = " "
			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}
	return bestMove
}

func minimax(board []string, depth int, isMaximizing bool, token string, alpha, beta float64) float64 {
	opponentToken := "o"
	if token == "o" {
		opponentToken = "x"
	}

	if isWinning(board, token) {
		return 1
	}
	if isWinning(board, opponentToken) {
		return -1
	}
	if isBoardFull(board) {
		return 0
	}

	if isMaximizing {
		bestScore := math.Inf(-1)
		for i := 0; i < 9; i++ {
			if board[i] == " " {
				board[i] = token
				score := minimax(board, depth+1, false, token, alpha, beta)
				board[i] = " "
				bestScore = math.Max(bestScore, score)
				alpha = math.Max(alpha, bestScore)
				if beta <= alpha {
					break
				}
			}
		}
		return bestScore
	} else {
		bestScore := math.Inf(1)
		for i := 0; i < 9; i++ {
			if board[i] == " " {
				board[i] = opponentToken
				score := minimax(board, depth+1, true, token, alpha, beta)
				board[i] = " "
				bestScore = math.Min(bestScore, score)
				beta = math.Min(beta, bestScore)
				if beta <= alpha {
					break
				}
			}
		}
		return bestScore
	}
}

func isBoardFull(board []string) bool {
	for _, cell := range board {
		if cell == " " {
			return false
		}
	}
	return true
}
