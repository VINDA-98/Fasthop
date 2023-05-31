package config

import "reflect"

// @Title  config
// @Description  MyGO
// @Author  WeiDa  2023/5/5 11:28
// @Update  WeiDa  2023/5/5 11:28

type SMTP struct {
	// 邮件服务器地址
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// 端口
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// 发送邮件用户账号
	User string `mapstructure:"user" json:"user" yaml:"user"`
	// 授权密码
	Password string `mapstructure:"password" json:"pwd" yaml:"password"`
	// 发送邮件昵称
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"`
	// 发送的目标用户
	Target string `mapstructure:"target" json:"target" yaml:"target"`
}

func (s SMTP) IsEmpty() bool {
	return reflect.DeepEqual(s, SMTP{})
}
