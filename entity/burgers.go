package entity

type Burger struct {
	ID          interface{} `json:"id,omitempty" bson:"id,omitempty"`
	Name        string      `json:"name,omitempty" bson:"name,omitempty"`
	Restaurant  string      `json:"restaurant,omitempty" bson:`
	Web         string      `json:"web,omitempty" bson:"web,omitempty"`
	Description string      `json:"description,omitempty" bson:"description,omitempty"`
	Ingredients []string    `json:"ingredients,omitempty" bson:"ingredients,omitempty"`
	Addresses   []string    `json:"addresses,omitempty" bson:"addresses,omitempty"`
	Image_url   string      `json:"image_url" bson:"image_url"`
}
