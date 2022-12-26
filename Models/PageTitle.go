package Models

import "github.com/jinzhu/gorm"

type PageData struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Data         struct {
		Size       string `json:"size"`
		Number     string `json:"number"`
		TotalPages int    `json:"total-pages"`
		TotalItems int    `json:"total_items"`
		PageData   []struct {
			gorm.Model
			PosterId      int         `json:"poster_id"`
			CollegeId     interface{} `json:"college_id"`
			Content       string      `json:"content"`
			Attachments   []string    `json:"attachments"`
			Topic         string      `json:"topic"`
			Type          int         `json:"type"`
			Status        int         `json:"status"`
			Private       int         `json:"private"`
			CommentNumber int         `json:"comment_number"`
			PraiseNumber  int         `json:"praise_number"`
			Mobile        interface{} `json:"mobile"`
			NewColumn     interface{} `json:"new_column"`
			Poster        struct {
				Id        int    `json:"id"`
				Nickname  string `json:"nickname"`
				Avatar    string `json:"avatar"`
				Gender    int    `json:"gender"`
				CreatedAt string `json:"created_at"`
				Type      int    `json:"type"`
			} `json:"poster"`
			Praises []struct {
				Id        int         `json:"id"`
				OwnerId   int         `json:"owner_id"`
				ObjType   int         `json:"obj_type"`
				CollegeId interface{} `json:"college_id"`
				UserId    int         `json:"user_id"`
				Nickname  string      `json:"nickname"`
				Avatar    string      `json:"avatar"`
			} `json:"praises"`
			Comments []struct {
				Id           int         `json:"id"`
				CommenterId  int         `json:"commenter_id"`
				ObjId        int         `json:"obj_id"`
				CollegeId    interface{} `json:"college_id"`
				Content      string      `json:"content"`
				Attachments  interface{} `json:"attachments"`
				RefCommentId interface{} `json:"ref_comment_id"`
				ObjType      int         `json:"obj_type"`
				Type         int         `json:"type"`
				Status       int         `json:"status"`
				CreatedAt    string      `json:"created_at"`
				UpdatedAt    string      `json:"updated_at"`
				DeletedAt    interface{} `json:"deleted_at"`
				Author       int         `json:"author"`
				Commenter    struct {
					Id        int    `json:"id"`
					Nickname  string `json:"nickname"`
					Avatar    string `json:"avatar"`
					Text      string `json:"text"`
					Supertube int    `json:"supertube"`
				} `json:"commenter"`
				RefComment string `json:"ref_comment"`
				CanDelete  bool   `json:"can_delete"`
			} `json:"comments"`
			Follow    bool `json:"follow"`
			CanDelete bool `json:"can_delete"`
			CanChat   bool `json:"can_chat"`
			Supertube int  `json:"supertube"`
		} `json:"page_data"`
	} `json:"data"`
}

type HomeData struct {
	Id            int         `json:"id"`
	PosterId      int         `json:"poster_id"`
	CollegeId     interface{} `json:"college_id"`
	Content       string      `json:"content"`
	Attachments   []string    `json:"attachments"`
	Topic         string      `json:"topic"`
	Type          int         `json:"type"`
	Status        int         `json:"status"`
	Private       int         `json:"private"`
	CommentNumber int         `json:"comment_number"`
	PraiseNumber  int         `json:"praise_number"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	DeletedAt     interface{} `json:"deleted_at"`
	Mobile        interface{} `json:"mobile"`
	NewColumn     interface{} `json:"new_column"`
	Poster        struct {
		Id        int    `json:"id"`
		Nickname  string `json:"nickname"`
		Avatar    string `json:"avatar"`
		Gender    int    `json:"gender"`
		CreatedAt string `json:"created_at"`
		Type      int    `json:"type"`
	} `json:"poster"`
	Praises []struct {
		Id        int         `json:"id"`
		OwnerId   int         `json:"owner_id"`
		ObjType   int         `json:"obj_type"`
		CollegeId interface{} `json:"college_id"`
		UserId    int         `json:"user_id"`
		Nickname  string      `json:"nickname"`
		Avatar    string      `json:"avatar"`
	} `json:"praises"`
	Comments []struct {
		Id           int         `json:"id"`
		CommenterId  int         `json:"commenter_id"`
		ObjId        int         `json:"obj_id"`
		CollegeId    interface{} `json:"college_id"`
		Content      string      `json:"content"`
		Attachments  interface{} `json:"attachments"`
		RefCommentId interface{} `json:"ref_comment_id"`
		ObjType      int         `json:"obj_type"`
		Type         int         `json:"type"`
		Status       int         `json:"status"`
		CreatedAt    string      `json:"created_at"`
		UpdatedAt    string      `json:"updated_at"`
		DeletedAt    interface{} `json:"deleted_at"`
		Author       int         `json:"author"`
		Commenter    struct {
			Id        int    `json:"id"`
			Nickname  string `json:"nickname"`
			Avatar    string `json:"avatar"`
			Text      string `json:"text"`
			Supertube int    `json:"supertube"`
		} `json:"commenter"`
		RefComment string `json:"ref_comment"`
		CanDelete  bool   `json:"can_delete"`
	} `json:"comments"`
	Follow    bool `json:"follow"`
	CanDelete bool `json:"can_delete"`
	CanChat   bool `json:"can_chat"`
	Supertube int  `json:"supertube"`
}

type SendTitle struct {
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Data         UserToPic `json:"data"`
}

//type Title struct {
//	gorm.Model
//	UserId        int      `json:"user_id"`
//	AppId         int      `json:"app_id"`
//	UserType      int      `json:"user_type"`
//	Title         string   `json:"title"`
//	Content       string   `json:"content"`
//	Attachments   []string `json:"attachments"`
//	PraiseNumber  int      `json:"praise_number"`
//	ViewNumber    int      `json:"view_number"`
//	CommentNumber int      `json:"comment_number"`
//	Status        int      `json:"status"`
//}
