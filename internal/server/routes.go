package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kapralovs/warehouse/internal/users"
)

func (s *server) homepage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := authorization(s.storage, w, r)
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, "Введите логин и пароль через пробел:")

			return
		}

		switch user.General.Role {
		case users.Manager:
			// fmt.Fprintln(w, view.Show(view.ManagerHomePage))
		case users.Dispatcher:
		case users.OrderPicker:

		}
	}
}
