package migrator

import (
	"fmt"
	"strings"
	"time"

	"github.com/trangmaiq/kgs/pkg/permutation"
	"gorm.io/gorm"
)

const standard = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

type BeforehandKeyGenerator struct {
	keyLength int
	batchSize int

	db *gorm.DB
}

func NewBeforehandKeyGenerator(keyLength, batchSize int, db *gorm.DB) *BeforehandKeyGenerator {
	return &BeforehandKeyGenerator{
		keyLength: keyLength,
		batchSize: batchSize,
		db:        db,
	}
}

func (g *BeforehandKeyGenerator) GenerateAndInsert() {
	out := permutation.PermutationChan(strings.Split(standard, ""), g.keyLength)

	var (
		batch []string
	)
	for key := range out {
		batch = append(batch, key)
		if len(batch) == g.batchSize {
			err := g.insert(batch)
			if err != nil {
				fmt.Println("insert keys failed: ", err)
			}
			batch = []string{}
		}
	}
}

func (g *BeforehandKeyGenerator) insert(data []string) error {
	q, args := g.buildQuery(data)
	return g.db.Table("keys").Exec(q, args...).Error
}

func (g *BeforehandKeyGenerator) buildQuery(data []string) (string, []interface{}) {
	now := time.Now()

	switch len(data) {
	case 0:
		return "", nil
	case 1:
		return `INSERT INTO keys (key, is_used, created_at, updated_at) VALUES (?, ?, ?, ?)`, []interface{}{data[0], false, now, now}
	}

	var (
		args      = make([]interface{}, 0)
		builder   strings.Builder
		separator = ","
		str       = "(?, ?, ?, ?)"
		n         = len(separator)*(len(data)-1) + (len(str) * len(data))
	)

	builder.Grow(n)
	builder.WriteString(str)
	args = append(args, data[0], false, now, now)

	for i := 1; i < len(data); i++ {
		builder.WriteString(separator)
		builder.WriteString(str)

		args = append(args, data[i], false, now, now)
	}

	return fmt.Sprintf("INSERT INTO `keys` (`key`, is_used, created_at, updated_at) VALUES %s", builder.String()), args
}
