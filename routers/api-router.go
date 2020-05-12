package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"osapp/handlers"
)

// SlashSeparator - slash separator.
const SlashSeparator = "/"

var globalDomainNames []string // Root domains for virtual host style requests

// registerAPIRouter - registers S3 compatible APIs.
func RegisterAPIRouter(router *mux.Router) {
	// Initialize API.
	api := handlers.APIHandlers{}

	// API Router
	apiRouter := router.PathPrefix(SlashSeparator).Subrouter()
	var routers []*mux.Router
	for _, domainName := range globalDomainNames {
		routers = append(routers, apiRouter.Host("{bucket:.+}."+domainName).Subrouter())
		routers = append(routers, apiRouter.Host("{bucket:.+}."+domainName+":{port:.*}").Subrouter())
	}
	routers = append(routers, apiRouter.PathPrefix("/{bucket}").Subrouter())

	for _, bucket := range routers {
		// Object operations
		// HeadObject
		bucket.Methods(http.MethodHead).Path("/{object:.+}").HandlerFunc(api.HeadObjectHandler)

		/*
			// CopyObjectPart
			bucket.Methods(http.MethodPut).Path("/{object:.+}").HeadersRegexp(xhttp.AmzCopySource, ".*?(\\/|%2F).*?").HandlerFunc(httpTraceAll(api.CopyObjectPartHandler)).Queries("partNumber", "{partNumber:[0-9]+}", "uploadId", "{uploadId:.*}")
			// PutObjectPart
			bucket.Methods(http.MethodPut).Path("/{object:.+}").HandlerFunc(httpTraceHdrs(api.PutObjectPartHandler)).Queries("partNumber", "{partNumber:[0-9]+}", "uploadId", "{uploadId:.*}")
			// ListObjectPxarts
			bucket.Methods(http.MethodGet).Path("/{object:.+}").HandlerFunc(httpTraceAll(api.ListObjectPartsHandler)).Queries("uploadId", "{uploadId:.*}")
			// CompleteMultipartUpload
			bucket.Methods(http.MethodPost).Path("/{object:.+}").HandlerFunc(httpTraceAll(api.CompleteMultipartUploadHandler)).Queries("uploadId", "{uploadId:.*}")
			// NewMultipartUpload
			bucket.Methods(http.MethodPost).Path("/{object:.+}").HandlerFunc(httpTraceAll(api.NewMultipartUploadHandler)).Queries("uploads", "")
			// AbortMultipartUpload
			bucket.Methods(http.MethodDelete).Path("/{object:.+}").HandlerFunc(httpTraceAll(api.AbortMultipartUploadHandler)).Queries("uploadId", "{uploadId:.*}")
			// GetObjectACL - this is a dummy call.
			bucket.Methods(http.MethodGet).Path("/{object:.+}").HandlerFunc(httpTraceHdrs(api.GetObjectACLHandler)).Queries("acl", "")
			// GetObjectTagging - this is a dummy call.
			bucket.Methods(http.MethodGet).Path("/{object:.+}").HandlerFunc(httpTraceHdrs(api.GetObjectTaggingHandler)).Queries("tagging", "")
			// SelectObjectContent
			bucket.Methods(http.MethodPost).Path("/{object:.+}").HandlerFunc(httpTraceHdrs(api.SelectObjectContentHandler)).Queries("select", "").Queries("select-type", "2")
		*/
		// GetObject
		bucket.Methods(http.MethodGet).Path("/{object:.+}").HandlerFunc(handlers.HttpTraceAll(api.GetObjectHandler))
		// CopyObject
		// bucket.Methods(http.MethodPut).Path("/{object:.+}").HeadersRegexp(xhttp.AmzCopySource, ".*?(\\/|%2F).*?").HandlerFunc(httpTraceAll(api.CopyObjectHandler))
		// PutObject
		bucket.Methods(http.MethodPut).Path("/{object:.+}").HandlerFunc(handlers.HttpTraceAll(api.PutObjectHandler))
		// DeleteObject
		bucket.Methods(http.MethodDelete).Path("/{object:.+}").HandlerFunc(handlers.HttpTraceAll(api.DeleteObjectHandler))

		/// Bucket operations
		/*
			// GetBucketLocation
			bucket.Methods(http.MethodGet).HandlerFunc(handlers.HttpTraceAll(api.GetBucketLocationHandler)).Queries("location", "")
			// GetBucketPolicy
			bucket.Methods("GET").HandlerFunc(handlers.HttpTraceAll(api.GetBucketPolicyHandler)).Queries("policy", "")
			// GetBucketLifecycle
			bucket.Methods("GET").HandlerFunc(handlers.HttpTraceAll(api.GetBucketLifecycleHandler)).Queries("lifecycle", "")


			// Dummy Bucket Calls
			// GetBucketACL -- this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketACLHandler)).Queries("acl", "")
			// GetBucketCors - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketCorsHandler)).Queries("cors", "")
			// GetBucketWebsiteHandler - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketWebsiteHandler)).Queries("website", "")
			// GetBucketVersioningHandler - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketVersioningHandler)).Queries("versioning", "")
			// GetBucketAccelerateHandler - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketAccelerateHandler)).Queries("accelerate", "")
			// GetBucketRequestPaymentHandler - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketRequestPaymentHandler)).Queries("requestPayment", "")
			// GetBucketLoggingHandler - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketLoggingHandler)).Queries("logging", "")
			// GetBucketLifecycleHandler - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketLifecycleHandler)).Queries("lifecycle", "")
			// GetBucketReplicationHandler - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketReplicationHandler)).Queries("replication", "")
			// GetBucketTaggingHandler - this is a dummy call.
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketTaggingHandler)).Queries("tagging", "")
			//DeleteBucketWebsiteHandler
			bucket.Methods(http.MethodDelete).HandlerFunc(httpTraceAll(api.DeleteBucketWebsiteHandler)).Queries("website", "")
			// DeleteBucketTaggingHandler
			bucket.Methods(http.MethodDelete).HandlerFunc(httpTraceAll(api.DeleteBucketTaggingHandler)).Queries("tagging", "")

			// GetBucketNotification
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.GetBucketNotificationHandler)).Queries("notification", "")
			// ListenBucketNotification
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.ListenBucketNotificationHandler)).Queries("events", "{events:.*}")
			// ListMultipartUploads
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.ListMultipartUploadsHandler)).Queries("uploads", "")
			// ListObjectsV2
			bucket.Methods(http.MethodGet).HandlerFunc(httpTraceAll(api.ListObjectsV2Handler)).Queries("list-type", "2")
			// ListObjectsV1 (Legacy)
			bucket.Methods("GET").HandlerFunc(httpTraceAll(api.ListObjectsV1Handler))
			// PutBucketLifecycle
			bucket.Methods("PUT").HandlerFunc(httpTraceAll(api.PutBucketLifecycleHandler)).Queries("lifecycle", "")
			// PutBucketPolicy
			bucket.Methods("PUT").HandlerFunc(httpTraceAll(api.PutBucketPolicyHandler)).Queries("policy", "")

			// PutBucketNotification
			bucket.Methods(http.MethodPut).HandlerFunc(httpTraceAll(api.PutBucketNotificationHandler)).Queries("notification", "")
		*/

		// PutBucket
		bucket.Methods(http.MethodPut).HandlerFunc(handlers.HttpTraceAll(api.PutBucketHandler))
		// HeadBucket
		bucket.Methods(http.MethodHead).HandlerFunc(handlers.HttpTraceAll(api.HeadBucketHandler))
		/*
			// PostPolicy
			bucket.Methods(http.MethodPost).HeadersRegexp(xhttp.ContentType, "multipart/form-data*").HandlerFunc(httpTraceHdrs(api.PostPolicyBucketHandler))
			// DeleteMultipleObjects
			bucket.Methods(http.MethodPost).HandlerFunc(httpTraceAll(api.DeleteMultipleObjectsHandler)).Queries("delete", "")
			// DeleteBucketPolicy
			bucket.Methods("DELETE").HandlerFunc(httpTraceAll(api.DeleteBucketPolicyHandler)).Queries("policy", "")
			// DeleteBucketLifecycle
			bucket.Methods("DELETE").HandlerFunc(httpTraceAll(api.DeleteBucketLifecycleHandler)).Queries("lifecycle", "")
		*/
		// DeleteBucket
		bucket.Methods(http.MethodDelete).HandlerFunc(handlers.HttpTraceAll(api.DeleteBucketHandler))
	}

	/// Root operation

	// ListBuckets
	apiRouter.Methods(http.MethodGet).Path(SlashSeparator).HandlerFunc(handlers.HttpTraceAll(api.ListBucketsHandler))

	// If none of the routes match.
	apiRouter.NotFoundHandler = http.HandlerFunc(handlers.HttpTraceAll(handlers.NotFoundHandler))
}
