package models

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    //"github.com/micro/go-log"
	//"errors"
	//"time"

)

var (
	SysSetting Setting
)

func init() {

	orm.RegisterModel(new(Setting))

	//err := orm.RunSyncdb("univotedb", false, true)
	//if err != nil {
	//	log.Log(err)
	//}
}

type Setting struct{
	Id string `json:"id" orm:"column(id);pk;unique;size(32)"`
	Language string `json:"language, omitempty" orm:"column(language)"`
	CurrencyUnit string `json:"currency_unit, omitempty" orm:"column(currency_unit)"`
	Web3Url string `json:"web3_url, omitempty" orm:"column(web3_url)"`
}
//
//func (s *Setting) TableName() string {
//	return TableName("setting")
//}

func Settings() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User))
}

//check if exist user's setting
func CheckSetting(id string) bool {
	exist := Settings().Filter("Id", id).Exist()
	return exist
}

func GetSettingById(id string) (s *Setting, err error) {
	o := orm.NewOrm()
	s = &Setting{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(s); err == nil {
		return s, nil
	}
	return nil, err
}

func AddSetting(ss *Setting) (s *Setting, err error) {
	//o := orm.NewOrm()
	s = &Setting{Id: ss.Id}
	return s, nil
}

func UpdateSetting(userid string, us *Setting) (s *Setting) {
	if us.Language != "" {
		SysSetting.Language = us.Language
	}
	if us.CurrencyUnit != "" {
		SysSetting.CurrencyUnit = us.CurrencyUnit
	}
	if us.Web3Url != "" {
		SysSetting.Web3Url = us.Web3Url
	}
	return &SysSetting
}



