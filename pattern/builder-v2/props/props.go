package props

var WithCustomerId = func(s string) string { return s }

var WithAdsId = func(s string) string { return s }

type SonyProps struct {
	customerId string
	adsId      string
}

func NewSonyProps(
	customerId string,
	adsId string,
) *SonyProps {
	return &SonyProps{
		customerId: customerId,
		adsId:      adsId,
	}
}

func (sony *SonyProps) GetCustomerId() string {
	return sony.customerId
}

func (sony *SonyProps) GetAdsId() string {
	return sony.adsId
}
