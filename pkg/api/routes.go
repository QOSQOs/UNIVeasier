package api

/* The routes for basic methods look like following:

/<model>/{ID} GET a <record> by its ID
/<model>/{ID}  DELETE a <record> by its ID
/<model>  POST add a <record>
/<model>  PUT update a <record>
/<model>  GET list all <records> */

func (s *Server) RoutesTest() {
	// Test routes
	s.Router.HandleFunc("/test/{name}", s.getTest).Methods("GET")
	s.Router.HandleFunc("/test", s.getListTest).Methods("GET")
	s.Router.HandleFunc("/test", s.addTest).Methods("POST")

	// Person routes
	s.Router.HandleFunc("/person/{id}", s.getPersonById).Methods("GET")
	s.Router.HandleFunc("/person", s.getListPerson).Methods("GET")
	s.Router.HandleFunc("/person", s.addPerson).Methods("POST")

	// TypeUniversity routes
	s.Router.HandleFunc("/typeUniversity/{id}", s.getTypeUniversityById).Methods("GET")
	s.Router.HandleFunc("/typeUniversity", s.getListTypeUniversity).Methods("GET")
	s.Router.HandleFunc("/typeUniversity", s.addTypeUniversity).Methods("POST")
}
