package controllers

import (
	"fmt"
	"net/http"
	"server/db"
	"server/dto"
	"server/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	userService *services.UserService
}

type Claims struct {
	UserRole int
	UserName string
	UserId   uint
	jwt.RegisteredClaims
}

func (controller *UserController) GetUsers(c *gin.Context) {
	users := controller.userService.GetUsers()

	c.IndentedJSON(http.StatusOK, gin.H{"users": users})
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var userDto dto.CreateUserDto
	err := c.ShouldBindJSON(&userDto)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Username already taken!"})
		return
	}

	user, err := controller.userService.CreateUser(userDto)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"user": user})
}

func (controller *UserController) GetUser(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.userService.GetUser(id)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	println(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	controller.userService.DeleteUser(id)
}

func (controller *UserController) SignIn(c *gin.Context) {
	var userDto dto.SignInDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.userService.SignIn(userDto)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	claims := Claims{
		UserName:         user.Username,
		UserId:           user.ID,
		UserRole:         user.Role,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": tokenString, "user": user})
}

func (controller *UserController) GetCurrentUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var userIdString string

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdString = fmt.Sprintf("%v", claims["UserId"])
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	userId, err := strconv.Atoi(userIdString)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "USER NOT FOUND AAAAAAAA"})
		return
	}

	user, err := controller.userService.GetUser(userId)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}

func NewUserController() *UserController {
	userService := services.NewUserService(db.DB)
	return &UserController{userService: userService}
}
