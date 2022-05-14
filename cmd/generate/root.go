package generate

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/trangmaiq/kgs/migration/migrator"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGenerateCmd() *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate letter strings beforehand and stores them in database",
		Run: func(cmd *cobra.Command, args []string) {
			println("generating...")

			dsn := "<user>:<password>@tcp(127.0.0.1:3306)/keys?charset=utf8mb4&parseTime=True&loc=Local"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Println("unable to initialize db connection")
				os.Exit(1)
			}

			generator := migrator.NewBeforehandKeyGenerator(6, 2, db)
			generator.GenerateAndInsert()
		},
	}

	return generateCmd
}
