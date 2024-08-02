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
	Title          string
	Slug           string
	PagePath       string
	MainMenuPage   string
	InMenu         bool
	InPluginPage   bool
	PageID         int64
	PageReferences []string
	Role           int
	BoundFunctions []string
}

type Route struct {
	ID             int64
	Title          string // ex Home
	Slug           string // ex /index
	PagePath       string // ex /files/index.html
	PageID         int64
	InMenu         bool
	Role           int // minimum role to access this route
	PageReferences map[string]int64
	URLS           map[string]string
	BoundFunctions map[string][]string
}

type User struct {
	ID       int64
	Username string
	Name     string
	Surname  string
	Email    string
	Role     int
}

type BaseServer struct {
	DB         *sql.DB
	HttpServer *http.Server
}
