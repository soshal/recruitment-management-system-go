package utils

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "os"
)

type ResumeData struct {
    Skills     string `json:"skills"`
    Education  string `json:"education"`
    Experience string `json:"experience"`
    Phone      string `json:"phone"`
}

func ParseResume(filePath string) (*ResumeData, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    request, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", file)
    if err != nil {
        return nil, err
    }

    request.Header.Set("Content-Type", "application/octet-stream")
    request.Header.Set("apikey", "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1")

    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to parse resume: status code %d", response.StatusCode)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    var resumeData ResumeData
    if err := json.Unmarshal(body, &resumeData); err != nil {
        return nil, err
    }

    return &resumeData, nil
}
