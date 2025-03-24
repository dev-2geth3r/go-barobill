package barobill

import "github.com/rushysloth/go-tsid"

const customEpoch int64 = 1735689600000 // 2025-01-01 00:00:00 +0000

var (
	tsidFactory = NewTsidFactory()
)

func NewTsidFactory() *tsid.TsidFactory {
	return Must(tsid.TsidFactoryBuilder().
		WithCustomEpoch(customEpoch).
		WithNodeBits(10).
		WithRandom(tsid.NewIntRandom(tsid.NewCryptoRandomSupplier())).
		Build())
}

func Must[T any](a T, _ error) T {
	return a
}

func NewTSID() string {
	return Must(tsidFactory.Generate()).ToString()
}
