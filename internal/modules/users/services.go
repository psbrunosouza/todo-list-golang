package users

import "todo-list/internal/models"

func CreateUserService(user *models.User) error {
	if result := CreateUser(user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteUserService(user *models.User) error {
	if result := DeleteUser(user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindUserService(id int, user *models.User) error {
	if result := FindUser(id, user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListUsersService(users *[]models.User) error {
	if result := ListUsers(users); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateUserService(id int, user *models.User) error {
	if result := UpdateUser(id, user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
