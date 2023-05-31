package config

// @Title  config
// @Description  MyGO
// @Author  WeiDa  2023/5/10 11:05
// @Update  WeiDa  2023/5/10 11:05

type ONES struct {
	Base      string `mapstructure:"base" json:"base" yaml:"base"`
	UserUUID  string `mapstructure:"user_uuid" json:"user_uuid" yaml:"user_uuid"`
	UserToken string `mapstructure:"user_token" json:"user_token" yaml:"user_token"`
	SpaceUUID string `mapstructure:"space_uuid" json:"space_uuid" yaml:"space_uuid"`
}
