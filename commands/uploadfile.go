package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/storvik/pcloud-cli/models"
)

var (
	uploadfileCmd = &cobra.Command{
		Use:   "upload [path to remote folder] [path to local file]",
		Short: "Upload local file to remote folder.",
		Long: `Upload given file to remote folder in pCloud.
Paths containing spaces should be wrapped in double quotes.`,

		Run: uploadfile,
	}
)

var (
	renameifexists bool
)

func init() {
	FileCmd.AddCommand(uploadfileCmd)
	uploadfileCmd.Flags().BoolVarP(&renameifexists, "renameifexists", "", false, "rename file if it exists")

	// Hidden / Aliased
}

func uploadfile(cmd *cobra.Command, args []string) {
	path := args[0]

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("filename", fi.Name())
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}

	parameters := url.Values{}
	parameters.Add("nopartial", strconv.Itoa(1))
	if renameifexists {
		parameters.Add("renameifexists", strconv.Itoa(1))
	}
	if len(args) < 2 {
		parameters.Add("path", "/")
	} else {
		if args[1][0] != 47 {
			parameters.Add("path", "/"+args[1])
		} else {
			parameters.Add("path", args[1])
		}
	}

	pcloud := NewPcloud()
	pcloud.Endpoint = "/uploadfile"
	pcloud.Parameters = parameters
	pcloud.AccessToken = ACCESS_TOKEN
	pcloud.Body = body
	pcloud.Headers["Content-Type"] = writer.FormDataContentType()

	resp, err := pcloud.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var response models.UploadfileResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Could not decode server response.")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("File successfully uploaded")

}
