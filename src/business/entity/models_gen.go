// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package entity

type Mutation struct {
}

type QlCreateUserParam struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	DisplayName     string `json:"displayName"`
}

type QlLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type QlUser struct {
	ID          int    `json:"id"`
	Roleid      int    `json:"roleid"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
}

type QlUserLoginResponse struct {
	Email        string `json:"email"`
	DisplayName  string `json:"displayName"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Query struct {
}
