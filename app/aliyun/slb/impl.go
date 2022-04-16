package slb

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	slb "github.com/alibabacloud-go/slb-20140515/v3/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
)

var (
	validate = validator.New()
)

func NewConfig() *Config {
	return &Config{}
}

type Config struct {
	*AliSlb `toml:"aliSlb"`
}

type AliSlb struct {
	Ak       string `validate:"required" toml:"AccessKey"`
	Sk       string `validate:"required" toml:"AccessKeySecret"`
	RegionId string `validate:"required" toml:"RegionId"`
	Endpoint string `validate:"required" toml:"Endpoint""`
}

func (a *AliSlb) Validate() error {
	return validate.Struct(a)
}

func (a *AliSlb) Upload(CertName, CertificatePath, PrivateKeyPath string) (err error) {
	// 获取证书文本内容
	publicKeyPath, err := GetCert(CertificatePath)
	if err != nil {
		return err
	}
	privateKeyPath, err := GetCert(PrivateKeyPath)
	if err != nil {
		return err
	}
	// 构建请求参数
	uploadServerCertificateRequest := &slb.UploadServerCertificateRequest{
		RegionId:                    tea.String(a.RegionId),
		AliCloudCertificateRegionId: tea.String(a.RegionId),
		ServerCertificate:           tea.String(publicKeyPath),
		PrivateKey:                  tea.String(privateKeyPath),
		ServerCertificateName:       tea.String(CertName),
	}

	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &a.Ak,
		// 您的AccessKey Secret
		AccessKeySecret: &a.Sk,
	}
	// 访问的域名
	config.Endpoint = tea.String(a.Endpoint)
	//slb.Client{}
	client, err := slb.NewClient(config)
	if err != nil {
		return err
	}

	resp, err := client.UploadServerCertificate(uploadServerCertificateRequest)
	if err != nil {
		return err
	}
	console.Log(util.ToJSONString(tea.ToMap(resp)))

	return nil
}

// 读取证书，并返回 string
func GetCert(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
