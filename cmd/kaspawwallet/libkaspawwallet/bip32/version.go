package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// KaspawMainnetPrivate is the version that is used for
// kaspaw mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var KaspawMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// KaspawMainnetPublic is the version that is used for
// kaspaw mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var KaspawMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// KaspawTestnetPrivate is the version that is used for
// kaspaw testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var KaspawTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// KaspawTestnetPublic is the version that is used for
// kaspaw testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var KaspawTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// KaspawDevnetPrivate is the version that is used for
// kaspaw devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var KaspawDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// KaspawDevnetPublic is the version that is used for
// kaspaw devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var KaspawDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// KaspawSimnetPrivate is the version that is used for
// kaspaw simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var KaspawSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// KaspawSimnetPublic is the version that is used for
// kaspaw simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var KaspawSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case KaspawMainnetPrivate:
		return KaspawMainnetPublic, nil
	case KaspawTestnetPrivate:
		return KaspawTestnetPublic, nil
	case KaspawDevnetPrivate:
		return KaspawDevnetPublic, nil
	case KaspawSimnetPrivate:
		return KaspawSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case KaspawMainnetPrivate:
		return true
	case KaspawTestnetPrivate:
		return true
	case KaspawDevnetPrivate:
		return true
	case KaspawSimnetPrivate:
		return true
	}

	return false
}
