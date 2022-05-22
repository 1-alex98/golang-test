package db

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
)

var db *gorm.DB
var pepper = "pepper" //obviously pepper should be set to something else here

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=5432 sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{}, &Good{}, &DataPoint{}, &AccountEntry{}, &Offer{})
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
	db.Preload(clause.Associations).First(&good, id)
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
	db.Where("email = ?", email).Preload("AccountEntrys.Good").First(&user)
	return user.AccountEntrys
}

func GetUser(email string) User {
	user := User{}
	db.Where("email = ?", email).First(&user)
	return user
}

func UpdateAccount(value float64, goodId uint, userId uint) {
	accountEntry := AccountEntry{Value: value, GoodID: goodId, UserID: userId}
	if db.Model(&accountEntry).Where("good_id = ? AND user_id = ?", goodId, userId).Update("Value", value).RowsAffected == 0 {
		db.Create(&accountEntry)
	}
}

func UpdateCredit(email string, value float64) {
	db.Model(&User{}).Where("email = ?", email).Update("credit", value)
}

func CreateOffer(email string, value float64, quantity float64, goodId uint) {
	user := Offer{GoodID: goodId, Value: value, Quantity: quantity, UserID: GetUser(email).ID}
	db.Create(&user)
}
