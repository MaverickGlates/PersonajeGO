package controller

import (
	"aprendiendo/domain"
	"aprendiendo/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type ItemController struct{ service *service.ItemService }

func NewItemController(service *service.ItemService) *ItemController {
	ItemController := &ItemController{}
	ItemController.service = service
	ItemController.itemRequestHandler()
	return ItemController
}

func (i *ItemController) itemRequestHandler() {
	http.HandleFunc("/item",
		func(w http.ResponseWriter, r *http.Request) {

			switch m := r.Method; m {
			case "POST":
				w.WriteHeader(200)
				var Item domain.Item
				err := json.NewDecoder(r.Body).Decode(&Item)
				if err != nil {
					fmt.Fprintf(w, "400 bad request")
					return
				}
				err = i.service.Save(Item)
				if err != nil {
					fmt.Fprint(w, "error de guardado: "+err.Error())
					return
				}
				fmt.Fprintf(w, "El dato ha sido guardado exitosamente")
			case "GET":
				w.WriteHeader(200)

				elements, err := i.service.Getall()

				elementAsBytes, _ := json.Marshal(elements)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Error retrieving data", err)
					return
				}
				w.Header().Add("Content-Type", "application/json")
				fmt.Fprint(w, string(elementAsBytes))
			default:
				w.WriteHeader(405)
				fmt.Fprintf(w, "405 Method Not Allowed")
			}

		})
}
