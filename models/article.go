package models

type Article struct {
	Idx			int
	// 标题
	Title		string
	// 缩略图
	ThumbImg	string
	// 描述
	Description string
	Url			string
	// 正文内容
	Content 	string
	// 标签
	Tag			string
	// 分类
	Category	string
	// 阅读数
	ViewCount	int
	// 时间戳
	CreateDate	int
}