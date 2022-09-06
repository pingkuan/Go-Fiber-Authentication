package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pingkuan/go-fiber-api/database"
	"github.com/pingkuan/go-fiber-api/models"
	"github.com/pingkuan/go-fiber-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct{ 
	ID uint32 `json:"id"`
	Username    string `json:"username"`
	Email string `json:"email"`
}

func Hash(password string) (string, error) {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes),err
}

func VerifyPassword(hash, password string) bool {
	err:= bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(c *fiber.Ctx) error {

	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}
    
	db :=database.Db
	user :=new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error","message":"Invalid input","data":err})		
	}
 	
	hash,err := Hash(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error","message":"Cannot hash password","data":err})
	}
	
	user.Password = hash

	if err:=db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error","message":"Couldn't create user","data":err})
	}
 
	t,err := utils.GenerateToken(user.ID)
    if err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error","message":"Couldn't generate token","data":err})
	}

	newUser :=NewUser{
		Username: user.Username,
		Email: user.Email,
		Token: t,
	}

	return c.JSON(fiber.Map{"status":"success","message":"Created user","data":newUser})
}

func AuthUser(c *fiber.Ctx)error{
	type Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type LoginUser struct{
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}
	
	var input Input
	
	db := database.Db
	var user models.User

	if err:=c.BodyParser(&input); err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Error on request","data":err})
	}
    
	if err:= db.Where(&models.User{Email: input.Email}).Find(&user).Error; err !=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found", "data": nil})
	}

	if !VerifyPassword(user.Password, input.Password){
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error","message":"Invalid password","data":nil})
	}
	
	t,err :=utils.GenerateToken(user.ID)
    if err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error","message":"Couldn't generate token","data":err})
	}
	
	loginUser :=LoginUser{
		Username: user.Username,
		Email: user.Email,
		Token: t,
	}

	return c.JSON(fiber.Map{"status": "success","message":"Success login","data":loginUser})
}

func GetUser(c *fiber.Ctx)error{

	id:=c.Locals("ID")
	db:= database.Db
	var user models.User

	db.Find(&user, id)
	if user.Username == ""{
		return c.Status(404).JSON(fiber.Map{"status":"error","message":"No user found","data":nil})
	}

	userData :=UserData{
		ID:user.ID,
		Username:user.Username,
		Email:user.Email,
	}
	return c.JSON(fiber.Map{"status": "success","message":"user found","data":userData})
}

func UpdateUser(c *fiber.Ctx)error{
	type UpdateInput struct{
		Username string `json:"username"`
	}
	
	var ui UpdateInput
	if err:=c.BodyParser(&ui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error","message":"Invalid input","data":err})
	}

	id:=c.Locals("ID")
	db:=database.Db
	var user models.User

	db.First(&user, id)
	user.Username = ui.Username
	db.Save(&user)

	userData :=UserData{
		ID:user.ID,
		Username:user.Username,
		Email:user.Email,
	}

	return c.JSON(fiber.Map{"status":"success","message":"User updated","data":userData})

}

func DeleteUser(c *fiber.Ctx)error{
	type PasswordInput struct{
		Password string `json:"password"`
	}
	
	var pi PasswordInput
	if err:=c.BodyParser(&pi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status":"error","message":"Invalid input","data":err})
	}
	
	id := c.Locals("ID")
	db := database.Db
	var user models.User

	db.First(&user, id)
	if !VerifyPassword(user.Password, pi.Password){
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error","message":"Invalid password","data":nil})
	}

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success","message":"User deleted","data":nil})
}