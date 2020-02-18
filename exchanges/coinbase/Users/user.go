package Users

import (
	"encoding/json"
	"fmt"
	"github.com/ivandonyk/Crypto-Trader/exchanges/coinbase"
	"github.com/ivandonyk/Crypto-Trader/lib"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID                                    string        `json:"id"`
	Name                                  string        `json:"name"`
	UserName                              string        `json:"username"`
	ProfileLocation                       string        `json:"profile_location"`
	ProfileBio                            string        `json:"profile_bio"`
	ProfileURL                            string        `json:"profile_url"`
	AvatarURL                             string        `json:"avatar_url"`
	Resource                              string        `json:"resource"`
	ResourcePath                          string        `json:"resource_path"`
	Email                                 string        `json:"email"`
	LegacyID                              string        `json:"legacy_id"`
	TimeZone                              string        `json:"time_zone"`
	NativeCurrency                        string        `json:"native_currency"`
	BitcoinUnit                           string        `json:"bitcoin_unit"`
	State                                 string        `json:"state"`
	Country                               County        `json:"country"`
	RegionSupportFiatTransfers            bool          `json:"region_support_fiat_transfers"`
	RegionSupportsCryptoToCryptoTransfers bool          `json:"region_supports_crypto_to_crypto_transfers"`
	CreatedAt                             string        `json:"created_at"`
	SupportRewards                        bool          `json:"support_rewards"`
	Tiers                                 Tiers         `json:"tiers"`
	ReferralMoney                         ReferralMoney `json:"referral_money"`
	SecondFactor                          SecondFactor  `json:"second_factor"`
	HasBlockingBuyRestrictions            bool          `json:"has_blocking_buy_restrictions"`
	HasMadeAPurchase                      bool          `json:"has_made_a_purchase"`
	HasBuyDepositPaymentMethods           bool          `json:"has_buy_deposit_payment_methods"`
	HasUnverifiedBuyPaymentMethods        bool          `json:"has_unverified_buy_payment_methods"`
	Warnings                              []*Warning    `json:"warnings,omitempty"`
}

type County struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	IsInEurope bool   `json:"is_in_europe"`
}

type ReferralMoney struct {
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	CurrencySymbol    string `json:"currency_symbol"`
	ReferralThreshold string `json:"referral_threshold"`
}

type SecondFactor struct {
	Method string `json:"method"`
	Totp   Totp   `json:"totp"`
	Sms    SMS    `json:"sms"`
	Authy  Authy  `json:"authy"`
	U2F    U2F    `json:"u2f"`
}

type Totp struct {
	Digits int `json:"digits"`
}

type SMS struct {
	Digits int `json:"digits"`
}

type Authy struct {
	MinDigits int `json:"minDigits"`
	MaxDigits int `json:"maxDigits"`
}

type U2F struct {
}

type Tiers struct {
	CompletedDescription string `json:"completed_description"`
	UpgradeButtonText    string `json:"upgrade_button_text"`
	Header               string `json:"header"`
	Body                 string `json:"body"`
}
type Warning struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

type CurrentUserResponse struct {
	Data User `json:"data"`
}

func ShowCurrentUser(baseURL, endpoint, apiKey, secretKey string) (*CurrentUserResponse, error) {
	var currUserResp *CurrentUserResponse

	reqUrl := lib.URLJoin(baseURL, endpoint)

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("could not build new request due to %v", err)
	}

	epoch, err := coinbase.GetTimestamp(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error getting timestamp %v", err)
	}

	sig := lib.CoinbaseAuthorize(strconv.FormatInt(epoch, 10), secretKey, http.MethodGet, endpoint)
	req.Header.Set("content-type", "application/json")
	req.Header.Add("CB-ACCESS-SIGN", sig)
	req.Header.Add("CB-ACCESS-TIMESTAMP", strconv.FormatInt(epoch, 10))
	req.Header.Add("CB-ACCESS-KEY", apiKey)

	fmt.Println(reqUrl)
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not create new response type %v", err)
	}

	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body %v", err)
	}

	if err := json.Unmarshal(f, &currUserResp); err != nil {
		return nil, fmt.Errorf("error serliazing user due to %v", err)
	}

	defer resp.Body.Close()

	return currUserResp, nil
}
