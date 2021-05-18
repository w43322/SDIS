package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
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

// func main() {
// 	parseCVE()
// }

func parseCVE() {
	data, _ := os.Open("./total.json")

	byteValue, _ := ioutil.ReadAll(data)

	var vul Vul
	json.Unmarshal(byteValue, &vul)
	total := 0.0
	for _, v := range vul.Vulnerabilities {
		total += v.CvssScore

		fmt.Println("Title :", v.Title, " Severity :", v.Severity, " CvssScore :", v.CvssScore, "CVE :", v.Identifiers.CVE)
	}
	fmt.Println(total)
}
