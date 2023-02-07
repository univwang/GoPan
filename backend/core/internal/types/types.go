// Code generated by goctl. DO NOT EDIT.
package types

type CommonReply struct {
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type FileUploadPrepareRequest struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
}

type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ParentId           int64  `json:"parent_id"`
}

type ShareBasicCreateReply struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               string `json:"size"`
	Path               string `json:"path"`
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity"`
}

type ShareBasicCreateRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ExpiredTime        int    `json:"expired_time"`
}

type UserFileMoveRequest struct {
	Identity       string `json:"identity"`
	ParentIdentity string `json:"parent_identity"`
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFolderCreateRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileListRequest struct {
	Id   int64 `json:"id, optional"`
	Page int   `json:"page, optional"`
	Size int   `json:"size, optional"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parent_id"`
	RepositoryIdentity string `json:"repository_identity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type FileUploadRequest struct {
	Hash string `json:"hash, optional"`
	Name string `json:"name, optional"`
	Ext  string `json:"ext, optional"`
	Size int64  `json:"size, optional"`
	Path string `json:"path, optional"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserDetailsRequest struct {
	Identity string `json:"identity"`
}
