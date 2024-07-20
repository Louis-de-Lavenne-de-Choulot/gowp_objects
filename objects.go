package lib

import (
	"database/sql"
	"net/http"
)

type RouteImport struct {
	Name           string
	PagePath       string
	PageID         int64
	PageReferences []string
	Role           int64
}

type Route struct {
	ID             int64
	Name           string // ex index
	PagePath       string // ex /files/index.html
	PageID         int64
	PageReferences map[string]int64
	Role           int64 // minimum role to access this route
}

type BaseServer struct {
	DB         *sql.DB
	HttpServer *http.Server
}
