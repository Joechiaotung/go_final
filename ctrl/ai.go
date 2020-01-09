package ctrl

import (
	"github.com/Joechiaotung/go_final/model"
	"fmt"
)

// Convenience method to check if a piece on a board equals a type of piece
func check(i int, j int, board [][]model.Stone, piece model.Stone) bool {
	if i < len(board) && j < len(board) && i >= 0 && j >= 0 {
		return board[i][j] == piece
	} else {
		return false
	}
}

func defenseScore(count int, targetCount int) int {
	if count >= targetCount {
		if targetCount == 1 {
			return 10
		} else if targetCount == 2 {
			return 30
		} else if targetCount == 3 {
			return 40
		} else if targetCount == 4 {
			return 111111
		}
	}
	return 0
}

func offenseScore(count int, targetCount int) int {
	if count >= targetCount {
		if targetCount == 1 {
			return 20
		} else if targetCount == 2 {
			return 60
		} else if targetCount == 3 {
			return 80
		} else if targetCount == 4 {
			return 99999999
		}
	}
	return 0
}

func hHeatDefense(i int, j int, targetCount int, board [][]model.Stone) int {

	// targetStone := model.StoneBlack
	// notTargetStone := model.StoneWhite
	// scoreFunc := offenseScore

	// if piece == model.StoneWhite {
	// 	targetStone = model.StoneWhite
	// 	notTargetStone = model.StoneBlack
	// 	scoreFunc = defenseScore
	// }

	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {

		if j < len(board)-k {
			if board[i][j+k] == model.StoneBlack {
				count0 -= 1
			}

			if board[i][j+k] == model.StoneWhite {
				count0 += 1
			}
		}

		if j-k > 0 {
			if board[i][j-k] == model.StoneBlack {
				count1 -= 1
			}

			if board[i][j-k] == model.StoneWhite {
				count1 += 1
			}
		}
	}

	if count0 > count1 {
		return defenseScore(count0, targetCount)
	} else {
		return defenseScore(count1, targetCount)
	}

}

func dHeat(i int, j int, targetCount int, piece model.Stone, board [][]model.Stone) int {
	count0, count1 := 0, 0
	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore
	}

	for k := 1; k <= targetCount; k++ {

		if check(i+k, j+k, board, notTargetStone) {
			count0 -= 1
		}

		if check(i+k, j+k, board, targetStone) {
			count0 += 1
		}

		if check(i-k, j-k, board, notTargetStone) {
			count1 -= 1
		}

		if check(i-k, j-k, board, targetStone) {
			count1 += 1
		}
	}

	if count0 > count1 {
		return scoreFunc(count0, targetCount)
	} else {
		return scoreFunc(count1, targetCount)
	}
}

func dHeat2(i int, j int, targetCount int, piece model.Stone, board [][]model.Stone) int {
	count0, count1 := 0, 0
	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore
	}

	for k := 1; k <= targetCount; k++ {
		if check(i-k, j+k, board, notTargetStone) {
			count0 -= 1
		}

		if check(i-k, j+k, board, targetStone) {
			count0 += 1
		}

		if check(i+k, j-k, board, notTargetStone) {
			count1 -= 1
		}

		if check(i+k, j-k, board, targetStone) {
			count1 += 1
		}
	}

	if count0 > count1 {
		return scoreFunc(count0, targetCount)
	} else {
		return scoreFunc(count1, targetCount)
	}
}

func vHeatDefense(i int, j int, targetCount int, board [][]model.Stone) int {
	count0, count1 := 0, 0
	for k := 1; k <= targetCount; k++ {

		if (i < len(board)-k) && (board[i+k][j] == model.StoneBlack) {
			count0 -= 1
		}

		if (i-k > 0) && (board[i-k][j] == model.StoneBlack) {
			count1 -= 1
		}

		if (i < len(board)-k) && (board[i+k][j] == model.StoneWhite) {
			count0 += 1
		}

		if (i-k > 0) && (board[i-k][j] == model.StoneWhite) {
			count1 += 1
		}

	}

	if count0 > count1 {
		return defenseScore(count0, targetCount)
	} else {
		return defenseScore(count1, targetCount)
	}

}

func vPotential(i int, j int, piece model.Stone, board [][]model.Stone) int {
	count, maxCount := 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore
	}

	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}

			if check(i+k-m, j, board, targetStone) {
				count += 1
			}

			if check(i+k-m, j, board, notTargetStone) {
				count = 0
				break
			}

		}

		if count > maxCount {
			maxCount = count
		}
	}

	return scoreFunc(maxCount, maxCount)
}

func hPotential(i int, j int, piece model.Stone, board [][]model.Stone) int {
	count, maxCount := 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore
	}

	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}

			if check(i, j+k-m, board, targetStone) {
				count += 1
			}

			if check(i, j+k-m, board, notTargetStone) {
				count = 0
				break
			}

		}
		if count > maxCount {
			maxCount = count
		}
	}
	return scoreFunc(maxCount, maxCount)
}

func dPotential(i int, j int, piece model.Stone, board [][]model.Stone) int {
	count, maxCount := 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore
	}

	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}

			if check(i+k-m, j+k-m, board, targetStone) {
				count += 1
			}

			if check(i+k-m, j+k-m, board, notTargetStone) {
				count = 0
				break
			}

		}
		if count > maxCount {
			maxCount = count
		}
	}
	return scoreFunc(maxCount, maxCount)
}

func dPotential2(i int, j int, piece model.Stone, board [][]model.Stone) int {
	count, maxCount := 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	scoreFunc := offenseScore

	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
		scoreFunc = defenseScore
	}

	for m := 0; m < 5; m++ {
		count = 0
		for k := 0; k < 5; k++ {
			if k-m == 0 {
				continue
			}

			if check(i-k+m, j+k-m, board, targetStone) {
				count += 1
			}

			if check(i-k-m, j+k-m, board, notTargetStone) {
				count = 0
				break
			}

		}
		if count > maxCount {
			maxCount = count
		}
	}
	return scoreFunc(maxCount, maxCount)
}

// --ooo

// --ooo--
// -o-oo-
// -oo-o-

// Check for three horizontal with a) two empy spaces surrounding for consecutive pieces.
// and b) one empty space surrounding for non-consecutive pieces
// a) x-ooo-x
// b) -o-oo-
//    -oo-o-
// c) -*ooo-

// This format guarantees a one-move away from win
func threeHorizontal(i int, j int, piece model.Stone, board [][]model.Stone, score int) int {
	count, count0, count1 := 0, 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
	}

	// a) x-ooo-x
	for k := 0; k < 6; k++ {
		if (k == 0) || (k == 4) || (k == 5) {
			targetStone = model.StoneEmpty
		} else {
			targetStone = piece
		}

		// Check once that -1 stone is empty
		if (k == 0) && check(i, j-1, board, model.StoneEmpty) {
			count0 += 1
		}

		if check(i, j+k, board, targetStone) {
			count0 += 1
		}

		// check once that +1 stone is empty
		if (k == 0) && check(i, j+1, board, model.StoneEmpty) {
			count1 += 1
		}

		if check(i, j-k, board, targetStone) {
			count1 += 1
		}
	}

	if count0 == 7 || count1 == 7 {
		return score
	}

	// b)
	// -oo-o-
	// -o-oo-
	count = 0
	for k := 1; k < 3; k++ {
		if (j < len(board)-k) && (j-k >= 0) {
			if (k == 1) && check(i, j+k, board, piece) && check(i, j-k, board, piece) {
				count = 2
			} else if check(i, j+k, board, piece) {
				count += 1
			} else if check(i, j-k, board, piece) {
				count += 1
			} else if check(i, j+k, board, notTargetStone) {
				count = 0
			} else if check(i, j-k, board, notTargetStone) {
				count = 0
			}
		}
	}
	if count == 3 {
		return score
	}

	// c)
	// -*ooo-
	// -ooo*-
	count0, count1 = 0, 0
	for k := 1; k < 5; k++ {
		if (k == 1) && check(i, j+k, board, piece) && check(i, j-k, board, model.StoneEmpty) {
			count0 = 2
		} else if (k == 4) && check(i, j+k, board, model.StoneEmpty) {
			count0 += 1
		} else if check(i, j+k, board, piece) {
			count0 += 1
		}

		if (k == 1) && check(i, j-k, board, piece) && check(i, j+k, board, model.StoneEmpty) {
			count1 = 2
		} else if (k == 4) && check(i, j-k, board, model.StoneEmpty) {
			count1 += 1
		} else if check(i, j-k, board, piece) {
			count1 += 1
		}
	}

	if count0 == 5 || count1 == 5 {
		return score
	}

	return 0

}

func threeVertical(i int, j int, piece model.Stone, board [][]model.Stone, score int) int {
	count, count0, count1 := 0, 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
	}

	// x-ooo-x
	for k := 0; k < 6; k++ {
		if (k == 0) || (k == 4) || (k == 5) {
			targetStone = model.StoneEmpty
		} else {
			targetStone = piece
		}

		// Check once that -1 stone is empty
		if (k == 0) && check(i-1, j, board, model.StoneEmpty) {
			count0 += 1
		}

		if check(i+k, j, board, targetStone) {
			count0 += 1
		}

		// check once that +1 stone is empty
		if (k == 0) && check(i+1, j, board, model.StoneEmpty) {
			count1 += 1
		}

		if check(i-k, j, board, targetStone) {
			count1 += 1
		}
	}

	if count0 == 7 || count1 == 7 {
		return score
	}

	// -oo-o-
	// -o-oo-
	count = 0
	for k := 1; k < 3; k++ {
		if (k == 1) && check(i+k, j, board, piece) && check(i-k, j, board, piece) {
			count = 2
		} else if check(i+k, j, board, piece) {
			count += 1
		} else if check(i-k, j, board, piece) {
			count += 1
		} else if check(i+k, j, board, notTargetStone) {
			count = 0
		} else if check(i-k, j, board, notTargetStone) {
			count = 0
		}
	}
	if count == 3 {
		return score
	}

	// c)
	// -*ooo-
	// -ooo*-
	count0, count1 = 0, 0
	for k := 1; k < 5; k++ {
		if (k == 1) && check(i+k, j, board, piece) && check(i-k, j, board, model.StoneEmpty) {
			count0 = 2
		} else if (k == 4) && check(i+k, j, board, model.StoneEmpty) {
			count0 += 1
		} else if check(i+k, j, board, piece) {
			count0 += 1
		}

		if (k == 1) && check(i-k, j, board, piece) && check(i+k, j, board, model.StoneEmpty) {
			count1 = 2
		} else if (k == 4) && check(i-k, j, board, model.StoneEmpty) {
			count1 += 1
		} else if check(i-k, j, board, piece) {
			count1 += 1
		}
	}

	if count0 == 5 || count1 == 5 {
		return score
	}

	return 0
}

func threeDiagonal(i int, j int, piece model.Stone, board [][]model.Stone, score int) int {
	count, count0, count1 := 0, 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
	}

	// x-ooo-x
	for k := 0; k < 6; k++ {
		if (k == 0) || (k == 4) || (k == 5) {
			targetStone = model.StoneEmpty
		} else {
			targetStone = piece
		}

		// Check once that -1 stone is empty
		if (k == 0) && check(i-1, j-1, board, model.StoneEmpty) {
			count0 += 1
		}

		if check(i+k, j+k, board, targetStone) {
			count0 += 1
		}

		// check once that +1 stone is empty
		if (k == 0) && check(i+1, j+1, board, model.StoneEmpty) {
			count1 += 1
		}

		if check(i-k, j-k, board, targetStone) {
			count1 += 1
		}
	}

	if count0 == 7 || count1 == 7 {
		return score
	}

	// -oo*o-
	// -o*oo-
	count = 0
	for k := 1; k < 3; k++ {
		if (k == 1) && check(i+k, j+k, board, piece) && check(i-k, j-k, board, piece) {
			count = 2
		} else if check(i+k, j+k, board, piece) {
			count += 1
		} else if check(i-k, j-k, board, piece) {
			count += 1
		} else if check(i+k, j+k, board, notTargetStone) {
			count = 0
		} else if check(i-k, j-k, board, notTargetStone) {
			count = 0
		}
	}
	if count == 3 {
		return score
	}

	// -*ooo-
	// -ooo*-
	count0, count1 = 0, 0
	for k := 1; k < 5; k++ {
		if (k == 1) && check(i+k, j+k, board, piece) && check(i-k, j-k, board, model.StoneEmpty) {
			count0 = 2
		} else if (k == 4) && check(i+k, j+k, board, model.StoneEmpty) {
			count0 += 1
		} else if check(i+k, j+k, board, piece) {
			count0 += 1
		}

		if (k == 1) && check(i-k, j-k, board, piece) && check(i+k, j+k, board, model.StoneEmpty) {
			count1 = 2
		} else if (k == 4) && check(i-k, j-k, board, model.StoneEmpty) {
			count1 += 1
		} else if check(i-k, j-k, board, piece) {
			count1 += 1
		}
	}

	if count0 == 5 || count1 == 5 {
		return score
	}

	return 0
}

func threeDiagonal2(i int, j int, piece model.Stone, board [][]model.Stone, score int) int {
	count, count0, count1 := 0, 0, 0

	targetStone := model.StoneBlack
	notTargetStone := model.StoneWhite
	if piece == model.StoneWhite {
		targetStone = model.StoneWhite
		notTargetStone = model.StoneBlack
	}

	// x-ooo-x
	for k := 0; k < 6; k++ {
		if (k == 0) || (k == 4) || (k == 5) {
			targetStone = model.StoneEmpty
		} else {
			targetStone = piece
		}

		// Check once that -1 stone is empty
		if (k == 0) && check(i+1, j-1, board, model.StoneEmpty) {
			count0 += 1
		}

		if check(i-k, j+k, board, targetStone) {
			count0 += 1
		}

		// check once that +1 stone is empty
		if (k == 0) && check(i-1, j+1, board, model.StoneEmpty) {
			count1 += 1
		}

		if check(i+k, j-k, board, targetStone) {
			count1 += 1
		}
	}

	if count0 == 7 || count1 == 7 {
		return score
	}

	// -*ooo-
	// -ooo*-
	count0, count1 = 0, 0
	for k := 1; k < 5; k++ {
		if (k == 1) && check(i-k, j+k, board, piece) && check(i+k, j-k, board, model.StoneEmpty) {
			count0 = 2
		} else if (k == 4) && check(i-k, j+k, board, model.StoneEmpty) {
			count0 += 1
		} else if check(i-k, j+k, board, piece) {
			count0 += 1
		}

		if (k == 1) && check(i+k, j-k, board, piece) && check(i-k, j+k, board, model.StoneEmpty) {
			count1 = 2
		} else if (k == 4) && check(i+k, j-k, board, model.StoneEmpty) {
			count1 += 1
		} else if check(i+k, j-k, board, piece) {
			count1 += 1
		}
	}
	if count0 == 5 || count1 == 5 {
		return score
	}

	// -oo*o-
	// -o*oo-
	count = 0
	for k := 1; k < 3; k++ {
		if (k == 1) && check(i+k, j-k, board, piece) && check(i-k, j+k, board, piece) {
			count = 2
		} else if check(i-k, j+k, board, piece) {
			count += 1
		} else if check(i+k, j-k, board, piece) {
			count += 1
		} else if check(i-k, j+k, board, notTargetStone) {
			count = 0
		} else if check(i-k, j+k, board, notTargetStone) {
			count = 0
		}
	}
	if count == 3 {
		return score
	}

	return 0

}

func hHeatOffense(i int, j int, targetCount int, board [][]model.Stone) int {
	count := 0
	for k := 1; k <= targetCount; k++ {

		if (j < len(board)-k) && (board[i][j+k] == model.StoneWhite) {
			count -= 1
		}

		if (j < len(board)-k) && (board[i][j+k] == model.StoneBlack) {
			count += 1
		}

		if (j-k > 0) && (board[i][j-k] == model.StoneWhite) {
			count -= 1
		}

		if (j-k > 0) && (board[i][j-k] == model.StoneBlack) {
			count += 1
		}
	}

	return offenseScore(count, targetCount)
}

func vHeatOffense(i int, j int, targetCount int, board [][]model.Stone) int {
	count := 0
	for k := 1; k <= targetCount; k++ {
		if (i < len(board)-k) && (board[i+k][j] == model.StoneWhite) {
			count -= 1
		}

		if (i < len(board)-k) && (board[i+k][j] == model.StoneBlack) {
			count += 1
		}

		if (i-k > 0) && (board[i-k][j] == model.StoneWhite) {
			count -= 1
		}

		if (i-k > 0) && (board[i-k][j] == model.StoneBlack) {
			count += 1
		}
	}

	return offenseScore(count, targetCount)
}

func dHeatOffense2(i int, j int, targetCount int, board [][]model.Stone) int {
	count := 0
	for k := 1; k <= targetCount; k++ {

		if (i-k > 0) && (j < len(board)-k) && (board[i-k][j+k] == model.StoneBlack) {
			count -= 1
		}

		if (i-k > 0) && (j < len(board)-k) && (board[i-k][j+k] == model.StoneWhite) {
			count += 1
		}

		if (i < len(board)-k) && (j-k > 0) && (board[i+k][j-k] == model.StoneBlack) {
			count -= 1
		}

		if (i < len(board)-k) && (j-k > 0) && (board[i+k][j-k] == model.StoneWhite) {
			count += 1
		}
	}

	return offenseScore(count, targetCount)

}

func NextMove(board_id int) (bestMove model.PiecePos) {
	board := model.Boards[board_id].Board

	tmp := make([][]int, len(board))
	for i := range board {
		tmp[i] = make([]int, len(board))
	}
	// copy(tmp, board)
	maxCount := 0

	// maxScore, bMaxScore, wMaxScore := 0, 0, 0
	maxScore, wMaxScore := 0, 0
	bestMove = model.PiecePos{-1, -1}
	bestMoveScore := model.PiecePos{-1, -1}

	for i := range board {
		for j := range board {

			if board[i][j] == model.StoneEmpty {
				for k := 1; k < 4; k++ {
					tmp[i][j] += hHeatDefense(i, j, k, board)
					tmp[i][j] += vHeatDefense(i, j, k, board)

					tmp[i][j] += dHeat(i, j, k, model.StoneWhite, board)
					tmp[i][j] += dHeat2(i, j, k, model.StoneWhite, board)

					tmp[i][j] += hHeatOffense(i, j, k, board)
					tmp[i][j] += vHeatOffense(i, j, k, board)

					tmp[i][j] += dHeat(i, j, k, model.StoneBlack, board)
					tmp[i][j] += dHeat2(i, j, k, model.StoneBlack, board)
				}

				tmp[i][j] += hPotential(i, j, model.StoneBlack, board)
				tmp[i][j] += hPotential(i, j, model.StoneWhite, board)
				tmp[i][j] += vPotential(i, j, model.StoneBlack, board)
				tmp[i][j] += vPotential(i, j, model.StoneWhite, board)

				tmp[i][j] += dPotential(i, j, model.StoneBlack, board)
				tmp[i][j] += dPotential(i, j, model.StoneWhite, board)
				tmp[i][j] += dPotential2(i, j, model.StoneBlack, board)
				tmp[i][j] += dPotential2(i, j, model.StoneWhite, board)

				tmp[i][j] += threeHorizontal(i, j, model.StoneWhite, board, 200)
				tmp[i][j] += threeHorizontal(i, j, model.StoneBlack, board, 500)
				tmp[i][j] += threeVertical(i, j, model.StoneWhite, board, 200)
				tmp[i][j] += threeVertical(i, j, model.StoneBlack, board, 500)

				tmp[i][j] += threeDiagonal(i, j, model.StoneWhite, board, 200)
				tmp[i][j] += threeDiagonal(i, j, model.StoneBlack, board, 500)
				tmp[i][j] += threeDiagonal2(i, j, model.StoneWhite, board, 200)
				tmp[i][j] += threeDiagonal2(i, j, model.StoneBlack, board, 500)

				if tmp[i][j] > wMaxScore {
					wMaxScore = tmp[i][j]
					bestMoveScore = model.PiecePos{i, j}
				}

			}

		}

	}

	for i := range board {
		fmt.Println(tmp[i])
	}

	// fmt.Println(tmp)
	fmt.Printf("BY MAX SCORE %v - %v\n", bestMoveScore, maxScore)
	fmt.Printf("highest: %v - %v\n", bestMove, maxCount)
	if bestMove.X == -1 {
		return bestMoveScore
	}
	return bestMove
}
