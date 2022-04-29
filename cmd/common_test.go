package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSessionIdWithoutError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/security/login" {
			t.Errorf("Expected to request '/api/v3/security/login', got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":"27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f"}`))
	}))
	defer server.Close()

	expectedSessionId := "27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f"
	sessionId, err := getSessionId(server.URL, mock.Anything, mock.Anything)

	assert.Nil(t, err)
	assert.Equal(t, expectedSessionId, sessionId)
}

func TestGetUserIdWithoutError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v3/security/me" {
			t.Errorf("Expected to request '/api/v3/security/me', got: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer 27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f" {
			t.Errorf("Expected Authorization: Bearer 27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f, got: %s", r.Header.Get("Accept"))
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":[{"id":"ecb42151-1f78-4e23-8f62-40b88df9f2c4","groupId":null,"groupName":null,"group":null,"companyId":"495dc019-8d9e-449d-a5fd-a6828e226d88","companyName":"Mercadona Tech","company":{"id":"495dc019-8d9e-449d-a5fd-a6828e226d88","name":"Mercadona Tech","betaVersion":null,"referredId":null,"logo":"https:\/\/sesame-v2-prod-companies.s3.eu-west-3.amazonaws.com\/821ffa77-edad-4a11-aed6-8810c8e256b2\/public-read\/5eb5d26d-03c8-44ad-acc3-6d8896533dad.svg","icon":"https:\/\/sesame-v2-prod-companies.s3.eu-west-3.amazonaws.com\/821ffa77-edad-4a11-aed6-8810c8e256b2\/public-read\/d656b968-de43-4c2c-9133-2da68b17b90d.png"},"companyEmailNotification":"kbednari@mercadona.es","myCompany":{"id":"495dc019-8d9e-449d-a5fd-a6828e226d88","name":"Mercadona Tech","country":"ES","vat":null,"status":"active","createdAt":"2022-03-10T19:06:17+00:00","updatedAt":"2022-04-06T09:29:45+00:00","language":"es-ES","logo":"https:\/\/sesame-v2-prod-companies.s3.eu-west-3.amazonaws.com\/821ffa77-edad-4a11-aed6-8810c8e256b2\/public-read\/5eb5d26d-03c8-44ad-acc3-6d8896533dad.svg","icon":"https:\/\/sesame-v2-prod-companies.s3.eu-west-3.amazonaws.com\/821ffa77-edad-4a11-aed6-8810c8e256b2\/public-read\/d656b968-de43-4c2c-9133-2da68b17b90d.png","emailNotification":"kbednari@mercadona.es","phone":"684000129","numEmployees":117,"origin":"web","legalAccepted":true,"isSubscribedToNewsletter":true,"betaVersion":null,"groupId":null},"onBoardingMetadata":{"employee":true,"office":true,"holiday":true,"schedule":true,"skipped":true},"groupFrontMetadata":[],"employeeFrontMetadata":[],"firstName":"Miguel","lastName":"Villar Juan","workStatus":"offline","phone":null,"departments":[{"id":"3849d42c-1b40-4776-9e75-accba92d1e33","name":"Engineering","numberEmployees":0,"companyName":null}],"offices":[{"id":"35f4ea54-5de5-41a6-88ed-e721c2635291","name":"Tech Hub Valencia","address":"Pl. d\u0027Am\u00e8rica, 2, 46004 Val\u00e8ncia, Valencia, Espa\u00f1a","coordinates":{"latitude":39.4695297,"longitude":-0.3649067},"description":null,"radio":150,"defaultEmployeesDateTimeZone":"Europe\/Madrid","numberEmployees":0,"companyName":null}],"workDayProgress":0,"createdAt":"2022-04-01T10:26:16+00:00","updatedAt":"2022-04-29T06:53:11+00:00","email":"mvillju@mercadona.es","isAdmin":false,"isSuperAdmin":false,"accumulatedSeconds":0,"dailySchedule":32400,"status":"active","lastOpenedCheck":"2022-04-28T14:06:08+02:00","lastCheck":{"checkId":"87c5249b-f3b6-4234-9feb-8b98acdea50d","checkInDatetime":"2022-04-28T14:06:08+02:00","checkOutDatetime":"2022-04-28T17:56:52+02:00","checkInCoordinates":{"latitude":null,"longitude":null},"checkOutCoordinates":{"latitude":null,"longitude":null},"workCheckTypeColor":null,"workCheckTypeName":null},"entityRoles":[],"dni":"","ssn":null,"gender":null,"birthday":"1991-11-15","code":63,"pin":1947,"imageProfileURL":"https:\/\/sesame-v2-prod-companies.s3.eu-west-3.amazonaws.com\/821ffa77-edad-4a11-aed6-8810c8e256b2\/public-read\/6ca3c7d2-cff9-4bc0-a7d6-206f276ad77e.jpg","imageProfileRoundedURL":"https:\/\/back.sesametime.com\/api\/v3\/images\/ecb43050-1e68-4e23-8f62-40b88df9f2b4","language":null,"companyLanguage":"es-ES","invitedAt":"2022-04-20T11:10:51+00:00","invitationAcceptedAt":null,"invitationStatus":"resend_available","timer":null,"subscriptionLevel":10,"permissionVisibilities":{"vacations":true,"otherEmployeesHours":true,"otherEmployeesExtraHours":true,"whosin":true,"timeline":true,"document":true,"shift":true,"timer":true,"employeeViewOtherVacations":true,"suprema":false,"whosin_for_department_center":false,"timeline_for_department_center":false,"employeeViewOtherVacations_for_department_center":false,"employee_view_signings":true,"employee_home":true,"statistics":true,"companyLogo":false},"configurations":{"RequireCoordsForMobileCheck":false,"RequireCoordsForWebCheck":false,"CheckValidationConfiguration":false,"FaceIdConfiguration":false,"LoginOnlyWithSSOConfiguration":false,"ShowOrganizationChart":false},"subscriptionStatus":"active","subscriptionTrialEndDate":"2022-03-24T19:06:18+00:00","mainOfficeId":"35f4ea54-5de5-41a6-88ed-e721c2635291","mainOffice":{"id":"35f4ea54-5de5-41a6-88ed-e721c2635291","name":"Tech Hub Valencia","address":"Pl. d\u0027Am\u00e8rica, 2, 46004 Val\u00e8ncia, Valencia, Espa\u00f1a","coordinates":{"latitude":39.4695297,"longitude":-0.3649067},"description":null,"radio":150,"defaultEmployeesDateTimeZone":"Europe\/Madrid","numberEmployees":0,"companyName":null},"contractId":null,"employeeTimeZone":"Europe\/Madrid","companyModules":{"ticket":-10,"modular_report":-10,"documents":-10,"hours_bag":-10,"shift":-10,"evaluation":-10,"organization_chart":10,"core":10,"announcement":10,"talent":10,"okr":10,"on_boarding":10,"digital_signature":10,"room":10,"split_payroll":10,"payroll":10,"time":10,"recruitment":10,"one_to_one":10,"schedule":10,"chat":10,"api":10,"office":10,"vacation":10,"permission":10,"workbreak":10,"report":10,"group":10,"role":10,"poll":10,"timeline":10,"device":10,"employee_portal":10,"check":10,"day_off":10,"timer":10,"utilities":10},"marketPlaceModules":{"benefit":-10,"travel":-10,"advanced-signature":-10,"payflow":-10,"sign3g":-10,"travelperk":-10,"hastee":-10},"accessibleCompanies":[],"notAccessibleCompanies":[],"isGroupOwner":false,"highestRole":null,"hasWorkBreaks":true,"numActiveEmployees":117,"maxEmployees":120,"customer":{"id":"1c60a3a9-4a29-4958-97b2-92588a9eac53","entityId":"495dc019-8d9e-449d-a5fd-a6828e226d88","entityType":"company","firstName":"jose ramon","lastName":"perez aguera","companyName":"Mercadona Tech","email":"jpereag@mercadona.es","locale":"es-ES","preferredCurrency":"EUR","phone":null,"vatNumber":null,"billingAddress":{"line1":null,"city":null,"state":null,"country":null,"zip":null},"currentActiveEmployees":117,"createdAt":"2022-03-10T19:06:17+00:00","updatedAt":"2022-03-10T19:06:17+00:00","subscription":{"id":"6619a7b0-8b55-4edc-bff3-fb579b289ef7","planName":"Time Yearly","maxEmployees":120,"quantity":120,"freeQuantity":0,"subscriptionStatus":"active","nextBillingAt":"2023-03-31T13:01:11+02:00","trialEndDate":"2022-03-24T20:06:18+01:00","memberSince":"2022-03-10T20:06:17+01:00","createdAt":"2022-03-10T19:06:17+00:00","updatedAt":"2022-03-31T11:01:22+00:00","amount":"2.880,00","currency":"EUR","billingPeriodUnit":"month","billingPeriod":12,"planId":"6a7349f5-e8b4-4029-8479-eac6d62d7562","companiesSharedWith":[],"addons":[],"subscriptionLevel":10,"discount":null,"originId":"v2-yearly-time-2021-11-10","cancelledAt":null}},"chatToken":null,"pricePerHour":0,"identityNumberType":"nic","favorite":false,"accountNumber":null,"companyUrlOrigin":{"value":"sesametime"},"entityPermissions":[]}],"meta":{"currentPage":1,"lastPage":1,"total":1,"perPage":1}}`))
	}))
	defer server.Close()

	expectedUserId := "ecb42151-1f78-4e23-8f62-40b88df9f2c4"
	userId, err := getUserId(server.URL, "27fdfc761b323943a1349a220fe57c0f1bcf4d6345f574b160cfc7090e4006c5f")
	assert.Nil(t, err)
	assert.Equal(t, expectedUserId, userId)
}
