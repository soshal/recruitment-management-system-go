package models

import (
    "github.com/jinzhu/gorm"
)

type Profile struct {
    gorm.Model
    ApplicantID      uint   `json:"applicant_id"`
    ResumeFileAddress string `json:"resume_file_address"`
    Skills           string `json:"skills"`
    Education        string `json:"education"`
    Experience       string `json:"experience"`
    Phone            string `json:"phone"`
}
