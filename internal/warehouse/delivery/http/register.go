package http

s.router.HandleFunc("/cell", s.getUsers()).Methods("GET")
s.router.HandleFunc("/cell", s.createUser()).Methods("POST")
s.router.HandleFunc("/cell/{id}", s.editUser()).Methods("POST")
s.router.HandleFunc("/cell/{id}", s.deleteUser()).Methods("DELETE")
s.router.HandleFunc("/cell/{id}", s.getUserByID()).Methods("GET")