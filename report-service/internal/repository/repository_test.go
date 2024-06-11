package repository_test

import (
	"os"
	"report-service/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Save(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "data-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	repo := repository.NewRepository(tmpFile.Name())
	err = repo.Save("data1")
	assert.NoError(t, err)

	content, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.Contains(t, string(content), "data1")
}

func TestRepository_Load(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "data-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	repo := repository.NewRepository(tmpFile.Name())
	err = repo.Save("data1")
	assert.NoError(t, err)
	err = repo.Save("data2")
	assert.NoError(t, err)

	data, err := repo.Load()
	assert.NoError(t, err)
	assert.Equal(t, []string{"data1", "data2"}, data)
}

func TestRepository_Delete(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "data-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	repo := repository.NewRepository(tmpFile.Name())
	err = repo.Save("plate1,Red,SUV,A1,2024-06-11 10:00")
	assert.NoError(t, err)
	err = repo.Save("plate2,Blue,MPV,B2,2024-06-11 10:30")
	assert.NoError(t, err)

	err = repo.Delete("plate1")
	assert.NoError(t, err)

	data, err := repo.Load()
	assert.NoError(t, err)
	assert.Equal(t, []string{"plate2,Blue,MPV,B2,2024-06-11 10:30"}, data)
}
