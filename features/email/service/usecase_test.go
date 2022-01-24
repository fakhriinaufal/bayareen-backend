package service_test

import (
	"bayareen-backend/features/email"
	"bayareen-backend/features/email/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	emailService email.Service
	emailRequest *email.Request
	emailData    interface{}
)

func setup() {
	emailService = service.NewEmailService("../../../config/config.toml")
	emailRequest = email.NewEmailRequest([]string{"fha.naufal06@gmail.com"}, "Test Email")
	emailData = email.NewInvoiceMailData("lorem", 20000, "mango", "http://google.com")
}

func TestSend(t *testing.T) {
	setup()

	t.Run("Test Case 1  | Success send email", func(t *testing.T) {
		err := emailService.Send("../../transaction/service/template/invoice.html", emailRequest, emailData)

		assert.Nil(t, err)
	})
}
