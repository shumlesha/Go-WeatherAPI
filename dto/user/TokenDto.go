package user

type TokenDto struct {
	Token string `json:"token"`
}

func NewTokenDto(token string) TokenDto {
	return TokenDto{
		Token: token,
	}
}
