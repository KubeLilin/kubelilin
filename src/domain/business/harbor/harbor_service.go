package harbor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kubelilin/domain/database/models"
	"kubelilin/domain/dto"
	"net/http"
	"net/url"
	"strings"
)

//创建harbor项目
func CreateProject(projectName string, harborDetail models.ServiceConnectionDetails) error {
	reqData := dto.CreateHarborProjectDTO{
		ProjectName:  projectName,
		StorageLimit: -1,
		Metadata: dto.CreateHarborProjectMetadata{
			Public: "true",
		},
	}
	var scd dto.ServiceConnectionDetails
	json.Unmarshal([]byte(harborDetail.Detail), &scd)
	jsonStr, _ := json.Marshal(reqData)
	url, err := url.ParseRequestURI("https://" + scd.Repo + "/api/v2.0/projects")
	if err != nil {
		return err
	}
	req := http.Request{
		URL:    url,
		Header: map[string][]string{},
		Method: "POST",
		Body:   ioutil.NopCloser(bytes.NewReader(jsonStr)),
	}
	req.Header.Add("Authorization", "Basic "+scd.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(&req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func QueryProjectPage(projectName string, harborDetail models.ServiceConnectionDetails) (error, []dto.HarborProjectDTO) {
	var scd dto.ServiceConnectionDetails
	json.Unmarshal([]byte(harborDetail.Detail), &scd)
	reqUrlBuilder := strings.Builder{}
	reqUrlBuilder.WriteString("https://")
	reqUrlBuilder.WriteString(scd.Repo)
	reqUrlBuilder.WriteString("/api/v2.0/projects")
	reqUrlBuilder.WriteString("?page=1&page_size=15")
	if projectName != "" {
		reqUrlBuilder.WriteString("&name=")
		reqUrlBuilder.WriteString(projectName)
	}
	reqUrl, err := url.ParseRequestURI(reqUrlBuilder.String())
	if err != nil {
		return err, nil
	}
	req := http.Request{
		URL:    reqUrl,
		Header: map[string][]string{},
		Method: "GET",
	}
	req.Header.Add("Authorization", "Basic "+scd.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(&req)
	if err != nil {
		return err, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	fmt.Println(string(body))
	var pageData []dto.HarborProjectDTO
	json.Unmarshal(body, &pageData)
	return nil, pageData
}
