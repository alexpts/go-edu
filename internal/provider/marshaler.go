package provider

import (
	"github.com/bytedance/sonic"

	"github.com/alexpts/edu-go/pkg/convert"
)

func ProvideStdSonicJsonMarshaler() sonic.API {
	return sonic.ConfigStd
}

func ProvideDefaultSonicJsonMarshaler() sonic.API {
	return sonic.ConfigDefault
}

func ProvideFastestSonicJsonMarshaler() sonic.API {
	return sonic.ConfigFastest
}

func ProvideStdEncodingJsonMarshaler() *convert.StdMarshaler {
	return &convert.StdMarshaler{}
}
