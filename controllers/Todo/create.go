// controllers/Todo/create.go

package Todo

import (
	"database/sql"
	"fmt"
	"ginjing/structs"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/pusher/pusher-http-go"
)

// Create handles creating a new todo
func Create(ctx iris.Context) {
	var t structs.Todo
	err := ctx.ReadJSON(&t)

	if err != nil {
		ctx.StopWithProblem(
			iris.StatusBadRequest,
			iris.NewProblem().Title("Todo creation failure").DetailErr(err))
		return
	}

	// *------*|*------*
	token := os.Getenv("TURSO_TOKEN")
	url := os.Getenv("TURSO_URL") + token

	// Open database connection
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return
	}
	defer db.Close()

	// Dummy insert query
	insertQuery := `INSERT INTO jokowi (username, email, password) VALUES (?, ?, ?)`

	// Execute the insert query with dummy values
	username := "rizki"
	email := "john.doe@example.com"
	password := "password123"
	_, err = db.Exec(insertQuery, username, email, password)
	if err != nil {
		fmt.Println("Error inserting data into 'jokowi' table:", err)
		return
	}

	fmt.Println("Dummy data inserted into 'jokowi' table successfully.")
	// *------*|*------*

	// Remove WS change to Pusher
	pusherClient := pusher.Client{
		AppID:   "1784889",
		Key:     "7c218173bec5716bfbc6",
		Secret:  "1d49abc4b7b219d97ef6",
		Cluster: "ap1",
		Secure:  true,
	}

	data := map[string]string{"message": "hello world you ll receive it when hit create"}
	errPC := pusherClient.Trigger("my-channel", "my-event", data)
	if errPC != nil {
		fmt.Println(errPC.Error())
	}

	// shfaiofioahiuazsdh

	ctx.JSON(t)
}
