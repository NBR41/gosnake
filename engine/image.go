package engine

//ImageType type for snake body image
type ImageType int

// List of type of body parts
const (
	HeadNorth      ImageType = iota // 0
	HeadEast                        // 1
	HeadSouth                       // 2
	HeadWest                        // 3
	TailNorth                       // 4
	TailEast                        // 5
	TailSouth                       // 6
	TailWest                        // 7
	BodyHorizontal                  // 8
	BodyVertical                    // 9
	BodyNorthWest                   // 10
	BodyNorthEast                   // 11
	BodySouthEast                   // 12
	BodySouthWest                   // 13
)

//BodyPart struct to represent a body part
type BodyPart struct {
	img ImageType
	pos *Position
}

func newBodyPart(img ImageType, pos Position) *BodyPart {
	return &BodyPart{img: img, pos: &pos}
}

func getHeadImageType(dir Direction) ImageType {
	switch dir {
	case North:
		return HeadNorth
	case East:
		return HeadEast
	case South:
		return HeadSouth
	default:
		return HeadWest
	}
}

func getTailImageType(dir Direction) ImageType {
	switch dir {
	case North:
		return TailNorth
	case East:
		return TailEast
	case South:
		return TailSouth
	default:
		return TailWest
	}
}

func getBodyImageType(dir Direction) ImageType {
	if dir == North || dir == South {
		return BodyVertical
	}
	return BodyHorizontal
}

func getCurveBodyImageType(d1, d2 Direction) ImageType {
	switch {
	case d1 == East && d2 == South:
		return BodySouthWest
	case d1 == North && d2 == West:
		return BodySouthWest
	case d1 == East && d2 == North:
		return BodyNorthWest
	case d1 == South && d2 == West:
		return BodyNorthWest
	case d1 == West && d2 == South:
		return BodySouthEast
	case d1 == North && d2 == East:
		return BodySouthEast
	case d1 == West && d2 == North:
		return BodyNorthEast
	case d1 == South && d2 == East:
		return BodyNorthEast
	default:
		return BodyHorizontal
	}
}
