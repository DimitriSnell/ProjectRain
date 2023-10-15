package game

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/math/f64"
)

type AABB struct {
	min f64.Vec2
	max f64.Vec2
}

func newAABB(min, max f64.Vec2) AABB {
	box := AABB{min, max}
	return box
}

type Sprite struct {
	img             *ebiten.Image
	numOfFrames     int
	framesPerSecond int
	ImgWidth        int
	ImgHeight       int
	Name            string
	count           int
	boundingBox     *AABB
	originX         int
	originY         int
	spriteIndex     int
	imageXScale     int
}

func NewSprite(filename string, numOfFrames int, framesPerSecond int, imgWidth int, imgHeigt int, name string, min, max f64.Vec2, originX, originY int) *Sprite {
	var image *ebiten.Image
	image = nil
	if filename != "" {
		img, _, err := ebitenutil.NewImageFromFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		image = img
	}

	//boundingBox := calculateSingleAABBFromImage(AABBImg)
	//fmt.Println(boundingBox.min, boundingBox.max)
	bb := newAABB(min, max)
	s := Sprite{image, numOfFrames, framesPerSecond, imgWidth, imgHeigt, name, 0, &bb, originX, originY, 0, 1}
	return &s
}

//func calculateSingleAABBFromImage(img *ebiten.Image) AABB {
/*minPoint := f64.Vec2{0, 0}
maxPoint := f64.Vec2{0, 0}

rgbaImage := image.NewRGBA(img.Bounds())
ebiten
for x := 0; x < img.Bounds().Dx(); x++ {
	for y := 0; y < img.Bounds().Dy(); y++ {
		r, _, _, _ := img.At(x, y).RGBA()
		if r > 0 {
			minPoint[0] = math.Min(minPoint[0], float64(x))
			minPoint[1] = math.Min(minPoint[1], float64(y))
			maxPoint[0] = math.Max(maxPoint[0], float64(x))
			maxPoint[1] = math.Max(maxPoint[1], float64(y))
		}
	}
}
result := newAABB(minPoint, maxPoint)
return result*/
//}

func (s *Sprite) Step() {

	s.count++
	i := (s.count / s.framesPerSecond) % int(s.numOfFrames)
	s.spriteIndex = i
}

func (s *Sprite) Draw(img *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	adjustedX := math.Floor(x - float64(s.originX))
	adjustedY := math.Floor(y - float64(s.originY))
	op.GeoM.Translate(adjustedX*float64(s.imageXScale), adjustedY)
	op.GeoM.Scale(float64(s.imageXScale), 1)
	if s.imageXScale == -1 {
		op.GeoM.Translate(math.Round(s.boundingBox.max[0]+s.boundingBox.min[0]), 0)
	}
	//op.GeoM.Translate(500, 0)
	i := s.spriteIndex
	//fmt.Println(s.spriteIndex)
	//fmt.Println(i)
	sx, sy := i*int(s.ImgWidth), 0
	minPointX := adjustedX + s.boundingBox.min[0]
	minPointY := adjustedY + s.boundingBox.min[1]
	maxPointX := adjustedX + s.boundingBox.max[0]
	maxPointY := adjustedY + s.boundingBox.max[1]
	//fmt.Println(x, y)
	//fmt.Println(s.boundingBox.min, s.boundingBox.max)
	//fmt.Println(minPointX, minPointY, maxPointX, maxPointY)
	//vector.DrawFilledRect(img, float32(minPointX), float32(minPointY), float32(maxPointX), float32(maxPointY), color.RGBA{255, 0, 0, 0xff}, false)
	//ebitenutil.DrawRect(img, minPointX, minPointY, maxPointX, maxPointY, color.RGBA{255, 0, 0, 0xff})
	if s.img != nil {
		img.DrawImage(s.img.SubImage(image.Rect(sx, sy, sx+int(s.ImgWidth), sy+int(s.ImgWidth))).(*ebiten.Image), op)
		//img.DrawImage(s.img, op)
	}
	vector.StrokeLine(img, float32(minPointX), float32(minPointY), float32(maxPointX), float32(minPointY), 1, color.RGBA{255, 0, 0, 0xff}, false)
	vector.StrokeLine(img, float32(minPointX), float32(minPointY), float32(minPointX), float32(maxPointY), 1, color.RGBA{255, 0, 0, 0xff}, false)
	vector.StrokeLine(img, float32(minPointX), float32(maxPointY), float32(maxPointX), float32(maxPointY), 1, color.RGBA{255, 0, 0, 0xff}, false)
	vector.StrokeLine(img, float32(maxPointX), float32(minPointY), float32(maxPointX), float32(maxPointY), 1, color.RGBA{255, 0, 0, 0xff}, false)

	//img.DrawImage(s.img.SubImage(image.Rect(sx, sy, sx+64, sy+64)).(*ebiten.Image), op)
}
