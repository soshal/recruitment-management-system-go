package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "recruitment-management-system/controllers"
    "recruitment-management-system/models"
    "recruitment-management-system/utils"
)

func main() {
    r := gin.Default()

    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Job{})

    controllers.InitAuthController(db)

    authRoutes := r.Group("/api/auth")
    {
        authRoutes.POST("/signup", controllers.SignUp)
        authRoutes.POST("/login", controllers.Login)
    }

    applicantRoutes := r.Group("/api/applicant")
    applicantRoutes.Use(utils.AuthMiddleware())
    {
        applicantRoutes.POST("/uploadResume", controllers.UploadResume)
        applicantRoutes.GET("/jobs", controllers.GetJobs)
        applicantRoutes.GET("/jobs/apply", controllers.ApplyForJob)
    }

    adminRoutes := r.Group("/api/admin")
    adminRoutes.Use(utils.AuthMiddleware())
    {
        adminRoutes.POST("/job", controllers.CreateJob)
        adminRoutes.GET("/job/:job_id", controllers.GetJobDetails)
        adminRoutes.GET("/applicants", controllers.GetAllApplicants)
        adminRoutes.GET("/applicant/:applicant_id", controllers.GetApplicantDetails)
    }

    r.Run()
}
