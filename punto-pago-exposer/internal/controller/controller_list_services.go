package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	config "punto-pago-exposer/configs"

	"github.com/gin-gonic/gin"
)

type ControllerPuntoPagoGetServices struct {
	Context *gin.Context
}

// entry point
func (cb ControllerPuntoPagoGetServices) GetListServices() {
	cb.GetListServicesXML(config.SettingsPuntoPago.PathPuntoPagoBackend, EncryptedHeaderGetListServices())
}

func (cb ControllerPuntoPagoGetServices) GetListServicesXML(url string, signatureHeader string) {

	// request
	xmlReq := "<request><command>servicelist</command></request>"
	xmlByte := []byte(xmlReq)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlByte))
	if err != nil {
		cb.Context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// header
	req.Header.Set("Signature", signatureHeader)

	// content type
	req.Header.Set("Content-Type", "application/xml")
	req.ContentLength = int64(len(xmlByte))

	// transaction
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		cb.Context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// response
	xmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cb.Context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cb.Context.Data(http.StatusOK, "application/xml", xmlData)
}

func EncryptedHeaderGetListServices() string {
	input := "<request><command>servicelist</command></request>" + config.SettingsPuntoPago.HeaderEncryptedToken
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
