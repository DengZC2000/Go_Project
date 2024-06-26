package config

type Config struct {
	Mysql       Mysql       `yaml:"mysql"`
	Logger      Logger      `yaml:"logger"`
	System      System      `yaml:"system"`
	SiteInfo    SiteInfo    `yaml:"site_info"`
	QQ          QQ          `yaml:"qq"`
	Email       Email       `yaml:"email"`
	QiNiu       QiNiu       `yaml:"qiniu"`
	Jwt         Jwt         `yaml:"jwt"`
	LocalUpload LocalUpload `yaml:"local_upload"`
	Redis       Redis       `yaml:"redis"`
}
