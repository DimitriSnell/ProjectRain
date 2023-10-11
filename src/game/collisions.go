package game

import (
	"reflect"
)

func placeMeeting(entity Entity, x float64, y float64, t reflect.Type) bool {

	for _, e := range EntityList {
		if reflect.TypeOf(e) == t {
			s1 := entity.GetCurrentSprite()
			s2 := e.GetCurrentSprite()
			adjustedX1 := x - float64(s1.originX)
			adjustedY1 := y - float64(s1.originY)
			x2, y2 := e.GetPosition()
			adjustedX2 := x2 - float64(s2.originX)
			adjustedY2 := y2 - float64(s2.originY)
			//fmt.Println("COLLISION TEST")
			x1World := float64(adjustedX1) + s1.boundingBox.min[0]
			y1World := float64(adjustedY1) + s1.boundingBox.min[1]
			x2World := float64(adjustedX1) + s1.boundingBox.max[0]
			y2World := float64(adjustedY1) + s1.boundingBox.max[1]

			x3World := float64(adjustedX2) + s2.boundingBox.min[0]
			y3World := float64(adjustedY2) + s2.boundingBox.min[1]
			x4World := float64(adjustedX2) + s2.boundingBox.max[0]
			y4World := float64(adjustedY2) + s2.boundingBox.max[1]

			if (x2World >= x3World) && (x1World <= x4World) && (y2World >= y3World) && (y1World <= y4World) {
				//fmt.Println("COLLISION DETECTED")
				return true
			}
			//fmt.Println(s1.boundingBox.max, s1.boundingBox.min, x1, y1)
			//fmt.Println(s2.boundingBox.max, s2.boundingBox.min, x2, y2)

		}
	}
	return false
}
