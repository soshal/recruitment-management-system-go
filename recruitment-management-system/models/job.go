package models

import (
    "github.com/jinzhu/gorm"
)

type Job struct {
    gorm.Model
    Title            string `json:"title"`
    Description      string `json:"description"`
    PostedOn         string `json:"posted_on"`
    TotalApplications int    `json:"total_applications"`
    CompanyName      string `json:"company_name"`
    PostedBy         uint   `json:"posted_by"`
}
