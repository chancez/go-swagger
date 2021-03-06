package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/go-swagger/go-swagger/examples/todo-list/models"
)

/*UpdateOneOK OK

swagger:response updateOneOK
*/
type UpdateOneOK struct {

	// In: body
	Payload *models.Item `json:"body,omitempty"`
}

// WriteResponse to the client
func (o *UpdateOneOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateOneDefault error

swagger:response updateOneDefault
*/
type UpdateOneDefault struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// WriteResponse to the client
func (o *UpdateOneDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
