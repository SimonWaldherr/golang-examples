package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth      = 320  // Width of the playing field
	screenHeight     = 240  // Height of the playing field
	gridSize         = 5    // Size of each grid cell
	initialSpeed     = 3    // Frames between opponent moves
	winPercentage    = 75.0 // Percentage to win the game
	initialOpponentX = screenWidth / (2 * gridSize)
	initialOpponentY = screenHeight / (2 * gridSize)
)

// Grid cell states
const (
	Empty          = 0 // Empty cell
	Border         = 1 // Border cell
	Captured       = 2 // Captured area
	Opponent       = 3 // Opponent cell
	PlayerTrail    = 4 // Player trail
	AccessibleArea = 5 // Used during flood fill
)

var gameOver bool

// Directions for movement
var (
	JOYSTICK_UP    = ebiten.KeyArrowUp
	JOYSTICK_DOWN  = ebiten.KeyArrowDown
	JOYSTICK_LEFT  = ebiten.KeyArrowLeft
	JOYSTICK_RIGHT = ebiten.KeyArrowRight
)

// QixGame structure holds the game state
type QixGame struct {
	playerX, playerY       int     // Player position
	opponentX, opponentY   int     // Opponent position
	opponentDX, opponentDY int     // Opponent direction
	occupiedPercentage     float64 // Percentage of the playfield occupied
	grid                   [][]int // Game grid
	prevPlayerPos          int     // Previous player position
	width, height          int     // Grid dimensions
	frameCount             int     // Frame counter
	moveInterval           int     // Frames between opponent moves
}

// Initialize a new game
func NewGame() *QixGame {
	game := &QixGame{
		playerX:            0,
		playerY:            0,
		opponentX:          initialOpponentX,
		opponentY:          initialOpponentY,
		opponentDX:         1,
		opponentDY:         1,
		width:              screenWidth / gridSize,
		height:             screenHeight / gridSize,
		prevPlayerPos:      1,
		occupiedPercentage: 0.0,
		moveInterval:       initialSpeed,
	}

	// Initialize the grid with Empty
	game.grid = make([][]int, game.width)
	for i := range game.grid {
		game.grid[i] = make([]int, game.height)
	}

	game.initializeGame()
	return game
}

// Initialize game by drawing borders, placing player and opponent
func (g *QixGame) initializeGame() {
	g.drawFrame()
	g.placePlayer()
	g.placeOpponent()
}

// Draw frame (border)
func (g *QixGame) drawFrame() {
	for x := 0; x < g.width; x++ {
		g.grid[x][0] = Border
		g.grid[x][g.height-1] = Border
	}
	for y := 0; y < g.height; y++ {
		g.grid[0][y] = Border
		g.grid[g.width-1][y] = Border
	}
}

// Place player at a random position on the edge
func (g *QixGame) placePlayer() {
	edgePositions := []struct{ x, y int }{}
	for x := 0; x < g.width; x++ {
		edgePositions = append(edgePositions, struct{ x, y int }{x, 0})
		edgePositions = append(edgePositions, struct{ x, y int }{x, g.height - 1})
	}
	for y := 1; y < g.height-1; y++ {
		edgePositions = append(edgePositions, struct{ x, y int }{0, y})
		edgePositions = append(edgePositions, struct{ x, y int }{g.width - 1, y})
	}

	pos := edgePositions[rand.Intn(len(edgePositions))]
	g.playerX, g.playerY = pos.x, pos.y
}

// Place opponent at random position inside the playfield
func (g *QixGame) placeOpponent() {
	for {
		x := rand.Intn(g.width-2) + 1
		y := rand.Intn(g.height-2) + 1
		if g.grid[x][y] == Empty {
			g.opponentX, g.opponentY = x, y
			g.grid[g.opponentX][g.opponentY] = Opponent
			break
		}
	}
}

// Update is called on every frame update
func (g *QixGame) Update() error {
	if gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			*g = *NewGame()
			gameOver = false
		}
		return nil
	}

	g.frameCount++

	// Move opponent every 'moveInterval' frames
	if g.frameCount%g.moveInterval == 0 {
		g.moveOpponent()
	}

	// Move player based on input
	g.movePlayer()

	// Check win condition
	if g.occupiedPercentage >= winPercentage {
		gameOver = true
	}

	return nil
}

// Draw the game frame
func (g *QixGame) Draw(screen *ebiten.Image) {
	// Clear screen
	screen.Fill(color.Black)

	// Draw grid and game objects
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			switch g.grid[x][y] {
			case Border:
				ebitenutil.DrawRect(screen, float64(x*gridSize), float64(y*gridSize), gridSize, gridSize, color.RGBA{0, 0, 255, 255})
			case Captured:
				ebitenutil.DrawRect(screen, float64(x*gridSize), float64(y*gridSize), gridSize, gridSize, color.RGBA{0, 0, 255, 255})
			case Opponent:
				ebitenutil.DrawRect(screen, float64(x*gridSize), float64(y*gridSize), gridSize, gridSize, color.RGBA{255, 0, 0, 255})
			case PlayerTrail:
				ebitenutil.DrawRect(screen, float64(x*gridSize), float64(y*gridSize), gridSize, gridSize, color.RGBA{0, 255, 0, 255})
			}
		}
	}

	// Draw player
	ebitenutil.DrawRect(screen, float64(g.playerX*gridSize), float64(g.playerY*gridSize), gridSize, gridSize, color.RGBA{0, 255, 0, 255})

	// Display score (occupied percentage)
	scoreMsg := fmt.Sprintf("Occupied: %.2f%%", g.occupiedPercentage)
	ebitenutil.DebugPrintAt(screen, scoreMsg, 10, screenHeight-20)

	// Display win message
	if g.occupiedPercentage >= winPercentage {
		msg := "YOU WIN! Press 'Space' to Restart"
		ebitenutil.DebugPrintAt(screen, msg, screenWidth/2-100, screenHeight/2-10)
	}

	// Display game over message
	if gameOver && g.occupiedPercentage < winPercentage {
		msg := "GAME OVER! Press 'Space' to Restart"
		ebitenutil.DebugPrintAt(screen, msg, screenWidth/2-100, screenHeight/2-10)
	}
}

// Layout defines the game window size
func (g *QixGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Move the player based on keyboard input
func (g *QixGame) movePlayer() {
	var newX, newY int = g.playerX, g.playerY
	if ebiten.IsKeyPressed(JOYSTICK_UP) {
		newY--
	} else if ebiten.IsKeyPressed(JOYSTICK_DOWN) {
		newY++
	} else if ebiten.IsKeyPressed(JOYSTICK_LEFT) {
		newX--
	} else if ebiten.IsKeyPressed(JOYSTICK_RIGHT) {
		newX++
	} else {
		return // No movement key pressed
	}

	// Ensure player stays within the bounds
	if newX >= 0 && newX < g.width && newY >= 0 && newY < g.height {
		cellValue := g.grid[newX][newY]
		switch cellValue {
		case Empty:
			// Moving into empty space, mark trail
			g.grid[newX][newY] = PlayerTrail
			g.prevPlayerPos = 0
		case Border, Captured:
			if g.prevPlayerPos == 0 {
				g.closeArea(newX, newY)
			}
			g.prevPlayerPos = 1
		case Opponent, PlayerTrail:
			// Collision, game over
			gameOver = true
			return
		}
		g.playerX, g.playerY = newX, newY
	}
}

// Close the area when player reconnects to the border or their trail
func (g *QixGame) closeArea(x, y int) {
	// Convert all player trails to permanent borders
	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height; j++ {
			if g.grid[i][j] == PlayerTrail {
				g.grid[i][j] = Border
			}
		}
	}

	// Perform flood fill from the opponent's position
	g.floodFill(g.opponentX, g.opponentY)

	// Mark areas not accessible from the opponent as captured
	for i := 0; i < g.width; i++ {
		for j := 0; j < g.height; j++ {
			switch g.grid[i][j] {
			case Empty:
				g.grid[i][j] = Captured
			case AccessibleArea:
				g.grid[i][j] = Empty
			}
		}
	}

	// Re-mark the opponent's position
	g.grid[g.opponentX][g.opponentY] = Opponent

	// Recalculate occupied percentage
	g.calculateOccupiedPercentage()
}

// Flood fill to determine the opponent's accessible area
func (g *QixGame) floodFill(x, y int) {
	queue := []struct{ x, y int }{{x, y}}
	visited := make(map[[2]int]bool)
	visited[[2]int{x, y}] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		directions := []struct{ dx, dy int }{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		}

		for _, dir := range directions {
			nx, ny := current.x+dir.dx, current.y+dir.dy
			if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height {
				if !visited[[2]int{nx, ny}] && (g.grid[nx][ny] == Empty || g.grid[nx][ny] == Opponent) {
					g.grid[nx][ny] = AccessibleArea
					visited[[2]int{nx, ny}] = true
					queue = append(queue, struct{ x, y int }{nx, ny})
				}
			}
		}
	}
}

// Calculate how much of the playfield has been occupied
func (g *QixGame) calculateOccupiedPercentage() {
	total := g.width * g.height
	occupied := 0
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			if g.grid[x][y] == Captured {
				occupied++
			}
		}
	}
	g.occupiedPercentage = (float64(occupied) / float64(total)) * 100
}

// Move the opponent and handle collisions
func (g *QixGame) moveOpponent() {
	nextX := g.opponentX + g.opponentDX
	nextY := g.opponentY + g.opponentDY

	// Check for collision with borders or captured areas
	if g.grid[nextX][nextY] == Border || g.grid[nextX][nextY] == Captured {
		// Reverse direction upon collision
		if g.grid[nextX][g.opponentY] == Border || g.grid[nextX][g.opponentY] == Captured {
			g.opponentDX = -g.opponentDX
		}
		if g.grid[g.opponentX][nextY] == Border || g.grid[g.opponentX][nextY] == Captured {
			g.opponentDY = -g.opponentDY
		}
		nextX = g.opponentX + g.opponentDX
		nextY = g.opponentY + g.opponentDY
	}

	// Check for collision with player trail or player
	if g.grid[nextX][nextY] == PlayerTrail || (nextX == g.playerX && nextY == g.playerY) {
		gameOver = true
		return
	}

	// Clear current position
	g.grid[g.opponentX][g.opponentY] = Empty

	// Update position
	g.opponentX = nextX
	g.opponentY = nextY

	// Set new position
	g.grid[g.opponentX][g.opponentY] = Opponent
}

// Main loop
func (g *QixGame) Run() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Qix/Xonix Game")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	game.Run()
}
