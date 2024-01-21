package test

import (
	"context"
	"news_blogs_service/config"
	"news_blogs_service/models"
	"news_blogs_service/pkg/helper"
	"news_blogs_service/storage/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBlog(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewBlogsRepo(db)

	// Create a test blog
	testBlog := &models.CreateBlog{
		Title:   "Test Blog",
		Content: "Test Content",
	}

	// Call the CreateBlog function
	createdBlog, err := repo.CreateBlog(context.Background(), testBlog)
	assert.NoError(t, err)
	assert.NotNil(t, createdBlog)
	assert.NotEmpty(t, createdBlog.Id)
	assert.Equal(t, testBlog.Title, createdBlog.Title)
	assert.Equal(t, testBlog.Content, createdBlog.Content)
	// Additional assertions for createdBlog fields

	// Clean up (delete the test blog)
	err = repo.DeleteBlog(context.Background(), createdBlog.Id)
	assert.NoError(t, err)
}

func TestGetSingleBlog(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewBlogsRepo(db)

	// Create a test blog
	testBlog := &models.CreateBlog{
		Title:   "Test Blog",
		Content: "Test Content",
	}

	// Create the test blog
	createdBlog, err := repo.CreateBlog(context.Background(), testBlog)
	assert.NoError(t, err)
	assert.NotNil(t, createdBlog)

	// Call the GetSingleBlog function
	retrievedBlog, err := repo.GetSingleBlog(context.Background(), createdBlog.Id)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedBlog)
	assert.Equal(t, createdBlog.Id, retrievedBlog.Id)
	assert.Equal(t, createdBlog.Title, retrievedBlog.Title)
	assert.Equal(t, createdBlog.Content, retrievedBlog.Content)
	// Additional assertions for retrievedBlog fields

	// Clean up (delete the test blog)
	err = repo.DeleteBlog(context.Background(), createdBlog.Id)
	assert.NoError(t, err)
}

func TestGetListBlogs(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewBlogsRepo(db)

	// Create a test blog
	testBlog := &models.CreateBlog{
		Title:   "Test Blog",
		Content: "Test Content",
	}

	_, err = repo.CreateBlog(context.Background(), testBlog)
	assert.NoError(t, err)

	// Call the GetListBlogs function with a request
	req := &models.GetListBlogsRequest{
		Limit:  1,
		Offset: 0,
		Title:  "Test",
	}

	res, err := repo.GetListBlogs(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Blogs)
	// Additional assertions for res fields

	// Clean up (delete the test blog)
	err = repo.DeleteBlog(context.Background(), res.Blogs[0].Id)
	assert.NoError(t, err)
}

func TestUpdateBlog(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewBlogsRepo(db)

	// Create a test blog
	testBlog := &models.CreateBlog{
		Title:   "Test Blog",
		Content: "Test Content",
	}

	createdBlog, err := repo.CreateBlog(context.Background(), testBlog)
	assert.NoError(t, err)

	// Update the test blog
	updateReq := &models.UpdateBlog{
		Id:      createdBlog.Id,
		Title:   "Updated Test Blog",
		Content: "Updated Test Content",
	}

	updatedBlog, err := repo.UpdateBlog(context.Background(), updateReq)
	assert.NoError(t, err)
	assert.NotNil(t, updatedBlog)
	assert.Equal(t, updateReq.Id, updatedBlog.Id)
	assert.Equal(t, updateReq.Title, updatedBlog.Title)
	assert.Equal(t, updateReq.Content, updatedBlog.Content)
	// Additional assertions for updatedBlog fields

	// Clean up (delete the test blog)
	err = repo.DeleteBlog(context.Background(), createdBlog.Id)
	assert.NoError(t, err)
}

func TestDeleteBlog(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewBlogsRepo(db)

	// Create a test blog
	testBlog := &models.CreateBlog{
		Title:   "Test Blog",
		Content: "Test Content",
	}

	createdBlog, err := repo.CreateBlog(context.Background(), testBlog)
	assert.NoError(t, err)

	// Call the DeleteBlog function
	err = repo.DeleteBlog(context.Background(), createdBlog.Id)
	assert.NoError(t, err)

	// Try to get the deleted blog
	deletedBlog, err := repo.GetSingleBlog(context.Background(), createdBlog.Id)
	assert.NoError(t, err)
	assert.Nil(t, deletedBlog)
}
