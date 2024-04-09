// controllers/Todo/list.go

package Todo

import (
	"encoding/json"
	"fmt"
	auth "ginjing/controllers/authentication"
	"ginjing/structs"
	"net/http"

	"github.com/kataras/iris/v12"
)

// List handles listing of todos
func List(ctx iris.Context) {
	userJSON := ctx.Values().Get("user").([]byte)
	var user auth.JwtBody
	errz := json.Unmarshal(userJSON, &user)
	if errz != nil {
		fmt.Println("Error unmarshalling JSON:", errz)
	}
	fmt.Println("User->Username from middleware is:", user.Username)

	// ------------------

	getAuthToken, err := auth.Login("admin", "admin")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User auth is: %s.\n", *getAuthToken)

	verifyingToken, err := auth.VerifyToken(*getAuthToken)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonData, err := json.Marshal(verifyingToken.Claims)
	if err != nil {
		fmt.Println("Error marshalling token claims to JSON:", err)
		return
	}

	jsonString := string(jsonData)
	fmt.Println("JSON string:", jsonString)

	// ------------------

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	var todos []structs.Todo
	if err := json.NewDecoder(response.Body).Decode(&todos); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	ctx.JSON(todos)
}
