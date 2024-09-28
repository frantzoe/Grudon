package utlt

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// LoadUsersFromFile loads users from a JSON file
func LoadUsersFromFile() (*Data, error) {
	var data Data
	
	// Open JSON file
	jsonFile, openErr := os.Open(FileName)
	if openErr != nil {
		if errors.Is(openErr, os.ErrNotExist) {
			data.Users = []User{}
		} else {
			log.Fatal("Error opening JSON file:", openErr)
			return &data, openErr
		}
	} else {
		// Close JSON file when done
		defer jsonFile.Close()

		// Read JSON file
		users, dataErr := io.ReadAll(jsonFile)
		if dataErr != nil {
			log.Fatal("Error reading JSON file:", dataErr)
			return &data, dataErr
		}

		// Unmarshal JSON data
		if jsonErr := json.Unmarshal(users, &data); jsonErr != nil {
			log.Fatal("Error unmarshalling JSON data:", jsonErr)
			return &data, jsonErr
		}
	}

	return &data, nil
}

// SaveUsersToFile saves users to a JSON file
func SaveUsersToFile(data *Data) bool {
	// Create JSON file
	jsonFile, creatErr := os.Create(FileName)
	if creatErr != nil {
		log.Fatal("Error creating JSON file:", creatErr)
		return false
	}
	defer jsonFile.Close()

	// Marshal JSON data
	userData, jsonErr := json.MarshalIndent(data, "", "\t")
	if jsonErr != nil {
		log.Fatal("Error marshalling JSON data:", jsonErr)
		return false
	}

	// Write JSON data
	if _, writErr := jsonFile.Write(userData); writErr != nil {
		log.Fatal("Error writing JSON file:", writErr)
		return false
	}

	return true
}

// EncodeResponse encodes data as JSON and writes it to the response writer
func EncodeResponse(writer http.ResponseWriter, data interface{}) {
	if encodErr := json.NewEncoder(writer).Encode(data); encodErr != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("Error encoding response:", encodErr)
	}
}

// DecodeUser decodes JSON data from the reader and returns a User object
func DecodeUser(reader io.Reader) (*User, error) {
	var user User

	if decodErr := json.NewDecoder(reader).Decode(&user); decodErr != nil {
		return nil, decodErr
	}

	return &user, nil
}

// GetNextUserId returns the next user ID
func GetNextUserId(data *Data) string {
	if len(data.Users) < 1  {
		return "6d7b1001"
	}

	lastIdStr := data.Users[len(data.Users)-1].Id
	lastIdInt, _ := strconv.Atoi(lastIdStr[4:8])

	return "6d7b" + strconv.Itoa(lastIdInt + 1)
}
