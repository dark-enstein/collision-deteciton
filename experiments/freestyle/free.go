package freestyle

import (
	"github.com/dark-enstein/collission-detection/engine/object"
	"github.com/dark-enstein/collission-detection/experiments"
	"log"
)

type FreeCollissionManager struct {
	objects []experiments.Object
}

func NewFreeCollissionManager(objs ...experiments.Object) *FreeCollissionManager {
	manager := new(FreeCollissionManager)
	for i := 0; i < len(objs); i++ {
		manager.objects = append(manager.objects, objs[i])
	}
	return manager
}

func (f *FreeCollissionManager) WillCollide(one experiments.Object) bool {
	for i := 0; i < len(f.objects); i++ {
		if one.ID() == f.objects[i].ID() {
			continue
		}
		//bounds := f.objects[i].Bounds()
		//askBounds := one.Bounds()
		////var highX, lowX, highY, lowY float32
		//if askBounds[0].X()
		//for i := 0; i < len(askBounds); i++ {
		//}
	}
	return false
}

func lineCrosses(one [2]*object.Coordinates, two [2]*object.Coordinates) bool {
	return false
}

// CoordCrosses checks if when mover moves by delta it crosses the x or y axis of still
func CoordCrosses(still, mover *object.Coordinates, axis int, delta float32) bool {
	defer mover.Displace(axis, -delta) // test
	switch axis {
	case experiments.AXIS_X:
		log.Println("axis X detected")
		if still.X() > mover.X() {
			fmrmove := mover.X()
			log.Printf("X: coordinate %v is greater than %v\n", still.X(), mover.X())
			afterMove := mover.Displace(experiments.AXIS_X, delta).(*object.Coordinates).X()
			log.Printf("X: after displace, mover = %v | %v => %v", afterMove, fmrmove, afterMove)
			if mover.Displace(experiments.AXIS_X, delta).(*object.Coordinates).X() > still.X() {
				log.Println("displace crossed over")
				return true
			}
		} else if still.X() < mover.X() {
			fmrmove := mover.X()
			log.Printf("X: coordinate %v is less than %v\n", still.X(), mover.X())
			afterMove := mover.Displace(experiments.AXIS_X, delta).(*object.Coordinates).X()
			log.Printf("X: after displace, mover = %v | %v => %v", afterMove, fmrmove, afterMove)
			if afterMove < still.X() {
				log.Println("displace crossed over")
				return true
			}
		}
	case experiments.AXIS_Y:
		log.Println("axis Y detected")
		if still.Y() > mover.Y() {
			fmrmove := mover.Y()
			log.Printf("Y: coordinate %v is greater than %v\n", still.Y(), mover.Y())
			afterMove := mover.Displace(experiments.AXIS_X, delta).(*object.Coordinates).Y()
			log.Printf("Y: after displace, mover = %v | %v => %v", afterMove, fmrmove, afterMove)
			if afterMove > still.Y() {
				log.Println("displace crossed over")
				return true
			}
		} else if still.Y() < mover.Y() {
			fmrmove := mover.Y()
			log.Printf("Y: coordinate %v is less than %v", still.Y(), mover.Y())
			afterMove := mover.Displace(experiments.AXIS_X, delta).(*object.Coordinates).Y()
			log.Printf("Y: after displace, mover = %v | %v => %v", afterMove, fmrmove, afterMove)
			if afterMove < still.Y() {
				log.Println("displace crossed over")
				return true
			}
		}
	}
	return false
}
