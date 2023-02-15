package sqlite

import (
	"context"

	"github.com/Akmyrzza/simple-service/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

func (d *St) UserCreate(ctx *gin.Context, obj *entities.UserCUSt) string {
	var result string
	var lastId int

	statement, _ := d.db.Prepare("SELECT max(id) FROM users")
	statement.QueryRow().Scan(&lastId)

	ctx.ShouldBindJSON(obj)
	statementInsert, _ := d.db.Prepare("INSERT INTO users (id, name, phone) VALUES (?, ?, ?)")
	statementInsert.Exec(lastId+1, obj.Name, obj.Phone)

	result = "User successfully created."
	return result
}

func (d *St) UserGet(ctx context.Context, id int) (*entities.UserSt, error) {
	result := &entities.UserSt{}

	statement, err := d.db.Prepare("SELECT id, name, phone FROM users WHERE id = ?")
	statement.QueryRow(id).Scan(&result.Id, &result.Name, &result.Phone)

	return result, err
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
