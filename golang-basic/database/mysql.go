package database

var connection string

func init() {
	connection = "mysql"
}

func GetMysql() string {
	return connection
}
