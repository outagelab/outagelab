package httpserver

type HttpServer struct {
	accountService  accountService
	datapageService datapageService
	authService     authService
}

func New(accountService accountService, datapageService datapageService, authService authService) *HttpServer {
	return &HttpServer{
		accountService:  accountService,
		datapageService: datapageService,
		authService:     authService,
	}
}

func (hs *HttpServer) Start() {
	hs.router().Run(":8080")
}
