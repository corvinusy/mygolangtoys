package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "time"
    "runtime"
)

const DB_CONNECT_STRING =
    "host=localhost port=5432 user=your_role password=your_password dbname=your_database sslmode=disable"

const TEST_NUMBER = 100000 //inserts for one thread

func main() {

    const THREAD_COUNT = 4
    runtime.GOMAXPROCS(4)

    db, err := sql.Open("postgres", DB_CONNECT_STRING)
    defer db.Close()

    if err != nil {
        fmt.Printf("Database opening error -->%v\n", err)
        panic("Database error")
    }

    init_database(&db)
    done := make(chan bool, THREAD_COUNT)

    t1 := time.Now()

    for i := 0; i < THREAD_COUNT; i++ {
        go make_insertion(&db, done)
    }

    for i := 0; i < THREAD_COUNT; i++ {
        <-done
    }

    t2 := time.Since(t1)

    fmt.Printf("%v queries executed for %v seconds (%v per second)\n",
        TEST_NUMBER*THREAD_COUNT, t2.Seconds(), TEST_NUMBER*THREAD_COUNT/t2.Seconds())

}

/*-----------------------------------------------------------------------------*/
func init_database(pdb **sql.DB) {

    db := *pdb

    init_db_strings := []string{
        "DROP SCHEMA IF EXISTS sb CASCADE;",
        "CREATE SCHEMA sb;",
        //be careful - next multiline string is quoted by backquote symbol
        `CREATE TABLE sb.test_data(
         id serial,
         device_id integer,
         parameter_id integer,
         value varchar(100),
	 event_ctime timestamp default current_timestamp,
         constraint id_pk primary key(id)
	);`
	}

    for _, qstr := range init_db_strings {
        _, err := db.Exec(qstr)

        if err != nil {
            fmt.Printf("Database init error -->%v\n", err)
            panic("Query error")
        }
    }
    fmt.Println("Database rebuilded successfully")
}

/*-----------------------------------------------------------------------------*/
func make_insertion(pdb **sql.DB, done chan bool) {

    db := *pdb
//	const TRANS_SIZE = 200;

    // backquotes for next multiline string
    const INSERT_QUERY = `insert into sb.test_data(device_id, parameter_id, value)
                                  values ($1, $2, $3);`

    insert_query, err := db.Prepare(INSERT_QUERY)
    defer insert_query.Close()

    if err != nil {
        fmt.Printf("Query preparation error -->%v\n", err)
        panic("Preparation error")
    }

    tx, err := db.Begin();
    t_stmt := tx.Stmt(insert_query)

    for i := 0; i < TEST_NUMBER; i++ {

        _, err = t_stmt.Exec(i, i, "0")

        if err != nil {
            fmt.Printf("Query execution error -->%v\n", err)
            panic("Execution error")
        }
    }

    err = tx.Commit();

    if err != nil {
        fmt.Printf("Transaction commit error -->%v\n", err)
        panic("Commit error")
    }

    done <-true

    // do not forget to clean up after work done )
    //_, err = db.Query("TRUNCATE sb.test_data;")
}
