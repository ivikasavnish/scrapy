package cloud

// Path: cloud/aws_spot.go

type AWS_SPOT struct {
	Price              float64
	InstanceType       string
	AvailabilityZone   string
	ProductDescription string
	SpotPrice          string
}

func NewAWS_SPOT() *AWS_SPOT {
	return &AWS_SPOT{}
}

func (a *AWS_SPOT) StartAndRun() {

}
func (a *AWS_SPOT) Stop() {

}

type GCP_SPOT struct {
	Price              float64
	InstanceType       string
	AvailabilityZone   string
	ProductDescription string
	SpotPrice          string
}

func New_GCPSPOT() *GCP_SPOT {
	return &GCP_SPOT{}
}

func (a *GCP_SPOT) StartAndRun() {

}
func (a *GCP_SPOT) Stop() {

}

type AZURE_SPOT struct {
	Price              float64
	InstanceType       string
	AvailabilityZone   string
	ProductDescription string
	SpotPrice          string
}

func New_AZURESPOT() *AZURE_SPOT {
	return &AZURE_SPOT{}
}

func (a *AZURE_SPOT) StartAndRun() {}

func (a *AZURE_SPOT) Stop() {}

type SpotManagerInterface interface {
	StartAndRun()
	Stop()
}
