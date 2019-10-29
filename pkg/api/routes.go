package api

/* The routes for basic methods look like following:

/person/{ID} GET a <record> by its ID
/person/{ID}  DELETE a <record> by its ID
/person  POST add a <record>
/person  PUT update a <record>
/person  GET list all <records> */

func (s *Server) RoutesTest() {
	// Test routes
	s.Router.HandleFunc("/test/{name}", s.getTest).Methods("GET")
	s.Router.HandleFunc("/test", s.getListTest).Methods("GET")
	s.Router.HandleFunc("/test", s.addTest).Methods("POST")

	// Person routes
	s.Router.HandleFunc("/person/{id}", s.getPersonById).Methods("GET")
	s.Router.HandleFunc("/person/{id}", s.deletePersonById).Methods("DELETE")
	s.Router.HandleFunc("/person", s.getListPerson).Methods("GET")
	s.Router.HandleFunc("/person", s.addPerson).Methods("POST")
}
