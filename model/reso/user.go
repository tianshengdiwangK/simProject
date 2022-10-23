package reso

// GetUser GET "/user" response object
type GetUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Gender   int64  `json:"gender"`
	Age      int64  `json:"age"`
	Interest string `json:"interest"`
}
