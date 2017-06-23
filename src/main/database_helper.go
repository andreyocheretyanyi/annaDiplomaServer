package main


import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "111111"
	dbname   = "postgres"
)

func open_db() *sql.DB {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//db, err := sql.Open("postgres",psqlInfo)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	fmt.Print(db)
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()

	return db
}

//--------------------------------
func create_table ()  {
	db := open_db()
	defer db.Close()
	qv := "CREATE TABLE IF NOT EXISTS blocks("+
		"id SERIAL CONSTRAINT _id PRIMARY KEY,"+
		"name VARCHAR(64) NOT NULL"+
	");"

	qv1 := "CREATE TABLE IF NOT EXISTS rooms("+
		"id SERIAL CONSTRAINT _id_r PRIMARY KEY,"+
		"number INTEGER NOT NULL CHECK (number > -1),"+
		"block_id INTEGER REFERENCES blocks,"+
		"price INTEGER NOT NULL CHECK  (price > -1),"+
		"water INTEGER NOT NULL CHECK (water = 1 OR water =0),"+
		"free INTEGER NOT NULL CHECK (free = 1 OR free = 0),"+
		"date VARCHAR(255) NOT NULL"+
	");"

	_, err := db.Exec(qv)

	if err != nil {
		fmt.Println(err)
	}
	_, err1 := db.Exec(qv1)

	if err1 != nil {
		fmt.Println(err1)
	}
}

func wipe_table(){
	db := open_db()
	defer db.Close()
	_,err := db.Exec("TRUNCATE TABLE blocks " +
		"CASCADE;")
	if err != nil{
		fmt.Println(err)
	}

}

func add_block(block Block)  {
	db := open_db()
	defer db.Close()
	_, err := db.Exec("INSERT INTO blocks (name)" +
		" VALUES ($1);",
		block.Name)
	if err != nil{
		fmt.Println(err)
	}
}

func add_room(room Room)  {
	db := open_db()
	defer db.Close()
	_, err := db.Exec("INSERT INTO rooms (number, block_id, price, water, free,date)" +
		"VALUES ($1,$2,$3,$4,$5,$6);",
		room.Number,room.Block_id,room.Price,room.Water,room.Free,room.Date)
	if err != nil{
		fmt.Println(err)
	}

}

func get_rooms() []Room  {
	var rooms []Room
	db := open_db()
	defer db.Close()
	rows, err := db.Query("SELECT *" +
		" FROM rooms;")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var block_id int
		var number int
		var price int
		var free int
		var water int
		var date string
			err = rows.Scan(&id,&number,&block_id,&price,&water,&free,&date)
		if err != nil {
			fmt.Println(err)
		}
		room := Room {id,block_id,number,
					  price,free,water,date}
		rooms = append(rooms,room)
	}
	rows.Close()

	return rooms
}


func get_blocks() []Block  {
	var blocks []Block
	db := open_db()
	defer db.Close()
	rows, err := db.Query("SELECT *" +
		" FROM blocks;")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id,&name)
		if err != nil {
			fmt.Println(err)
		}
		blck := Block{id,name}
		blocks = append(blocks,blck)
	}
	rows.Close()

	return blocks
}

//----------------------------------------------------



func delete_block(id int)  {
	db := open_db()
	defer db.Close()
	_,err := db.Exec("DELETE FROM blocks " +
		"WHERE id = ($1)",id)
	if err != nil{
		fmt.Println(err)
	}
}



