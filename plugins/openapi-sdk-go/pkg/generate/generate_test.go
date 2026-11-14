package generate

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testdataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
}

func TestGenerateSDK_Full(t *testing.T) {
	inputDir := filepath.Join(testdataDir(), "input")
	outputDir := t.TempDir()

	cfg := Config{
		InputDir:       inputDir,
		OutputDir:      outputDir,
		SpecFileName:   "openapi.yaml",
		PackageName:    "client",
		GenerateTypes:  true,
		GenerateClient: true,
	}

	err := GenerateSDK(cfg)
	require.NoError(t, err)

	typesPath := filepath.Join(outputDir, "types.go")
	typesContent, err := os.ReadFile(typesPath)
	require.NoError(t, err)
	typesStr := string(typesContent)

	assert.Contains(t, typesStr, "package client")
	assert.Contains(t, typesStr, "type Company struct")
	assert.Contains(t, typesStr, "type CompanyCreate struct")
	assert.Contains(t, typesStr, "type CompanyUpdate struct")
	assert.Contains(t, typesStr, "Name")
	assert.Contains(t, typesStr, "time.Time")

	clientPath := filepath.Join(outputDir, "client.go")
	clientContent, err := os.ReadFile(clientPath)
	require.NoError(t, err)
	clientStr := string(clientContent)

	assert.Contains(t, clientStr, "package client")
	assert.Contains(t, clientStr, "type Client struct")
	assert.Contains(t, clientStr, "func NewClient")
	assert.Contains(t, clientStr, "ListCompany")
	assert.Contains(t, clientStr, "CreateCompany")
	assert.Contains(t, clientStr, "GetCompany")
	assert.Contains(t, clientStr, "UpdateCompany")
	assert.Contains(t, clientStr, "DeleteCompany")
}

func TestGenerateSDK_TypesOnly(t *testing.T) {
	inputDir := filepath.Join(testdataDir(), "input")
	outputDir := t.TempDir()

	cfg := Config{
		InputDir:       inputDir,
		OutputDir:      outputDir,
		SpecFileName:   "openapi.yaml",
		PackageName:    "models",
		GenerateTypes:  true,
		GenerateClient: false,
	}

	err := GenerateSDK(cfg)
	require.NoError(t, err)

	typesPath := filepath.Join(outputDir, "types.go")
	_, err = os.ReadFile(typesPath)
	require.NoError(t, err)

	clientPath := filepath.Join(outputDir, "client.go")
	_, err = os.Stat(clientPath)
	assert.True(t, os.IsNotExist(err))
}

func TestParseSpec(t *testing.T) {
	inputDir := filepath.Join(testdataDir(), "input")
	specBytes, err := os.ReadFile(filepath.Join(inputDir, "openapi.yaml"))
	require.NoError(t, err)

	doc, err := ParseSpec(specBytes)
	require.NoError(t, err)
	assert.NotNil(t, doc)
	assert.Equal(t, "Test API", doc.Info.Title)
	assert.Greater(t, len(doc.Paths), 0)
	assert.Greater(t, len(doc.Schemas), 0)
}

func TestGenerateTypes_Schemas(t *testing.T) {
	inputDir := filepath.Join(testdataDir(), "input")
	specBytes, err := os.ReadFile(filepath.Join(inputDir, "openapi.yaml"))
	require.NoError(t, err)

	doc, err := ParseSpec(specBytes)
	require.NoError(t, err)

	types, err := GenerateTypes(doc, "testpkg")
	require.NoError(t, err)

	assert.Contains(t, types, "package testpkg")
	assert.Contains(t, types, "type Company struct")
	assert.Contains(t, types, "Id")
	assert.Contains(t, types, "Name")
	assert.Contains(t, types, "CreatedAt")
	assert.Contains(t, types, "UpdatedAt")
}

func TestGenerateClient_Operations(t *testing.T) {
	inputDir := filepath.Join(testdataDir(), "input")
	specBytes, err := os.ReadFile(filepath.Join(inputDir, "openapi.yaml"))
	require.NoError(t, err)

	doc, err := ParseSpec(specBytes)
	require.NoError(t, err)

	client, err := GenerateClient(doc, "testpkg")
	require.NoError(t, err)

	assert.Contains(t, client, "package testpkg")
	assert.Contains(t, client, "func (c *Client) ListCompany")
	assert.Contains(t, client, "func (c *Client) CreateCompany")
	assert.Contains(t, client, "func (c *Client) GetCompany")
	assert.Contains(t, client, "func (c *Client) UpdateCompany")
	assert.Contains(t, client, "func (c *Client) DeleteCompany")
}

func TestConfig_Resolve(t *testing.T) {
	cfg := Config{}
	cfg.Resolve()
	assert.Equal(t, "openapi.yaml", cfg.SpecFileName)
	assert.Equal(t, "client", cfg.PackageName)
	assert.Equal(t, "generated/client", cfg.PackagePath)
}
