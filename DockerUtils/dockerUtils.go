package DockerUtils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/fsouza/go-dockerclient"
)

type Vul struct {
	Vulnerabilities []struct {
		Title          string   `json:"title"`
		Credit         []string `json:"credit"`
		PackageName    string   `json:"packageName"`
		Language       string   `json:"language"`
		PackageManager string   `json:"packageManager"`
		Description    string   `json:"description"`
		Identifiers    struct {
			ALTERNATIVE []interface{} `json:"ALTERNATIVE"`
			CVE         []string      `json:"CVE"`
			CWE         []string      `json:"CWE"`
		} `json:"identifiers"`
		Severity             string        `json:"severity"`
		SeverityWithCritical string        `json:"severityWithCritical"`
		CvssScore            float64       `json:"cvssScore"`
		CVSSv3               string        `json:"CVSSv3"`
		Patches              []interface{} `json:"patches"`
		References           []struct {
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"references"`
		CreationTime       time.Time `json:"creationTime"`
		ModificationTime   time.Time `json:"modificationTime"`
		PublicationTime    time.Time `json:"publicationTime"`
		DisclosureTime     time.Time `json:"disclosureTime"`
		ID                 string    `json:"id"`
		NvdSeverity        string    `json:"nvdSeverity"`
		RelativeImportance string    `json:"relativeImportance"`
		Semver             struct {
			Vulnerable []string `json:"vulnerable"`
		} `json:"semver"`
		Exploit      string        `json:"exploit"`
		From         []string      `json:"from"`
		UpgradePath  []interface{} `json:"upgradePath"`
		IsUpgradable bool          `json:"isUpgradable"`
		IsPatchable  bool          `json:"isPatchable"`
		Name         string        `json:"name"`
		Version      string        `json:"version"`
	} `json:"vulnerabilities"`
}

type ListImagesOptions struct {
	Filters map[string][]string
	All     bool
	Digest  bool
	Filter  string
	Context context.Context
}

func ListDockerImage() []docker.APIImages {
	fmt.Println("Listing Docker image")

	client, err := docker.NewClientFromEnv()

	opts := docker.ListImagesOptions{
		All: true,
	}

	if err != nil {
		log.Println(err)
	}

	images, err := client.ListImages(opts)
	if err != nil {
		log.Println(err)
	}
	return images
}

func DockerImagePull(imageName string) {
	// now := time.Now()
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	createResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   []string{"touch", "/helloworld"},
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, createResp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	commitResp, err := cli.ContainerCommit(ctx, createResp.ID, types.ContainerCommitOptions{Reference: "helloworld"})
	if err != nil {
		panic(err)
	}

	fmt.Println(commitResp.ID)
}

// DockrImageTarchanges the docker image to tar file
func DockerImageToTar(exportName, dockerImageName string) bool {
	fmt.Printf("Converting docker image to tar file: %s\n", exportName)
	f, err := os.Create(exportName)

	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	client, err := docker.NewClientFromEnv()

	if err != nil {
		log.Println(err)
	}

	opts := docker.ExportImageOptions{
		Name:         dockerImageName,
		OutputStream: f,
	}

	if err := client.ExportImage(opts); err != nil {
		log.Println(err)
	}

	return true
}

// image to docker local repository
// docker image into local repository
func DockerImageLoad(loadImage string) bool {
	fmt.Println("Importing docker image into the repository")

	f, err := os.Open(loadImage)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	client, err := docker.NewClientFromEnv()

	if err != nil {
		log.Println(err)
	}

	opts := docker.LoadImageOptions{
		InputStream: f,
	}

	if err := client.LoadImage(opts); err != nil {
		log.Println(err)
		return false
	}

	return true
}

func RemoveDockerImage(imageName string) {
	fmt.Printf("Removing image %s \n", imageName)
	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Println(err)
	}
	client.RemoveImage(imageName)
	fmt.Printf("Remove completed\n")
}

// RunDockerSynkScan needs to be deprecated.
// Using command to use docker client

func RunDockerSynkScan(imageName string) string {

	// now := time.Now()
	// DockerImageLoad(imageName)
	// DockerImagePull(imageName)
	fmt.Printf("Scanning the docker image: %s\n", imageName)

	dockerScanCmd := exec.Command("docker", "scan", "--json", imageName)
	dockerScanCmdOut, err := dockerScanCmd.Output()

	if err != nil {
		fmt.Println("Error in dockerScanCmdout")
		log.Println(err)
	} else {
		fmt.Println("Succesfully printed out CVE")
	}

	fmt.Println("Changing the docker cve list to json file ")
	dockerVulnerString := BytesToString(dockerScanCmdOut)

	f, err := os.Create("./cveFile.json")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Successfully resulted in json file.")
	}
	defer f.Close()

	f.WriteString(dockerVulnerString)
	if parseCVE("cveFile") {
		return dockerVulnerString
	} else {
		return "Denied"
	}

}

func BytesToString(data []byte) string {
	return string(data[:])
}

func parseCVE(cveFileName string) bool {
	data, _ := os.Open("./" + cveFileName + ".json")

	byteValue, _ := ioutil.ReadAll(data)

	var vul Vul
	json.Unmarshal(byteValue, &vul)
	total := 0.0
	for _, v := range vul.Vulnerabilities {
		total += v.CvssScore

		fmt.Println("Title :", v.Title, " Severity :", v.Severity, " CvssScore :", v.CvssScore, " CVE :", v.Identifiers.CVE)
	}
	defer data.Close()
	fmt.Println(total)
	if total > 2000 {
		return false
	} else {
		return true
	}

}
