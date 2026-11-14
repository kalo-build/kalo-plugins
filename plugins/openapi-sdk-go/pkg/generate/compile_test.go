package generate_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/kalo-build/go-util/assertfile"
	"github.com/kalo-build/kalo-plugins/plugins/openapi-sdk-go/pkg/generate"
	"github.com/stretchr/testify/suite"
)

type CompileTestSuite struct {
	assertfile.FileSuite

	TestDataPath    string
	GroundTruthPath string
}

func TestCompileTestSuite(t *testing.T) {
	suite.Run(t, new(CompileTestSuite))
}

func (suite *CompileTestSuite) SetupTest() {
	_, filename, _, _ := runtime.Caller(0)
	suite.TestDataPath = filepath.Join(filepath.Dir(filename), "..", "..", "testdata")
	suite.GroundTruthPath = filepath.Join(suite.TestDataPath, "ground-truth")
}

func (suite *CompileTestSuite) TestGenerateSDK() {
	inputDir := filepath.Join(suite.TestDataPath, "input")
	outputDir := suite.T().TempDir()

	cfg := generate.Config{
		InputDir:       inputDir,
		OutputDir:      outputDir,
		SpecFileName:   "openapi.yaml",
		PackageName:    "client",
		GenerateTypes:  true,
		GenerateClient: true,
	}

	err := generate.GenerateSDK(cfg)
	suite.NoError(err)

	typesPath := filepath.Join(outputDir, "types.go")
	gtTypesPath := filepath.Join(suite.GroundTruthPath, "types.go")
	suite.FileExists(typesPath)
	suite.FileEquals(typesPath, gtTypesPath)

	clientPath := filepath.Join(outputDir, "client.go")
	gtClientPath := filepath.Join(suite.GroundTruthPath, "client.go")
	suite.FileExists(clientPath)
	suite.FileEquals(clientPath, gtClientPath)
}

func (suite *CompileTestSuite) TestGenerateSDK_TypesOnly() {
	inputDir := filepath.Join(suite.TestDataPath, "input")
	outputDir := suite.T().TempDir()

	cfg := generate.Config{
		InputDir:       inputDir,
		OutputDir:      outputDir,
		SpecFileName:   "openapi.yaml",
		PackageName:    "client",
		GenerateTypes:  true,
		GenerateClient: false,
	}

	err := generate.GenerateSDK(cfg)
	suite.NoError(err)

	typesPath := filepath.Join(outputDir, "types.go")
	gtTypesPath := filepath.Join(suite.GroundTruthPath, "types.go")
	suite.FileExists(typesPath)
	suite.FileEquals(typesPath, gtTypesPath)

	clientPath := filepath.Join(outputDir, "client.go")
	_, err = os.Stat(clientPath)
	suite.True(os.IsNotExist(err), "client.go should not exist when GenerateClient is false")
}
