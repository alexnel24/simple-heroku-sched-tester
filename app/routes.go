package app



func (s *Server) routes(){

	MakeAndRunScheduler(s)
	
	s.router.HandleFunc("/health-check", s.handleHealthCheck())
	s.router.HandleFunc("/time-grab", s.handleTimeGrab())
}