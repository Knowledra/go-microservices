package services

import (
	"context"
	"errors"
	"user/models"
)

// UpdateUserByRole dispatches an admin update to the correct role-specific service.
func UpdateUserByRole(c context.Context, role, userId string, input any) error {
	switch role {
	case "teacher":
		typed, ok := input.(models.UpdateTeacher)
		if !ok {
			return errors.New("invalid input type for teacher")
		}
		return UpdateTeacherUser(c, userId, typed)
	case "student":
		typed, ok := input.(models.UpdateStudent)
		if !ok {
			return errors.New("invalid input type for student")
		}
		return UpdateStudentUser(c, userId, typed)
	case "parent":
		typed, ok := input.(models.UpdateParent)
		if !ok {
			return errors.New("invalid input type for parent")
		}
		return UpdateParentUser(c, userId, typed)
	default:
		return errors.New("invalid role: must be teacher, student, or parent")
	}
}

// UpdateSelfByRole dispatches a self-update to the correct role-specific service.
func UpdateSelfByRole(c context.Context, role, userId string, input any) error {
	switch role {
	case "admin":
		typed, ok := input.(models.UpdateAdmin)
		if !ok {
			return errors.New("invalid input type for admin")
		}
		return UpdateAdminUser(c, userId, typed)
	case "teacher":
		typed, ok := input.(models.UpdateTeacher)
		if !ok {
			return errors.New("invalid input type for teacher")
		}
		return UpdateTeacherUser(c, userId, typed)
	case "student":
		typed, ok := input.(models.UpdateStudent)
		if !ok {
			return errors.New("invalid input type for student")
		}
		return UpdateStudentUser(c, userId, typed)
	case "parent":
		typed, ok := input.(models.UpdateParent)
		if !ok {
			return errors.New("invalid input type for parent")
		}
		return UpdateParentUser(c, userId, typed)
	default:
		return errors.New("unsupported role for self-update")
	}
}
