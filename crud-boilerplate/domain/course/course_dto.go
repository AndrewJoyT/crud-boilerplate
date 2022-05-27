package course

import "time"

type CreateCourseRequest struct {
	CompanyId string `json:"company_id"`
	Name      string `json:"name"`
}

type JoinCourseRequest struct {
	RoomId  string `json:"room_id"`
	UserId  string `json:"user_id"`
	ByAdmin bool   `json:"by_admin"`
}

type JoinCourseResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CourseResponse struct {
	CourseId  string                      `json:"course_id"`
	CompanyId string                      `json:"company_id"`
	Name      string                      `json:"name"`
	CreatedAt time.Time                   `json:"created_at"`
	CreatedBy string                      `json:"created_by"`
	UpdateAt  time.Time                   `json:"updated_at"`
	UpdatedBy string                      `json:"updated_by"`
	File      []*CourseFileUploadResponse `json:"file"`
}

type CourseFileUploadRequest struct {
	CourseId string `json:"course_id"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
}

type CourseFileUploadResponse struct {
	FileId    string    `json:"file_id"`
	CourseId  string    `json:"course_id"`
	FileName  string    `json:"file_name"`
	FileUrl   string    `json:"file_url"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdateAt  time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
