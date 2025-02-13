package entity

type Measures struct {
	Gender         string  `json:"gender"`
	BikeType       string  `json:"bike-type"`
	Inseam         float32 `json:"inseam"`
	Forearm        float32 `json:"forearm"`
	Flexibility    float32 `json:"flexibility"`
	FootSize       float32 `json:"foot-size"`
	SeatLenght     float32 `json:"seat-length"`
	HorizontalSize float32 `json:"horizontal-size"`
}

type MtbSize struct {
	Inches      float32 `json:"inch"`
	Centimetres float32 `json:"cm"`
}

type HandlerbarSeatDistance struct {
	MTB   float32 `json:"mtb"`
	Speed float32 `json:"speed"`
}

type ResultCalculation struct {
	MtbSize                  MtbSize                `json:"mtb-size"`
	HandlerbarSeatDistance   HandlerbarSeatDistance `json:"handlerbar-seat-distance"`
	HandlerbarSeatDifference float32                `json:"handler-seat-difference"`
	SeatHeight               float32                `json:"seat-height"`
	SeatCrankOffset          float32                `json:"seat-crank-offset"`
	CrankLenght              float32                `json:"crank-lenght"`
	HorizontalSpeed          float32                `json:"horizontal-speed"`
	VerticalSpeed            float32                `json:"vertical-speed"`
	StraightSeatLenghtSpeed  float32                `json:"straight-seat-lenght-speed"`
	StraightSeatLengthMtb    float32                `json:"straight-seat-length-mtb"`
	SeatbackLengthSpeed      float32                `json:"seatback-length-speed"`
	SeatbackLengthMtb        float32                `json:"seatback-length-mtb"`
	HorizontalReal           float32                `json:"horizontal-size"`
}
