package tukhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	cnst "github.com/ipthomas/tukcnst"
	util "github.com/ipthomas/tukutil"
)

type CGLRequest struct {
	Request    string
	X_Api_Key  string
	CGL_User   CGL_User
	StatusCode int
	Response   []byte
}
type CGL_User struct {
	Data struct {
		Client struct {
			BasicDetails struct {
				Address struct {
					AddressLine1 string `json:"addressLine1"`
					AddressLine2 string `json:"addressLine2"`
					AddressLine3 string `json:"addressLine3"`
					AddressLine4 string `json:"addressLine4"`
					AddressLine5 string `json:"addressLine5"`
					PostCode     string `json:"postCode"`
				} `json:"address"`
				BirthDate                    string `json:"birthDate"`
				Disability                   string `json:"disability"`
				LastEngagementByCGLDate      string `json:"lastEngagementByCGLDate"`
				LastFaceToFaceEngagementDate string `json:"lastFaceToFaceEngagementDate"`
				LocalIdentifier              int    `json:"localIdentifier"`
				Name                         struct {
					Family string `json:"family"`
					Given  string `json:"given"`
				} `json:"name"`
				NextCGLAppointmentDate interface{} `json:"nextCGLAppointmentDate"`
				NhsNumber              string      `json:"nhsNumber"`
				SexAtBirth             string      `json:"sexAtBirth"`
			} `json:"basicDetails"`
			BbvInformation struct {
				BbvTested        string      `json:"bbvTested"`
				HepCLastTestDate interface{} `json:"hepCLastTestDate"`
				HepCResult       interface{} `json:"hepCResult"`
				HivPositive      interface{} `json:"hivPositive"`
			} `json:"bbvInformation"`
			DrugTestResults struct {
				DrugTestDate          interface{} `json:"drugTestDate"`
				DrugTestSample        interface{} `json:"drugTestSample"`
				DrugTestStatus        interface{} `json:"drugTestStatus"`
				InstantOrConfirmation interface{} `json:"instantOrConfirmation"`
				Results               struct {
					Amphetamine     interface{} `json:"amphetamine"`
					Benzodiazepine  interface{} `json:"benzodiazepine"`
					Buprenorphine   interface{} `json:"buprenorphine"`
					Cannabis        interface{} `json:"cannabis"`
					Cocaine         interface{} `json:"cocaine"`
					Eddp            interface{} `json:"eddp"`
					Fentanyl        interface{} `json:"fentanyl"`
					Ketamine        interface{} `json:"ketamine"`
					Methadone       interface{} `json:"methadone"`
					Methamphetamine interface{} `json:"methamphetamine"`
					Morphine        interface{} `json:"morphine"`
					Opiates         interface{} `json:"opiates"`
					SixMam          interface{} `json:"sixMam"`
					Tramadol        interface{} `json:"tramadol"`
				} `json:"results"`
			} `json:"drugTestResults"`
			PrescribingInformation []interface{} `json:"prescribingInformation"`
			RiskInformation        struct {
				LastSelfReportedDate interface{} `json:"lastSelfReportedDate"`
				MentalHealthDomain   struct {
					AttemptedSuicide                            interface{} `json:"attemptedSuicide"`
					CurrentOrPreviousSelfHarm                   interface{} `json:"currentOrPreviousSelfHarm"`
					DiagnosedMentalHealthCondition              interface{} `json:"diagnosedMentalHealthCondition"`
					FrequentLifeThreateningSelfHarm             interface{} `json:"frequentLifeThreateningSelfHarm"`
					Hallucinations                              interface{} `json:"hallucinations"`
					HospitalAdmissionsForMentalHealth           interface{} `json:"hospitalAdmissionsForMentalHealth"`
					NoIdentifiedRisk                            interface{} `json:"noIdentifiedRisk"`
					NotEngagingWithSupport                      interface{} `json:"notEngagingWithSupport"`
					NotTakingPrescribedMedicationAsInstructed   interface{} `json:"notTakingPrescribedMedicationAsInstructed"`
					PsychiatricOrPreviousCrisisTeamIntervention interface{} `json:"psychiatricOrPreviousCrisisTeamIntervention"`
					Psychosis                                   interface{} `json:"psychosis"`
					SelfReportedMentalHealthConcerns            string      `json:"selfReportedMentalHealthConcerns"`
					ThoughtsOfSuicideOrSelfHarm                 interface{} `json:"thoughtsOfSuicideOrSelfHarm"`
				} `json:"mentalHealthDomain"`
				RiskOfHarmToSelfDomain struct {
					AssessedAsNotHavingMentalCapacity  interface{} `json:"assessedAsNotHavingMentalCapacity"`
					BeliefTheyAreWorthless             string      `json:"beliefTheyAreWorthless"`
					Hoarding                           interface{} `json:"hoarding"`
					LearningDisability                 interface{} `json:"learningDisability"`
					MeetsSafeguardingAdultsThreshold   interface{} `json:"meetsSafeguardingAdultsThreshold"`
					NoIdentifiedRisk                   interface{} `json:"noIdentifiedRisk"`
					OngoingConcernsRelatingToOwnSafety interface{} `json:"ongoingConcernsRelatingToOwnSafety"`
					ProblemsMaintainingPersonalHygiene interface{} `json:"problemsMaintainingPersonalHygiene"`
					ProblemsMeetingNutritionalNeeds    interface{} `json:"problemsMeetingNutritionalNeeds"`
					RequiresIndependentAdvocacy        interface{} `json:"requiresIndependentAdvocacy"`
					SelfNeglect                        string      `json:"selfNeglect"`
				} `json:"riskOfHarmToSelfDomain"`
				SocialDomain struct {
					FinancialProblems         interface{} `json:"financialProblems"`
					HomelessRoughSleepingNFA  interface{} `json:"homelessRoughSleepingNFA"`
					HousingAtRisk             interface{} `json:"housingAtRisk"`
					NoIdentifiedRisk          string      `json:"noIdentifiedRisk"`
					SociallyIsolatedNoSupport interface{} `json:"sociallyIsolatedNoSupport"`
				} `json:"socialDomain"`
				SubstanceMisuseDomain struct {
					ConfusionOrDisorientation interface{} `json:"ConfusionOrDisorientation"`
					AdmissionToAandE          interface{} `json:"admissionToAandE"`
					BlackoutOrSeizures        interface{} `json:"blackoutOrSeizures"`
					ConcurrentUse             interface{} `json:"concurrentUse"`
					HigherRiskDrinking        interface{} `json:"higherRiskDrinking"`
					InjectedByOthers          interface{} `json:"injectedByOthers"`
					Injecting                 string      `json:"injecting"`
					InjectingInNeckOrGroin    string      `json:"injectingInNeckOrGroin"`
					NoIdentifiedRisk          interface{} `json:"noIdentifiedRisk"`
					PolyDrugUse               string      `json:"polyDrugUse"`
					PreviousOverDose          interface{} `json:"previousOverDose"`
					RecentPrisonRelease       interface{} `json:"recentPrisonRelease"`
					ReducedTolerance          interface{} `json:"reducedTolerance"`
					SharingWorks              interface{} `json:"sharingWorks"`
					Speedballing              interface{} `json:"speedballing"`
					UsingOnTop                string      `json:"usingOnTop"`
				} `json:"substanceMisuseDomain"`
			} `json:"riskInformation"`
			SafeguardingInformation struct {
				LastReviewDate     interface{} `json:"lastReviewDate"`
				RiskHarmFromOthers string      `json:"riskHarmFromOthers"`
				RiskToAdults       interface{} `json:"riskToAdults"`
				RiskToChildrenOrYP interface{} `json:"riskToChildrenOrYP"`
				RiskToSelf         string      `json:"riskToSelf"`
			} `json:"safeguardingInformation"`
		} `json:"client"`
		KeyWorker struct {
			LocalIdentifier int `json:"localIdentifier"`
			Name            struct {
				Family string `json:"family"`
				Given  string `json:"given"`
			} `json:"name"`
			Telecom string `json:"telecom"`
		} `json:"keyWorker"`
	} `json:"data"`
}
type PIXmRequest struct {
	URL        string
	PID_OID    string
	PID        string
	Timeout    int64
	StatusCode int
	Response   []byte
}
type SOAPRequest struct {
	URL        string
	SOAPAction string
	Timeout    int64
	StatusCode int
	Body       []byte
	Response   []byte
}
type AWS_APIRequest struct {
	URL        string
	Act        string
	Resource   string
	Timeout    int64
	StatusCode int
	Body       []byte
	Response   []byte
}
type ClientRequest struct {
	HttpRequest  *http.Request
	ServerURL    string `json:"serverurl"`
	Act          string `json:"act"`
	User         string `json:"user"`
	Org          string `json:"org"`
	Orgoid       string `json:"orgoid"`
	Role         string `json:"role"`
	NHS          string `json:"nhs"`
	PID          string `json:"pid"`
	PIDOrg       string `json:"pidorg"`
	PIDOID       string `json:"pidoid"`
	FamilyName   string `json:"familyname"`
	GivenName    string `json:"givenname"`
	DOB          string `json:"dob"`
	Gender       string `json:"gender"`
	ZIP          string `json:"zip"`
	Status       string `json:"status"`
	XDWKey       string `json:"xdwkey"`
	ID           int    `json:"id"`
	Task         string `json:"task"`
	Pathway      string `json:"pathway"`
	Version      int    `json:"version"`
	ReturnFormat string `json:"returnformat"`
}
type TukHTTPInterface interface {
	newRequest() error
}

func NewRequest(i TukHTTPInterface) error {
	return i.newRequest()
}
func (i *ClientRequest) newRequest() error {
	req := i.HttpRequest
	req.ParseForm()
	i.Act = req.FormValue(cnst.ACT)
	i.User = req.FormValue("user")
	i.Org = req.FormValue("org")
	i.Orgoid = util.GetCodeSystemVal(req.FormValue("org"))
	i.Role = req.FormValue("role")
	i.NHS = req.FormValue("nhs")
	i.PID = req.FormValue("pid")
	i.PIDOrg = req.FormValue("pidorg")
	i.PIDOID = util.GetCodeSystemVal(req.FormValue("pidorg"))
	i.FamilyName = req.FormValue("familyname")
	i.GivenName = req.FormValue("givenname")
	i.DOB = req.FormValue("dob")
	i.Gender = req.FormValue("gender")
	i.ZIP = req.FormValue("zip")
	i.Status = req.FormValue("status")
	i.ID = util.GetIntFromString(req.FormValue("id"))
	i.Task = req.FormValue(cnst.TASK)
	i.Pathway = req.FormValue(cnst.PATHWAY)
	i.Version = util.GetIntFromString(req.FormValue("version"))
	i.XDWKey = req.FormValue("xdwkey")
	i.ReturnFormat = req.Header.Get(cnst.ACCEPT)
	if len(i.XDWKey) > 12 {
		i.Pathway, i.NHS = util.SplitXDWKey(i.XDWKey)
	}
	return nil
}
func (i *SOAPRequest) newRequest() error {
	if i.Timeout == 0 {
		i.Timeout = 5
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.Timeout)*time.Second)
	defer cancel()
	req, err := http.NewRequest(http.MethodPost, i.URL, strings.NewReader(string(i.Body)))
	if err != nil {
		return err
	}
	if i.SOAPAction != "" {
		req.Header.Set(cnst.SOAP_ACTION, i.SOAPAction)
	}
	req.Header.Set(cnst.CONTENT_TYPE, cnst.SOAP_XML)
	req.Header.Set(cnst.ACCEPT, cnst.ALL)
	req.Header.Set(cnst.CONNECTION, cnst.KEEP_ALIVE)
	i.logRequest(req.Header)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	i.StatusCode = resp.StatusCode
	i.Response, err = io.ReadAll(resp.Body)
	i.logResponse()
	return err
}
func (i *PIXmRequest) newRequest() error {
	if i.Timeout == 0 {
		i.Timeout = 5
	}
	i.URL = i.URL + "?identifier=" + i.PID_OID + "%7C" + i.PID + cnst.FORMAT_JSON_PRETTY
	req, _ := http.NewRequest(cnst.HTTP_GET, i.URL, nil)
	req.Header.Set(cnst.CONTENT_TYPE, cnst.APPLICATION_JSON)
	req.Header.Set(cnst.ACCEPT, cnst.ALL)
	req.Header.Set(cnst.CONNECTION, cnst.KEEP_ALIVE)
	i.logRequest(req.Header)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.Timeout)*time.Second)
	defer cancel()
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	i.StatusCode = resp.StatusCode
	i.Response, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	i.logResponse()
	return err
}
func (i *CGLRequest) newRequest() error {
	req, _ := http.NewRequest(cnst.HTTP_GET, i.Request, nil)
	req.Header.Set(cnst.ACCEPT, cnst.APPLICATION_JSON)
	req.Header.Set("X-API-KEY", i.X_Api_Key)
	i.logRequest(req.Header)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	i.StatusCode = resp.StatusCode
	i.Response, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(i.Response, &i.CGL_User)
	i.logResponse()
	return err
}
func (i *AWS_APIRequest) newRequest() error {
	if i.Timeout == 0 {
		i.Timeout = 5
	}
	var err error
	var req *http.Request
	var resp *http.Response
	client := &http.Client{}
	if req, err = http.NewRequest(strings.ToUpper(i.Act), i.URL+i.Resource, bytes.NewBuffer(i.Body)); err == nil {
		req.Header.Add(cnst.CONTENT_TYPE, cnst.APPLICATION_JSON_CHARSET_UTF_8)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i.Timeout)*time.Second)
		defer cancel()
		i.logRequest(req.Header)
		if resp, err = client.Do(req.WithContext(ctx)); err == nil {
			log.Printf("Response Status Code %v\n", resp.StatusCode)
			if resp.StatusCode == http.StatusOK {
				i.Response, err = io.ReadAll(resp.Body)
			}
		}
	}
	defer resp.Body.Close()
	i.StatusCode = resp.StatusCode
	i.logResponse()
	return err
}
func (i *AWS_APIRequest) logRequest(headers http.Header) {
	log.Println("HTTP " + strings.ToUpper(i.Act) + " Request Headers")
	util.Log(headers)
	log.Printf("HTTP Request\nURL = %s\nTimeout = %v\nMessage body\n%s", i.URL, i.Timeout, string(i.Body))
}
func (i *AWS_APIRequest) logResponse() {
	log.Printf("HTML Response - Status Code = %v\n%s", i.StatusCode, string(i.Response))
}
func (i *SOAPRequest) logRequest(headers http.Header) {
	log.Println("SOAP Request Headers")
	util.Log(headers)
	log.Printf("SOAP Request\nURL = %s\nAction = %s\nTimeout = %v\n\n%s", i.URL, i.SOAPAction, i.Timeout, string(i.Body))
}
func (i *SOAPRequest) logResponse() {
	log.Printf("SOAP Response - Status Code = %v\n%s", i.StatusCode, string(i.Response))
}
func (i *PIXmRequest) logRequest(headers http.Header) {
	log.Println("HTTP GET Request Headers")
	util.Log(headers)
	log.Printf("HTTP Request\nURL = %s\nTimeout = %v", i.URL, i.Timeout)
}
func (i *CGLRequest) logRequest(headers http.Header) {
	log.Println("HTTP GET Request Headers")
	util.Log(headers)
	log.Printf("HTTP Request\nURL = %s\nTimeout = %v", i.Request, 5)
}
func (i *PIXmRequest) logResponse() {
	log.Printf("HTML Response - Status Code = %v\n%s", i.StatusCode, string(i.Response))
}
func (i *CGLRequest) logResponse() {
	log.Printf("HTML Response - Status Code = %v\n%s", i.StatusCode, string(i.Response))
}
