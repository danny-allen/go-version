package version

// Version interface.
type VInterface interface {

	// Public methods.
	SetLogUrl()
	SetTag()
	GetTag() 		string
	FindByTag() 	LogItem
	GetCurrent() 	LogItem
	GetLatest() 	LogItem
	GetLast() 		LogItem
}

// Version data struct.
type Version struct {

	// Interface.
	VInterface

	// Fields.
	log 	[]LogItem
	tag		string
	latest	string
	logUrl 	string
}

// Version log struct.
type LogItem struct {
	Tag		  	string	`yaml:"tag"`
	Filename 	string	`yaml:"filename"`
}

// Create an array of log items.
var Log = []LogItem{}