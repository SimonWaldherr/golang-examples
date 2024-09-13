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
	screenWidth  = 320 // Width of the playing field
	screenHeight = 240 // Height of the playing field
	gridSize     = 10  // Size of each segment of the snake and the food
	initialSpeed = 8   // Initial speed of the snake
	margin       = 2   // Margin to prevent food from appearing at the edge
)

type Game struct {
	snake      []Position // Position of the snake
	direction  Direction  // Current movement direction
	food       Position   // Position of the food
	score      int        // Player's score
	gameOver   bool       // Game status
	speed      int        // Game speed
	frameCount int        // Counts frames to control the speed
}

type Position struct {
	X int
	Y int
}

type Direction struct {
	X int
	Y int
}

// Layout defines the size of the game window
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Update contains the game logic that is updated every frame
func (g *Game) Update() error {
	// If "Escape" is pressed, the game ends
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("Game ended")
	}

	// If the game is over, it can be restarted by pressing the spacebar
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			*g = *NewGame() // Reset the game
		}
		return nil
	}

	// Process input (change direction)
	g.handleInput()

	// Count frames and control snake movement
	g.frameCount++
	if g.frameCount%g.speed != 0 {
		return nil
	}

	// Move the snake
	g.moveSnake()

	// Check for collisions (with the wall or itself)
	g.checkCollisions()

	return nil
}

// handleInput processes the player's keyboard input
func (g *Game) handleInput() {
	// Change direction, but prevent the snake from reversing directly
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.direction.Y == 0 {
		g.direction = Direction{X: 0, Y: -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.direction.Y == 0 {
		g.direction = Direction{X: 0, Y: 1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.direction.X == 0 {
		g.direction = Direction{X: -1, Y: 0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.direction.X == 0 {
		g.direction = Direction{X: 1, Y: 0}
	}
}

// moveSnake moves the snake and checks if it has reached the food
func (g *Game) moveSnake() {
	// New head position based on the current direction
	newHead := Position{
		X: g.snake[0].X + g.direction.X*gridSize,
		Y: g.snake[0].Y + g.direction.Y*gridSize,
	}

	// Add the new head to the snake
	g.snake = append([]Position{newHead}, g.snake...)

	// Check if the snake has reached the food
	if newHead == g.food {
		g.score++ // Increase score
		g.spawnFood()

		// Increase speed every 5 points until a limit is reached
		if g.score%5 == 0 && g.speed > 2 {
			g.speed--
		}
	} else {
		// Remove the last tail segment to move the snake
		g.snake = g.snake[:len(g.snake)-1]
	}
}

// checkCollisions checks if the snake collides with the wall or itself
func (g *Game) checkCollisions() {
	head := g.snake[0]

	// Collision with the wall
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth || head.Y >= screenHeight {
		g.gameOver = true
	}

	// Collision with itself
	for _, segment := range g.snake[1:] {
		if head == segment {
			g.gameOver = true
			break
		}
	}
}

// Draw draws the snake, food, and score on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Set background to black
	screen.Fill(color.Black)

	// Draw the snake (white color)
	for _, segment := range g.snake {
		ebitenutil.DrawRect(screen, float64(segment.X), float64(segment.Y), gridSize, gridSize, color.White)
	}

	// Draw the food (red color)
	ebitenutil.DrawRect(screen, float64(g.food.X), float64(g.food.Y), gridSize, gridSize, color.RGBA{255, 0, 0, 255})

	// Display score and game-over message
	if g.gameOver {
		msg := fmt.Sprintf("Game Over! Score: %d\nPress 'Space' to restart.", g.score)
		ebitenutil.DebugPrintAt(screen, msg, screenWidth/2-80, screenHeight/2)
	} else {
		msg := fmt.Sprintf("Score: %d", g.score)
		ebitenutil.DebugPrint(screen, msg)
	}
}

// spawnFood generates a new food object at a random valid position
func (g *Game) spawnFood() {
	rand.Seed(time.Now().UnixNano())

	// Random position within the playing field, but not directly at the edge (with margin)
	for {
		g.food = Position{
			X: rand.Intn(screenWidth/gridSize-2*margin)*gridSize + margin*gridSize,
			Y: rand.Intn(screenHeight/gridSize-2*margin)*gridSize + margin*gridSize,
		}

		// Ensure the food doesn't appear on the snake
		collision := false
		for _, segment := range g.snake {
			if g.food == segment {
				collision = true
				break
			}
		}
		// If there's no collision, the position is valid
		if !collision {
			break
		}
	}
}

// NewGame initializes a new game
func NewGame() *Game {
	return &Game{
		snake: []Position{{
			X: screenWidth / 2,
			Y: screenHeight / 2},
		}, // Start the snake in the middle
		direction: Direction{X: 0, Y: -1}, // Start moving upwards
		speed:     initialSpeed,           // Initial speed
		food: Position{
			X: rand.Intn(screenWidth/gridSize-2*margin)*gridSize + margin*gridSize,
			Y: rand.Intn(screenHeight/gridSize-2*margin)*gridSize + margin*gridSize,
		},
	}
}

// main starts the game
func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2) // Window size
	ebiten.SetWindowTitle("Snake Game")                 // Window title

	// Run the game
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
