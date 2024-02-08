package freestyle

import (
	"github.com/dark-enstein/collission-detection/engine/object"
	"github.com/dark-enstein/collission-detection/experiments"
	"github.com/davecgh/go-spew/spew"
	"log"
	"math/rand"
	"testing"
)

type CoordCrossCases struct {
	still    *object.Coordinates
	mover    *object.Coordinates
	expected bool
	delta    float32
	axis     int
}

func TestCoordCrosses(t *testing.T) {
	//testCoordCrosses := []CoordCrossCases{
	//		object.NewCoordinates(3, 4),
	//		object.NewCoordinates(5, 6),
	//		true
	//	},
	//	{
	//		object.NewCoordinates(3, 4),
	//		object.NewCoordinates(6, 6),
	//	},
	//	{
	//		object.NewCoordinates(3, 4),
	//		object.NewCoordinates(7, 6),
	//	},
	//	{
	//		object.NewCoordinates(150, 4),
	//		object.NewCoordinates(101, 6),
	//	},
	//	{
	//		object.NewCoordinates(31000, 4),
	//		object.NewCoordinates(3111, 6),
	//	},
	//	{
	//		object.NewCoordinates(55, 4),
	//		object.NewCoordinates(57, 6),
	//	},
	//	{
	//		object.NewCoordinates(3, 4),
	//		object.NewCoordinates(5, 6),
	//	},
	//}
	testGenerated := gentestCoordCrosses(13)
	log.Printf("Generated: %v\n", spew.Sdump(testGenerated))
	//if true != false {
	//	t.Errorf("expected %v, got %v", true, false)
	//}

	for i := 0; i < len(testGenerated); i++ {
		res := CoordCrosses(testGenerated[i].still, testGenerated[i].mover, testGenerated[i].axis, testGenerated[i].delta)
		if res != testGenerated[i].expected {
			t.Errorf("expected %v, got %v", testGenerated[i].expected, res)
		}
	}
}

func gentestCoordCrosses(num int) []*CoordCrossCases {
	var cc []*CoordCrossCases
	axis := experiments.AXIS_X
	for i := 0; i < num; i++ {
		if i%2 == 0 {
			log.Println("odd")
			axis = experiments.AXIS_Y
		}
		var expected bool
		sX, sY, mX, mY, delta := rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()
		if ((mX > sX) && (mX+delta < sX)) || ((mX < sX) && (mX+delta > sX)) {
			expected = true
		}
		if ((mY > sY) && (mY+delta < sY)) || ((mY < sY) && (mY+delta > sY)) {
			expected = true
		}
		cc = append(cc, &CoordCrossCases{
			object.NewCoordinates(sX, sY),
			object.NewCoordinates(mX, mY),
			expected,
			delta,
			axis,
		})
	}
	return cc
}
