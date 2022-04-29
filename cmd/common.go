package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type ApiLoginResponse struct {
	Data string `json:"data"`
	Meta struct {
		CurrentPage int `json:"currentPage"`
		LastPage    int `json:"lastPage"`
		Total       int `json:"total"`
		PerPage     int `json:"perPage"`
	} `json:"meta"`
}

type ApiMeResponse struct {
	Data []struct {
		ID          string `json:"id"`
		GroupID     string `json:"groupId"`
		GroupName   string `json:"groupName"`
		Group       string `json:"group"`
		CompanyID   string `json:"companyId"`
		CompanyName string `json:"companyName"`
		Company     struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			BetaVersion string `json:"betaVersion"`
			ReferredID  string `json:"referredId"`
			Logo        string `json:"logo"`
			Icon        string `json:"icon"`
		} `json:"company"`
		CompanyEmailNotification string `json:"companyEmailNotification"`
		MyCompany                struct {
			ID                       string    `json:"id"`
			Name                     string    `json:"name"`
			Country                  string    `json:"country"`
			Vat                      string    `json:"vat"`
			Status                   string    `json:"status"`
			CreatedAt                time.Time `json:"createdAt"`
			UpdatedAt                time.Time `json:"updatedAt"`
			Language                 string    `json:"language"`
			Logo                     string    `json:"logo"`
			Icon                     string    `json:"icon"`
			EmailNotification        string    `json:"emailNotification"`
			Phone                    string    `json:"phone"`
			NumEmployees             int       `json:"numEmployees"`
			Origin                   string    `json:"origin"`
			LegalAccepted            bool      `json:"legalAccepted"`
			IsSubscribedToNewsletter bool      `json:"isSubscribedToNewsletter"`
			BetaVersion              string    `json:"betaVersion"`
			GroupID                  string    `json:"groupId"`
		} `json:"myCompany"`
		OnBoardingMetadata struct {
			Employee bool `json:"employee"`
			Office   bool `json:"office"`
			Holiday  bool `json:"holiday"`
			Schedule bool `json:"schedule"`
			Skipped  bool `json:"skipped"`
		} `json:"onBoardingMetadata"`
		GroupFrontMetadata    []string `json:"groupFrontMetadata"`
		EmployeeFrontMetadata []string `json:"employeeFrontMetadata"`
		FirstName             string   `json:"firstName"`
		LastName              string   `json:"lastName"`
		WorkStatus            string   `json:"workStatus"`
		Phone                 string   `json:"phone"`
		Departments           []struct {
			ID              string `json:"id"`
			Name            string `json:"name"`
			NumberEmployees int    `json:"numberEmployees"`
			CompanyName     string `json:"companyName"`
		} `json:"departments"`
		Offices []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Address     string `json:"address"`
			Coordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"coordinates"`
			Description                  string `json:"description"`
			Radio                        int    `json:"radio"`
			DefaultEmployeesDateTimeZone string `json:"defaultEmployeesDateTimeZone"`
			NumberEmployees              int    `json:"numberEmployees"`
			CompanyName                  string `json:"companyName"`
		} `json:"offices"`
		WorkDayProgress    int       `json:"workDayProgress"`
		CreatedAt          time.Time `json:"createdAt"`
		UpdatedAt          time.Time `json:"updatedAt"`
		Email              string    `json:"email"`
		IsAdmin            bool      `json:"isAdmin"`
		IsSuperAdmin       bool      `json:"isSuperAdmin"`
		AccumulatedSeconds int       `json:"accumulatedSeconds"`
		DailySchedule      int       `json:"dailySchedule"`
		Status             string    `json:"status"`
		LastOpenedCheck    time.Time `json:"lastOpenedCheck"`
		LastCheck          struct {
			CheckID            string    `json:"checkId"`
			CheckInDatetime    time.Time `json:"checkInDatetime"`
			CheckOutDatetime   time.Time `json:"checkOutDatetime"`
			CheckInCoordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"checkInCoordinates"`
			CheckOutCoordinates struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"checkOutCoordinates"`
			WorkCheckTypeColor string `json:"workCheckTypeColor"`
			WorkCheckTypeName  string `json:"workCheckTypeName"`
		} `json:"lastCheck"`
		EntityRoles            []string  `json:"entityRoles"`
		Dni                    string    `json:"dni"`
		Ssn                    string    `json:"ssn"`
		Gender                 string    `json:"gender"`
		Birthday               string    `json:"birthday"`
		Code                   int       `json:"code"`
		Pin                    int       `json:"pin"`
		ImageProfileURL        string    `json:"imageProfileURL"`
		ImageProfileRoundedURL string    `json:"imageProfileRoundedURL"`
		Language               string    `json:"language"`
		CompanyLanguage        string    `json:"companyLanguage"`
		InvitedAt              time.Time `json:"invitedAt"`
		InvitationAcceptedAt   string    `json:"invitationAcceptedAt"`
		InvitationStatus       string    `json:"invitationStatus"`
		Timer                  string    `json:"timer"`
		SubscriptionLevel      int       `json:"subscriptionLevel"`
		PermissionVisibilities struct {
			Vacations                                     bool `json:"vacations"`
			OtherEmployeesHours                           bool `json:"otherEmployeesHours"`
			OtherEmployeesExtraHours                      bool `json:"otherEmployeesExtraHours"`
			Whosin                                        bool `json:"whosin"`
			Timeline                                      bool `json:"timeline"`
			Document                                      bool `json:"document"`
			Shift                                         bool `json:"shift"`
			Timer                                         bool `json:"timer"`
			EmployeeViewOtherVacations                    bool `json:"employeeViewOtherVacations"`
			Suprema                                       bool `json:"suprema"`
			WhosinForDepartmentCenter                     bool `json:"whosin_for_department_center"`
			TimelineForDepartmentCenter                   bool `json:"timeline_for_department_center"`
			EmployeeViewOtherVacationsForDepartmentCenter bool `json:"employeeViewOtherVacations_for_department_center"`
			EmployeeViewSignings                          bool `json:"employee_view_signings"`
			EmployeeHome                                  bool `json:"employee_home"`
			Statistics                                    bool `json:"statistics"`
			CompanyLogo                                   bool `json:"companyLogo"`
		} `json:"permissionVisibilities"`
		Configurations struct {
			RequireCoordsForMobileCheck   bool `json:"RequireCoordsForMobileCheck"`
			RequireCoordsForWebCheck      bool `json:"RequireCoordsForWebCheck"`
			CheckValidationConfiguration  bool `json:"CheckValidationConfiguration"`
			FaceIDConfiguration           bool `json:"FaceIdConfiguration"`
			LoginOnlyWithSSOConfiguration bool `json:"LoginOnlyWithSSOConfiguration"`
			ShowOrganizationChart         bool `json:"ShowOrganizationChart"`
		} `json:"configurations"`
		SubscriptionStatus       string    `json:"subscriptionStatus"`
		SubscriptionTrialEndDate time.Time `json:"subscriptionTrialEndDate"`
		MainOfficeID             string    `json:"mainOfficeId"`
		MainOffice               struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Address     string `json:"address"`
			Coordinates struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"coordinates"`
			Description                  string `json:"description"`
			Radio                        int    `json:"radio"`
			DefaultEmployeesDateTimeZone string `json:"defaultEmployeesDateTimeZone"`
			NumberEmployees              int    `json:"numberEmployees"`
			CompanyName                  string `json:"companyName"`
		} `json:"mainOffice"`
		ContractID       string `json:"contractId"`
		EmployeeTimeZone string `json:"employeeTimeZone"`
		CompanyModules   struct {
			Ticket            int `json:"ticket"`
			ModularReport     int `json:"modular_report"`
			Documents         int `json:"documents"`
			HoursBag          int `json:"hours_bag"`
			Shift             int `json:"shift"`
			Evaluation        int `json:"evaluation"`
			OrganizationChart int `json:"organization_chart"`
			Core              int `json:"core"`
			Announcement      int `json:"announcement"`
			Talent            int `json:"talent"`
			Okr               int `json:"okr"`
			OnBoarding        int `json:"on_boarding"`
			DigitalSignature  int `json:"digital_signature"`
			Room              int `json:"room"`
			SplitPayroll      int `json:"split_payroll"`
			Payroll           int `json:"payroll"`
			Time              int `json:"time"`
			Recruitment       int `json:"recruitment"`
			OneToOne          int `json:"one_to_one"`
			Schedule          int `json:"schedule"`
			Chat              int `json:"chat"`
			API               int `json:"api"`
			Office            int `json:"office"`
			Vacation          int `json:"vacation"`
			Permission        int `json:"permission"`
			Workbreak         int `json:"workbreak"`
			Report            int `json:"report"`
			Group             int `json:"group"`
			Role              int `json:"role"`
			Poll              int `json:"poll"`
			Timeline          int `json:"timeline"`
			Device            int `json:"device"`
			EmployeePortal    int `json:"employee_portal"`
			Check             int `json:"check"`
			DayOff            int `json:"day_off"`
			Timer             int `json:"timer"`
			Utilities         int `json:"utilities"`
		} `json:"companyModules"`
		MarketPlaceModules struct {
			Benefit           int `json:"benefit"`
			Travel            int `json:"travel"`
			AdvancedSignature int `json:"advanced-signature"`
			Payflow           int `json:"payflow"`
			Sign3G            int `json:"sign3g"`
			Travelperk        int `json:"travelperk"`
			Hastee            int `json:"hastee"`
		} `json:"marketPlaceModules"`
		AccessibleCompanies    []string `json:"accessibleCompanies"`
		NotAccessibleCompanies []string `json:"notAccessibleCompanies"`
		IsGroupOwner           bool     `json:"isGroupOwner"`
		HighestRole            string   `json:"highestRole"`
		HasWorkBreaks          bool     `json:"hasWorkBreaks"`
		NumActiveEmployees     int      `json:"numActiveEmployees"`
		MaxEmployees           int      `json:"maxEmployees"`
		Customer               struct {
			ID                string `json:"id"`
			EntityID          string `json:"entityId"`
			EntityType        string `json:"entityType"`
			FirstName         string `json:"firstName"`
			LastName          string `json:"lastName"`
			CompanyName       string `json:"companyName"`
			Email             string `json:"email"`
			Locale            string `json:"locale"`
			PreferredCurrency string `json:"preferredCurrency"`
			Phone             string `json:"phone"`
			VatNumber         string `json:"vatNumber"`
			BillingAddress    struct {
				Line1   string `json:"line1"`
				City    string `json:"city"`
				State   string `json:"state"`
				Country string `json:"country"`
				Zip     string `json:"zip"`
			} `json:"billingAddress"`
			CurrentActiveEmployees int       `json:"currentActiveEmployees"`
			CreatedAt              time.Time `json:"createdAt"`
			UpdatedAt              time.Time `json:"updatedAt"`
			Subscription           struct {
				ID                  string    `json:"id"`
				PlanName            string    `json:"planName"`
				MaxEmployees        int       `json:"maxEmployees"`
				Quantity            int       `json:"quantity"`
				FreeQuantity        int       `json:"freeQuantity"`
				SubscriptionStatus  string    `json:"subscriptionStatus"`
				NextBillingAt       time.Time `json:"nextBillingAt"`
				TrialEndDate        time.Time `json:"trialEndDate"`
				MemberSince         time.Time `json:"memberSince"`
				CreatedAt           time.Time `json:"createdAt"`
				UpdatedAt           time.Time `json:"updatedAt"`
				Amount              string    `json:"amount"`
				Currency            string    `json:"currency"`
				BillingPeriodUnit   string    `json:"billingPeriodUnit"`
				BillingPeriod       int       `json:"billingPeriod"`
				PlanID              string    `json:"planId"`
				CompaniesSharedWith []string  `json:"companiesSharedWith"`
				Addons              []string  `json:"addons"`
				SubscriptionLevel   int       `json:"subscriptionLevel"`
				Discount            string    `json:"discount"`
				OriginID            string    `json:"originId"`
				CancelledAt         string    `json:"cancelledAt"`
			} `json:"subscription"`
		} `json:"customer"`
		ChatToken          string `json:"chatToken"`
		PricePerHour       int    `json:"pricePerHour"`
		IdentityNumberType string `json:"identityNumberType"`
		Favorite           bool   `json:"favorite"`
		AccountNumber      string `json:"accountNumber"`
		CompanyURLOrigin   struct {
			Value string `json:"value"`
		} `json:"companyUrlOrigin"`
		EntityPermissions []string `json:"entityPermissions"`
	} `json:"data"`
	Meta struct {
		CurrentPage int `json:"currentPage"`
		LastPage    int `json:"lastPage"`
		Total       int `json:"total"`
		PerPage     int `json:"perPage"`
	} `json:"meta"`
}

func getSessionId(sesameUrl, sesameUsername, sesamePassword string) (string, error) {
	values := map[string]string{"origin": "web"}
	jsonData, _ := json.Marshal(values)

	req, _ := http.NewRequest(http.MethodPost, sesameUrl+"/api/v3/security/login", bytes.NewBuffer(jsonData))
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	post := &ApiLoginResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	return post.Data, nil
}

func getUserId(sesameUrl, sessionId string) (string, error) {
	values := map[string]string{"origin": "web"}
	jsonData, _ := json.Marshal(values)
	req, _ := http.NewRequest(http.MethodPost, sesameUrl+"/api/v3/security/me", bytes.NewBuffer(jsonData))
	req.Header.Add("Authorization", "Bearer "+sessionId)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	post := &ApiMeResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}
	return post.Data[0].ID, nil
}
