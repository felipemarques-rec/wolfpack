/*
Package app is the main runtime package for Efs2. This package holds all of the logging and task execution controls.
*/
package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"wolfpack/config"

	"github.com/fatih/color"
	"github.com/tgglv/wc-api-go/client"
	"github.com/tgglv/wc-api-go/options"
)

// Run is the primary execution function for this application.
func Run(cfg config.Config) error {
	// Execute
	// var wg sync.WaitGroup
	// var errCount int

	factory := client.Factory{}
	c := factory.NewClient(options.Basic{
		URL:    cfg.Url,
		Key:    cfg.Key,
		Secret: cfg.Secret,
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v1",
		},
	})

	if r, err := c.Get("subscriptions", nil); err != nil {
		color.Red("%s", err)
		log.Fatal(err)
	} else if r.StatusCode != http.StatusOK {
		color.Red("Unexpected StatusCode: %s", r)
	} else {
		defer r.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
			color.Red("Error: %s", err)
		} else {
			fmt.Println(string(bodyBytes))
		}
	}

	// for _, h := range xxx {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 	}()

	// }
	// wg.Wait()

	// if errCount > 0 {
	// 	return fmt.Errorf("Execution failed with %d errors", errCount)
	// }

	return nil
}
