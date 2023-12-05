package common

type Message struct {
	Code string `json:"code"`
	Data string `json:"message"`
}

type Room struct {
	RoomId          int
	RoomName        string
	RoomStatus      int
	RoomType        int
	RoomOwner       string
	RoomClientCount int
}
