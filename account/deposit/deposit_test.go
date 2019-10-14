package deposit_test

import (
	"github.com/iotaledger/iota.go/account/deposit"
	"github.com/iotaledger/iota.go/checksum"
	"github.com/iotaledger/iota.go/consts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"strings"
	"time"
)

var _ = Describe("Deposit", func() {

	addr := strings.Repeat("A", 81)
	addrWithChecksum, err := checksum.AddChecksum(addr, true, consts.AddressChecksumTrytesSize)
	if err != nil {
		panic(err)
	}
	timeoutAt := time.Date(2019, time.March, 17, 18, 34, 0, 0, time.UTC)

	actualMagnetLink := "iota://AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAWOWRMBLMD/?timeout_at=1552847640&multi_use=0&expected_amount=1000&message="

	var expAmount uint64 = 1000

	cda := &deposit.CDA{
		Address: addrWithChecksum,
		Conditions: deposit.Conditions{
			MultiUse:       false,
			ExpectedAmount: &expAmount,
			TimeoutAt:      &timeoutAt,
		},
	}

	msg := "そほは - sensei"
	cdaWithMessage := &deposit.CDA{
		Address: addrWithChecksum,
		Conditions: deposit.Conditions{
			MultiUse:       false,
			ExpectedAmount: &expAmount,
			TimeoutAt:      &timeoutAt,
			Message:        &msg,
		},
	}
	magnetLinkWithMessage := "iota://AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANVXKPIYGU/?timeout_at=1552847640&multi_use=0&expected_amount=1000&message=%E3%81%9D%E3%81%BB%E3%81%AF+-+sensei"

	Context("Creating a magnet-link from a CDA", func() {
		invalidDueToDefinedExpectedAndMultiCDA := &deposit.CDA{
			Address: addrWithChecksum,
			Conditions: deposit.Conditions{
				MultiUse:       true,
				ExpectedAmount: &expAmount,
				TimeoutAt:      &timeoutAt,
			},
		}

		It("returns an error when both multi use and expected amount are set", func() {
			_, err := invalidDueToDefinedExpectedAndMultiCDA.AsMagnetLink()
			Expect(errors.Cause(err)).To(Equal(deposit.ErrInvalidDepositAddressOptions))
		})

		It("works", func() {
			magnetLink, err := cda.AsMagnetLink()
			Expect(err).ToNot(HaveOccurred())
			Expect(magnetLink).To(Equal(actualMagnetLink))
		})

		It("works also with non ASCII message", func() {
			_, err := cdaWithMessage.AsMagnetLink()
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("A valid CDA in magnet-link form", func() {

		It("parses", func() {
			condsFromMangetLink, err := deposit.ParseMagnetLink(actualMagnetLink)
			Expect(err).ToNot(HaveOccurred())
			Expect(condsFromMangetLink.Address).To(Equal(cda.Address))
			Expect(*condsFromMangetLink.ExpectedAmount).To(Equal(*cda.ExpectedAmount))
			Expect(*condsFromMangetLink.TimeoutAt).To(Equal(*cda.TimeoutAt))
		})

		It("parses also with non ASCII message", func() {
			condsFromMangetLink, err := deposit.ParseMagnetLink(magnetLinkWithMessage)
			Expect(err).ToNot(HaveOccurred())
			Expect(condsFromMangetLink.Address).To(Equal(cdaWithMessage.Address))
			Expect(*condsFromMangetLink.ExpectedAmount).To(Equal(*cdaWithMessage.ExpectedAmount))
			Expect(*condsFromMangetLink.TimeoutAt).To(Equal(*cdaWithMessage.TimeoutAt))
			Expect(condsFromMangetLink.Message).To(Equal(cdaWithMessage.Message))
		})

	})

})
