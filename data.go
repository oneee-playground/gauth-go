package gauth

type gender string

const (
	GenderMale   gender = "MALE"
	GenderFemale gender = "FEMALE"
)

type role string

const (
	RoleStudent  role = "ROLE_STUDENT"
	RoleTeacher  role = "ROLE_TEACHER"
	RoleGraduate role = "ROLE_GRADUATE"
)

type UserInfo struct {
	Email      string
	Name       *string
	Grade      *int
	ClassNum   *int
	Num        *int
	Gender     gender
	ProfileURL *string
	Role       role
}

type issueCodeRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type codeResponse struct {
	Code string `json:"code"`
}

type issueTokenRequest struct {
	Code         string `json:"code"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	RedirectURI  string `json:"redirectUri"`
}

type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type userInfoResponse struct {
	Email      string  `json:"email"`
	Name       *string `json:"name"`
	Grade      *int    `json:"grade"`
	ClassNum   *int    `json:"classNum"`
	Num        *int    `json:"num"`
	Gender     gender  `json:"gender"`
	ProfileURL *string `json:"profileUrl"`
	Role       role    `json:"role"`
}
