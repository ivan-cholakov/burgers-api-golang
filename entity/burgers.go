package entity

type Burger struct {
	ID          interface{} `json:"id,omitempty"`
	Name        string      `json:"name"`
	Restaurant  string      `json:"restaurant"`
	Web         string      `json:"web"`
	Description string      `json:"description"`
	Ingredients []string    `json:"ingredients"`
	Addresses   []string    `json:"addresses"`
	Image_url   string      `json:"image_url"`
}
