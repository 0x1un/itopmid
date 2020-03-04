package util

import "fmt"

// return a string like "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
func GenDBUrl(host, port, dbname, uname, passwd string) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, uname, dbname, passwd)
}
