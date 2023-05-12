package observer

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func Init(fileChannel chan string) {
	w, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println("Erro ao criar o watcher:", err.Error())
		panic(err)

	}

	watcher = w

	go observer(fileChannel)

	fmt.Println("Observer OK")

}

func Add(path string) {
	err := watcher.Add(path)

	if err != nil {
		fmt.Println("Erro ao adicionar diretório ao watcher:", err)
	}

}

func Remove(path string) {
	err := watcher.Remove(path)

	if err != nil {
		fmt.Println("Erro ao remover diretório ao watcher:", err)
	}
}

func observer(channel chan string) {

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Create == fsnotify.Write {
				//Envia a informacao para o channel do dispatcher
				channel <- event.Name
			}

		case event := <-watcher.Errors:
			fmt.Println("Error on ewatcher event", event)

		}
	}

}
