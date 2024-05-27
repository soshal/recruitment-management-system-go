package controllers

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "recruitment-management-system/models"
    "recruitment-management-system/utils"
)

var applicantDB *gorm.DB

func InitApplicantController(database *gorm.DB) {
    applicantDB = database
}

func UploadResume(c *gin.Context) {
    file, _ := c.FormFile("resume")
    userID := c.GetUint("user_id")

    dst := filepath.Join("uploads", fmt.Sprintf("%d_%s", userID, filepath.Base(file.Filename)))
    if err := c.SaveUploadedFile(file, dst); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
        return
    }

    resumeData, err := utils.ParseResume(dst)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse resume"})
        return
    }

    var profile models.Profile
    profile.ApplicantID = userID
    profile.ResumeFileAddress = dst
    profile.Skills = resumeData.Skills
    profile.Education = resumeData.Education
    profile.Experience = resumeData.Experience
    profile.Phone = resumeData.Phone

    if err := applicantDB.Create(&profile).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save profile"})
        return
    }

    c.JSON(http.StatusOK, profile)
}

func GetJobs(c *gin.Context) {
    var jobs []models.Job
    if err := applicantDB.Find(&jobs).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs"})
        return
    }
    c.JSON(http.StatusOK, jobs)
}

func ApplyForJob(c *gin.Context) {
    jobID := c.Query("job_id")
    userID := c.GetUint("user_id")

    var job models.Job
    if err := applicantDB.First(&job, jobID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
        return
    }

    job.TotalApplications++
    if err := applicantDB.Save(&job).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply for job"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Applied successfully"})
}
