package course

import (
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"
)

func (s *CourseFileUploadRequest) CourseFileUpload() (*CourseFileUploadResponse, *errors.RestErr) {
	file, err := s.CourseFileUploadToDB()
	if err != nil {
		return nil, err
	}
	return &CourseFileUploadResponse{
		FileId:    file.ID,
		CourseId:  file.CourseId,
		FileName:  file.Name,
		FileUrl:   file.FileUrl,
		CreatedAt: file.CreatedAt,
		CreatedBy: file.CreatedBy,
		UpdateAt:  file.UpdatedAt,
		UpdatedBy: file.UpdatedBy,
	}, nil
}

func (s *CreateCourseRequest) CreateCourse() (*CourseResponse, *errors.RestErr) {
	course, err := s.CreateCourseToDB()
	if err != nil {
		return nil, err
	}
	return &CourseResponse{
		CourseId:  course.ID,
		CompanyId: course.CompanyID,
		Name:      course.Name,
		CreatedAt: course.CreatedAt,
		CreatedBy: course.CreatedBy,
		UpdateAt:  course.UpdatedAt,
		UpdatedBy: course.UpdatedBy,
	}, nil
}

func GetCourseByID(courseId string) (*CourseResponse, *errors.RestErr) {
	course, err := GetCourseByIDFromDB(courseId)
	if err != nil {
		return nil, err
	}
	return &CourseResponse{
		CourseId:  course.CourseId,
		CompanyId: course.CompanyId,
		Name:      course.Name,
		CreatedAt: course.CreatedAt,
		CreatedBy: course.CreatedBy,
		UpdateAt:  course.UpdateAt,
		UpdatedBy: course.UpdatedBy,
		File:      course.File,
	}, nil
}

func (s *JoinCourseRequest) JoinCourse() (*JoinCourseResponse, *errors.RestErr) {
	return nil, nil
}
