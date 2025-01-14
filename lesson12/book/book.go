package book

import "errors"

func BookCheck(BookName, AuthorName string) (string, error) {
	//创建一个新的异常类型
	if BookName == "" {
		return BookName, errors.New("error图书名称不能为空")
	}
	if AuthorName == "" {
		return AuthorName, errors.New("error图书作者名称不能为空")
	}
	return "图书名称是" + BookName + "图书的作者是" + AuthorName, nil
}
