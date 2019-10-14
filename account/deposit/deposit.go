package deposit

import (
	"fmt"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/checksum"
	"github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/curl"
	. "github.com/iotaledger/iota.go/trinary"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
	"time"
)

// ErrAddressInvalid is returned when an address is invalid when parsed from a serialized form.
var ErrAddressInvalid = errors.New("invalid address")
var ErrMagnetLinkChecksumInvalid = errors.New("magnet-link checksum is invalid")
var ErrInvalidDepositAddressOptions = errors.New("invalid address options are invalid")

// CDA defines a conditional deposit address.
type CDA struct {
	Conditions
	Address Hash `json:"address"`
}

// Defines the names of the condition fields in a magnet link.
const (
	MagnetLinkTimeoutField        = "timeout_at"
	MagnetLinkMultiUseField       = "multi_use"
	MagnetLinkExpectedAmountField = "expected_amount"
	MagnetLinkMessageField        = "message"
	MagnetLinkFormat              = "iota://%s%s/?%s=%d&%s=%d&%s=%d&%s=%s"
)

// AsMagnetLink converts the conditions into a magnet link URL.
func (cda *CDA) AsMagnetLink() (string, error) {
	var expectedAmount uint64
	if cda.ExpectedAmount != nil {
		expectedAmount = *cda.ExpectedAmount
	}
	checksum, err := cda.Checksum()
	if err != nil {
		return "", err
	}
	var multiUse int
	if cda.MultiUse {
		multiUse = 1
	}

	return fmt.Sprintf(MagnetLinkFormat,
		cda.Address[:consts.HashTrytesSize],
		checksum[consts.HashTrytesSize-consts.AddressChecksumTrytesSize:consts.HashTrytesSize],
		MagnetLinkTimeoutField, cda.TimeoutAt.Unix(),
		MagnetLinkMultiUseField, multiUse,
		MagnetLinkExpectedAmountField, expectedAmount,
		MagnetLinkMessageField, func() string {
			if cda.Message == nil {
				return ""
			}
			return url.QueryEscape(*cda.Message)
		}()), nil
}

// AsTransfer converts the conditional deposit address into a transfer object.
func (cda *CDA) AsTransfer() bundle.Transfer {
	return bundle.Transfer{
		Address: cda.Address,
		Value: func() uint64 {
			if cda.ExpectedAmount == nil {
				return 0
			}
			return *cda.ExpectedAmount
		}(),
		Message: func() Trytes {
			if cda.Message == nil {
				return ""
			}
			trytes, err := BytesToTrytes([]byte(*cda.Message))
			if err != nil {
				panic(err)
			}
			return trytes
		}(),
	}
}

// Checksum returns the checksum of the the CDA.
func (cda *CDA) Checksum() (Trytes, error) {
	// checksum formula:
	// Checksum = CurlHash(
	// 	CurlHash(address_trits)[:134] +
	// 	PadTrits27(timeout_value_trits) +
	// 	MultiUse(0/1) +
	// 	PadTrits81(amount_value_trits) +
	//  PadTo243Multiple(bytes_to_trits(message))
	// )
	if err := ValidateConditions(&cda.Conditions); err != nil {
		return "", err
	}
	addrTrits, err := TrytesToTrits(cda.Address[:consts.HashTrytesSize])
	if err != nil {
		return "", err
	}
	c := curl.NewCurlP81()
	if err := c.Absorb(addrTrits); err != nil {
		return "", err
	}
	addrChecksumTrits, err := c.Squeeze(consts.HashTrinarySize)
	if err != nil {
		return "", err
	}
	timeoutAtTrits := PadTrits(IntToTrits(cda.TimeoutAt.Unix()), 27)
	var expectedAmountTrits Trits
	if cda.ExpectedAmount != nil {
		expectedAmountTrits = PadTrits(IntToTrits(int64(*cda.ExpectedAmount)), 81)
	} else {
		expectedAmountTrits = PadTrits(expectedAmountTrits, 81)
	}
	var multiUse int8
	if cda.MultiUse {
		multiUse = 1
	}
	input := make(Trits, 0)
	input = append(input, addrChecksumTrits[:134]...)
	input = append(input, timeoutAtTrits...)
	input = append(input, multiUse)
	input = append(input, expectedAmountTrits...)

	if cda.Message != nil && len(*cda.Message) != 0 {
		msgTrits, err := BytesToTrits([]byte(*cda.Message))
		if err != nil {
			return "", err
		}

		// pad to multiple of 243
		msgTritLen := len(msgTrits)
		paddedMsgTrits := PadTrits(msgTrits, msgTritLen+(consts.HashTrinarySize-msgTritLen%consts.HashTrinarySize))
		input = append(input, paddedMsgTrits...)
	}
	c.Reset()
	if err := c.Absorb(input); err != nil {
		return "", err
	}
	checksumTrits, err := c.Squeeze(consts.HashTrinarySize)
	if err != nil {
		return "", err
	}
	return TritsToTrytes(checksumTrits)
}

// ParseMagnetLink parses the given magnet link URL.
func ParseMagnetLink(cdaMagnetLink string) (*CDA, error) {
	link, err := url.Parse(cdaMagnetLink)
	if err != nil {
		return nil, err
	}
	query := link.Query()
	cda := &CDA{}
	if len(link.Host) != consts.AddressWithChecksumTrytesSize {
		return nil, errors.Wrap(ErrAddressInvalid, "host/address part of magnet-link must be 90 trytes long")
	}
	addrWithChecksum, err := checksum.AddChecksum(link.Host[:consts.HashTrytesSize], true, consts.AddressChecksumTrytesSize)
	if err != nil {
		return nil, err
	}
	cda.Address = addrWithChecksum
	expiresSeconds, err := strconv.ParseInt(query.Get(MagnetLinkTimeoutField), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "invalid expire timestamp")
	}
	expire := time.Unix(expiresSeconds, 0).UTC()
	cda.TimeoutAt = &expire
	cda.MultiUse = query.Get(MagnetLinkMultiUseField) == "1"
	expectedAmount, err := strconv.ParseUint(query.Get(MagnetLinkExpectedAmountField), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "invalid expected amount")
	}
	cda.ExpectedAmount = &expectedAmount

	msg, err := url.QueryUnescape(query.Get(MagnetLinkMessageField))
	if err != nil {
		return nil, errors.Wrap(err, "invalid message")
	}
	if len(msg) != 0 {
		cda.Message = &msg
	}

	computedChecksum, err := cda.Checksum()
	if err != nil {
		return nil, err
	}
	magnetLinkChecksum := link.Host[consts.HashTrytesSize:]
	if computedChecksum[consts.HashTrytesSize-consts.AddressChecksumTrytesSize:consts.HashTrytesSize] != magnetLinkChecksum {
		return nil, ErrMagnetLinkChecksumInvalid
	}
	if err := ValidateConditions(&cda.Conditions); err != nil {
		return nil, err
	}
	return cda, nil
}

// Conditions define conditions for a new deposit address generated by an account.
type Conditions struct {
	// The time after this deposit address becomes invalid.
	TimeoutAt *time.Time `json:"timeout_at,omitempty" bson:"timeout_at,omitempty"`
	// Whether to expect multiple deposits to this address in the given timeout.
	// If this flag is false, the deposit address is considered
	// in the input selection as soon as one deposit is available.
	// ExpectedAmount and MultiUse are mutually exclusive: MultiUse must be false if an ExpectedAmount over 0 is set.
	MultiUse bool `json:"multi_use,omitempty" bson:"multi_use,omitempty"`
	// The expected amount which gets deposited.
	// If the timeout is hit, the address is automatically
	// considered in the input selection.
	// ExpectedAmount and MultiUse are mutually exclusive: MultiUse must be false if an ExpectedAmount over 0 is set.
	ExpectedAmount *uint64 `json:"expected_amount,omitempty" bson:"expected_amount,omitempty"`
	// An arbitrary message with an arbitrary size.
	Message *string `json:"message,omitempty" bson:"message,omitempty"`
}

// ValidateConditions validates the deposit conditions.
func ValidateConditions(conds *Conditions) error {
	if conds.ExpectedAmount != nil {
		if *conds.ExpectedAmount > 0 && conds.MultiUse {
			return errors.Wrap(ErrInvalidDepositAddressOptions, "expected amount and multi use are mutually exclusive")
		}
	}
	return nil
}
