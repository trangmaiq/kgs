package serve

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/trangmaiq/kgs/registry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewServeCmd() *cobra.Command {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Run the Short server",
		Run: func(cmd *cobra.Command, args []string) {
			println("serving...")

			dsn := "<user>:<password>@tcp(127.0.0.1:3306)/keys?charset=utf8mb4&parseTime=True&loc=Local"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Println("unable to initialize db connection")
				os.Exit(1)
			}

			routes := gin.Default()

			_, err = registry.New(routes, db)
			if err != nil {
				log.Println("unable to create new registry")
				os.Exit(1)
			}

			err = routes.Run(":9009")
			if err != nil {
				log.Println("unable to run `kgs` service")
				os.Exit(1)
			}
		},
	}

	return serveCmd
}
