package game

//Referenced from https://github.com/hajimehoshi/ebiten/blob/main/examples/camera/main.go
import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type Camera struct {
	ViewPort   f64.Vec2
	Position   f64.Vec2
	ZoomFactor int
	Rotation   int
	target     *Player
}

func (c *Camera) String() string {
	return fmt.Sprintf(
		"T: %.1f, R: %d, S: %d",
		c.Position, c.Rotation, c.ZoomFactor,
	)
}

func (c *Camera) viewportCenter() f64.Vec2 {
	return f64.Vec2{
		c.ViewPort[0] * .5,
		c.ViewPort[1] * .5,
	}
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position[0], -c.Position[1])
	// We want to scale and rotate around center of image / screen
	m.Translate(-c.viewportCenter()[0], -c.viewportCenter()[1])
	m.Scale(
		math.Pow(1.01, float64(c.ZoomFactor)),
		math.Pow(1.01, float64(c.ZoomFactor)),
	)
	m.Rotate(float64(c.Rotation) * 2 * math.Pi / 360)
	m.Translate(c.viewportCenter()[0], c.viewportCenter()[1])
	return m
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = c.worldMatrix()
	screen.DrawImage(world, op)
}

func (c *Camera) ScreenToWorld(posX, posY int) (float64, float64) {
	inverseMatrix := c.worldMatrix()
	if inverseMatrix.IsInvertible() {
		inverseMatrix.Invert()
		return inverseMatrix.Apply(float64(posX), float64(posY))
	} else {
		// When scaling it can happened that matrix is not invertable
		return math.NaN(), math.NaN()
	}
}

func (c *Camera) Target() {
	if c.target != nil && InstanceExists(c.target.getUID()) {
		x, y := c.target.GetPosition()
		currentSprite := c.target.GetCurrentSprite()
		adjustedX := math.Floor(x - float64(currentSprite.originX))
		adjustedY := math.Floor(y - float64(currentSprite.originY))

		targetX := currentSprite.boundingBox.max[0]

		targetY := currentSprite.boundingBox.max[1]
		targetY += adjustedY
		targetX += adjustedX

		TargetPositionX := (targetX - (c.ViewPort[0] / 2))
		TargetPositionY := (targetY - (c.ViewPort[1] / 2))
		//c.Position[0] = math.Floor(TargetPositionX)
		//c.Position[1] = math.Floor(TargetPositionY)
		c.Position[0] += math.Floor(.5 * (TargetPositionX - c.Position[0]))
		c.Position[1] += math.Floor(.5 * (TargetPositionY - c.Position[1]))
	}
}

func (c *Camera) SetTarget(p *Player) {
	c.target = p
}
