package split

import (
	"fmt"
	"os"
	"sort"
	"log"
	"io"
	"path/filepath"
)

// byDatTime implements sort.Interface.
type byDateTime []os.FileInfo

func (f byDateTime) Len() int           { return len(f) }
func (f byDateTime) Less(i, j int) bool { return f[i].ModTime().Before(f[j].ModTime()) }
func (f byDateTime) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func SplitFolder(options Options, folderLocation string) ([]os.FileInfo, error) {

	if folderLocation == "" {
		fmt.Println("Nothing to do. Folder not provided.")
	}

	file, err := os.Open(folderLocation)
	if err != nil {
		log.Printf("Could not open folder %s", err.Error())
		return []os.FileInfo{}, err
	}
	fileInfos, err := file.Readdir(10000)
	if err != nil {
		return []os.FileInfo{}, err
	}
	sort.Sort(byDateTime(fileInfos))

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		toFolderName := filepath.Join(options.Folder, fileInfo.ModTime().Format("2006-01-02"))
		_, err := findOrCreateFolder(toFolderName, options)
		if err != nil {
			log.Printf("Not create folder %s for error %s", toFolderName, err.Error())
			return []os.FileInfo{}, nil
		}

		inputFileName := filepath.Join(folderLocation, fileInfo.Name())
		err = copyFile(fileInfo, inputFileName, toFolderName, options)

		//err = os.Chtimes(outputFileName, fileInfo.ModTime(), fileInfo.ModTime())

		if err != nil {
			log.Printf("Error writing file %s for error: %s", inputFileName, err)
			return []os.FileInfo{}, err
		}
		outputFileName := filepath.Join(toFolderName, fileInfo.Name());
		err = os.Chtimes(outputFileName, fileInfo.ModTime(), fileInfo.ModTime())
		if err != nil {
			log.Printf("Error chaning file %s for error: %s", outputFileName, err)
			return []os.FileInfo{}, err
		}

	}
	return []os.FileInfo{}, nil


}

func copyFile(fileInfo os.FileInfo, inputFileName string, toFolderName string, options Options) error {
	r, err := os.Open(inputFileName)
	if (err != nil) {
		log.Printf("Could not open file %s. Error occured: %s", inputFileName, err)
		return err
	}
	defer r.Close()
	if options.Verbose {
		log.Printf("writing to folder %s", toFolderName)
	}
	outputFileName := filepath.Join(toFolderName, fileInfo.Name())

	w, err := os.Create(outputFileName)
	defer w.Close()
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.Copy(w, r)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if options.Verbose {
		log.Printf("bytes written %d", b)
	}
	err = w.Sync()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}


func findOrCreateFolder(folderName string, options Options) (os.FileInfo, error) {
	fileInfo, err := os.Stat(folderName)

	if err == nil {
		if options.Verbose {
			log.Printf("Folder %s found", folderName)
		}

		return fileInfo, nil
	}
	// if there is an error, we assume folder does not exists

	if options.Verbose {
		log.Printf("Folder %s needs to be created", folderName)
	}
	merr := os.Mkdir(folderName, 0777)
	if merr != nil {
		return fileInfo, merr
	}
	log.Printf("Folder %s created with permission 777", folderName)
	return fileInfo, nil



}