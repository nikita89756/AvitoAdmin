package client

type Client struct {
	Price           int `json:"price"`
	LocationId      int `json:"location_id"`
	MicrocategoryId int `json:"microcategory_id"`
	UserId          int `json:"user_id"`
	MatrixId        int `json:"matrix_id"`
	UserSegmentId   int `json:"user_segment_id"`
}
