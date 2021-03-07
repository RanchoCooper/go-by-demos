package reader

import (
	"log"
	"time"

	"github.com/RanchoCooper/go-by-demos/articles/singleton-database/pkg/db"
	"github.com/RanchoCooper/go-by-demos/articles/singleton-database/pkg/model"
)

func dumpEmployee() {
	var rs []model.Employee
	d := db.DB()
	d.Find(&rs)
	log.Println(rs)
}

func Run(quit <-chan struct{}) {
	tk := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-tk.C:
			dumpEmployee()
		case <-quit:
			return
		}
	}
}