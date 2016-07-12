package rest

// Server represents information about a rest server.
type Server struct {
	port        string
	environment string
}

func (server Server) Make(config Config) (Server, error) {
	return server, nil
}

func (server *Server) Routes(r Routes) {

}

func (server Server) Listen() {

}
