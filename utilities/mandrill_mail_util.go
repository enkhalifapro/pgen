package utilities

import (
	"fmt"
	"github.com/mattbaird/gochimp"
	"github.com/spf13/viper"
)

type MandrillMailUtil struct {
}

func (m *MandrillMailUtil) SendTemplate(templateName string, toEmail string, fromEmail string, fromName string, subject string, globalVars []gochimp.Var) ([]gochimp.SendResponse, error) {
	mandrillKey := viper.GetString("mandrill.key")
	mandrillApi, err := gochimp.NewMandrill(mandrillKey)

	if err != nil {
		fmt.Println("Error instantiating client")
	}

	renderedTemplate, err := mandrillApi.TemplateRender(templateName, nil, globalVars)
	fmt.Println(renderedTemplate)
	if err != nil {
		fmt.Println("Error rendering template: %v", err)
		return nil, err
	}
	recipients := []gochimp.Recipient{
		{Email: toEmail},
	}

	message := gochimp.Message{
		Html:      renderedTemplate,
		Subject:   subject,
		FromEmail: fromEmail,
		FromName:  fromName,
		To:        recipients,
	}
	res, err := mandrillApi.MessageSend(message, false)

	if err != nil {
		fmt.Println("Error sending message ", err)
	}
	return res, err
}
