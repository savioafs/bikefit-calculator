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

	straightSeatLenghtSpeed := c.calculateHandlebarSeatDistance().Speed + (m.SeatLenght / 2) - m.HorizontalSize
	straightSeatLengthMtb := c.calculateHandlebarSeatDistance().MTB + +(m.SeatLenght / 2) - m.HorizontalSize
	seatbackLengthSpeed := straightSeatLenghtSpeed - 1.65
	seatbackLengthMtb := (straightSeatLengthMtb - 1.65) + 1.7

	resultCalculation := entity.ResultCalculation{
		MtbSize:                  c.calculateMtbSize(),
		HandlerbarSeatDistance:   c.calculateHandlebarSeatDistance(),
		HandlerbarSeatDifference: c.calculateHandlerbarSeatDifference(),

		StraightSeatLenghtSpeed: straightSeatLenghtSpeed,
		StraightSeatLengthMtb:   straightSeatLengthMtb,
		SeatbackLengthSpeed:     seatbackLengthSpeed,
		SeatbackLengthMtb:       seatbackLengthMtb,
	}

	return resultCalculation
}

func (c *Calculation) calculateMtbSize() entity.MtbSize {
	inches := c.measures.Inseam * 0.22

	return entity.MtbSize{
		Inches:      inches,
		Centimetres: inches * 2.54,
	}
}

func (c *Calculation) calculateHandlebarSeatDistance() entity.HandlerbarSeatDistance {
	flexibilityFactor := 4 * (c.measures.Flexibility - 5) / 10

	return entity.HandlerbarSeatDistance{
		MTB:   (c.measures.Forearm + 5.7) + flexibilityFactor,
		Speed: (c.measures.Forearm + 4) + flexibilityFactor,
	}
}

func (c *Calculation) calculateHandlerbarSeatDifference() float32 {
	return (c.measures.Inseam-76.5)/4 + 5.5
}
