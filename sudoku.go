package main
import (
    "fmt"
)
const gridSize = 9
// Function to solve the Sudoku puzzle
func solveSudoku(grid [][]int) bool {
    row, col, found := findEmptyCell(grid)
    if !found {
        // No empty cell found, the puzzle is solved
        return true
    }
    for num := 1; num <= gridSize; num++ {
        if isSafe(grid, row, col, num) {
            // If it's safe to place the number, assign it and recursively solve
            grid[row][col] = num
            // Continue solving the rest of the puzzle
            if solveSudoku(grid) {
                return true
            }
            // If placing the number leads to an invalid solution, backtrack
            grid[row][col] = 0
        }
    }
    // If no number can be placed, backtrack to the previous cell
    return false
}
// Function to find an empty cell in the grid
func findEmptyCell(grid [][]int) (int, int, bool) {
    for row := 0; row < gridSize; row++ {
        for col := 0; col < gridSize; col++ {
            if grid[row][col] == 0 {
                return row, col, true
            }
        }
    }
    return 0, 0, false
}
// Function to check if it's safe to place a number in a cell
func isSafe(grid [][]int, row, col, num int) bool {
    // Check if the number is not in the same row or column
    for i := 0; i < gridSize; i++ {
        if grid[row][i] == num || grid[i][col] == num {
            return false
        }
    }
    // Check if the number is not in the same 3x3 subgrid
    startRow, startCol := row-row%3, col-col%3
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[startRow+i][startCol+j] == num {
                return false
            }
        }
    }
    return true
}
// Function to print the Sudoku grid
func printGrid(grid [][]int) {
    for row := 0; row < gridSize; row++ {
        for col := 0; col < gridSize; col++ {
            fmt.Printf("%d ", grid[row][col])
        }
        fmt.Println()
    }
}
func main() {
    input := []string{
			".96.4...1",
			"1...6...4",
			"5.481.39.",
			"..795..43",
			".3..8....",
			"4.5.23.18",
			".1.63..59",
			".59.7.83.",
			"..359...7",
    }
    grid := parseSudoku(input)
    if grid == nil {
        fmt.Println("Invalid input")
        return
    }
    if solveSudoku(grid) {
        fmt.Println("Solved Sudoku:")
        printGrid(grid)
    } else {
        fmt.Println("No solution exists")
    }
}
// Function to parse the Sudoku input into a 2D grid
func parseSudoku(input []string) [][]int {
    if len(input) != gridSize {
        return nil // Invalid input, must have exactly 9 rows
    }
    grid := make([][]int, gridSize)
    for i, row := range input {
        if len(row) != gridSize {
            return nil // Invalid row length
        }
        grid[i] = make([]int, gridSize)
        for j, char := range row {
            if char == '.' {
                grid[i][j] = 0
            } else if char >= '1' && char <= '9' {
                grid[i][j] = int(char - '0')
            } else {
                return nil // Invalid character in the input
            }
        }
    }
    return grid
}
