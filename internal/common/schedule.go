package common

type Schedule struct {
	GraphDashboards DashBoards
	DataStores      DataStores
	Active          bool
	CadenceInHours  int
}

type DashBoards struct {
	GrafanaDashboards GrafanaDashboards
}

type GrafanaDashboards struct {
	Dashboards map[string]GrafanaDashBoard
	UserID     string
}

type DataStores struct {
	ConfluencePages map[string]ConfluencePage
}

//ConfluencePage defines the location in which pages will be created
type ConfluencePage struct {
	SpaceKey     string
	ParentPageID string
	UserID       string
}
