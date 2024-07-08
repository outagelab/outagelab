package authservice

type LoginRequest struct {
	AuthProvider string `json:"authProvider"`
	AuthToken    string `json:"authToken"`
}

type LoginResponse struct {
	AuthToken string `json:"authToken"`
}
