package app

type Uploader interface {
	Upload(CertName, CertificatePath, PrivateKeyPath string) error
}
