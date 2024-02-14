package user

type Model struct {
	PhoneNumber string   `json:"phone_number"`
	Address     string   `json:"address"`
	Friends     []string `json:"friends"`
	Name        string   `json:"name"`
	CountryCode string   `json:"country_code"`
}
