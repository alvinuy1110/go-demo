package main
import "fmt"
import "io/ioutil"
import "os"
func read_file(fileUri string)  ([]byte, error) {
	data, err := ioutil.ReadFile(fileUri)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}
	fmt.Println("Contents of file:", string(data))
	return data, nil
}
func write_file() error {
	f, err := os.Create("file1.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	l, err := f.WriteString("Write Line one")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	fmt.Println(l, "bytes written")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func main() {
	data, err := read_file("../data/data.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))

	// write file
	write_file()
}