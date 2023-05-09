package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Init() {

}

type Generator struct {
	cmd         string
	parameters  map[string]string
	storagePath string
}

type Application interface {
	AddParameter(name string, value string) *Generator
	Exec(filename string) (string, error)
}

func New() *Generator {
	//Software responsavel por manipular o arquivo
	g := Generator{
		parameters: map[string]string{},
	}

	g.cmd = os.Getenv("FILE_WORKER")
	g.storagePath = os.Getenv("FILE_STORAGE")

	g.AddParameter("c", "copy").AddParameter("f", "segment").
		AddParameter("segment_time", os.Getenv("FILE_SEGMENT_TIME")).AddParameter("reset_timestamps", "1").AddParameter("segment_format", "mpegts")

	return &g
}

func (g *Generator) AddParameter(name string, value string) *Generator {

	nameFormat := "-" + name

	g.parameters[nameFormat] = value

	return g
}

func (g *Generator) Exec(filePath string, exportPath string, fileID string) (string, error) {
	args := []string{}

	// Add o parametro indicando qual o path a ser lido
	args = append(args, "-i")
	args = append(args, filePath)

	for key, value := range g.parameters {
		args = append(args, key)
		args = append(args, value)
	}

	//Add output
	filFinalName := fmt.Sprintf("%s_%%03d.ts", fileID)

	args = append(args, filepath.Join(exportPath, filFinalName))

	command := exec.Command(g.cmd, args...)

	if err := command.Start(); err != nil {
		return err.Error(), err

	}

	go func() {

		if err := command.Wait(); err != nil {
			println("Erro no comando: " + command.String())
		}
	}()

	println(command.String() + " EM ANDAMENTO")
	return command.String(), nil
}
