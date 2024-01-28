/*
package main

import (
	"fmt"
	"net/http"
	"sync"
	"load-balancer/concqueue"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

router := gin.Default();

router.GET("/", func (c *gin.Context){
		c.JSON(http.StatusAccepted, msg);	
});

router.POST("/api/compress-url", func(c *gin.Context) {
		body := Body{}
		if err:=context.BindJSON(&body); err!=nil{
				context.AbortWithError(http.StatusBadRequest,err);
				return;
		}
		res := compress_url(body.url);
		c.JSON(http.StatusAccepted, &res);
});
*/
