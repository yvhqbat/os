package cmd

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)


var listBucketsCmd = &cobra.Command{
	Use:   "buckets",
	Short: "list buckets",
	Long:  `list buckets owned by user`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("list buckets cmd")

		c, err := getOssClient(cfg)
		if err != nil {
			return
		}

		res, err := c.ListBuckets()
		if err != nil {
			log.Errorf("list buckets failed, err: %v", err)
			return
		}
		for pos, bucket := range res.Buckets {
			log.Infof("pos:%d, bucket:%v\n", pos, bucket)
		}
	},
}

var listObjectsCmd = &cobra.Command{
	Use:   "objects",
	Short: "list objects",
	Long:  `list objects owned by user`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("list objects cmd")

		c, err := getOssClient(cfg)
		if err != nil {
			return
		}

		bucket, err := c.Bucket(cfg.bucket)
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		objects, err := bucket.ListObjects()
		if err != nil {
			log.Errorf("%v", err)
			return
		}
		for pos, object := range objects.Objects {
			log.Infof("pos:%d, object:%v", pos, object)
		}
	},
}

var putObjectCmd = &cobra.Command{
	Use:   "put",
	Short: "put object",
	Long:  `put an object from file`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("put object cmd")

		c, err := getOssClient(cfg)
		if err != nil {
			return
		}

		bucket, err := c.Bucket(cfg.bucket)
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		err = bucket.PutObjectFromFile(cfg.object, cfg.file)
		if err != nil {
			log.Errorf("put object failed, err: %v", err)
			return
		}
		log.Infof("put object success")
	},
}

var uploadObjectCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload object",
	Long:  `multipart upload an object from file`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("upload object cmd")

		c, err := getOssClient(cfg)
		if err != nil {
			return
		}

		bucket, err := c.Bucket(cfg.bucket)
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		err = bucket.UploadFile(cfg.object, cfg.file, 1*1024*124, oss.Routines(3), oss.Checkpoint(true, ""))
		if err != nil {
			log.Errorf("put object failed, err: %v", err)
			return
		}
		log.Infof("put object success")
	},
}

var downloadObjectCmd = &cobra.Command{
	Use:   "download",
	Short: "download object",
	Long:  `download an object to local file`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("downlaod object cmd")

		c, err := getOssClient(cfg)
		if err != nil {
			return
		}

		bucket, err := c.Bucket(cfg.bucket)
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		err = bucket.DownloadFile(cfg.object, cfg.file, 1*1024*124, oss.Routines(3), oss.Checkpoint(true, ""))
		if err != nil {
			log.Errorf("download object failed, err: %v", err)
			return
		}
		log.Infof("download object success")
	},
}

var getObjectCmd = &cobra.Command{
	Use:   "get",
	Short: "get object",
	Long:  `get an object to local file`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("get object cmd")

		c, err := getOssClient(cfg)
		if err != nil {
			return
		}

		bucket, err := c.Bucket(cfg.bucket)
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		err = bucket.GetObjectToFile(cfg.object, cfg.file)
		if err != nil {
			log.Errorf("get object failed, err: %v", err)
			return
		}
		log.Infof("get object success")
	},
}

var getObjectMetadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "get object metadata",
	Long:  `get object metadata`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("get object metadata cmd")

		c, err := getOssClient(cfg)
		if err != nil {
			return
		}

		bucket, err := c.Bucket(cfg.bucket)
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		res, err := bucket.GetObjectMeta(cfg.object)
		if err != nil {
			log.Errorf("get object metadata failed, err: %v", err)
			return
		}
		log.Infof("%v", res)
	},
}

var deleteObjectCmd = &cobra.Command{
	Use:   "del",
	Short: "delete object",
	Long:  `delete an object`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("delete object cmd")

		c, err := getOssClient(cfg)
		if err != nil {
			return
		}

		bucket, err := c.Bucket(cfg.bucket)
		if err != nil {
			log.Errorf("%v", err)
			return
		}

		err = bucket.DeleteObject(cfg.object)
		if err != nil {
			log.Errorf("delete object failed, err: %v", err)
			return
		}
		log.Infof("delete object success")
	},
}

type config struct {
	endpoint string
	ak       string
	sk       string
	bucket   string
	object   string
	file     string
}

var cfg config
var ossClient *oss.Client

func getOssClient(cfg config) (*oss.Client, error) {
	if ossClient != nil {
		return ossClient, nil
	}

	// 创建OSSClient实例。
	//tr := &http.Transport{
	//    TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	//}
	//c := &http.Client{Transport: tr}

	// client, err := oss.New(cfg.endpoint, cfg.ak, cfg.sk, oss.UseCname(true))
	client, err := oss.New(cfg.endpoint, cfg.ak, cfg.sk)
	if err != nil {
		log.Errorf("create oss client failed, err: %v", err)
		return nil, err
	}
	return client, nil
}

func RegisterOSSCmd(c *cobra.Command) {
	c.PersistentFlags().StringVar(&cfg.endpoint, "endpoint", "https://oss.hcs.com", "please input endpoint")
	c.PersistentFlags().StringVar(&cfg.sk, "sk", "14u3X8S7WZ1Z1h5VC5R49rq039183tT6Rw3yWq6A1b98337851S352lW9lX28q9", "please input sk")
	c.PersistentFlags().StringVar(&cfg.bucket, "bucket", "defaultBucket", "please input bucket name")
	c.PersistentFlags().StringVar(&cfg.object, "object", "defaultObject", "please input object key")

	putObjectCmd.Flags().StringVar(&cfg.file, "file", "test.jpg", "please input file path")
	getObjectCmd.Flags().StringVar(&cfg.file, "file", "test.jpg", "please input file path")
	uploadObjectCmd.Flags().StringVar(&cfg.file, "file", "test.jpg", "please input file path")
	downloadObjectCmd.Flags().StringVar(&cfg.file, "file", "test.jpg", "please input file path")

	c.AddCommand(listBucketsCmd)
	c.AddCommand(listObjectsCmd)
	c.AddCommand(putObjectCmd)
	c.AddCommand(getObjectCmd)
	c.AddCommand(getObjectMetadataCmd)
	c.AddCommand(deleteObjectCmd)
	c.AddCommand(uploadObjectCmd)
	c.AddCommand(downloadObjectCmd)
}

