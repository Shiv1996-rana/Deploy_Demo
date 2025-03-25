package Controllers

import (
	"deploy_demo/database"
	"deploy_demo/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestCode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Tested": http.StatusOK, "message": "starts for building all api..."})
}

func CreateData(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf("error on Request_Payload of CreateData %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	jsonAddress, err := json.Marshal(user.Address)
	if err != nil {
		log.Printf("Error marshalling address to JSON: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process address data"})
		return
	}
	query := `INSERT INTO crud(name,email,mobile_no,address) VALUES($1,$2,$3,$4) RETURNING id`
	var UserID int64
	err = database.DB.QueryRow(query, user.Name, user.Email, user.Mobile_No, jsonAddress).Scan(&UserID)
	if err != nil {
		log.Printf("Error inserting new record: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Create_newMessage": "new Created..",
		"status": http.StatusOK,
		"user":   UserID,
	})

}

func FetchAllData(ctx *gin.Context) {
	var users []models.User

	query := `SELECT name, email, mobile_no, address FROM crud`
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching records: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var jsonAddress []byte

		err := rows.Scan(&user.Name, &user.Email, &user.Mobile_No, &jsonAddress)
		if err != nil {
			log.Printf("Error scanning record: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process records"})
			return
		}

		// Convert JSONB address to Go struct
		err = json.Unmarshal(jsonAddress, &user.Address)
		if err != nil {
			log.Printf("Error unmarshalling address JSON: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode address data"})
			return
		}

		users = append(users, user)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": users})
}

func GetAllById(ctx *gin.Context) {
	id := ctx.Param("id") // Get ID from URL parameter

	query := `SELECT name, email, mobile_no, address FROM crud WHERE id=$1`
	var user models.User
	var jsonAddress []byte

	err := database.DB.QueryRow(query, id).Scan(&user.Name, &user.Email, &user.Mobile_No, &jsonAddress)
	if err != nil {
		log.Printf("Error fetching record by ID: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Convert JSONB address to Go struct
	err = json.Unmarshal(jsonAddress, &user.Address)
	if err != nil {
		log.Printf("Error unmarshalling address JSON: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode address data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "user": user})
}

func DeleteDataById(ctx *gin.Context) {
	id := ctx.Param("id")

	query := `DELETE FROM crud WHERE id=$1`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting record by ID: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully"})
}

func UpdateData(ctx *gin.Context) {
	id := ctx.Param("id") // Get ID from URL parameter

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Printf("Error parsing request payload: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonAddress, err := json.Marshal(user.Address)
	if err != nil {
		log.Printf("Error marshalling address to JSON: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process address data"})
		return
	}

	query := `UPDATE crud SET name=$1, email=$2, mobile_no=$3, address=$4 WHERE id=$5`
	_, err = database.DB.Exec(query, user.Name, user.Email, user.Mobile_No, jsonAddress, id)
	if err != nil {
		log.Printf("Error updating record: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully"})
}
