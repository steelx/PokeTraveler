package stage

// Warps ワープポイントのリスト
type Warps struct {
	List []*Warp `json:"warps"`
}

// Warp ワープポイント
type Warp struct {
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Dst   string `json:"dst"`
	DstID int    `json:"dstid"`
	Pos   [2]int `json:"pos"`
	InOut string `json:"inout"`
}
