package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gravityConst = 6.67430e-11 * 2e8 // Amplified Gravitational Constant for better visualization
	minDistance  = 5.0               // Minimum distance to prevent extreme forces
	maxTrails    = 500               // Max number of trails to show per body
)

// Vector represents a 2D vector or point.
type Vector struct {
	X, Y float64
}

// Add adds another vector to this vector.
func (v *Vector) Add(other Vector) {
	v.X += other.X
	v.Y += other.Y
}

// Sub subtracts another vector from this vector.
func (v *Vector) Sub(other Vector) {
	v.X -= other.X
	v.Y -= other.Y
}

// Mul multiplies the vector by a scalar.
func (v *Vector) Mul(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

// Length computes the magnitude of the vector.
func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize scales the vector to have a length of 1.
func (v *Vector) Normalize() {
	length := v.Length()
	if length != 0 {
		v.X /= length
		v.Y /= length
	}
}

// Body represents a celestial body with mass, position, and velocity.
type Body struct {
	Mass     float64
	Position Vector
	Velocity Vector
	Color    color.Color
	Radius   float64
}

// NewBody creates a new body with the given parameters.
func NewBody(mass float64, position, velocity Vector, clr color.Color) *Body {
	radius := math.Cbrt(mass) / 10 // Scale the radius based on the cube root of the mass
	return &Body{
		Mass:     mass,
		Position: position,
		Velocity: velocity,
		Color:    clr,
		Radius:   radius,
	}
}

// Simulation holds the state of the simulation.
type Simulation struct {
	Bodies []*Body
	Trails [][]Vector
}

// NewSimulation initializes a new simulation with the provided bodies.
func NewSimulation(bodies []*Body) *Simulation {
	trails := make([][]Vector, len(bodies))
	return &Simulation{
		Bodies: bodies,
		Trails: trails,
	}
}

// UpdatePhysics updates the positions and velocities of the bodies.
func (sim *Simulation) UpdatePhysics(dt float64) {
	forces := make([]Vector, len(sim.Bodies))

	// Calculate gravitational forces.
	for i, bodyA := range sim.Bodies {
		force := Vector{0, 0}
		for j, bodyB := range sim.Bodies {
			if i == j {
				continue
			}
			r := Vector{bodyB.Position.X - bodyA.Position.X, bodyB.Position.Y - bodyA.Position.Y}
			distance := r.Length()
			if distance == 0 {
				continue
			}
			r.Normalize()

			// Apply minimum distance to prevent extreme forces.
			if distance < minDistance {
				distance = minDistance
			}
			F := gravityConst * bodyA.Mass * bodyB.Mass / (distance * distance)
			r.Mul(F)
			force.Add(r)
		}
		forces[i] = force
	}

	// Update velocities and positions.
	for i, body := range sim.Bodies {
		acceleration := Vector{forces[i].X / body.Mass, forces[i].Y / body.Mass}
		body.Velocity.Add(Vector{acceleration.X * dt, acceleration.Y * dt})
		body.Position.Add(Vector{body.Velocity.X * dt, body.Velocity.Y * dt})

		// Record trail.
		sim.Trails[i] = append(sim.Trails[i], body.Position)
		if len(sim.Trails[i]) > maxTrails {
			sim.Trails[i] = sim.Trails[i][1:]
		}
	}
}

// DrawCircle is a helper function to draw a filled circle on the screen.
func DrawCircle(screen *ebiten.Image, x, y, radius float64, clr color.Color) {
	for dx := -int(radius); dx <= int(radius); dx++ {
		for dy := -int(radius); dy <= int(radius); dy++ {
			if dx*dx+dy*dy <= int(radius)*int(radius) {
				screen.Set(int(x)+dx, int(y)+dy, clr)
			}
		}
	}
}

// Game implements ebiten.Game interface.
type Game struct {
	sim *Simulation
}

// Update updates the game state.
func (g *Game) Update() error {
	dt := 1.0 / 30.0
	g.sim.UpdatePhysics(dt)
	return nil
}

// Draw renders the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen with black color.
	screen.Fill(color.Black)

	// Draw trails.
	for i, trail := range g.sim.Trails {
		for j := 1; j < len(trail); j++ {
			start := trail[j-1]
			end := trail[j]
			ebitenutil.DrawLine(screen, start.X, start.Y, end.X, end.Y, g.sim.Bodies[i].Color)
		}
	}

	// Draw bodies.
	for _, body := range g.sim.Bodies {
		DrawCircle(screen, body.Position.X, body.Position.Y, body.Radius, body.Color)
	}
}

// Layout specifies the game screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// CreateRandomBody generates a random celestial body with sensible parameters.
func CreateRandomBody() *Body {
	mass := rand.Float64()*9.9e5 + 1e5 // Random mass between 1e5 and 1e6
	position := Vector{rand.Float64() * screenWidth, rand.Float64() * screenHeight}
	velocity := Vector{(rand.Float64() - 0.5) * 50, (rand.Float64() - 0.5) * 50} // Random velocity
	color := color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}

	return NewBody(mass, position, velocity, color)
}

// CreateRandomBodies generates multiple random celestial bodies.
func CreateRandomBodies(numBodies int) []*Body {
	bodies := make([]*Body, numBodies)
	for i := 0; i < numBodies; i++ {
		bodies[i] = CreateRandomBody()
	}
	return bodies
}

// AddRandomBodies adds multiple random celestial bodies to the bodies slice.
func AddRandomBodies(bodies []*Body, numBodies int) []*Body {
	for i := 0; i < numBodies; i++ {
		bodies = append(bodies, CreateRandomBody())
	}
	return bodies
}

func main() {
	bodies := []*Body{
		NewBody(9e6, Vector{400, 300}, Vector{1, 1}, color.RGBA{255, 20, 20, 255}),
	}

	// Add random bodies.
	bodies = AddRandomBodies(bodies, 5)

	sim := NewSimulation(bodies)
	game := &Game{sim: sim}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Random Celestial Bodies Simulation")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
