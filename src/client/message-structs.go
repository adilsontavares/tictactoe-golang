package client

type msgPlay struct {
	
	Id int `json:"id"`
	Line int `json:"lin"`
	Column int `json:"col"`
}

type msgRequestNewGame struct {

	Id int `json:"id"`
}
