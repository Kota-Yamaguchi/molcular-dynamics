package domain

import (
	"fmt"
	"os"
	"strconv"
)

func exportCdview(atoms Atoms, count int) {
	filename := "./atoms-view/atoms.xyz"
	if err := writeByres(filename, strconv.Itoa(len(atoms.atoms))+"\n"); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	if err := writeByres(filename, strconv.FormatFloat(float64(count+1), 'f', 2, 64)+"ps\n"); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	for i, a := range atoms.atoms {
		if err := writeByres(filename, a.toString(i)); err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
	}

}
func writeByres(filename string, contents string) error {
	_, err := os.Stat(filename)
	if err != nil {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err = file.WriteString(contents)
	if err != nil {
		return err
	}
	return nil
}
