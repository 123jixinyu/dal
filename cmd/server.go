package cmd

import (
	"dal/process"
	"dal/process/source"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var ServerHandler = &server{
	DefaultPort: 8080,
	DefaultPath: "./samples",
}

type server struct {
	DefaultPath string
	DefaultPort int
}

// 启动服务
func (s *server) Server(ctx *cli.Context) error {

	path := ctx.String("path")
	if path == "" {
		path = s.DefaultPath
	}

	port := ctx.Int("port")
	if port == 0 {
		port = s.DefaultPort
	}

	dalFiles, err := s.findDalFiles(path)
	if err != nil {
		return err
	}
	if len(dalFiles) == 0 {
		return errors.New("no matching files")
	}

	e := echo.New()
	e.Use(middleware.CORS())

	for _, dalFile := range dalFiles {

		cnf, err := os.ReadFile(dalFile)
		if err != nil {
			return err
		}
		var content = string(cnf)

		kindVersion, err := process.GetKindVersion(content)
		if err != nil {
			return err
		}
		if strings.HasPrefix(kindVersion, "source") {
			if err := source.Syntax(content); err != nil {
				return err
			}
			continue
		}

		if err := process.Syntax(content); err != nil {
			return err
		}

		if err := s.addApi(e, content); err != nil {
			return err
		}

	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))

	return nil
}

// 添加API
func (*server) addApi(e *echo.Echo, content string) error {

	da := process.NewDal(content)

	e.Add(da.Method(), da.Route(), da.Handler(), da.MiddlewareFunc()...)

	if da.HasPreflight() {
		e.Add(http.MethodOptions, da.Route(), da.GetPreflightHandler())
	}
	return nil
}

// 查找dal文件
func (*server) findDalFiles(dir string) ([]string, error) {

	files := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			lastPoint := strings.LastIndex(info.Name(), ".")
			if lastPoint == -1 || lastPoint == len(info.Name())-1 {
				return nil
			}
			extension := info.Name()[lastPoint+1:]
			if extension == "dal" {
				files = append(files, path)
			}
		}
		return nil
	})

	return files, err
}
