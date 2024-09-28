package utlt

// AppendUser adds a user to a JSON file
func AppendUser(data *Data, user User) bool {
	user.Id = GetNextUserId(data)
	data.Users = append(data.Users, user)
	return SaveUsersToFile(data)
}

// FindUser finds a user by ID
func (data *Data) FindUser(id string) *User {
	for _, user := range data.Users {
		if user.Id == id {
			return &user
		}
	}

	return nil
}

// UpdateUser updates a user in a JSON file
func UpdateUser(data *Data, user User) bool {
	for index, userToUpdate := range data.Users {
		if userToUpdate.Id == user.Id {
			data.Users[index].Name = user.Name
			data.Users[index].Email = user.Email
			return SaveUsersToFile(data)
		}
	}
	
	return false
}

// DeleteUser deletes a user from a JSON file
func DeleteUser(data *Data, id string) bool {
	for index, user := range data.Users {
		if user.Id == id {
			data.Users = append(data.Users[:index], data.Users[index+1:]...)
			return SaveUsersToFile(data)
		}
	}

	return false
}
