package config

type Api_admin struct {
	Api_port   string
	Api_secret string
}
type Mysql struct {
	Parsetime string
	Loc       string
	Driver    string
	User      string
	Port      string
	Db_name   string
	Show_sql  string
	Pass_word string
	Host      string
	Charset   string
}
type Config struct {
	Api_admin Api_admin
	Mysql     Mysql
}
