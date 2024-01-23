// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Data -url /api/v5/trade/orders-pending -type GetOpenOrdersRequest -responseDataType []OpenOrder"; DO NOT EDIT.

package okexapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"time"
)

func (g *GetOpenOrdersRequest) InstrumentType(instrumentType InstrumentType) *GetOpenOrdersRequest {
	g.instrumentType = instrumentType
	return g
}

func (g *GetOpenOrdersRequest) InstrumentID(instrumentID string) *GetOpenOrdersRequest {
	g.instrumentID = &instrumentID
	return g
}

func (g *GetOpenOrdersRequest) OrderType(orderType OrderType) *GetOpenOrdersRequest {
	g.orderType = &orderType
	return g
}

func (g *GetOpenOrdersRequest) State(state OrderState) *GetOpenOrdersRequest {
	g.state = &state
	return g
}

func (g *GetOpenOrdersRequest) Category(category string) *GetOpenOrdersRequest {
	g.category = &category
	return g
}

func (g *GetOpenOrdersRequest) After(after string) *GetOpenOrdersRequest {
	g.after = &after
	return g
}

func (g *GetOpenOrdersRequest) Before(before string) *GetOpenOrdersRequest {
	g.before = &before
	return g
}

func (g *GetOpenOrdersRequest) Begin(begin time.Time) *GetOpenOrdersRequest {
	g.begin = &begin
	return g
}

func (g *GetOpenOrdersRequest) End(end time.Time) *GetOpenOrdersRequest {
	g.end = &end
	return g
}

func (g *GetOpenOrdersRequest) Limit(limit string) *GetOpenOrdersRequest {
	g.limit = &limit
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetOpenOrdersRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}
	// check instrumentType field -> json key instType
	instrumentType := g.instrumentType

	// TEMPLATE check-valid-values
	switch instrumentType {
	case InstrumentTypeSpot, InstrumentTypeSwap, InstrumentTypeFutures, InstrumentTypeOption, InstrumentTypeMARGIN:
		params["instType"] = instrumentType

	default:
		return nil, fmt.Errorf("instType value %v is invalid", instrumentType)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of instrumentType
	params["instType"] = instrumentType
	// check instrumentID field -> json key instId
	if g.instrumentID != nil {
		instrumentID := *g.instrumentID

		// assign parameter of instrumentID
		params["instId"] = instrumentID
	} else {
	}
	// check orderType field -> json key ordType
	if g.orderType != nil {
		orderType := *g.orderType

		// TEMPLATE check-valid-values
		switch orderType {
		case OrderTypeMarket, OrderTypeLimit, OrderTypePostOnly, OrderTypeFOK, OrderTypeIOC:
			params["ordType"] = orderType

		default:
			return nil, fmt.Errorf("ordType value %v is invalid", orderType)

		}
		// END TEMPLATE check-valid-values

		// assign parameter of orderType
		params["ordType"] = orderType
	} else {
	}
	// check state field -> json key state
	if g.state != nil {
		state := *g.state

		// TEMPLATE check-valid-values
		switch state {
		case OrderStateCanceled, OrderStateLive, OrderStatePartiallyFilled, OrderStateFilled:
			params["state"] = state

		default:
			return nil, fmt.Errorf("state value %v is invalid", state)

		}
		// END TEMPLATE check-valid-values

		// assign parameter of state
		params["state"] = state
	} else {
	}
	// check category field -> json key category
	if g.category != nil {
		category := *g.category

		// assign parameter of category
		params["category"] = category
	} else {
	}
	// check after field -> json key after
	if g.after != nil {
		after := *g.after

		// assign parameter of after
		params["after"] = after
	} else {
	}
	// check before field -> json key before
	if g.before != nil {
		before := *g.before

		// assign parameter of before
		params["before"] = before
	} else {
	}
	// check begin field -> json key begin
	if g.begin != nil {
		begin := *g.begin

		// assign parameter of begin
		params["begin"] = begin
	} else {
	}
	// check end field -> json key end
	if g.end != nil {
		end := *g.end

		// assign parameter of end
		params["end"] = end
	} else {
	}
	// check limit field -> json key limit
	if g.limit != nil {
		limit := *g.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetOpenOrdersRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetOpenOrdersRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetOpenOrdersRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetOpenOrdersRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetOpenOrdersRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetOpenOrdersRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetOpenOrdersRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetOpenOrdersRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

// GetPath returns the request path of the API
func (g *GetOpenOrdersRequest) GetPath() string {
	return "/api/v5/trade/orders-pending"
}

// Do generates the request object and send the request object to the API endpoint
func (g *GetOpenOrdersRequest) Do(ctx context.Context) ([]OpenOrder, error) {

	// no body params
	var params interface{}
	query, err := g.GetQueryParameters()
	if err != nil {
		return nil, err
	}

	var apiURL string

	apiURL = g.GetPath()

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}

	type responseValidator interface {
		Validate() error
	}
	validator, ok := interface{}(apiResponse).(responseValidator)
	if ok {
		if err := validator.Validate(); err != nil {
			return nil, err
		}
	}
	var data []OpenOrder
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return data, nil
}
