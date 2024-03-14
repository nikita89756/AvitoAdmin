package client

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hack/servies"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	//"github.com/redis/go-redis/v9"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Start(addr string) error {
	r := mux.NewRouter()
	r.HandleFunc("/json", Jsontouch).Methods("POST")
	go http.ListenAndServe(addr, r)
	return nil

}
func Search(c *Client, mic int, loc int, db *sql.DB) error {
	var price sql.NullInt32
	var comback string
	err := db.QueryRow("SELECT price,comeback FROM baseline1 WHERE microcategory_id = $1 AND location_id = $2;", mic, loc).Scan(&price, &comback)
	if err != nil {
		return err
	}
	if price.Valid {
		c.Price = int(price.Int32)
		c.LocationId = loc
		c.MicrocategoryId = mic
		fmt.Print(price)
		return nil
	} else {
		mas := strings.Split(comback, ":")
		l, _ := strconv.Atoi(mas[1])
		ca, _ := strconv.Atoi(mas[0])
		if mas[1] == "0" {
			Search(c, ca, loc, db)
			mic = ca
		} else {
			Search(c, mic, l, db)
		}
	}
	return nil

}

func Jsontouch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var C Client
	json.NewDecoder(r.Body).Decode(&C)
	db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=tree sslmode=disable")
	SearchDiscount(&C, C.MicrocategoryId, C.LocationId, db)
	fmt.Print(err)
	json.NewEncoder(w).Encode(C)
}

func SearchDiscount(c *Client, mic int, loc int, db *sql.DB) {
	mas := servies.GetSegmentsByUserIDs([]int64{int64(c.UserId)})[int64(c.UserId)]
	fmt.Print(mas)
	var wg sync.WaitGroup
	m := make(map[int][]int)
	if len(mas) != 0 {
		for _, k := range mas {
			wg.Add(1)
			go SearchD(c, mic, loc, db, &wg, k, m)
		}
		wg.Wait()
		if len(m) != 0 {
			keys := make([]int, 0, len(m))
			for k := range m {
				keys = append(keys, int(k))
			}

			sort.Ints(keys)

			p := keys[len(keys)-1]
			c.Price = m[p][0]
			c.LocationId = m[p][2]
			c.MicrocategoryId = m[p][1]
			c.MatrixId = keys[len(keys)-1]
		} else {
			Search(c, c.MicrocategoryId, c.LocationId, db)
		}
	}
}

func SearchD(c *Client, mic int, loc int, db *sql.DB, wg *sync.WaitGroup, index int, m map[int][]int) {
	defer wg.Done()
	var matrixName string
	err := db.QueryRow("SELECT matrix_name FROM discounts WHERE id = $1", index).Scan(&matrixName)
	if err != nil {
		fmt.Print(err)
	} else {
		SearchByDiscount(mic, loc, db, matrixName, m, index)
	}
}

func SearchByDiscount(mic int, loc int, db *sql.DB, matrixName string, m map[int][]int, ind int) {
	var price sql.NullInt32
	var comeback string

	go db.QueryRow("SELECT comeback FROM comeback WHERE microcategory_id = $1 AND location_id = $2;", mic, loc).Scan(&comeback)
	err := db.QueryRow(fmt.Sprintf("SELECT price FROM %s WHERE microcategory_id = $1 AND location_id = $2;", matrixName), mic, loc).Scan(&price)
	if err != nil {
		fmt.Print(err)
	} else {
		if price.Valid {
			m[ind] = []int{int(price.Int32), mic, loc}
			fmt.Print(price)
		} else {
			mas := strings.Split(comeback, ":")
			l, _ := strconv.Atoi(mas[1])
			ca, _ := strconv.Atoi(mas[0])
			if mas[1] == "0" {
				SearchByDiscount(ca, loc, db, matrixName, m, ind)
			} else if mas[0] == "0" {
				SearchByDiscount(mic, l, db, matrixName, m, ind)
			} else {
				m[ind] = []int{}
			}
		}
	}
}
