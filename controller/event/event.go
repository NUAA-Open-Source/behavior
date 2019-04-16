package event

type CreateRequest struct {
	Name string `json:"name"`
	Src  string `json:"src"`
}
