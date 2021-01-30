package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[string]*user{}
	seq   = 1
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[strconv.Itoa(u.ID)] = u
	seq++
	a, err := json.Marshal(users)
	if err != nil {
		return err
	}
	MyJson := []byte(a)
	error := ioutil.WriteFile("Users.json", MyJson, 0777)
	if error != nil {
		return err
	}
	b, err := json.Marshal(seq)
	if err != nil {
		return err
	}
	SeqJson := []byte(b)
	error2 := ioutil.WriteFile("Seq.json", SeqJson, 0777)
	if error2 != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)

}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id := c.Param("id")
	users[id].Name = u.Name
	a, err := json.Marshal(users)
	if err != nil {
		return err
	}
	MyJson := []byte(a)
	error := ioutil.WriteFile("Users.json", MyJson, 0777)
	if error != nil {
		return err
	}
	b, err := json.Marshal(seq)
	if err != nil {
		return err
	}
	SeqJson := []byte(b)
	error2 := ioutil.WriteFile("Seq.json", SeqJson, 0777)
	if error2 != nil {
		return err
	}
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	delete(users, id)
	a, err := json.Marshal(users)
	if err != nil {
		return err
	}
	MyJson := []byte(a)
	error := ioutil.WriteFile("Users.json", MyJson, 0777)
	if error != nil {
		return err
	}
	b, err := json.Marshal(seq)
	if err != nil {
		return err
	}
	SeqJson := []byte(b)
	error2 := ioutil.WriteFile("Seq.json", SeqJson, 0777)
	if error2 != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func main() {
	jsonFile, err := os.Open("Users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Users.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	error := json.Unmarshal([]byte(byteValue), &users)
	if error != nil {
		fmt.Println(error)
	}
	jsonFile2, err := os.Open("Seq.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Seq.json")
	defer jsonFile2.Close()
	byteValue2, _ := ioutil.ReadAll(jsonFile2)
	error2 := json.Unmarshal([]byte(byteValue2), &seq)
	if error2 != nil {
		fmt.Println(error)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
