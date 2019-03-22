package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // import with _ for only register driver
)

func main() {
	r := gin.Default()
	r.GET("/", defaultHandler)

	r.POST("/bookings/", createBooking)
	r.GET("/bookings/", getAllBookings)
	r.GET("/bookings/:id", getBookingByID)
	r.DELETE("/bookings/:id", deleteBooking)

	r.Run(":8000")
}

type Booking struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Room  string    `json:"room"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`

	//Done bool   `json:"-"` //ignore field ที่ไม่ต้องการให้ออกเป็น json
	//Done bool   `json:"done,omitempty"` //ignore field ที่เป็น zero value ที่ไม่ต้องการให้ออกเป็น json
}

func defaultHandler(c *gin.Context) { //รวมทั้ง request , responnse ใน param เดียว
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, `<!doctype html><html><body><h1 style="color:red">Hello</h1></body></html>`)
}

func createBooking(c *gin.Context) {
	var b Booking
	if err := c.Bind(&b); err != nil { //ใช้ method Bind โดยโยน pointer ของ struct(&) ที่ต้องการรับค่า
		return
	}
	log.Printf("%+v\n", b) //format +v print struct จะดูง่าย

	ctx := context.Background()
	check := func(err error) {
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	// Open connection to MySQL Database
	db, err := sql.Open("mysql", "root:cityhunter@/bookingdb")
	check(err)

	stmt, err := db.PrepareContext(ctx, "INSERT INTO tb_booking(NAME, ROOM, START, END) values (?,?,?,?)")
	check(err)

	//startDate, err := time.Parse("2019-03-22T07:00:00Z", "2019-03-22 07:00:00")
	//endDate, err := time.Parse("2019-03-22T07:00:00Z", "2019-03-22 08:00:00")

	result, err := stmt.ExecContext(ctx, b.Name, b.Room, b.Start, b.End)
	check(err)
	lastID, _ := result.LastInsertId()
	fmt.Println("New Record ID:", lastID)

	id, err := result.LastInsertId()
	b.ID = int(id)
	c.JSON(http.StatusOK, b)
}

func getAllBookings(c *gin.Context) {
	booking := []Booking{
		// {ID: 1, Name: "Michael", Room: "Room1", Start: time.Now().Format("2006-01-02 15:04:05"), End: time.Now().Format("2006-01-02 15:04:05")},
		// {ID: 2, Name: "John", Room: "Room2", Start: time.Now().Format("2006-01-02 15:04:05"), End: time.Now().Format("2006-01-02 15:04:05")},
		// {ID: 3, Name: "Jason", Room: "Room3", Start: time.Now().Format("2006-01-02 15:04:05"), End: time.Now().Format("2006-01-02 15:04:05")},
	}
	c.JSON(http.StatusOK, booking)
}

func getBookingByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"msg": "receive_id " + id,
	})
}
func deleteBooking(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"msg": "receive_id " + id,
	})
}
