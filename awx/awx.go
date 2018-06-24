package awx

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	. "./structs"
	"github.com/urfave/cli"
)

func createClient() (client *http.Client) {
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true},
		},
	}
	return client
}

func createURL(c *cli.Context) string {
	var url string
	if c.Bool("ssl") {
		url = "https://" + c.String("host")
	} else {
		url = "http://" + c.String("host")
	}
	return url
}

func getLaunchUrl(c *cli.Context, base_url string) (string, error) {
	client := createClient()
	url := base_url + "/api/v1/job_templates"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.String("username"), c.String("password"))
	res, _ := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode == 200 {
		data := new(JobTemplates)
		body, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(body, data)

		for _, v := range data.Results {
			if v.Name == c.String("template") {
				launch_url := base_url + v.Related.Launch
				return launch_url, nil
			}
		}
	}

	return "", errors.New("1")
}

func postUrl(c *cli.Context, url string) error {
	var body []byte
	if c.String("extra") != "" {
		body = []byte(c.String("extra"))
	} else {
		body = nil
	}

	client := createClient()
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.String("username"), c.String("password"))
	res, _ := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode == 201 {
		return nil
	}

	return errors.New("1")
}

func doMain(c *cli.Context) {
	base_url := createURL(c)
	if u, err := getLaunchUrl(c, base_url); err != nil {
		os.Exit(1)
	} else {
		if err := postUrl(c, u); err != nil {
			os.Exit(1)
		}
	}
}

func Do() {
	app := cli.NewApp()
	app.Name = "awx-exec-job"
	app.Version = version
	app.Usage = "Tool to execute AWX job"
	app.Author = "sky_joker"
	app.Flags = flags
	app.Action = doMain
	app.Run(os.Args)
}
