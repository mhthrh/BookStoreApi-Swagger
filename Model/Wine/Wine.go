package Wine

import (
	"database/sql"
	"fmt"
	"github.com/mhthrh/ApiStore/Utility/DbUtil/PgSql"
	"time"
)

type Wines []Wine
type Wine struct {
	ID       int
	Name     string
	MadeDate time.Time
	Price    float32
	Volume   float32
}

type Tool struct {
	db *sql.DB
}

func New(db *sql.DB) *Tool {
	return &Tool{db: db}
}
func (t *Tool) List(i, j int) (Wines, error) {
	var result Wines
	rows, err := PgSql.RunQuery(t.db, "select * from Wines")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var w Wine
		rows.Scan(&w.ID, &w.Name, &w.MadeDate, &w.Price, &w.Volume)
		result = append(result, w)
	}

	return result, nil
}

func (t *Tool) Find(id int) (Wine, error) {
	var w Wine
	rows, err := PgSql.RunQuery(t.db, fmt.Sprintf("select * from Wines where id=%d", id))
	if err != nil {
		return w, err
	}
	rows.Next()
	rows.Scan(&w.ID, &w.Name, &w.MadeDate, &w.Price, &w.Volume)

	return w, err
}

func (t *Tool) Add(p Wine) (interface{}, error) {
	_, err := PgSql.ExecuteCommand(fmt.Sprintf("INSERT INTO public.\"Wines\"(\"ID\", \"Name\", \"MadeDate\", \"Price\", \"Volume\")VALUES (%d, %s, %s,%F , %2f)", p.ID, p.Name, p.MadeDate, p.Price, p.Volume), t.db)
	if err != nil {
		return nil, err
	}
	return 0, nil
}

func (t *Tool) Delete(id int) (interface{}, error) {
	_, err := PgSql.ExecuteCommand(fmt.Sprintf("Delete from Wine where Id=%d", id), t.db)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
func (t *Tool) Update(id int, p float32) (interface{}, error) {
	_, err := PgSql.ExecuteCommand(fmt.Sprintf("Update Wine set price=%2f where Id=%d", p, id), t.db)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
