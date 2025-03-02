package dto

import "mime/multipart"

type UpdateUserRequest struct {
	FullName string                `form:"full_name"`
	Phone    string                `form:"phone"`
	Photo    *multipart.FileHeader `form:"photo"` // ✅ Untuk upload foto profil
}
