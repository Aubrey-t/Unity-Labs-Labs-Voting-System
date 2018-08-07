package models

import(

	"strconv"
	"time"
	"errors"
	"github.com/univote/utils"
	"github.com/astaxie/beego"
	"github.com/micro/go-log"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


func init(){
	log.Log("Creating Votes Table")
	orm.RegisterModel(new(Vote))
}

type Vote struct{

	Title       string `json:"title" orm:"column(title);pk;size(64)"`
	Description string `json:"description" orm:"column(description );unique;size(32)"`
	Publisher   string `json:"publisher" orm:"column(publisher);size(128)"`// Publisher name as used on registration
	Votes       int64  `json:"votes" orm:"column(votes);size(128)"` // number of votes received
	Image       string `json:"phone" orm:"column(phone);size(128)"` // store a path to the image as a string
	Status      int    `json:"status" orm:"column(status);size(1)"` // 0: enabled, 1:disabled
}

func Votes() orm.QuerySeter{
	return orm.NewOrm().QueryTable(new(Vote))
}

//check if vote exits by Title
func CheckTitle(title string) bool {
	exist := Votes().Filter("Title", title).Exist()
	return exist
}

// Query all votes by Publisher
func GetVotesByPublisher(publisher string) (v *Vote, err error) {
	o := orm.NewOrm()
	v = &Vote{Publisher: publisher}
	if err = o.QueryTable(new(Vote)).Filter("Publisher", publisher).RelatedSel().One(v); err == nil { // not clear what the "One" is
		return v, nil
	}
	return nil, err
}

func GetVoteByTitle(title string) (v *Vote, err error) {
	o := orm.NewOrm()
	v = &Vote{Publisher: title}
	if err = o.QueryTable(new(Vote)).Filter("Publisher", title).RelatedSel().One(v); err == nil { // not clear what the "One" is
		return v, nil
	}
	return nil, err
}

func CreateVote(m *Vote) (title string, e error){
	log.Log("Creating Vote request received")
	o := orm.NewOrm()

	status := 0
	votes := RecordVote()

	vote := Vote{

		Title: m.Title,
		Description: m.Description,
		Image: m.Image,
		Publisher:m.Publisher,
		Status:status,
		Votes:votes,
	}

	_, err := o.Insert(&vote)
	if err != nil {
		return "", err
	}
	return vote.Title,nil
}


