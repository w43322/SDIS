package IPFSUtils

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func InitIPFSShell() {
	sh = shell.NewShell("localhost:5001")
}

// AddFileOneIPFS Adds the file into ipfs
func AddFileOnIPFS(filename string) string {

	fmt.Printf("Adding %s on IPFS \n", filename)

	// sh := shell.NewShell("localhost:5001")
	// opens file
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	// cid, err := sh.Add(strings.NewReader("Hello IPFS it`s been a long time"))
	var r = f

	cid, err := sh.Add(r)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		// os.Exit(1)
	}
	defer f.Close()

	err = os.Remove(filename)
	if err != nil {
		fmt.Println("The file could not be removed. Error message:")
		fmt.Println(err)
	}

	fmt.Printf("added %s\n", cid)

	return cid

}

func GetFileFromIPFS(hash, oFilename string) string {
	outputName := "./" + oFilename
	err := sh.Get(hash, outputName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Successfully downloaded docker image and saved as %s\n", outputName)
	}
	return outputName
}
