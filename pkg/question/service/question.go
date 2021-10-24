package question

import (
	"context"
	"fmt"
	question "zkzj_resource_project/pkg/question/api"
	repo "zkzj_resource_project/pkg/question/repo"
)

type Server struct {
	question.QuestionServer
	repo *repo.Repo
}

func NewServer() *Server{

	repo := repo.NewRepo()

	return &Server{repo:repo}
}

//获取试题数目
func (s *Server) GetQuestionNum(ctx context.Context,in *question.QuestionNumRequest) (*question.QuestionNumResponse, error) {

	fmt.Println("我在执行了！！！")

	count := s.repo.GetQuestionNum(in.QuestionId)

	return &question.QuestionNumResponse{Count: count}, nil
}
