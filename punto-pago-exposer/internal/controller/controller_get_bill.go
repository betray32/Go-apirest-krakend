package controllers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	config "punto-pago-exposer/configs"

	"github.com/gin-gonic/gin"
)

type ControllerPuntoPagoGetBill struct {
	Context *gin.Context
}

// entry point
func (cb ControllerPuntoPagoGetBill) GetBill() {
	// Obtener el valor de los parámetros de la URL
	serviceProvider := cb.Context.Param("serviceProvider")
	paymentReference := cb.Context.Query("paymentReference")

	fmt.Println("serviceProvider:", serviceProvider)
	fmt.Println("paymentReference:", paymentReference)

	// response
	cb.GetBillMockResponse()
	//cb.PuntoPagoGetBillIntegration(config.SettingsPuntoPago.PathPuntoPagoBackend)
}

// mocked response
func (cb ControllerPuntoPagoGetBill) GetBillMockResponse() {
	var mockData = gin.H{
		"bills": []gin.H{
			{
				"billReference": gin.H{
					"key":   "UUID",
					"value": "9ee79ab5-c846-433a-ae93-e2513d868614",
				},
				"billdescription":  "Cuenta Telefonia",
				"billStatus":       "unpaid",
				"currency":         "PAB",
				"amountDue":        "65",
				"minimumAmountDue": "65",
				"dueDate":          "2021-02-17",
				"metadata": []gin.H{
					{
						"key":   "billingCycle",
						"value": "19",
					},
					{
						"key":   "frequency",
						"value": "monthly",
					},
					{
						"key":   "accountId",
						"value": "3549050",
					},
					{
						"key":   "BalanceType",
						"value": "Debt",
					},
				},
			},
		},
	}
	cb.Context.JSON(http.StatusOK, mockData)
}

// Punto Pago Integrator
func (cb ControllerPuntoPagoGetBill) PuntoPagoGetBillIntegration(url string) {

	// request
	body, err := ioutil.ReadAll(cb.Context.Request.Body)
	if err != nil {
		cb.Context.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo leer el body"})
		return
	}
	buf := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		cb.Context.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la petición"})
		return
	}

	// header
	req.Header.Add("Signature", BuildSignatureHeaderFromBody(body))
	req.Header.Set("Content-Type", "application/xml")

	// transaction
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		cb.Context.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo ejecutar el request"})
		return
	}
	defer resp.Body.Close()

	// read response
	xmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cb.Context.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo leer el body de la respuesta"})
		return
	}

	cb.Context.Data(http.StatusOK, "application/xml", xmlData)
}

func BuildSignatureHeaderFromBody(body []byte) string {
	bodyString := string(body)
	input := bodyString + config.SettingsPuntoPago.HeaderEncryptedToken
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
