package user

import "app/pkg/platform/middleware"

type UserInfoResponse struct {
	Name          string `json:"name"`
	GivenName     string `json:"givenName"`
	FamilyName    string `json:"familyName"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
	Username      string `json:"username"`
}

func ToUserInfoResponse(tokenContext middleware.TokenContext) UserInfoResponse {
	return UserInfoResponse{
		Name:          tokenContext.Name,
		GivenName:     tokenContext.GivenName,
		FamilyName:    tokenContext.FamilyName,
		Email:         tokenContext.Email,
		EmailVerified: tokenContext.EmailVerified,
		Username:      tokenContext.PreferredUsername,
	}
}
