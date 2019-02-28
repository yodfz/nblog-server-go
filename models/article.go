package models

type Article struct {
	idx			int
	// 标题
	title		string
	// 缩略图
	thumbImg	string
	// 描述
	description string
	url			string
	// 正文内容
	content 	string
	// 标签
	tag			string
	// 分类
	category	string
	// 阅读数
	viewCount	int
	// 时间戳
	createDate	int
}