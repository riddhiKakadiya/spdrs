package models

import (
    "io"
    "net/http"
    "os"
    "os/exec"
)

func DownloadData() {
    //delete already present file
    exec.Command("sh","-c","rm SPY_All_Holdings2.csv").Output() 
    // fileUrl := "https://us.spdrs.com/site-content/xls/SPY_All_Holdings.xls?fund=SPY&docname=All+Holdings&onynx_code1=&onyx_code2="
    fileUrl := "https://us.spdrs.com/site-content/xls/SPY_All_Holdings.xls"
    if err := DownloadFile("SPY_All_Holdings2.csv", fileUrl); err != nil {panic(err)}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}

