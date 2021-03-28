package models

type Burger struct {
	ID          interface{} `json:"id,omitempty"`
	Name        string      `json:"name"`
	Restaurant  string      `json:"restaurant"`
	Web         string      `json:"web"`
	Description string      `json:"description"`
	Ingredients []string    `json:"ingredients"`
	Addresses   []string    `json:"addresses"`
}

//  {
// 	"id": 0,
// 	"name": "Tribute Burger",
// 	"restaurant": "Honest Burgers",
// 	"web": "www.honestburgers.co.uk",
// 	"description": "A mouth-watering honest beef burger",
// 	"ingredients": [
// 		"beef",
// 		"american cheese",
// 		"burger sauce",
// 		"french mustard",
// 		"pickes",
// 		"onion",
// 		"lettuce"
// 	],
// 	"addresses": [
// 		{
// 			"addressId": 0,
// 			"number": "75",
// 			"line1": "Venn Street",
// 			"line2": "Clapham",
// 			"postcode": "SW4 0BD",
// 			"country": "United Kingdom"
// 		}
// 	]
// },2.
