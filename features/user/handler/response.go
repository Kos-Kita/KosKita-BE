package handler

import "KosKita/features/user"

type UserResponse struct {
	ID uint `json:"name" form:"name"`
	Name         string `json:"name" form:"name"`
	UserName     string `json:"user_name" form:"user_name"`
	Email        string `json:"email" form:"email"`
	Gender       string `json:"gender" form:"gender"`
	Role         string `json:"role" form:"role"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

type UserKosDetailResponse struct {
	Name         string `json:"name" form:"name"`
	UserName     string `json:"user_name" form:"user_name"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

func CoreToResponse(data *user.Core) UserResponse {
	var result = UserResponse{
		ID: data.ID,
		Name:         data.Name,
		UserName:     data.UserName,
		Email:        data.Email,
		Gender:       data.Gender,
		Role:         data.Role,
		PhotoProfile: data.PhotoProfile,
	}
	return result
}
