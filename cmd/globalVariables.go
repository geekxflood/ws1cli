// cmd/globalVariables.go

package cmd

var insecure bool
var prettyPrint bool

var forceRecreate bool

var showDetails bool

var lgid int
var inventory bool

var productID int
var startStopProduct bool

var deviceUuid string
var sessionType string

var sessionTypes = []string{
	"ScreenShare",
	"FileManager",
	"RemoteShell",
	"RegistryEditor",
}
