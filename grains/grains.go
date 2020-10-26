package grains
import (
	"math"
	"errors"
)
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

func Total() uint64{
var grainCount uint64
	for i:= 0;i<64; i++{
		grainCount += uint64(math.Pow(2,float64(i)))
}
return grainCount
}
