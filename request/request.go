package request

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Print() {
	url := "https://reetika.pr2.harness.io/gateway/ssca-manager/v1/orgs/default/projects/ssca/orchestration/tlSUzKq1RE2ubz1SMB_5SQ/sbom-upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	mimeHeader1 := make(map[string][]string)
	mimeHeader1["Content-Disposition"] = append(mimeHeader1["Content-Disposition"], "form-data; name=\"metadata\"")
	mimeHeader1["Content-Type"] = append(mimeHeader1["Content-Type"], "application/json")
	fieldWriter1, _ := writer.CreatePart(mimeHeader1)
	fieldWriter1.Write([]byte("{\"format\":\"\",\"name\":\"node_sbom_tlSUzKq1RE2ubz1SMB_5SQ\",\"pipeline_execution_id\":\"_majnQU3QDu_tfVViVpYYg\",\"pipeline_id\":\"syft\",\"sequence_id\":\"12\",\"stage_execution_id\":\"\",\"stage_id\":\"sbom\",\"step_execution_id\":\"\",\"step_id\":\"\",\"tool_name\":\"\",\"tool_version\":\"\"}"))

	file, errFile2 := os.Open("/Users/uday/Downloads/sbom-cy.json")
	defer file.Close()
	part2,
		errFile2 := writer.CreateFormFile("sbom", filepath.Base("/Users/uday/Downloads/sbom-cy.json"))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "reetika.pr2.harness.io")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdXRoVG9rZW4iOiI2NWYxYWNmYmI0ODExZTM0MmE3MzQ0MjMiLCJpc3MiOiJIYXJuZXNzIEluYyIsImV4cCI6MTcxMDU1NTU0NSwiZW52IjoiZ2F0ZXdheSIsImlhdCI6MTcxMDQ2OTA4NX0.1EVAfAGC5FjfZa4_Cv5PdFH7wxWkWpyWadkuxjpfzY4")
	req.Header.Add("cookie", "mutiny.user.token=63b873f3-9f47-4649-b3b1-5ad8b45dc121; _hjSessionUser_2892596=eyJpZCI6IjVkMjQyN2Y4LTFkMjMtNWIzMS1iY2RiLTJjNGRlYzg0NDRjYiIsImNyZWF0ZWQiOjE3MDg1ODY3NDc2OTksImV4aXN0aW5nIjp0cnVlfQ==; intercom-device-id-s96ohxu3=aa4d9138-a502-409f-9a58-2030a7a6cce0; _gcl_au=1.1.48451475.1708664630; _ga=GA1.1.1007088312.1708664631; intercom-id-iw4sgfzm=19da7c24-e178-4d67-840d-71ef72151ba4; intercom-device-id-iw4sgfzm=3ac23421-8d74-47af-ab64-d70ac2192c20; ajs_anonymous_id=b35d3acf-cae7-4c4d-b9e8-d4954e0e3232; cb_user_id=null; cb_group_id=null; cb_anonymous_id=%228e5e86cc-baed-403c-aabf-5b1d4b32ab47%22; ajs_user_id=uday.vunnam@harness.io; OptanonAlertBoxClosed=2024-02-29T12:28:29.014Z; intercom-session-s96ohxu3=bWtmMnA0aFg5Zlp3cUZVQ3ZBNzFhMUhUM3lTSzM5Vm5VNURxSzJOSFRNbmpqb2g2bnhMZ2RoT1BSeGZRSmhIbC0tUWhpMkFuNGpCNG0xT2RqS1pSVzRndz09--7eecd3c0e701fcc6b015a18a8d1ca29cfe49a9e5; _ga_246280343=GS1.1.1710137811.9.0.1710137865.0.0.0; routingId=65f1acfbb4811e342a734424; OptanonConsent=isGpcEnabled=0&datestamp=Wed+Mar+13+2024+20%3A19%3A51+GMT%2B0530+(India+Standard+Time)&version=202402.1.0&browserGpcFlag=0&isIABGlobal=false&hosts=&consentId=248a4e9b-1288-4f65-a532-d5ccd4134c4b&interactionCount=2&isAnonUser=1&landingPath=NotLandingPage&groups=C9990%3A1%2CC0001%3A1%2CC0003%3A1%2CC0004%3A1%2CC0002%3A1&AwaitingReconsent=false&geolocation=IN%3BKA; _ga_46758J5H8P=GS1.1.1710343447.10.0.1710343447.60.0.0; ADRUM=s=1710417058100&r=https%3A%2F%2Fapp.harness.io%2Fng%2Faccount%2FvpCkHKsDSxK9_KYfjCTMKA%2Fmodule%2Fci%2Forgs%2Fdefault%2Fprojects%2FSSCS%2Fpipelines%2Fssca_plugin_build%2Fexecutions%2F6GZpTYSER5aIpKOk64nrTg%2Fpipeline; intercom-session-iw4sgfzm=a3VlQWVab09aZk4xMzNFY0NLdkdpTUJEdWxnZ25kNTgrUWZ3UEdVaVJnWHZxSGVKRStjTEt3SWpWUmdRbVZQdC0tMEJlREZITGlxeFlhbHZ4WUxNVURGQT09--d6cf971dc70cff20c70d8c53c83914f7266693e4; token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdXRoVG9rZW4iOiI2NWYxYWNmYmI0ODExZTM0MmE3MzQ0MjMiLCJpc3MiOiJIYXJuZXNzIEluYyIsImV4cCI6MTcxMDU1NTU0NSwiZW52IjoiZ2F0ZXdheSIsImlhdCI6MTcxMDQ2OTA4NX0.1EVAfAGC5FjfZa4_Cv5PdFH7wxWkWpyWadkuxjpfzY4")
	req.Header.Add("harness-account", "w0MlMf5oQiCZ8U_q9V7iFA")
	req.Header.Add("origin", "https://reetika.pr2.harness.io")
	req.Header.Add("referer", "https://reetika.pr2.harness.io/ng/account/w0MlMf5oQiCZ8U_q9V7iFA/ssca/orgs/default/projects/ssca/artifacts")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Google Chrome\";v=\"122\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdXRoVG9rZW4iOiI2NWYxYWNmYmI0ODExZTM0MmE3MzQ0MjMiLCJpc3MiOiJIYXJuZXNzIEluYyIsImV4cCI6MTcxMDU1NTU0NSwiZW52IjoiZ2F0ZXdheSIsImlhdCI6MTcxMDQ2OTA4NX0.1EVAfAGC5FjfZa4_Cv5PdFH7wxWkWpyWadkuxjpfzY4")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
