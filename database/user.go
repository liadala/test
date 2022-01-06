package database

import "fmt"

func WriteUser(id string, login string, displayName string) {
	lock.Lock()
	defer lock.Unlock()
	stmt, err := db.Prepare(`INSERT OR REPLACE INTO users (id, login, displayName)VALUES(?,?,?);`)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, login, displayName)
	if err != nil {
		fmt.Println(err)
	}
}
