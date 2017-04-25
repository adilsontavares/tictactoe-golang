package server

type msgGameOver struct {

	Id int `json:"id"`
	Winner int `json:"winner"`
}

type msgPlacePlayer struct {

	Id int `json:"id"`
	Player int `json:"player"`
	Line int `json:"lin"`
	Column int `json:"col"`
}

type msgInvalidPosition struct {

	Id int `json:"id"`
	Line int `json:"lin"`
	Column int `json:"col"`
}

type msgRequestPlay struct {
	Id int `json:"id"`
}

type msgNewGame struct {
	Id int `json:"id"`
	Player int `json:"player"`
}
