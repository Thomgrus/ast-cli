// +build integration

package integration

import (
	"fmt"
	"github.com/checkmarxDev/ast-cli/internal/commands"
	"github.com/checkmarxDev/ast-cli/internal/wrappers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	astSchema    = "AST_SCHEMA"
	astHost      = "AST_HOST"
	astPort      = "80"
	scansPath    = "SCANS_PATH"
	projectsPath = "PROJECTS_PATH"
	resultsPath  = "RESULTS_PATH"
	uploadsPath  = "UPLOADS_PATH"
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandomizeString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[seededRand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestMain(m *testing.M) {
	log.Println("CLI integration tests started")
	// Run all tests
	exitVal := m.Run()
	log.Println("CLI integration tests done")
	os.Exit(exitVal)
}

func createASTIntegrationTestCommand() *cobra.Command {
	log.Println("Reading env variables")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	_ = viper.ReadInConfig()

	viper.SetDefault(astSchema, "http")
	viper.SetDefault(astHost, "localhost")
	viper.SetDefault(astPort, "80")
	viper.SetDefault(scansPath, "scans")
	viper.SetDefault(projectsPath, "projects")
	viper.SetDefault(uploadsPath, "uploads")
	viper.SetDefault(resultsPath, "results")

	schema := viper.GetString(astSchema)
	host := viper.GetString(astHost)
	port := viper.GetString(astPort)
	ast := fmt.Sprintf("%s://%s:%s/api", schema, host, port)

	scans := viper.GetString(scansPath)
	uploads := viper.GetString(uploadsPath)
	projects := viper.GetString(projectsPath)
	results := viper.GetString(resultsPath)

	scansURL := fmt.Sprintf("%s/%s", ast, scans)
	uploadsURL := fmt.Sprintf("%s/%s", ast, uploads)
	projectsURL := fmt.Sprintf("%s/%s", ast, projects)
	resultsURL := fmt.Sprintf("%s/%s", ast, results)

	scansWrapper := wrappers.NewHTTPScansWrapper(scansURL)
	uploadsWrapper := wrappers.NewUploadsHTTPWrapper(uploadsURL)
	projectsWrapper := wrappers.NewHTTPProjectsWrapper(projectsURL)
	resultsWrapper := wrappers.NewHTTPResultsWrapper(resultsURL)

	return commands.NewAstCLI(scansWrapper, uploadsWrapper, projectsWrapper, resultsWrapper)
}

func execute(cmd *cobra.Command, args ...string) error {
	cmd.SetArgs(args)
	return cmd.Execute()
}