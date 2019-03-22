package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // import with _ for only register driver
)

func main() {
	r := gin.Default()

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
}

func createBooking(c *gin.Context) {
	var b Booking
	if err := c.Bind(&b); err != nil {
		return
	}
	log.Printf("%+v\n", b)

	ctx := context.Background()
	check := func(err error) {
		if err != nil {
			log.Println(err)
		}
	}
	// Open connection to MySQL Database
	db, err := sql.Open("mysql", "root:cityhunter@/bookingdb?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci")
	check(err)
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO tb_booking(NAME, ROOM, START, END) values (?,?,?,?)")
	check(err)

	result, err := stmt.ExecContext(ctx, b.Name, b.Room, b.Start, b.End)
	check(err)
	lastID, _ := result.LastInsertId()
	fmt.Println("New Record ID:", lastID)

	id, err := result.LastInsertId()
	b.ID = int(id)
	c.JSON(http.StatusOK, b)
}

func getAllBookings(c *gin.Context) {
	ctx := context.Background()
	check := func(err error) {
		if err != nil {
			log.Println(err)
		}
	}
	db, err := sql.Open("mysql", "root:cityhunter@/bookingdb?parseTime=true")
	check(err)
	defer db.Close()

	qry := "SELECT id, name, room, start, end FROM tb_booking ORDER BY start "
	stmt, err := db.PrepareContext(ctx, qry)
	check(err)

	rows, err := stmt.QueryContext(ctx)
	check(err)
	var bookingSlice []Booking
	var booking Booking
	for rows.Next() {
		rows.Scan(&booking.ID, &booking.Name, &booking.Room, &booking.Start, &booking.End)
		bookingSlice = append(bookingSlice, booking)
	}
	fmt.Println(bookingSlice)

	c.JSON(http.StatusOK, bookingSlice)
}

func getBookingByID(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	check := func(err error) {
		if err != nil {
			log.Println(err)
		}
	}
	// Open connection to MySQL Database
	db, err := sql.Open("mysql", "root:cityhunter@/bookingdb?parseTime=true")
	check(err)
	defer db.Close()

	qry := "SELECT ID, NAME, ROOM, START, END FROM tb_booking WHERE id = ?"
	stmt, err := db.PrepareContext(ctx, qry)
	check(err)

	row := stmt.QueryRowContext(ctx, id)
	var booking Booking
	err = row.Scan(&booking.ID, &booking.Name, &booking.Room, &booking.Start, &booking.End)
	check(err)
	c.JSON(http.StatusOK, booking)
}
func deleteBooking(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()
	check := func(err error) {
		if err != nil {
			log.Println(err)
		}
	}
	// Open connection to MySQL Database
	db, err := sql.Open("mysql", "root:cityhunter@/bookingdb")
	check(err)
	defer db.Close()

	stmt, err := db.PrepareContext(ctx, "DELETE FROM tb_booking WHERE id = ?")
	check(err)

	_, err = stmt.ExecContext(ctx, id)
	check(err)

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, `""`)
}
