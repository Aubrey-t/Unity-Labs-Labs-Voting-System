package controllers

import (
	"encoding/json"
	"fmt"
	"log"

    "github.com/astaxie/beego/logs"
	//"github.com/micro/go-log"
	"github.com/univote/models"
	"github.com/univote/common"
	"github.com/astaxie/beego"
)

//Operations about Users
type  UserController  struct {
	beego.Controller
}

var(logger *log.Logger)

func init(){
	logs.SetLogger("console")
	logger = logs.GetLogger("UserController")
}

func (u *UserController) URLMapping() {
	u.Mapping("create", u.Post)
	// @router /logout [get]
	u.Mapping("logout", u.Logout)
	// @router /info [get]
	u.Mapping("info", u.Get)
	// @router /update [put]
	u.Mapping("update", u.Put)
	// @router /login [post]
	u.Mapping("login", u.Login)

}


// @Title CreateUser
// @Description create user with the parameter in the post body
// @Param body body models.User true "body for user content arranged in json"
// @Success 200 Register user success!
// @Failure 400 Register user fail!
// @router /create [post]
func(u *UserController) Post() {
	fmt.Println("welcome...")
	var user models.User

	//parse request body to User model
	e1 := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if e1 != nil {
		logs.Error(common.ErrParse)
		u.Data["json"] = common.Response{400, 400, common.ErrParse, "creating new account information incorrect"}
		u.ServeJSON()
		return
}

	//check if the user name had been registered
	exist := models.CheckUserName(user.Username)
	if exist == true {
		logs.Error(common.ErrUserNameExist)
		u.Data["json"] = common.Response{400, 400, common.ErrUserNameExist, ""}
		u.ServeJSON()
		return
	}

	//try adding user
	newUser, e4 := models.AddUser(&user)

	if e4 != nil {
		logs.Error(common.ErrAddUser)
		u.Data["json"] = common.Response{400, 400, common.ErrAddUser, ""}
		u.ServeJSON()
		return
	}

	models.AddSetting(&models.Setting{Id: newUser.Uid})

	u.Data["json"] = common.Response{Status:200, Data:&models.User{Uid: newUser.Uid}}
	u.ServeJSON()
}

//Get User struct by the ID
// @router /info [get]
func (u *UserController) Get() {
	token := u.Ctx.Input.Header("token")
	valid := models.ValidateToken(token)
	if !valid {
		u.Ctx.ResponseWriter.WriteHeader(401)
		u.Data["json"] = common.Response{401, 401, common.ErrInvalidToken, ""}
		u.ServeJSON()
		return
	}
	uid := u.GetString(":uid")
	user, err := models.GetUserById(uid)
	if err != nil {
		u.Data["json"] = common.Response{400, 400, err.Error(), ""}
	} else {
		u.Data["json"] = common.Response{200, 0, "", user}
	}
	u.ServeJSON()
}

//
// @router /update [put]
func (u *UserController) Put() {
	token := u.Ctx.Input.Header("token")
	valid := models.ValidateToken(token)
	if !valid {
		u.Ctx.ResponseWriter.WriteHeader(401)
		u.Data["json"] = common.Response{401, 401, common.ErrInvalidToken, ""}
		u.ServeJSON()
		return
	}
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		user.Uid = uid
		_, err := models.UpdateUser(&user)
		if err != nil {
			u.Data["json"] = common.Response{400, 400, err.Error(), ""}
		} else {
			u.Data["json"] = common.Response{200, 0, "", user}
		}
	}
	u.ServeJSON()
}

// @router /login [post]
func (u *UserController) Login() {

	fmt.Println("Request to login file received")

	type UserNamePassword struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var up UserNamePassword
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &up); err != nil {
		u.Data["json"] = common.Response{400, 400, common.ErrNoUserPass, ""}
		u.ServeJSON()
		return
	}

	fmt.Printf("username:%s password:%s", up.Username, up.Password)
	if status, user := models.Login(up.Username, up.Password); status == true {
		type Token struct {
			Token string `json:"token"`
		}
		u.Data["json"] = common.Response{200, 0, "", &Token{Token:user.Token}}
	} else {
		u.Data["json"] = common.Response{400, 400, common.ErrLoginFail, ""}
	}
	u.ServeJSON()
}

// @Title ValidateToken
// @Description Validate the token sent by the client.
// @Param token formData string true "token of the user client"
// @Success 200 Valid token.
// @Failure 400 Invalid token.
// @router /validate_token [post]
func (u *UserController) ValidateToken() {
	token := u.Ctx.Input.Header("token")
	valid := models.ValidateToken(token)
	if valid == true {
		u.Data["json"] =  common.Response{Status:200}
	} else {
		u.Data["json"] = common.Response{400, 400, common.ErrInvalidToken, ""}
	}
	u.ServeJSON()
}

// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] =  common.Response{Status:200}
	//u.Data["json"] = "logout success"
	u.ServeJSON()
}

//func (u *UserController) Delete() {
//	token := u.Ctx.Input.Header("token")
//	valid := models.ValidateToken(token)
//	if !valid {
//		u.Ctx.ResponseWriter.WriteHeader(401)
//		u.Data["json"] = common.Response{401, 401, common.ErrInvalidToken, ""}
//		u.ServeJSON()
//		return
//	}
//	uid := u.GetString(":uid")
//	models.DeleteUser(uid)
//	u.Data["json"] = common.Response{Status:200}
//	u.ServeJSON()
//}
