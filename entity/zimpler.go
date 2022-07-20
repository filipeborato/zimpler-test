package entity

type TopRate []StoreResp

type StoreResp struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnack"`
}
