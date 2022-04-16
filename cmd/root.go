package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wuennan/certupload/conf"
	"log"
	"strings"
	"time"
)

var (
	Cloud 	string
	CertName	string
	CertificatePath	string
	PrivateKeyPath	string
	ConfigPath	string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "certupload",
	Short: "Upload Cert Program.",
	Run: func(cmd *cobra.Command, args []string) {
		// 加载解析配置文件
		err := conf.LoadConfigFromToml("etc/config.toml")
		if err != nil {
			log.Println(err)
		}

		c := conf.C()
		// 校验必须项是否为空
		err = c.Validate()
		if err != nil{
			log.Println(err)
		}
		// 上传证书
		CertName=GetCertName()
		log.Println(CertName)
		err = c.Upload(CertName, CertificatePath, PrivateKeyPath)
		if err != nil{
			log.Println(err)
		}
	},

}

func GetCertName() string {
	res := strings.Split(CertificatePath, "_")
	return res[0]+time.Now().Format("-200601021504")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&CertificatePath,"certificate","c","","Certificate Path")
	rootCmd.PersistentFlags().StringVarP(&PrivateKeyPath,"keypath","k","","PrivateKey Path")
	//rootCmd.PersistentFlags().StringVarP(&CertName,"name","n","","Cert Name")
	rootCmd.PersistentFlags().StringVarP(&ConfigPath,"config","f","etc/config.toml","Config Path")
}


