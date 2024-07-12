package httpserver

type HttpServer struct {
	accountService     accountService
	datapageService    datapageService
	authService        authService
	customIndexHeaders []byte
}

func New(
	accountService accountService,
	datapageService datapageService,
	authService authService,
	customIndexHeaders []byte,
) *HttpServer {
	return &HttpServer{
		accountService:     accountService,
		datapageService:    datapageService,
		authService:        authService,
		customIndexHeaders: customIndexHeaders,
	}
}

func (hs *HttpServer) Start() {
	hs.router().Run(":8080")
}
