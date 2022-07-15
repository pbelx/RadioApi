package radiodb

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Rstations struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//get all stations from DB
func Getstations(c *gin.Context) {
	dbs, err := sql.Open("sqlite3", "./controllers/radiodb/rstat.db")

	checkErr(err)
	fmt.Println("found file")
	defer dbs.Close()
	rows, err := dbs.Query("SELECT name,url FROM Rstations;")
	checkErr(err)
	var xip = []Rstations{}
	for rows.Next() {

		var name string
		var url string
		rows.Scan(&name, &url)
		fmt.Printf(("%v,%v\n"), name, url)
		xip = append(xip, Rstations{Name: name, Url: url})

	}
	fmt.Println(xip)
	c.JSON(200, gin.H{
		"data": xip,
	})

}

//add stations insert
func Addstation(name, url string, c *gin.Context) {
	// var rr Rstations
	dbs, err := sql.Open("sqlite3", "./controllers/radiodb/rstat.db")
	checkErr(err)
	defer dbs.Close()

	stmt, err := dbs.Prepare("INSERT INTO Rstations(name,url) values(?,?)")
	checkErr(err)

	res, err := stmt.Exec(name, url)
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(id)
	c.JSON(200, gin.H{
		"data": "Saved info",
	})

	// var xip = []Rstations{}

}

//update stations url
func UpdateUrl(name, url string, c *gin.Context) {
	// var rr Rstations
	dbs, err := sql.Open("sqlite3", "./controllers/radiodb/rstat.db")
	checkErr(err)
	defer dbs.Close()

	stmt, err := dbs.Prepare("UPDATE Rstations SET url=? where name=?")
	checkErr(err)

	res, err := stmt.Exec(url, name)
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(id)

	if id != 0 {
		c.JSON(200, gin.H{
			"data": "Saved info",
		})
	} else {
		c.JSON(400, gin.H{
			"data": "Failed to change info",
		})
	}

	// var xip = []Rstations{}

}

//delete station by name
func Delstation(name string, c *gin.Context) {
	// var rr Rstations
	dbs, err := sql.Open("sqlite3", "./controllers/radiodb/rstat.db")
	checkErr(err)
	// defer dbs.Close()

	stmt, err := dbs.Prepare("DELETE FROM Rstations WHERE name = ?")
	checkErr(err)

	res, err := stmt.Exec(name)
	checkErr(err)

	id, err := res.RowsAffected()
	checkErr(err)

	if id != 0 {
		c.JSON(200, gin.H{
			"data": "ok " + name,
		})
	} else {
		c.JSON(400, gin.H{
			"data": "400 " + name,
		})
	}

	// var xip = []Rstations{}

}
