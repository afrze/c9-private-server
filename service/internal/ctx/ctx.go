package ctx

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Ctx struct {
	Mongo *mongo.Database
	MSSQL *sql.DB
}
