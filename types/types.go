package types

type X_LOGIN_TOKEN struct {
}

type LOGIN struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TOKEN struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
