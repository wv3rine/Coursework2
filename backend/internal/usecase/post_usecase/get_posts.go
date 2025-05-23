package post_usecase

import (
	"context"
	"texts/internal/repository/postgres/post_repository"
	"texts/pkg/constants/utils"

	"github.com/pkg/errors"
)

type GetPostsReq struct {
	PostIds   []int64
	Names     []string
	Authors   []string
	Genres    []string
	Contents  []string
	EditorIds []int64
	TagIds    []int64
	TagNames  []string
	Statuses  []string
	Deleted   []bool
}

type GetPostResp struct {
	PostId      int64
	Name        string
	Author      string
	Genre       string
	Content     string
	EditorId    *int64
	EditorLogin *string
	TagId       int64
	TagName     string
	Status      string
	Deleted     bool
}

func (u *PostUC) GetPosts(ctx context.Context, getPostsReq GetPostsReq) ([]GetPostResp, error) {
	spanName := "PostUC.GetPosts"

	posts, err := u.postPGRepo.SelectPost(ctx, post_repository.SelectPostReq{
		PostIds:   getPostsReq.PostIds,
		Names:     getPostsReq.Names,
		Authors:   getPostsReq.Authors,
		Genres:    getPostsReq.Genres,
		Contents:  getPostsReq.Contents,
		EditorIds: getPostsReq.EditorIds,
		TagIds:    getPostsReq.TagIds,
		TagNames:  getPostsReq.TagNames,
		Statuses:  getPostsReq.Statuses,
		Deleted:   getPostsReq.Deleted,
	})
	if err != nil {
		return []GetPostResp{}, errors.Wrap(err, spanName)
	}

	return utils.MapArr(posts, func(post post_repository.SelectPostResp) GetPostResp {
		return GetPostResp{
			PostId:      post.PostId,
			Name:        post.Name,
			Author:      post.Author,
			Genre:       post.Genre,
			Content:     post.Content,
			EditorId:    post.EditorId,
			EditorLogin: post.EditorLogin,
			TagId:       post.TagId,
			TagName:     post.TagName,
			Status:      post.Status,
			Deleted:     post.Deleted,
		}
	}), nil
}
