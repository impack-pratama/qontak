# qontak
Qontak is a software development kit (SDK) for whats app services . We built this sdk in order to help us integrate qontak-service api to our CRM Platform

# status
Currently this library under development

# Example
```
package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/impack-pratama/qontak/whatsapp"
	"github.com/impack-pratama/qontak/whatsapp/broadcasts/direct_message"
)

func main() {

    	token := "{your token from qontak}"

	service := whatsapp.NewService(token)
	broadcastService := service.GetBroadCastService()

	request := new(direct_message.SendDirectMessageRequest)
	request.ToNumber = "6281234567890" 
	request.ToName = "{Customer Name}"

	request.MessageTemplateId = uuid.MustParse({Message Template ID})
	request.ChannelIntegrationId = uuid.MustParse({Channel Integration ID})
	request.Language = &direct_message.Language{Code: "id"}

	param := new(direct_message.Parameter)
	param.Header = direct_message.HeaderParameter{}
	param.Header.Format = "DOCUMENT"

	param.Header.Parameters = append(param.Header.Parameters, direct_message.HeaderParameterKV{Key: "filename", Value: "sample.pdf"})
	param.Header.Parameters = append(param.Header.Parameters, direct_message.HeaderParameterKV{Key: "url", Value: "https://qontak-hub-development.s3.amazonaws.com/uploads/direct/files/01417dc5-9cd1-40b7-8900-d8b9fd6f250e/sample.pdf"})

	//Body Parameter depends on the template
	param.Body = append(param.Body, direct_message.BodyParameter{Key: "1", TextValue: "PT Angin Ribut", Value: "company_name"})
	param.Body = append(param.Body, direct_message.BodyParameter{Key: "2", TextValue: "Gembit Soultan", Value: "contact_name"})
	param.Body = append(param.Body, direct_message.BodyParameter{Key: "3", TextValue: "Indra", Value: "sales_name"})
	param.Body = append(param.Body, direct_message.BodyParameter{Key: "4", TextValue: "6281281231", Value: "phone_number"})

	request.Parameters = param

	var resp *direct_message.SendDirectMessageResponse
	var err error
	if resp, err = broadcastService.SendDirectMessage(request); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
```

# File Upload Example
```go
package main

import (
	"context"
	"fmt"
	"github.com/impack-pratama/qontak/whatsapp"
	"github.com/impack-pratama/qontak/whatsapp/file_uploader"
)

func main() {
	token := "{your token from qontak}"
	service := whatsapp.NewService(token)
	uploader := service.GetFileUploaderService()

	// Upload from S3
	s3Config := file_uploader.S3Config{
		AccessKey: "YOUR_ACCESS_KEY",
		SecretKey: "YOUR_SECRET_KEY",
		Region:    "ap-southeast-1",
		Bucket:    "my-bucket",
		Key:       "path/to/file.pdf",
	}
	resp, err := uploader.UploadFromS3(context.Background(), "file.pdf", s3Config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Data.Url)
}
```
