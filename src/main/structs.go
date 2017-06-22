package main


type ResponseGetAll struct {
	Status bool   `json:"status"`
	Blocks  []Block `json:"blocks"`
	Rooms [] Room `json:"rooms"`
}

type RequestForUpdate struct {
	Blocks  []Block `json:"blocks"`
	Rooms [] Room `json:"rooms"`
}

type ResponsePost struct {
	Status bool `json:"status"`
}

type Block struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Room struct
{
	Id int `json:"id"`
	Block_id int `json:"block_id"`
	Number int `json:"number"`
	Price int `json:"price"`
	Free int `json:"free"`
	Water int `json:"water"`
	Date string `json:"date"`
	
}
