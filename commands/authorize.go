package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
	"github.com/storvik/pcloud-cli/config"
)

func init() {
	RootCmd.AddCommand(authorizeCmd)
}

var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Authorize with pCloud.",
	Long: `Authorization is neccessary to be able to interact with the pCloud API.
Will re-authorize with pCloud and rewrite config file. This command will
also be run if noe config file is present when running pcloud-cli`,

	Run: authorize,
}

func authorize(cmd *cobra.Command, args []string) {
	fmt.Println("pCloud-cli authorization started.")
	fmt.Println("This will delete the old configuration file.")
	fmt.Println("Open URL below in browser and copy the code to authenticate.")
	fmt.Println("https://my.pcloud.com/oauth2/authorize?response_type=code&client_id=" + CLIENT_ID)

	var code string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Code: ")
	code, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The code you entered: %s", code)

	parameters := url.Values{}
	parameters.Add("client_id", CLIENT_ID)
	parameters.Add("client_secret", CLIENT_SECRET)
	parameters.Add("code", strings.Replace(code, "\n", "", -1))

	pcloud := NewPcloud()
	pcloud.Endpoint = "/oauth2_token"
	pcloud.Parameters = parameters

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(resp, &dat); err != nil {
		panic(err)
	}

	var conf config.ConfigFile

	conf.UserID = int(dat["userid"].(float64))
	conf.AccessToken = dat["access_token"].(string)

	usr, _ := user.Current()
	configPath := usr.HomeDir

	config.WriteConfig(configPath, ".pcloud-cli", &conf)
}
