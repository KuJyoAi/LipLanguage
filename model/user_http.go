package model

type RegisterParam struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginParam struct {
	Account  string `json:"account,required"`
	Password string `json:"password,required"`
}

type LoginResponse struct {
	Token         string `json:"token"`
	Name          string `json:"username"`
	Nickname      string `json:"nickname"`
	AvatarUrl     string `json:"avatar_url"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	UserID        int64  `json:"user_id"`
	HearingDevice bool   `json:"hearing_device"`
	Gender        int    `json:"gender"`
	BirthDay      string `gorm:"birthday"`
}
type UpdateInfoParam struct {
	Nickname string `json:"nickname"`
	//AvatarID      string `json:"avatar_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Birthday      string `json:"birthday"`
	Gender        int    `json:"gender"`
	HearingDevice bool   `json:"hearing_device"`
}

type UserVerifyParam struct {
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type ResetPasswordParam struct {
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}
type UpdatePhoneParam struct {
	Phone string `json:"phone,omitempty"`
}

type UpdatePasswordParam struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
