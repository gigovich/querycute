package querycute

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
)

// DB instance
var DB *sql.DB
