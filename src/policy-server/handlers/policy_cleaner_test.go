package handlers_test

import (
	"errors"
	"policy-server/fakes"
	"policy-server/handlers"
	"policy-server/models"

	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("PolicyCleaner", func() {
	var (
		policyCleaner *handlers.PolicyCleaner
		fakeStore     *fakes.Store
		fakeUAAClient *fakes.UAAClient
		fakeCCClient  *fakes.CCClient
		logger        *lagertest.TestLogger
		allPolicies   []models.Policy
	)

	BeforeEach(func() {
		allPolicies = []models.Policy{{
			Source: models.Source{ID: "live-guid", Tag: "tag"},
			Destination: models.Destination{
				ID:       "live-guid",
				Tag:      "tag",
				Protocol: "tcp",
				Port:     8080,
			},
		}, {
			Source: models.Source{ID: "dead-guid", Tag: "tag"},
			Destination: models.Destination{
				ID:       "live-guid",
				Tag:      "tag",
				Protocol: "udp",
				Port:     1234,
			},
		}, {
			Source: models.Source{ID: "live-guid", Tag: "tag"},
			Destination: models.Destination{
				ID:       "dead-guid",
				Tag:      "tag",
				Protocol: "udp",
				Port:     1234,
			},
		}}

		fakeStore = &fakes.Store{}
		fakeUAAClient = &fakes.UAAClient{}
		fakeCCClient = &fakes.CCClient{}
		logger = lagertest.NewTestLogger("test")

		policyCleaner = &handlers.PolicyCleaner{
			Logger:    logger,
			Store:     fakeStore,
			UAAClient: fakeUAAClient,
			CCClient:  fakeCCClient,
		}

		fakeUAAClient.GetTokenReturns("valid-token", nil)
		fakeStore.AllReturns(allPolicies, nil)
		fakeCCClient.GetAllAppGUIDsReturns(map[string]interface{}{"live-guid": nil}, nil)
	})

	It("Deletes policies that reference apps that do not exist", func() {
		policies, err := policyCleaner.DeleteStalePolicies()
		Expect(err).NotTo(HaveOccurred())

		Expect(fakeStore.AllCallCount()).To(Equal(1))
		Expect(fakeUAAClient.GetTokenCallCount()).To(Equal(1))
		Expect(fakeCCClient.GetAllAppGUIDsCallCount()).To(Equal(1))
		Expect(fakeCCClient.GetAllAppGUIDsArgsForCall(0)).To(Equal("valid-token"))
		stalePolicies := allPolicies[1:]
		Expect(fakeStore.DeleteCallCount()).To(Equal(1))
		Expect(fakeStore.DeleteArgsForCall(0)).To(Equal(stalePolicies))

		Expect(logger).To(gbytes.Say("deleting stale policies:.*policies.*dead-guid.*dead-guid.*total_policies\":2"))
		Expect(policies).To(Equal(stalePolicies))
	})

	Context("When retrieving policies from the db fails", func() {
		BeforeEach(func() {
			fakeStore.AllReturns(nil, errors.New("potato"))
		})

		It("returns a meaningful error", func() {
			policies, err := policyCleaner.DeleteStalePolicies()
			Expect(err).To(MatchError("database read failed: potato"))
			Expect(policies).To(BeNil())
		})

		It("logs the error", func() {
			policyCleaner.DeleteStalePolicies()
			Expect(logger).To(gbytes.Say("store-list-policies-failed.*potato"))
		})
	})

	Context("When getting the UAA token fails", func() {
		BeforeEach(func() {
			fakeUAAClient.GetTokenReturns("", errors.New("potato"))
		})

		It("returns a meaningful error", func() {
			policies, err := policyCleaner.DeleteStalePolicies()
			Expect(err).To(MatchError("get UAA token failed: potato"))
			Expect(policies).To(BeNil())
		})

		It("logs the full error", func() {
			policyCleaner.DeleteStalePolicies()
			Expect(logger).To(gbytes.Say("get-uaa-token-failed.*potato"))
		})
	})

	Context("When getting the apps from the Cloud-Controller fails", func() {
		BeforeEach(func() {
			fakeCCClient.GetAllAppGUIDsReturns(nil, errors.New("potato"))
		})

		It("returns a meaningful error", func() {
			policies, err := policyCleaner.DeleteStalePolicies()
			Expect(err).To(MatchError("get app guids from Cloud-Controller failed: potato"))
			Expect(policies).To(BeNil())
		})

		It("logs the full error", func() {
			policyCleaner.DeleteStalePolicies()
			Expect(logger).To(gbytes.Say("cc-get-app-guids-failed.*potato"))
		})
	})

	Context("When deleting the policies fails", func() {
		BeforeEach(func() {
			fakeStore.DeleteReturns(errors.New("potato"))
		})

		It("returns a meaningful error", func() {
			policies, err := policyCleaner.DeleteStalePolicies()
			Expect(err).To(MatchError("database write failed: potato"))
			Expect(policies).To(BeNil())
		})

		It("logs the full error", func() {
			policyCleaner.DeleteStalePolicies()
			Expect(logger).To(gbytes.Say("store-delete-policies-failed.*potato"))
		})
	})
})
