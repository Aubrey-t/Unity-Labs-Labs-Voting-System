package models 

import (

	"strconv"
	"time"
	"errors"
	"github.com/univote/utils"
	"github.com/astaxie/beego"
	"github.com/micro/go-log"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

var(
	
	database string
	expire int64
	//logger *log.Logger
)

func init(){
	log.Log("Init ORM ...")
	expire, _ = strconv.ParseInt(beego.AppConfig.String("expire"), 10 , 64 )
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:FBIb7n3ChxJ5qevz0CN1ddqSeJajUq3@tcp(host.docker.internal:3308)/univotedb?charset=utf8")
	orm.RegisterModel(new(User))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Logf("%v", err)
	}
}

type User struct {

	Uid string      `json:"uid" orm:"column(uid);pk;size(64)"`
	Username string `json:"username" orm:"column(username);unique;size(32)"`
	Password string `json:"password" orm:"column(password);size(128)"`
	Email string    `json:"email" orm:"column(email);size(50)"`
	Token string    `json:"token" orm:"column(token);size(256)"`
	Salt string     `json:"salt" orm:"column(salt);size(128)"`
	LastLogin int64 `json:"last_login" orm:"column(last_login);size(11)"`
	CreatedAt int64 `json:"created_at" orm:"column(created_at);size(11)"`
	UpdatedAt int64 `json:"updated_at" orm:"column(updated_at);size(11)"`
}

func GetUsers() orm.QuerySeter{
	return orm.NewOrm().QueryTable(new(User))
}

//Query for username existence
func CheckUserName(username string) bool{
	exist := GetUsers().Filter("Username", username).Exist()
	return exist
}

//Query for user email existence
func CheckEmail(email string) bool{
	exist := GetUsers().Filter("Email",email).Exist()
	return exist
}

//Query for user id existence
func CheckUserId(userId string) bool {
	exist := GetUsers().Filter("Uid", userId).Exist()
	return exist
}

//check if password is correct
func (u *User) CheckPassword(password string) (ok bool, err error){
	hash, err := utils.GeneratePassHash(password, u.Salt)
	if err != nil{
		return false, err
	}
	return u.Password == hash, nil
}

//get user by the id
func GetUserById(uid string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Uid: uid}
	if err = o.QueryTable(new(User)).Filter("Uid", uid).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

//get user by the username
func GetUserByUserName(username string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Username: username}
	if err = o.QueryTable(new(User)).Filter("Username", username).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

//get user by the user's token
func GetUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

//Login with username and password aquired from the user
func Login(username string, password string) (bool, *User){

	o := orm.NewOrm()

	//Query if user exist by username
	user, err := GetUserByUserName(username)
	if err != nil{
		errors.New("user name does not exist")
		return false, nil
	}

	//hash the password and compare with the hash in the Database
	passwordHash, err := utils.GeneratePassHash(password, user.Salt)
	if err != nil{
		errors.New("password hash generation failure")
		return false, nil
	}

	err = o.QueryTable(user).Filter("Username", username).Filter("Password", passwordHash).One(user)
	pass := err != orm.ErrNoRows
	if pass == true{

		//given the password is right , login successfully, update the token and Login timestamp
		et := utils.EasyToken{
			Username: user.Username,
			Uid: 0,
			Expires: time.Now().Unix() + expire,
		}
		token, _ := et.GetToken()
		user.Token = token
		lastLogin := time.Now().UTC().Unix()
		user.LastLogin = lastLogin
		o.Update(user)
	}else{
		 errors.New("password does not match")
		return false, nil
	}
	return pass, user
}

// add a user when registering
func AddUser(m *User) (*User, error){

	o := orm.NewOrm()
	salt, err := utils.GenerateSalt()
	if err != nil {
		return nil, err
	}
	passwordHash, err := utils.GeneratePassHash(m.Password, salt)
	if err != nil {
		return nil, err
	}
	CreatedAt := time.Now().UTC().Unix()
	UpdatedAt := CreatedAt

	uid, _ := utils.GenerateUUID()

	user := User{
		Uid: uid,
		Username: m.Username,
		Password: passwordHash,
		Salt: salt,
		Email: m.Email,
		CreatedAt: CreatedAt,
		UpdatedAt:UpdatedAt,
	}
	_, err = o.Insert(&user)
	if err == nil {
		return &user, err
	}
	return nil, err
}

func ValidateToken(token string) bool {
	et := utils.EasyToken{
		Username: "m.Username",
		Uid:      0,
		Expires:  0,
	}
	valid, err := et.ValidateToken(token)
	if err != nil {
		errors.New("token validation error")
	}
	return valid
}

func UpdateUser(user *User) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Update(user)
	if num == 1 {
		return num, nil
	}
	return num, err
}