package writer

import (
	"fmt"
	"time"

	"github.com/RanchoCooper/go-by-demos/articles/singleton-database/pkg/db"
	"github.com/RanchoCooper/go-by-demos/articles/singleton-database/pkg/model"
)

func updateEmployeeEmailByName(name, email string) {
	d := db.DB()
	d.Model(&model.Employee{}).Where("name = ?", name).Update("email", email)
}

func Run(quit <-chan struct{}) {
	tk := time.NewTicker(5 * time.Second)
	count := 0
	for {
		select {
		case <-tk.C:
			updateEmployeeEmailByName("rancho", fmt.Sprintf("rancho@example-%d.com", count))
			count++
			if count > 10 {
				count = 0
			}
		case <-quit:
			return
		}
	}
}