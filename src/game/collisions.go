package game

import (
	"fmt"
	"reflect"
)

func placeMeeting(entity Entity, x int, y int, t reflect.Type) bool {
	for _, e := range EntityList {
		if reflect.TypeOf(e) == t {
			s1 := entity.GetCurrentSprite()
			s2 := e.GetCurrentSprite()
			x2, y2 := e.GetPosition()
			fmt.Println("COLLISION TEST")
			x1World := float64(x) + s1.boundingBox.min[0]
			y1World := float64(y) + s1.boundingBox.min[1]
			x2World := float64(x) + s1.boundingBox.max[0]
			y2World := float64(y) + s1.boundingBox.max[1]

			x3World := float64(x2) + s2.boundingBox.min[0]
			y3World := float64(y2) + s2.boundingBox.min[1]
			x4World := float64(x2) + s2.boundingBox.max[0]
			y4World := float64(y2) + s2.boundingBox.max[1]

			if (x2World >= x3World) && (x1World <= x4World) && (y2World >= y3World) && (y1World <= y4World) {
				fmt.Println("COLLISION DETECTED")
				return true
			}
			//fmt.Println(s1.boundingBox.max, s1.boundingBox.min, x1, y1)
			//fmt.Println(s2.boundingBox.max, s2.boundingBox.min, x2, y2)

		}
	}
	return false
}
