package models

import (
	bt "ApiManager/app/bootstrap"
)

type Upload struct {
	Id             int    `json:"id"`
	FileName       string `json:"file_name"`
	FilePath       string `json:"file_path"`
	UploadFileName string `json:"upload_file_name"`
	Size           int    `json:"size"`
	FileExt        string `json:"file_ext"`
	FileContent    []byte `json:"-"`
	Addtime        int64  `json:"addtime"`
}

func (u *Upload) Insert() error {
	_sql := "INSERT INTO `upload` (`file_name`, `file_path`, `upload_file_name`, `size`, `file_ext`, `file_content`, `addtime`) VALUES (?,?,?,?,?,?,?)"
	_, err := bt.DbCon.Exec(_sql, u.FileName, u.FilePath, u.UploadFileName, u.Size, u.FileExt, u.FileContent, u.Addtime)
	return err
}

func (u *Upload) GetByUploadFileName(name string) (Upload, error) {
	var out Upload
	_sql := "SELECT `id`,`file_name`,`file_path`,`upload_file_name`,`size`,`file_ext`,`file_content`,`addtime` FROM `upload` WHERE `upload_file_name` = ? LIMIT 1"
	err := bt.DbCon.QueryRow(_sql, name).Scan(
		&out.Id,
		&out.FileName,
		&out.FilePath,
		&out.UploadFileName,
		&out.Size,
		&out.FileExt,
		&out.FileContent,
		&out.Addtime,
	)
	return out, err
}
