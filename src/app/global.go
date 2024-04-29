package app

import (
	"os"

	"github.com/aquasecurity/table"
	"nemesis-cli/src/database"
)

var (
	t          = table.New(os.Stdout)
	dbInstance = database.GetDB()
)
