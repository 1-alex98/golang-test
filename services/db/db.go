package db

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var pepper = "pepper" //obviously pepper should be set to something else here

func Init() {
	dsn := "host=localhost user=postgres password=banana dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{}, &Good{}, &DataPoint{}, &AccountEntry{})
	if err != nil {
		panic(err)
	}

}

func CheckCredentials(email string, pw string) (success bool) {
	user := User{}
	db.Where("email = ?", email).First(&user)
	passwordHash := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(pw+pepper))
	return err == nil
}

func CreateUser(email string, pw string) (id uint, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pw+pepper), 12)
	user := User{Email: email, Password: string(passwordHash)}
	result := db.Create(&user)
	err = result.Error
	id = user.ID
	return
}

func Goods() (goods []Good) {
	db.Find(&goods)
	return
}

func GoodById(id string) (good Good) {
	db.First(&good, id)
	return
}

func GoodByIdPreloaded(id string) (good Good) {
	db.Preload("DataPoints").First(&good, id)
	return
}

func SaveGood(good Good) {
	db.Updates(&good)
}

func SaveGoodDataPoint(good Good) {
	dataPoint := DataPoint{Value: good.CurrentCourse, GoodID: good.ID}
	db.Create(&dataPoint)
}

func GetAccount(email string) []AccountEntry {
	user := User{}
	db.Where("Email = ?", email).Preload("AccountEntrys.Good").First(&user)
	return user.AccountEntrys
}
func GetUser(email string) User {
	user := User{}
	db.Where("email = ?", email).First(&user)
	return user
}
