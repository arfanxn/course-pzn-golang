package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type User struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	PhoneNumber string   `json:"phone_number"`
	Addresses   []string `json:"addresses"`
}

func EncodeJSON(data any) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	stringified := string(bytes)
	fmt.Println(stringified)
	return stringified
}

func DecodeJSON(bytes []byte, data any) *any {
	err := json.Unmarshal(bytes, data)
	if err != nil {
		panic(err)
	}
	return &data
}

func TestEncodeJSON(t *testing.T) {
	EncodeJSON("Arfan")
	EncodeJSON(1000)
	EncodeJSON(true)
	EncodeJSON([]string{
		"Hello", "World", "Foo", "Bar",
	})
	EncodeJSON(map[string]string{
		"Hello": "World",
		"Foo":   "Bar",
	})

	user := new(User)
	user.Name = "Muhammad Arfan"
	user.Email = "arfan@gm.com"
	user.PhoneNumber = "0896-7878-9999"
	user.Addresses = []string{"Jakarta", "Tokyo"}
	users := &[]User{
		*user,
	}
	EncodeJSON(users)

}

func TestDecodeJSON(t *testing.T) {
	jsonStringified := `[{"name":"Muhammad Arfan","email":"arfan@gm.com","phone_number":"0896-7878-9999","addresses":["Jakarta","Tokyo"]}]`
	jsonBytes := []byte(jsonStringified)

	users := &[]User{}
	DecodeJSON(jsonBytes, users)

	fmt.Println(users)
}

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("../resources/User.json")
	decoder := json.NewDecoder(reader)

	user := new(User)
	decoder.Decode(user)

	fmt.Println(user)
}

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("../resources/EncodedUser.json")
	encoder := json.NewEncoder(writer)

	user := new(User)
	user.Name = "Muhammad Arfan"
	user.Email = "arfan@gm.com"
	user.PhoneNumber = "0896-7878-9999"
	user.Addresses = []string{"Jakarta", "Tokyo"}
	encoder.Encode(user)

	fmt.Println(user)

}
