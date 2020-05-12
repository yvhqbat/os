package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "oss",
	Short: "this is a simple oss cli",
	Long: `A Fast and Flexible oss cli in Go.
                Complete documentation is available at liudong11`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Infof("root cmd")
	},
}

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

var getObjectCmd = &cobra.Command{
	Use:   "get",
	Short: "get object",
	Long:  `download an object to local file`,
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

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.endpoint, "endpoint", "https://oss.hcs.com", "please input endpoint")
	rootCmd.PersistentFlags().StringVar(&cfg.ak, "ak", "user01xg21q3d1U4y3F3IH1fo8OF9k6vXRKA4y6f35ASfV14r974F03q11klf93", "please input ak")
	rootCmd.PersistentFlags().StringVar(&cfg.sk, "sk", "14u3X8S7WZ1Z1h5VC5R49rq039183tT6Rw3yWq6A1b98337851S352lW9lX28q9", "please input sk")
	rootCmd.PersistentFlags().StringVar(&cfg.bucket, "bucket", "defaultBucket", "please input bucket name")
	rootCmd.PersistentFlags().StringVar(&cfg.object, "object", "defaultObject", "please input object key")

	putObjectCmd.Flags().StringVar(&cfg.file, "file", "test.jpg", "please input file path")
	getObjectCmd.Flags().StringVar(&cfg.file, "file", "test.jpg", "please input file path")

	rootCmd.AddCommand(listBucketsCmd)
	rootCmd.AddCommand(listObjectsCmd)
	rootCmd.AddCommand(putObjectCmd)
	rootCmd.AddCommand(getObjectCmd)
	rootCmd.AddCommand(getObjectMetadataCmd)
	rootCmd.AddCommand(deleteObjectCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
