package models

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	// This is required for using postgres with gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User is the structure of the class being used for the database
type User struct {
	ID string `gorm:"primaryKey;type:uuid"`
	gorm.Model
	Username string `json:"username"`
	Email string `json:"-"`
	HashedPassword string `json:"-"`
	Jwt string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// InitialUserMigration will use GORM to migrate the tables in the database.
func InitialUserMigration() {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

func (user *User) Authorise(username, password string) bool {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	db.Where("Username = ?", username).Find(&user)
	jwt, _ := GenerateJWT(user.ID)

	user.Jwt = jwt

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	return err == nil 
}

func GenerateJWT(id string) (string, error) {

	secret := os.Getenv("SECRET")

    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
    claims["expiration"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString([]byte(secret))

    if err != nil {
        fmt.Println(err)
        return "", err
    }

    return tokenString, nil
}

func DecodeJWT(token string) string {
	
	decodedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC); 
		if !ok {
			return nil, fmt.Errorf("There was an error")
		}
		secret := os.Getenv("SECRET")
		return []byte(secret), nil
	})
	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
		id := fmt.Sprintf("%v", claims["id"])
		return id
	} else {
		return ""
	}
}

// GetAllUsers Queries the database and returns all users
func (user *User) GetAllUsers() *[]User {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	users := []User{}

	db.Find(&users)

	return &users
}

// FindUserByID will be given an id and will find the user based upon it
func (user *User) FindUserByID(id string) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	db.First(&user, id)
	jwt, _ := GenerateJWT(user.ID)

	user.Jwt = jwt
}

func (u *User) Create(username, email, password string) interface{} {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	newUserID := uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		panic(err)
	}

	jwt, _ := GenerateJWT(newUserID)
 
	user := db.Create(&User{ID: newUserID, Username: username, Email: email, HashedPassword: string(hashedPassword), Jwt: jwt})

	return user.Value
}

func (u *User) Update(id, email string) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	user := User{}

	db.Where("id = ?", id).Find(&user)
	
	user.Email = email

	db.Save(&user)

	fmt.Println("User successfully updated")
}


func (user *User) Destroy(id string) map[string]string {

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to DataBase")
	}

	defer db.Close()

	db.Where("id = ?", id).Find(&user)
	db.Delete(&user)

	
	fmt.Println("User successfully deleted")
	m := make(map[string]string)
    m["Message"] = "User Deleted!"
	
	return m
}

func (user *User) SendFriendRequest() {
	
}

func (user *User) PendingFriendRequests() {
	
}