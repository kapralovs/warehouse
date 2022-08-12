package http

s.router.HandleFunc("/product", s.getUsers()).Methods("GET")
	s.router.HandleFunc("/product", s.createUser()).Methods("POST")
	s.router.HandleFunc("/product/{id}", s.editUser()).Methods("POST")
	s.router.HandleFunc("/product/{id}", s.deleteUser()).Methods("DELETE")
	s.router.HandleFunc("/product/{id}", s.getUserByID()).Methods("GET")