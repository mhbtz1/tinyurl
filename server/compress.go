package main

import (
	"fmt"
	"time"
	"encoding/json"
	"os"
)

type URLEntry struct {
	Url string `json:"url"`
	Comp string `json:"comp"`
};

func CompressURL(s string) string {
		ctime := time.Now().Unix();
		hexstring := fmt.Sprintf("%X", ctime);
	  compressedUrl := "http://tinyurl.com/" + hexstring;
		
		urld := URLEntry {
				Url : s,
				Comp : compressedUrl
		};

		file, err := os.Create("local-url-mappings.json");
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close();
	
		// Create a JSON encoder
		encoder := json.NewEncoder(file);
	
		// Encode the data and write it to the file
		err := encoder.Encode(urld);
		if err != nil {
			fmt.Println("Error encoding and writing JSON:", err)
		}

	return compressedUrl;
}
