// cmd/types.go

package cmd

type Config struct {
	APIURL             string `yaml:"apiurl"`
	APIAuth            string `yaml:"apiusername"`
	APISecret          string `yaml:"apisecret"`
	APIPath            string `yaml:"apipath"`
	DecryptedAPIAuth   string `yaml:"-"`
	DecryptedAPISecret string `yaml:"-"`
}

type DeviceSearchResult struct {
	Devices  []DeviceDefinition `json:"devices"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	Total    int                `json:"total"`
}

type DeviceDefinition struct {
	EasIds struct {
		EasID []string `json:"EasId"`
	} `json:"EasIds"`
	TimeZone           string `json:"TimeZone"`
	Udid               string `json:"Udid"`
	SerialNumber       string `json:"SerialNumber"`
	MacAddress         string `json:"MacAddress"`
	Imei               string `json:"Imei"`
	EasID              string `json:"EasId"`
	AssetNumber        string `json:"AssetNumber"`
	DeviceFriendlyName string `json:"DeviceFriendlyName"`
	DeviceReportedName string `json:"DeviceReportedName"`
	LocationGroupID    struct {
		Name string `json:"Name"`
		UUID string `json:"Uuid"`
	} `json:"LocationGroupId"`
	LocationGroupName string `json:"LocationGroupName"`
	UserID            struct {
	} `json:"UserId"`
	UserName             string `json:"UserName"`
	DataProtectionStatus int    `json:"DataProtectionStatus"`
	UserEmailAddress     string `json:"UserEmailAddress"`
	Ownership            string `json:"Ownership"`
	PlatformID           struct {
	} `json:"PlatformId"`
	Platform string `json:"Platform"`
	ModelID  struct {
	} `json:"ModelId"`
	Model                  string     `json:"Model"`
	OperatingSystem        string     `json:"OperatingSystem"`
	PhoneNumber            string     `json:"PhoneNumber"`
	LastSeen               CustomTime `json:"LastSeen"`
	EnrollmentStatus       string     `json:"EnrollmentStatus"`
	ComplianceStatus       string     `json:"ComplianceStatus"`
	CompromisedStatus      bool       `json:"CompromisedStatus"`
	LastEnrolledOn         CustomTime `json:"LastEnrolledOn"`
	LastComplianceCheckOn  CustomTime `json:"LastComplianceCheckOn"`
	LastCompromisedCheckOn CustomTime `json:"LastCompromisedCheckOn"`
	ComplianceSummary      struct {
		DeviceCompliance []struct {
			CompliantStatus     bool       `json:"CompliantStatus"`
			PolicyName          string     `json:"PolicyName"`
			PolicyDetail        string     `json:"PolicyDetail"`
			LastComplianceCheck CustomTime `json:"LastComplianceCheck"`
			NextComplianceCheck CustomTime `json:"NextComplianceCheck"`
			ActionTaken         []struct {
				ActionType int `json:"ActionType"`
			} `json:"ActionTaken"`
			ID struct {
				Value int `json:"Value"`
			} `json:"Id"`
			UUID string `json:"Uuid"`
		} `json:"DeviceCompliance"`
	} `json:"ComplianceSummary"`
	IsSupervised bool `json:"IsSupervised"`
	DeviceMCC    struct {
		SIMMCC     string `json:"SIMMCC"`
		CurrentMCC string `json:"CurrentMCC"`
	} `json:"DeviceMCC"`
	IsRemoteManagementEnabled        string     `json:"IsRemoteManagementEnabled"`
	DataEncryptionYN                 string     `json:"DataEncryptionYN"`
	AcLineStatus                     int        `json:"AcLineStatus"`
	VirtualMemory                    int        `json:"VirtualMemory"`
	OEMInfo                          string     `json:"OEMInfo"`
	DeviceCapacity                   int        `json:"DeviceCapacity"`
	AvailableDeviceCapacity          int        `json:"AvailableDeviceCapacity"`
	LastSystemSampleTime             CustomTime `json:"LastSystemSampleTime"`
	IsDeviceDNDEnabled               bool       `json:"IsDeviceDNDEnabled"`
	IsDeviceLocatorEnabled           bool       `json:"IsDeviceLocatorEnabled"`
	IsCloudBackupEnabled             bool       `json:"IsCloudBackupEnabled"`
	IsActivationLockEnabled          bool       `json:"IsActivationLockEnabled"`
	IsNetworkTethered                bool       `json:"IsNetworkTethered"`
	BatteryLevel                     string     `json:"BatteryLevel"`
	IsRoaming                        bool       `json:"IsRoaming"`
	LastNetworkLANSampleTime         CustomTime `json:"LastNetworkLANSampleTime"`
	LastBluetoothSampleTime          CustomTime `json:"LastBluetoothSampleTime"`
	SystemIntegrityProtectionEnabled bool       `json:"SystemIntegrityProtectionEnabled"`
	ProcessorArchitecture            int        `json:"ProcessorArchitecture"`
	UserApprovedEnrollment           bool       `json:"UserApprovedEnrollment"`
	EnrolledViaDEP                   bool       `json:"EnrolledViaDEP"`
	TotalPhysicalMemory              int        `json:"TotalPhysicalMemory"`
	AvailablePhysicalMemory          int        `json:"AvailablePhysicalMemory"`
	OSBuildVersion                   string     `json:"OSBuildVersion"`
	HostName                         string     `json:"HostName"`
	LocalHostName                    string     `json:"LocalHostName"`
	SecurityPatchDate                CustomTime `json:"SecurityPatchDate"`
	SystemUpdateReceivedTime         CustomTime `json:"SystemUpdateReceivedTime"`
	IsSecurityPatchUpdate            bool       `json:"IsSecurityPatchUpdate"`
	DeviceManufacturerID             int        `json:"DeviceManufacturerId"`
	DeviceNetworkInfo                []struct {
		ConnectionType string `json:"ConnectionType"`
		IPAddress      string `json:"IPAddress"`
		MACAddress     string `json:"MACAddress"`
		Name           string `json:"Name"`
		Vendor         string `json:"Vendor"`
	} `json:"DeviceNetworkInfo"`
	DeviceCellularNetworkInfo []struct {
		CarrierName string `json:"CarrierName"`
		CardID      string `json:"CardId"`
		PhoneNumber string `json:"PhoneNumber"`
		DeviceMCC   struct {
		} `json:"DeviceMCC"`
		IsRoaming bool `json:"IsRoaming"`
	} `json:"DeviceCellularNetworkInfo"`
	EnrollmentUserUUID string `json:"EnrollmentUserUuid"`
	ManagedBy          int    `json:"ManagedBy"`
	WifiSsid           string `json:"WifiSsid"`
	DepTokenSource     int    `json:"DepTokenSource"`
	ID                 struct {
		Value int `json:"Value"`
	} `json:"Id"`
	UUID string `json:"Uuid"`
}
