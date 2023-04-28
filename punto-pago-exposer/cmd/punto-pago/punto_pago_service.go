package main

import (
	config "punto-pago-exposer/configs"
	"punto-pago-exposer/internal"
)

func main() {
	r := internal.SetupRouterPuntoPago()
	r.Run(config.SettingsPuntoPago.ServerPath + ":" + config.SettingsPuntoPago.ServerPort)
}
