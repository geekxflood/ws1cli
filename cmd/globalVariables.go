// cmd/globalVariables.go

package cmd

var insecure bool
var prettyPrint bool

var forceRecreate bool

var showDetails bool

var lgid int
var inventory bool
var command string
var inputJson string
var valueFilter string

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

var commandTypes = []string{
	"EnterpriseWipe",
	"LockDevice",
	"ScheduleOsUpdate",
	"SoftReset",
	"Shutdown",
}

var valueFilterTypes = []string{
	"Macaddress",
	"Udid",
	"Serialnumber",
	"ImeiNumber",
}
