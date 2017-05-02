package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	CommitHash   string
	BuildTime    string
	BaseURL      string
	ClientID     string
	ClientSecret string

	AccessToken string
	cfgFile     string
	verbose     bool
)

var RootCmd = &cobra.Command{
	Use:   "pCloud-cli",
	Short: "pCloud-cli is a command line interface to the pCloud API.",
	Long: `A command line interface to interact with pCloud storage.
More info can be found on github, http://github.com/storvik/pcloud-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute adds all child commands to the root command
func Execute(commithash, buildtime, baseurl, clientid, clientsecret string) {
	CommitHash = commithash
	BuildTime = buildtime
	BaseURL = baseurl
	ClientID = clientid
	ClientSecret = clientsecret

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pcloud-cli.json)")
	RootCmd.PersistentFlags().StringVar(&AccessToken, "token", "", "bearer token to access API, can be used when not using config file")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output for debugging")

	viper.SetDefault("author", "Petter S. Storvik <petterstorvik@gmail.com>")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.SetConfigName(".pcloud-cli")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if cfgFile != "" { // Use custom config file if --config flag set
		viper.SetConfigFile(cfgFile)
	}

	err := viper.ReadInConfig()
	if err != nil { // No config file found, authorize of token not set
		if AccessToken == "" {
			fmt.Println("Config file not found, authorize or pass token with --token")
		}
	} else {
		if verbose {
			fmt.Println("Configuration file, " + viper.ConfigFileUsed() + " found")
		}
		AccessToken = viper.GetString("access_token")
	}
}

type Pcloud struct {
	Endpoint    string
	Parameters  url.Values
	AccessToken string
	Body        io.Reader
	Headers     map[string]string
}

func NewPcloud() *Pcloud {
	return &Pcloud{Headers: make(map[string]string)}
}

// Query API endpoint with url parameters. If authorization is true, the authorization
// header is set. Returns json []byte and optional error from server.
func (p *Pcloud) Query() ([]byte, error) {
	var URL *url.URL
	URL, err := url.Parse(BaseURL)
	if err != nil {
		fmt.Println("Error: Could not parse base url")
		os.Exit(1)
	}

	URL.Path += p.Endpoint
	URL.RawQuery = p.Parameters.Encode()

	if verbose {
		fmt.Println("Query path: " + URL.String())
	}

	request, err := http.NewRequest("POST", URL.String(), p.Body)
	for key, value := range p.Headers {
		request.Header.Add(key, value)
	}
	if p.AccessToken != "" {
		request.Header.Add("Authorization", "Bearer "+p.AccessToken)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error: Could not query endpoint")
		os.Exit(1)
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	if verbose {
		fmt.Println("Response Status:", resp.Status)
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(responseBody, &dat); err != nil {
		panic(err)
	}

	if dat["result"].(float64) != 0 {
		return []byte{}, errors.New("Error " + strconv.FormatFloat(dat["result"].(float64), 'f', 0, 64) + ": " + dat["error"].(string))
	}

	return responseBody, nil
}
