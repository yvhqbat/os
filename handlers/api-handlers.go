package handlers

import (
	"net/http"
)

type APIHandlers struct {
}

// ListBucketsHandler - GET Service.
// -----------
// This implementation of the GET operation returns a list of all buckets
// owned by the authenticated sender of the request.
func (api APIHandlers) ListBucketsHandler(w http.ResponseWriter, r *http.Request) {

}

func (api APIHandlers) HeadObjectHandler(w http.ResponseWriter, r *http.Request) {

}

func (api APIHandlers) GetObjectHandler(w http.ResponseWriter, r *http.Request) {

}

func (api APIHandlers) PutObjectHandler(w http.ResponseWriter, r *http.Request) {

}

func (api APIHandlers) DeleteObjectHandler(w http.ResponseWriter, r *http.Request) {

}

func (api APIHandlers) PutBucketHandler(w http.ResponseWriter, r *http.Request) {

}

func (api APIHandlers) HeadBucketHandler(w http.ResponseWriter, r *http.Request) {

}

func (api APIHandlers) DeleteBucketHandler(w http.ResponseWriter, r *http.Request) {

}
