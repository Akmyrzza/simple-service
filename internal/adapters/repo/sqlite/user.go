package sqlite

import (
	"context"
	"fmt"

	"github.com/Akmyrzza/simple-service/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

func (d *St) UserCreate(ctx *gin.Context, obj *entities.UserCUSt) string {
	var result string
	
	ctx.ShouldBindJSON(obj)
	statementInsert, _ := d.db.Prepare("INSERT INTO users (name, phone) VALUES (?, ?)")
	fmt.Print(obj.Name)
	statementInsert.Exec(obj.Name, obj.Phone)

	result = "User successfully created."
	return result
}

func (d *St) UserGet(ctx context.Context, id string) *entities.UserSt {

	/*ctx.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
	})
	*/
	result := &entities.UserSt{}

	return result
}

/*
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
*/

/*
	rows, _ := database.Query("SELECT id, name, phoneNumber FROM users")
	var id int
	var name string
	var phoneNumber int

	for rows.Next() {
		rows.Scan(&id, &name, &phoneNumber)
		fmt.Printf("id: %d, name: %s, phone number: %d", id, name, phoneNumber)
	}
*/
