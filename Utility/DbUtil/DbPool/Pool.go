package DbPool

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type DB struct {
	Db      *sql.DB
	Working bool
}

type DBs []*DB

type DbInfo struct {
	Host            string
	Port            int32
	User            string
	Pass            string
	Dbname          string
	Driver          string
	ConnectionCount int
	RefreshPeriod   time.Duration
}

var (
	cnn string
)

func New(d *DbInfo) *DBs {
	var dbs DBs
	cnn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Pass, d.Dbname)
	for i := 0; i < d.ConnectionCount; i++ {
		if d := newConnection(cnn, d.Driver); d != nil {
			fmt.Println("kossss ", i)
			dbs = append(dbs, &DB{
				Db:      d,
				Working: false,
			})
		} else {
			fmt.Println("kirrrr ", i)
		}
	}

	go func() {
		for {
			for range time.Tick(time.Second * d.RefreshPeriod) {
				fmt.Println("start refresh")
				for i, i2 := range dbs {
					fmt.Println("start ping")
					if i2.Db.Ping() != nil {
						fmt.Println("finish ping")
						if new := newConnection(cnn, d.Driver); new != nil {
							dbs[i] = &DB{
								Db:      newConnection(cnn, d.Driver),
								Working: false,
							}
						} else {
							//must aware admin
						}

					}
				}
			}
		}
	}()
	return &dbs
}
func newConnection(cnn, driver string) *sql.DB {

	db, err := sql.Open(driver, cnn)
	if err != nil {
		return nil
	}
	return db
}

func (db *DBs) Pull() *DB {
	c1 := make(chan *DB)
	c2 := make(chan bool)
	go func() {
		for {
			for _, i2 := range *db {
				if i2.Working == false {
					c1 <- i2
				}
			}
		}

	}()
	go func() {
		time.Sleep(10 * time.Second)
		c2 <- false
	}()
	select {
	case msg := <-c1:
		msg.Working = true
		return msg
	case _ = <-c2:
		return &DB{
			Db:      nil,
			Working: false,
		}
	}
}

func (db *DBs) Push(cc *DB) {
	cc.Working = false
}
