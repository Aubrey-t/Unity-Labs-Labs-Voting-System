package controllers

import (
	"encoding/json"
	"fmt"
	"errors"

	"github.com/astaxie/beego/logs"
	//"github.com/micro/go-log"
	"github.com/univote/models"
	"github.com/univote/common"
	"github.com/astaxie/beego"
)

type VoteController  struct {
	beego.Controller
}

func init(){
	logs.SetLogger("console")
	logger = logs.GetLogger("UserController")
}

func (v *VoteController) URLMapping() {
	v.Mapping("createvote", v.CreateVote)
	// @router /record [put]
	v.Mapping("record", v.RecordVote)

}

// @Title CreateVote
// @Description create vote with the parameter in the post body
// @Param body body models.CreateVote true "body for user content arranged in json"
// @Success 200 Register user success!
// @Failure 400 Register user fail!
// @router /createvote [post]

func(v *VoteController) CreateVote() {
	fmt.Println("welcome to votes")
	var vote models.Vote

	//Validate Token
	token := v.Ctx.Input.Header("token")
	valid := models.ValidateToken(token)
	if !valid{
		v.Ctx.ResponseWriter.WriteHeader(401)
		v.Data["json"] = common.Response{401, 401, common.ErrInvalidToken, ""}
		v.ServeJSON()
		return
	}

	//Get User by token
	b, user := models.GetUserByToken(token)
	if b != true{
		errors.New("user not found")
	}

	vote.Publisher = user.Username// Name the publisher

	//parse request body to Vote model
	e1 := json.Unmarshal(v.Ctx.Input.RequestBody, &vote)
	if e1 != nil {
		logs.Error(common.ErrParse)
		v.Data["json"] = common.Response{400, 400, common.ErrParse, "New vote information not correct, please check again"}
		v.ServeJSON()
		return
	}

	//try creating the vote
	newVote, e4 := models.CreateVote(&vote)

	if e4 != nil {
		//logs.Error(common.ErrAddUser)
		v.Data["json"] = common.Response{400, 400, common.ErrAddUser, "New vote creation failed"}
		v.ServeJSON()
		return
	}

	models.AddSetting(&models.Setting{Id: newVote.Title})

	v.Data["json"] = common.Response{Status:200, Data:&models.Vote{Title: newUser.title}}
	v.ServeJSON()
}

