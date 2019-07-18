package models

type AutoGenerated struct {
	ID                 int         `json:"id"`
	Type               int         `json:"type"`
	Subject            string      `json:"subject"`
	Category           string      `json:"category"`
	Urgency            string      `json:"urgency"`
	Status             string      `json:"status"`
	BaseStatus         string      `json:"baseStatus"`
	Justification      interface{} `json:"justification"`
	Origin             int         `json:"origin"`
	CreatedDate        string      `json:"createdDate"`
	IsDeleted          bool        `json:"isDeleted"`
	OriginEmailAccount interface{} `json:"originEmailAccount"`
	Owner              struct {
		ID           string      `json:"id"`
		PersonType   int         `json:"personType"`
		ProfileType  int         `json:"profileType"`
		BusinessName string      `json:"businessName"`
		Email        string      `json:"email"`
		Phone        string      `json:"phone"`
		Address      interface{} `json:"address"`
		Complement   interface{} `json:"complement"`
		Cep          interface{} `json:"cep"`
		City         interface{} `json:"city"`
		Bairro       interface{} `json:"bairro"`
		Number       interface{} `json:"number"`
		Reference    interface{} `json:"reference"`
	} `json:"owner"`
	OwnerTeam string `json:"ownerTeam"`
	CreatedBy struct {
		ID           string      `json:"id"`
		PersonType   int         `json:"personType"`
		ProfileType  int         `json:"profileType"`
		BusinessName string      `json:"businessName"`
		Email        string      `json:"email"`
		Phone        string      `json:"phone"`
		Address      interface{} `json:"address"`
		Complement   interface{} `json:"complement"`
		Cep          interface{} `json:"cep"`
		City         interface{} `json:"city"`
		Bairro       interface{} `json:"bairro"`
		Number       interface{} `json:"number"`
		Reference    interface{} `json:"reference"`
	} `json:"createdBy"`
	ServiceFull              []string      `json:"serviceFull"`
	ServiceFirstLevelID      int           `json:"serviceFirstLevelId"`
	ServiceFirstLevel        string        `json:"serviceFirstLevel"`
	ServiceSecondLevel       string        `json:"serviceSecondLevel"`
	ServiceThirdLevel        string        `json:"serviceThirdLevel"`
	ContactForm              interface{}   `json:"contactForm"`
	Tags                     []interface{} `json:"tags"`
	Cc                       interface{}   `json:"cc"`
	ResolvedIn               string        `json:"resolvedIn"`
	ClosedIn                 string        `json:"closedIn"`
	CanceledIn               interface{}   `json:"canceledIn"`
	ActionCount              int           `json:"actionCount"`
	LifeTimeWorkingTime      int           `json:"lifeTimeWorkingTime"`
	StoppedTime              int           `json:"stoppedTime"`
	StoppedTimeWorkingTime   int           `json:"stoppedTimeWorkingTime"`
	ResolvedInFirstCall      bool          `json:"resolvedInFirstCall"`
	ChatWidget               interface{}   `json:"chatWidget"`
	ChatGroup                interface{}   `json:"chatGroup"`
	ChatTalkTime             interface{}   `json:"chatTalkTime"`
	ChatWaitingTime          interface{}   `json:"chatWaitingTime"`
	Sequence                 interface{}   `json:"sequence"`
	SLAAgreement             string        `json:"slaAgreement"`
	SLAAgreementRule         string        `json:"slaAgreementRule"`
	SLASolutionTime          int           `json:"slaSolutionTime"`
	SLAResponseTime          int           `json:"slaResponseTime"`
	SLASolutionChangedByUser bool          `json:"slaSolutionChangedByUser"`
	SLASolutionChangedBy     struct {
		ID           string      `json:"id"`
		PersonType   int         `json:"personType"`
		ProfileType  int         `json:"profileType"`
		BusinessName string      `json:"businessName"`
		Email        string      `json:"email"`
		Phone        string      `json:"phone"`
		Address      interface{} `json:"address"`
		Complement   interface{} `json:"complement"`
		Cep          interface{} `json:"cep"`
		City         interface{} `json:"city"`
		Bairro       interface{} `json:"bairro"`
		Number       interface{} `json:"number"`
		Reference    interface{} `json:"reference"`
	} `json:"slaSolutionChangedBy"`
	SLASolutionDate                string      `json:"slaSolutionDate"`
	SLASolutionDateIsPaused        bool        `json:"slaSolutionDateIsPaused"`
	JiraIssueKey                   interface{} `json:"jiraIssueKey"`
	RedmineIssueID                 interface{} `json:"redmineIssueId"`
	MovideskTicketNumber           interface{} `json:"movideskTicketNumber"`
	LinkedToIntegratedTicketNumber interface{} `json:"linkedToIntegratedTicketNumber"`
	ReopenedIn                     interface{} `json:"reopenedIn"`
	LastActionDate                 string      `json:"lastActionDate"`
	LastUpdate                     string      `json:"lastUpdate"`
	SLAResponseDate                string      `json:"slaResponseDate"`
	SLARealResponseDate            string      `json:"slaRealResponseDate"`
	Clients                        []struct {
		ID           string `json:"id"`
		PersonType   int    `json:"personType"`
		ProfileType  int    `json:"profileType"`
		BusinessName string `json:"businessName"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
		IsDeleted    bool   `json:"isDeleted"`
		Organization struct {
			ID           string      `json:"id"`
			PersonType   int         `json:"personType"`
			ProfileType  int         `json:"profileType"`
			BusinessName string      `json:"businessName"`
			Email        interface{} `json:"email"`
			Phone        string      `json:"phone"`
			Address      interface{} `json:"address"`
			Complement   interface{} `json:"complement"`
			Cep          interface{} `json:"cep"`
			City         interface{} `json:"city"`
			Bairro       interface{} `json:"bairro"`
			Number       interface{} `json:"number"`
			Reference    interface{} `json:"reference"`
		} `json:"organization"`
		Address    interface{} `json:"address"`
		Complement interface{} `json:"complement"`
		Cep        interface{} `json:"cep"`
		City       interface{} `json:"city"`
		Bairro     interface{} `json:"bairro"`
		Number     interface{} `json:"number"`
		Reference  interface{} `json:"reference"`
	} `json:"clients"`
	Actions []struct {
		ID              int         `json:"id"`
		Type            int         `json:"type"`
		Origin          int         `json:"origin"`
		Description     string      `json:"description"`
		HTMLDescription string      `json:"htmlDescription"`
		Status          string      `json:"status"`
		Justification   interface{} `json:"justification"`
		CreatedDate     string      `json:"createdDate"`
		CreatedBy       struct {
			ID           string      `json:"id"`
			PersonType   int         `json:"personType"`
			ProfileType  int         `json:"profileType"`
			BusinessName string      `json:"businessName"`
			Email        string      `json:"email"`
			Phone        string      `json:"phone"`
			Address      interface{} `json:"address"`
			Complement   interface{} `json:"complement"`
			Cep          interface{} `json:"cep"`
			City         interface{} `json:"city"`
			Bairro       interface{} `json:"bairro"`
			Number       interface{} `json:"number"`
			Reference    interface{} `json:"reference"`
		} `json:"createdBy"`
		IsDeleted        bool          `json:"isDeleted"`
		TimeAppointments []interface{} `json:"timeAppointments"`
		Attachments      []interface{} `json:"attachments"`
		Expenses         []interface{} `json:"expenses"`
		Tags             []interface{} `json:"tags"`
	} `json:"actions"`
	ParentTickets   []interface{} `json:"parentTickets"`
	ChildrenTickets []interface{} `json:"childrenTickets"`
	OwnerHistories  []struct {
		OwnerTeam string `json:"ownerTeam"`
		Owner     struct {
			ID           string      `json:"id"`
			PersonType   int         `json:"personType"`
			ProfileType  int         `json:"profileType"`
			BusinessName string      `json:"businessName"`
			Email        string      `json:"email"`
			Phone        string      `json:"phone"`
			Address      interface{} `json:"address"`
			Complement   interface{} `json:"complement"`
			Cep          interface{} `json:"cep"`
			City         interface{} `json:"city"`
			Bairro       interface{} `json:"bairro"`
			Number       interface{} `json:"number"`
			Reference    interface{} `json:"reference"`
		} `json:"owner"`
		ChangedBy struct {
			ID           string      `json:"id"`
			PersonType   int         `json:"personType"`
			ProfileType  int         `json:"profileType"`
			BusinessName string      `json:"businessName"`
			Email        string      `json:"email"`
			Phone        string      `json:"phone"`
			Address      interface{} `json:"address"`
			Complement   interface{} `json:"complement"`
			Cep          interface{} `json:"cep"`
			City         interface{} `json:"city"`
			Bairro       interface{} `json:"bairro"`
			Number       interface{} `json:"number"`
			Reference    interface{} `json:"reference"`
		} `json:"changedBy"`
		ChangedDate               string  `json:"changedDate"`
		PermanencyTimeFullTime    float64 `json:"permanencyTimeFullTime"`
		PermanencyTimeWorkingTime float64 `json:"permanencyTimeWorkingTime"`
	} `json:"ownerHistories"`
	StatusHistories []struct {
		Status        string      `json:"status"`
		Justification interface{} `json:"justification"`
		ChangedBy     struct {
			ID           string      `json:"id"`
			PersonType   int         `json:"personType"`
			ProfileType  int         `json:"profileType"`
			BusinessName string      `json:"businessName"`
			Email        string      `json:"email"`
			Phone        string      `json:"phone"`
			Address      interface{} `json:"address"`
			Complement   interface{} `json:"complement"`
			Cep          interface{} `json:"cep"`
			City         interface{} `json:"city"`
			Bairro       interface{} `json:"bairro"`
			Number       interface{} `json:"number"`
			Reference    interface{} `json:"reference"`
		} `json:"changedBy"`
		ChangedDate               string  `json:"changedDate"`
		PermanencyTimeFullTime    float64 `json:"permanencyTimeFullTime"`
		PermanencyTimeWorkingTime float64 `json:"permanencyTimeWorkingTime"`
	} `json:"statusHistories"`
	SatisfactionSurveyResponses []struct {
		ID          int `json:"id"`
		ResponsedBy struct {
			ID           string      `json:"id"`
			PersonType   int         `json:"personType"`
			ProfileType  int         `json:"profileType"`
			BusinessName string      `json:"businessName"`
			Email        string      `json:"email"`
			Phone        string      `json:"phone"`
			Address      interface{} `json:"address"`
			Complement   interface{} `json:"complement"`
			Cep          interface{} `json:"cep"`
			City         interface{} `json:"city"`
			Bairro       interface{} `json:"bairro"`
			Number       interface{} `json:"number"`
		Reference    interface{} `json:"reference"`
		} `json:"responsedBy"`
			ResponseDate                               string      `json:"responseDate"`
			SatisfactionSurveyModel                    int         `json:"satisfactionSurveyModel"`
			SatisfactionSurveyNetPromoterScoreResponse int         `json:"satisfactionSurveyNetPromoterScoreResponse"`
			SatisfactionSurveyPositiveNegativeResponse interface{} `json:"satisfactionSurveyPositiveNegativeResponse"`
			SatisfactionSurveySmileyFacesResponse      interface{} `json:"satisfactionSurveySmileyFacesResponse"`
		Comments                                   interface{} `json:"comments"`
	} `json:"satisfactionSurveyResponses"`
	CustomFieldValues []interface{} `json:"customFieldValues"`
	Assets            []interface{} `json:"assets"`
	WebhookEvents     interface{}   `json:"webhookEvents"`
}