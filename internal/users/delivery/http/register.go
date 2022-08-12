package http

s.router.HandleFunc("/user", s.getUsers()).Methods("GET")
	s.router.HandleFunc("/user", s.createUser()).Methods("POST")
	s.router.HandleFunc("/user/{id}", s.editUser()).Methods("POST")
	s.router.HandleFunc("/user/{id}", s.deleteUser()).Methods("DELETE")
	s.router.HandleFunc("/user/{id}", s.getUserByID()).Methods("GET")