package models

type ChangePasswordRequest struct {
	OldPassword             string `json:oldPassword`
	NewPassword             string `json:newPassword`
	NewPasswordConfirmation string `json:newPasswordConfirmation`
}
