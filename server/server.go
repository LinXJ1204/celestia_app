package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/cors"
)

func main() {
	/* submit := []byte("test")
	fmt.Println(len(submit))
	got := Submit(submit, "test") */
	//got := Get("test", 231452)
	//fmt.Println(got)
	/*
		 	err := saveBytesAsJPG(got, "test.jpg")
			if err != nil {
				fmt.Println(err)
			}
	*/

	httpserver()
}

func httpserver() {
	ctx := context.Background()
	handler := http.NewServeMux()
	handler.HandleFunc("/submit", a_HandleFunc_submit(ctx))
	handler.HandleFunc("/get", a_HandleFunc_get(ctx))
	// Replace "http://localhost:8080" with the actual URL of your frontend application
	allowedOrigins := []string{"http://localhost:3000", "http://34.222.228.215:3000"}

	// CORS configuration
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	// Use the CORS middleware with the handler
	handlerWithCORS := corsMiddleware.Handler(handler)

	addr := ":9453"
	fmt.Println("listen on:" + addr)
	http.ListenAndServe(addr, handlerWithCORS) //設定監聽的埠
}

func a_HandleFunc_submit(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req submit_req
		file, _, _ := r.FormFile("arr")
		namespace := r.FormValue("name")
		body, _ := io.ReadAll(file)

		req.arr = body
		req.name = namespace

		res := Submit(req.arr, req.name)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Println("res:")
		fmt.Println(res)
		jsonres, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(jsonres)
		w.Write(jsonres)
	}
}

func a_HandleFunc_get(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req get_req
		req.blockheight, _ = strconv.ParseUint(r.FormValue("height"), 10, 64)
		fmt.Println(req.blockheight)
		req.name = r.FormValue("name")

		res_blob := Get(req.name, req.blockheight)

		var res get_res
		res_base64 := base64.StdEncoding.EncodeToString(res_blob)
		res.Blob = res_base64
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		jsonres, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(jsonres)
	}
}

func Bytes2photo(arr []byte) {
	img, _, err := image.Decode(bytes.NewReader(arr))
	fmt.Println(img)
	if err != nil {
		fmt.Println(err)
	}

	out, _ := os.Create("./img.jpeg")
	defer out.Close()

	var opts jpeg.Options

	err = jpeg.Encode(out, img, &opts)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	}
}

func saveBytesAsJPG(byteArray []byte, filename string) error {
	// Convert the byte array to an image
	img, _, err := image.Decode(bytes.NewReader(byteArray))
	if err != nil {
		return err
	}

	// Create a new file to save the JPG image
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the image as JPG and save it to the file
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	fmt.Println("Image saved as", filename)
	return nil
}

type submit_req struct {
	arr  []byte
	name string
}

type get_req struct {
	blockheight uint64 `json:"blockheight"`
	name        string `json:"name"`
}
