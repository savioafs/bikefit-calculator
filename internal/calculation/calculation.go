package calculation

import "github.com/savioafs/bikefit-calculator/internal/entity"

type Calculation struct {
	measures entity.Measures
}

func NewCalculation(m entity.Measures) Calculation {
	return Calculation{
		measures: m,
	}
}

func Resultcalculation(m entity.Measures) entity.ResultCalculation {
	c := NewCalculation(m)

	straightSeatLenghtSpeed := c.handlebarSeatDistance().Speed + (m.SeatLenght / 2) - m.HorizontalSize
	straightSeatLengthMtb := c.handlebarSeatDistance().MTB + (m.SeatLenght / 2) - m.HorizontalSize
	seatbackLengthSpeed := straightSeatLenghtSpeed - 1.65
	seatbackLengthMtb := (straightSeatLengthMtb - 1.65) + 1.7

	resultCalculation := entity.ResultCalculation{
		MtbSize:                  c.mtbSize(),
		HandlerbarSeatDistance:   c.handlebarSeatDistance(),
		HandlerbarSeatDifference: c.handlerbarSeatDifference(),
		SeatHeight:               c.seatHeight(),
		SeatCrankOffset:          c.seatCrankOffset(),
		CrankLenght:              c.seatCrankLength(),
		HorizontalSpeed:          c.horizontalSpeed(),
		VerticalSpeed:            c.verticalSpeed(),
		StraightSeatLenghtSpeed:  straightSeatLenghtSpeed,
		StraightSeatLengthMtb:    straightSeatLengthMtb,
		SeatbackLengthSpeed:      seatbackLengthSpeed,
		SeatbackLengthMtb:        seatbackLengthMtb,
		HorizontalReal:           c.horizontalReal(),
	}

	return resultCalculation
}

func (c *Calculation) mtbSize() entity.MtbSize {
	inches := c.measures.Inseam * 0.22

	return entity.MtbSize{
		Inches:      inches,
		Centimetres: inches * 2.54,
	}
}

func (c *Calculation) handlebarSeatDistance() entity.HandlerbarSeatDistance {
	flexibilityFactor := 4 * (c.measures.Flexibility - 5) / 10

	return entity.HandlerbarSeatDistance{
		MTB:   (c.measures.Forearm + 5.7) + flexibilityFactor,
		Speed: (c.measures.Forearm + 4) + flexibilityFactor,
	}
}

func (c *Calculation) handlerbarSeatDifference() float32 {
	return (c.measures.Inseam-76.5)/4 + 5.5
}

func (c *Calculation) seatHeight() float32 {
	return c.measures.Inseam * 0.889
}

func (c *Calculation) seatCrankOffset() float32 {
	return (c.measures.Inseam-76.5)/4 + 6
}

func (c *Calculation) seatCrankLength() float32 {
	return ((c.measures.Inseam*1.5)+c.measures.FootSize)*0.279486 + 126.5186
}

func (c *Calculation) horizontalSpeed() float32 {
	return c.measures.Inseam + 6.5
}

func (c *Calculation) verticalSpeed() float32 {
	return c.measures.Inseam * 0.653
}

func (c *Calculation) horizontalReal() float32 {
	return c.measures.HorizontalSize
}
