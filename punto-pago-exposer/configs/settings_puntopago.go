package config

type SettingsStruct struct {
	//server settings
	ServerScheme string
	ServerPort   string
	ServerPath   string

	//paths
	PathGetServiceList string
	PathPostGetBill    string
	PathPostPayBill    string

	// punto pago url
	PathPuntoPagoBackend string

	HeaderEncryptedToken string

	//config
	ResponseOk    string
	ResponseError string
}

var SettingsPuntoPago = SettingsStruct{
	ServerScheme: "http://",
	ServerPort:   "8080",
	ServerPath:   "localhost",

	PathGetServiceList: "/punto-pago/get-list-services",
	PathPostGetBill:    "/punto-pago/get-bill",
	PathPostPayBill:    "/punto-pago/pay-bill",

	// parametrizar
	PathPuntoPagoBackend: "[URL]",
	HeaderEncryptedToken: "[TOKEN]",

	ResponseOk:    "ok",
	ResponseError: "error",
}
