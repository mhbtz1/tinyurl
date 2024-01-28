package main

import (
  "net/http"
	"github.com/gin-gonic/gin"
)

type Server struct {
		Port int `yaml:"Port"`
		Address string `yaml:"Address"`
};

type ServerConfig struct {
	Server1 Server `yaml:"Server1"`
	Server2 Server `yaml:"Server2"`
	
	LoadBalancer struct {
		Port int `yaml:"Port"`
		Address string `yaml:"Address"`
	} `yaml:"LoadBalancer"`

};

type Body struct {
	url string `json:"url"`
};

func main() {
	router := gin.Default();

	router.GET("/", func (c *gin.Context){
		msg := "Hello from root page of localhost!"
		c.JSON(http.StatusAccepted, msg);	
	});
	

	router.POST("/api/compress-url", func(c *gin.Context) {
		body := Body{};
		if err:=c.BindJSON(&body); err!=nil {
			 c.AbortWithError(http.StatusBadRequest, err);
			 return;
		}
		res := CompressURL(body.url);
		c.JSON(http.StatusAccepted, &res);
	});

	router.Run("localhost:8080");
}

/*
func startServer(port int, h func(w http.ResponseWriter, r *http.Request), wg *sync.WaitGroup) {
	defer wg.Done()
	
	mux := http.NewServeMux()
	// Register a simple handler for the server
	mux.HandleFunc("/", h)

	// Start the server on the specified port
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on %s...\n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Printf("Error starting server on port %d: %v\n", port, err)
	}
}
*/
