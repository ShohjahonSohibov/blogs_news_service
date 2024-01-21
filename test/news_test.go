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

func TestCreateNews(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewNewsRepo(db)

	// Create a test news
	testNews := &models.CreateNews{
		Title:   "Test News",
		Content: "Test Content",
	}

	// Call the CreateNews function
	createdNews, err := repo.CreateNews(context.Background(), testNews)
	assert.NoError(t, err)
	assert.NotNil(t, createdNews)
	assert.NotEmpty(t, createdNews.Id)
	assert.Equal(t, testNews.Title, createdNews.Title)
	assert.Equal(t, testNews.Content, createdNews.Content)
	// Additional assertions for createdNews fields

	// Clean up (delete the test news)
	err = repo.DeleteNews(context.Background(), createdNews.Id)
	assert.NoError(t, err)
}

func TestGetSingleNews(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewNewsRepo(db)

	// Create a test news
	testNews := &models.CreateNews{
		Title:   "Test News",
		Content: "Test Content",
	}

	// Create the test news
	createdNews, err := repo.CreateNews(context.Background(), testNews)
	assert.NoError(t, err)
	assert.NotNil(t, createdNews)

	// Call the GetSingleNews function
	retrievedNews, err := repo.GetSingleNews(context.Background(), createdNews.Id)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedNews)
	assert.Equal(t, createdNews.Id, retrievedNews.Id)
	assert.Equal(t, createdNews.Title, retrievedNews.Title)
	assert.Equal(t, createdNews.Content, retrievedNews.Content)
	// Additional assertions for retrievedNews fields

	// Clean up (delete the test news)
	err = repo.DeleteNews(context.Background(), createdNews.Id)
	assert.NoError(t, err)
}

func TestGetListNews(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewNewsRepo(db)

	// Create a test news
	testNews := &models.CreateNews{
		Title:   "Test News",
		Content: "Test Content",
	}

	_, err = repo.CreateNews(context.Background(), testNews)
	assert.NoError(t, err)

	// Call the GetListNews function with a request
	req := &models.GetListNewsRequest{
		Limit:  1,
		Offset: 0,
		Title:  "Test",
	}

	res, err := repo.GetListNews(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.News)
	// Additional assertions for res fields

	// Clean up (delete the test news)
	err = repo.DeleteNews(context.Background(), res.News[0].Id)
	assert.NoError(t, err)
}

func TestUpdateNews(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewNewsRepo(db)

	// Create a test news
	testNews := &models.CreateNews{
		Title:   "Test News",
		Content: "Test Content",
	}

	createdNews, err := repo.CreateNews(context.Background(), testNews)
	assert.NoError(t, err)

	// Update the test news
	updateReq := &models.UpdateNews{
		Id:      createdNews.Id,
		Title:   "Updated Test News",
		Content: "Updated Test Content",
	}

	updatedNews, err := repo.UpdateNews(context.Background(), updateReq)
	assert.NoError(t, err)
	assert.NotNil(t, updatedNews)
	assert.Equal(t, updateReq.Id, updatedNews.Id)
	assert.Equal(t, updateReq.Title, updatedNews.Title)
	assert.Equal(t, updateReq.Content, updatedNews.Content)
	// Additional assertions for updatedNews fields

	// Clean up (delete the test news)
	err = repo.DeleteNews(context.Background(), createdNews.Id)
	assert.NoError(t, err)
}

func TestDeleteNews(t *testing.T) {
	db, err := helper.Setup(context.Background(), &config.Config{})
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewNewsRepo(db)

	// Create a test news
	testNews := &models.CreateNews{
		Title:   "Test News",
		Content: "Test Content",
	}

	createdNews, err := repo.CreateNews(context.Background(), testNews)
	assert.NoError(t, err)

	// Call the DeleteNews function
	err = repo.DeleteNews(context.Background(), createdNews.Id)
	assert.NoError(t, err)

	// Try to get the deleted news
	deletedNews, err := repo.GetSingleNews(context.Background(), createdNews.Id)
	assert.NoError(t, err)
	assert.Nil(t, deletedNews)
}
