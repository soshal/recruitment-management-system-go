package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "recruitment-management-system/models"
)

var adminDB *gorm.DB

func InitAdminController(database *gorm.DB) {
    adminDB = database
}

func CreateJob(c *gin.Context) {
    var job models.Job
    if err := c.ShouldBindJSON(&job); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    job.PostedBy = c.GetUint("user_id")
    if err := adminDB.Create(&job).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
        return
    }
    c.JSON(http.StatusOK, job)
}

func GetJobDetails(c *gin.Context) {
    jobID := c.Param("job_id")

    var job models.Job
    if err := adminDB.First(&job, jobID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
        return
    }

    var applicants []models.Profile
    if err := adminDB.Where("applicant_id IN (SELECT applicant_id FROM profiles WHERE job_id = ?)", jobID).Find(&applicants).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch applicants"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"job": job, "applicants": applicants})
}

func GetAllApplicants(c *gin.Context) {
    var users []models.User
    if err := adminDB.Where("user_type = ?", "Applicant").Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch applicants"})
        return
    }
    c.JSON(http.StatusOK, users)
}

func GetApplicantDetails(c *gin.Context) {
    applicantID := c.Param("applicant_id")

    var profile models.Profile
    if err := adminDB.Where("applicant_id = ?", applicantID).First(&profile).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
        return
    }

    c.JSON(http.StatusOK, profile)
}
