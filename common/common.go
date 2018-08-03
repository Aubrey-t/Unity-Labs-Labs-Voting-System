package common

import (
	//"errors"
	//"regexp"
	//"strings"

	"github.com/astaxie/beego"
	//"github.com/Unity-Labs-Development/WalletBackend/utils"
)

// Predefined const error strings.
const (
	ErrInputData    = "Input Data Error"
	ErrDatabase     = "Database Error"
	ErrDupUser      = "Duplicate User"
	ErrNoUser       = "Not Found User"
	ErrPass         = "Password Error"
	ErrNoUserPass   = "Not Found User Or Password Error"
	ErrNoUserChange = "Not Found User Or No User Changed"
	ErrInvalidUser  = "Invalid User"
	ErrOpenFile     = "Open File Error"
	ErrWriteFile    = "Write File Error"
	ErrSystem       = "System Error"
	ErrCode400       = "400"
	ErrPhone         = "Invalid phone number."
	ErrPhoneExist    = "Phone number had been registered."
	ErrUserNameExist = "User name had been registered."
	ErrParse         = "User information parsed from request body was invalid."
	ErrAddUser       = "Adding user fail."
	ErrLoginFail     = "Login fail."
	ErrInvalidToken  = "Invalid token."
)

// UserData definition.
type UserSuccessLoginData struct {
	//AccessToken string `json:"access_token"`
	UserName string `json:"user_name"`
}

// Controller Response is controller error info struct.
type Response struct {
	Status int `json:"status"`
	ErrorCode int `json:"code"`
	ErrorMessage string `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseController struct {
	beego.Controller
}

// Predefined controller error/success values.
var (

	successReturn = &Response{200, 0, "ok", "ok"}
	err404 = &Response{404, 404, "Webpage not Found", "Webpage not Found"}
	errInputData = &Response{400, 10001, "Data input Error", "Client Parameter Error"}
	errDatabase = &Response{500, 10002, "Server Error", "Database Operation Error"}
	errDupUser = &Response{400, 10003, "User information already Exist", "Duplicate database records"}
	errNoUser = &Response{400, 10004, "User information does not Exist", "Database record does not exist"}
	errPass = &Response{400, 10005, "Wrong User name or password", "Wrong Password"}
	errNoUserOrPass = &Response{400, 10006, "User does not exist or password is incorrect", "Database record is non-existing or password is incorrect"}
	errNoUserChange = &Response{400, 10007, "User does not exist or no change in data", "Database record does not exist"}
	errInvalidUser = &Response{400, 10008, "User information is incorrect", "Session information incorrect"}
	errOpenFile = &Response{500, 10009, "Server Error", "Open File Error"}
	errWriteFile = &Response{500, 10010, "Server Error", "Write file Error"}
	errSystem = &Response{500, 10011, "Server Error", "Operating System Error"}
	errExpired = &Response{400, 10012, "Login has expired", " "}
	errPermission = &Response{400, 10013, "No Permissions", "No Permissions"}
)

// BaseController definition.
//type BaseController struct {
//	beego.Controller
//}

// RetError return error information in JSON.
func (base *BaseController) RetError(e *Response) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.Data = ""
	}

	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Ctx.ResponseWriter.WriteHeader(e.Status)
	base.Data["json"] = e
	base.ServeJSON()
	base.StopRun()
}

var sqlOp = map[string]string{

	"eq": "=",
	"ne": "<>",
	"gt": ">",
	"ge": ">=",
	"lt": "<",
	"le": "<=",
}
