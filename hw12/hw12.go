package hw12

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Key struct {
	key   string
	file  string
	value string
}

type MapKey map[string]Key

func (k *MapKey) getValEnv() {

	for i, v := range *k {
		result, err := loadValFromFile(v.file)
		if err != nil {
			log.Fatal("Ошибка при получения значения из файла", v.file, " ошибкa :", err)
			continue
		}
		v.value = result
		(*k)[i] = v
	}
}

func loadValFromFile(fileName string) (string, error) {

	var result strings.Builder
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data := make([]byte, 128)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		result.WriteString(string(data[:n]))
	}
	res := strings.TrimSpace(result.String())
	res = strings.Replace(res, "\n", ";", -1)
	res = strings.Replace(res, "\r", "", -1)
	return res, nil
}

func (k *MapKey) createMapEnvVar(pathDir string) error {

	err := filepath.Walk(pathDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			f := strings.Split(info.Name(), ".")
			if !info.IsDir() && f[len(f)-1] == "env" {
				key := info.Name()[0 : len(info.Name())-4]
				(*k)[key] = Key{key, path, ""}
			}
			return nil
		})
	if err != nil {
		return err
	}
	k.getValEnv()
	return nil
}

func ReadDir(dir string) (map[string]string, error) {

	keyEnv := MapKey{}
	res := make(map[string]string)
	err := keyEnv.createMapEnvVar(dir)
	if err != nil {
		return nil, err
	}
	for _, v := range keyEnv {
		res[v.key] = v.value
	}
	return res, nil
}

func getEnvString(env map[string]string) []string {
	s := make([]string, len(env))
	index := 0
	for i, v := range env {
		s[index] = fmt.Sprintf("%v=%v", i, v)
		index++
	}
	return s
}

func RunCmd(cmd []string, env map[string]string) int {

	s := getEnvString(env)
	c := exec.Command(cmd[0])
	c.Env = append(os.Environ(), s...)
	if len(cmd) > 1 {
		c.Args = cmd
	}
	c.Stdout = os.Stdout // <===
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
	return -1
}
