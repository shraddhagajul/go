package grains
import (
	"math"
	"errors"
)
/*Square computes the number of grains on a square of a chessboard
where the number on each square doubles.
Only valid for the number of squres on a chessboard 1-64*/
func Square(squareNo int) (uint64, error){
	var grainCount float64
	switch {
	case  squareNo < 1 || squareNo > 64:
		return 0,errors.New("Not a valid square")
	case squareNo > 0 || squareNo < 65 :
		grainCount = math.Pow(2,float64(squareNo - 1))	
	}
	return uint64(grainCount),nil
}

/*Total computes to total number of grains on a chessboard
where the number of grains on each successive square doubles*/
func Total() uint64{
var grainCount uint64
	for i:= 0;i<64; i++{
		grainCount += uint64(math.Pow(2,float64(i)))
}
return grainCount
}
