package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerPuntoPagoPayBill struct {
	Context *gin.Context
}

// request structure definition
type PaymentRequest struct {
	Currency                       string `json:"currency"`
	AmountPaid                     string `json:"amountPaid"`
	ServiceProvider                string `json:"serviceProvider"`
	FinancialTransactionIdentifier string `json:"financialTransactionIdentifier"`
	PaymentType                    string `json:"paymentType"`
	RequestingOrganisation         string `json:"requestingOrganisation"`
	BrokerReference                string `json:"brokerReference"`
	Metadata                       []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"metadata"`
}

// entry point
func (cb ControllerPuntoPagoPayBill) PayBill() {
	var paymentReq PaymentRequest

	if err := cb.Context.ShouldBindJSON(&paymentReq); err != nil {
		cb.Context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Datos del pago recibido:\n%+v\n", paymentReq)

	// Obtener el valor del parámetro de la URL
	billReference := cb.Context.Param("billReference")
	fmt.Printf("billReference:\n%+v\n", billReference)

	// response
	cb.PayBillMockResponse()
	//cb.PuntoPagoPayBillIntegration(config.SettingsPuntoPago.PathPuntoPagoBackend)
}

// mocked response
func (cb ControllerPuntoPagoPayBill) PayBillMockResponse() {
	var mockData = gin.H{
		"identifier":                     "CBSAR310030",
		"currency":                       "PAB",
		"amountPaid":                     "50000",
		"serviceProvider":                "1002",
		"financialTransactionIdentifier": "113819",
		"requestingOrganisation":         "55F5E81CCB8602A4",
		"brokerReference":                "MTS",
		"paymentType":                    "fullpayment",
		"metadata":                       []interface{}{},
	}

	cb.Context.JSON(http.StatusOK, mockData)
}

// Punto Pago Integrator
func (cb ControllerPuntoPagoPayBill) PuntoPagoPayBillIntegration(url string) {
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
