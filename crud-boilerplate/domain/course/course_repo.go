package course

import (
	"github.com/AndrewJoyT/crud-boilerplate/database/schema"
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"
)

func (s *CourseFileUploadRequest) CourseFileUploadToDB() (*schema.File, *errors.RestErr) {
	fileData := &schema.File{
		CourseId: s.CourseId,
		Name:     s.FileName,
		FileUrl:  s.FileUrl,
	}

	if err := schema.Database.Create(fileData); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to save file upload to db: " + err.Error.Error())
	}
	return fileData, nil
}

func (s *CreateCourseRequest) CreateCourseToDB() (*schema.Course, *errors.RestErr) {
	course := &schema.Course{
		CompanyID: s.CompanyId,
		Name:      s.Name,
	}

	if err := schema.Database.Create(course); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to save course to db: " + err.Error.Error())
	}
	return course, nil
}

func GetCourseByIDFromDB(courseId string) (*CourseResponse, *errors.RestErr) {
	var courseSchema schema.Course
	if err := schema.Database.Preload("Files").Find(&courseSchema, "id = ?", courseId); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to get course by id: " + err.Error.Error())
	}
	course := CourseResponse{
		CourseId:  courseSchema.ID,
		CompanyId: courseSchema.CompanyID,
		Name:      courseSchema.Name,
		CreatedAt: courseSchema.CreatedAt,
		CreatedBy: courseSchema.CreatedBy,
		UpdateAt:  courseSchema.UpdatedAt,
		UpdatedBy: courseSchema.UpdatedBy,
	}
	for _, file := range courseSchema.Files {
		course.File = append(course.File, &CourseFileUploadResponse{
			FileId:    file.ID,
			CourseId:  file.CourseId,
			FileName:  file.Name,
			FileUrl:   file.FileUrl,
			CreatedAt: file.CreatedAt,
			CreatedBy: file.CreatedBy,
			UpdateAt:  file.UpdatedAt,
			UpdatedBy: file.UpdatedBy,
		})
	}

	return &course, nil
}

func (s *JoinCourseRequest) JoinCourseToDB() *errors.RestErr {
	fileData := &schema.UserRoom{
		UserUUID: s.UserId,
		RoomID:   s.RoomId,
	}

	if err := schema.Database.Create(fileData); err.Error != nil {
		return errors.NewInternalServerError("error when trying to save file upload to db: " + err.Error.Error())
	}
	return nil
}
