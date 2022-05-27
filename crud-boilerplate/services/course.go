package services

import (
	courseDomain "github.com/AndrewJoyT/crud-boilerplate/domain/course"
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"
)

var (
	CourseService courseServiceInterface = &courseService{}
)

type courseService struct {
}

type courseServiceInterface interface {
	CourseFileUpload(*courseDomain.CourseFileUploadRequest) (interface{}, *errors.RestErr)
	CreateCourse(*courseDomain.CreateCourseRequest) (interface{}, *errors.RestErr)
	GetCourseByID(string) (interface{}, *errors.RestErr)
	JoinCourse(payload *courseDomain.JoinCourseRequest) (interface{}, *errors.RestErr)
}

func (s *courseService) CourseFileUpload(payload *courseDomain.CourseFileUploadRequest) (interface{}, *errors.RestErr) {
	result, err := payload.CourseFileUpload()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *courseService) CreateCourse(payload *courseDomain.CreateCourseRequest) (interface{}, *errors.RestErr) {
	result, err := payload.CreateCourse()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *courseService) GetCourseByID(courseId string) (interface{}, *errors.RestErr) {
	res, err := courseDomain.GetCourseByID(courseId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *courseService) JoinCourse(payload *courseDomain.JoinCourseRequest) (interface{}, *errors.RestErr) {
	return nil, nil
}
