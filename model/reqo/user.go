package reqo

// GetUser GET "/user" request object
type GetUser struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
}
