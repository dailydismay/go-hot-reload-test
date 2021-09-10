package handlers_test

import (
	"fmt"
	"gofirstapp/internal/client/mock"
	"gofirstapp/internal/handlers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {
	var (
		ctrl *gomock.Controller
		client *mock.MockClient
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		client = mock.NewMockClient(ctrl)
	})

	When("/kanye", func() {
		When("GET /kanye (200)", func ()  {
			quote := "moya citata yopta huli"

			mockResponse := fmt.Sprintf(`{ "quote": "%s" }`, quote)

			resp := &http.Response{
				Body: ioutil.NopCloser(strings.NewReader(mockResponse)),
				StatusCode: http.StatusOK,
			}

			BeforeEach(func ()  {
				client.EXPECT().Get(gomock.Eq("https://api.kanye.rest/")).Return(resp, nil)
			})

			It("Should respond correctly with status OK", func ()  {
				req, _ := http.NewRequest("GET", "/kanye", nil)

				rr := httptest.NewRecorder()

				handler := http.HandlerFunc(handlers.NewHandlers(handlers.WithCustomClient(client)).Kanye)
				handler.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
				Expect(rr.Body.String()).To(Equal(quote))
			})
		})
	})
})
