package lib

import (
	"database/sql"
	"net/http"
)

type RootImport struct {
	Routes                []RouteImport
	ExtPageBoundFunctions map[int64]string
	BoundFunctions        []string
}

type RouteImport struct {
	Name           string
	PagePath       string
	PageID         int64
	PageReferences []string
	Role           int64
	BoundFunctions []string
}

type Route struct {
	ID             int64
	Name           string // ex index
	PagePath       string // ex /files/index.html
	PageID         int64
	Role           int64 // minimum role to access this route
	PageReferences map[string]int64
	URLS           map[string]string
	BoundFunctions []string
}

type BaseServer struct {
	DB         *sql.DB
	HttpServer *http.Server
}
