package registration

type Permissions struct {
	Read   bool `json:"read"`
	Write  bool `json:"write"`
	Delete bool `json:"delete"`
}

type User struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
	Salt           string `json:"salt"`
	Role           string `json:"role"`
}

type UserCollection struct {
	Users []User `json:"users"`
}
