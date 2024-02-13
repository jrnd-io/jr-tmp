// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaign_finance.avsc
 *     click_stream.avsc
 *     click_stream_users.avsc
 *     clickstreamcodes.avsc
 *     credit_cards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizzaorders.avsc
 *     pizzaorderscancelled.avsc
 *     pizzaorderscompleted.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siemlogs.avsc
 *     stock_trades.avsc
 *     stores.avsc
 *     sysloglogs.avsc
 *     transactions.avsc
 *     user.avsc
 *     users.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type InsuranceOffers struct {
	Offer_id int32 `json:"offer_id"`

	Offer_name string `json:"offer_name"`

	Offer_url string `json:"offer_url"`
}

const InsuranceOffersAvroCRC64Fingerprint = ":\x18\xd3Ȼ֒x"

func NewInsuranceOffers() InsuranceOffers {
	r := InsuranceOffers{}
	return r
}

func DeserializeInsuranceOffers(r io.Reader) (InsuranceOffers, error) {
	t := NewInsuranceOffers()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInsuranceOffersFromSchema(r io.Reader, schema string) (InsuranceOffers, error) {
	t := NewInsuranceOffers()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInsuranceOffers(r InsuranceOffers, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Offer_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Offer_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Offer_url, w)
	if err != nil {
		return err
	}
	return err
}

func (r InsuranceOffers) Serialize(w io.Writer) error {
	return writeInsuranceOffers(r, w)
}

func (r InsuranceOffers) Schema() string {
	return "{\"arg.properties\":{\"options\":[{\"offer_id\":1,\"offer_name\":\"new_home_policy\",\"offer_url\":\"http://privacy.gov.au/in/faucibus/orci/luctus.js\"},{\"offer_id\":2,\"offer_name\":\"new_auto_policy\",\"offer_url\":\"https://reddit.com/id.html\"},{\"offer_id\":3,\"offer_name\":\"discount_existing\",\"offer_url\":\"https://cbc.ca/dui/proin/leo/odio/porttitor/id.aspx\"},{\"offer_id\":4,\"offer_name\":\"safe_driver_program\",\"offer_url\":\"http://seattletimes.com/ante/vel/ipsum/praesent.json\"},{\"offer_id\":5,\"offer_name\":\"no_offer\",\"offer_url\":\"https://sciencedaily.com/ante.aspx\"}]},\"fields\":[{\"name\":\"offer_id\",\"type\":\"int\"},{\"name\":\"offer_name\",\"type\":\"string\"},{\"name\":\"offer_url\",\"type\":\"string\"}],\"name\":\"insurance.InsuranceOffers\",\"type\":\"record\"}"
}

func (r InsuranceOffers) SchemaName() string {
	return "insurance.InsuranceOffers"
}

func (_ InsuranceOffers) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ InsuranceOffers) SetInt(v int32)       { panic("Unsupported operation") }
func (_ InsuranceOffers) SetLong(v int64)      { panic("Unsupported operation") }
func (_ InsuranceOffers) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ InsuranceOffers) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ InsuranceOffers) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ InsuranceOffers) SetString(v string)   { panic("Unsupported operation") }
func (_ InsuranceOffers) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *InsuranceOffers) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Offer_id}

		return w

	case 1:
		w := types.String{Target: &r.Offer_name}

		return w

	case 2:
		w := types.String{Target: &r.Offer_url}

		return w

	}
	panic("Unknown field index")
}

func (r *InsuranceOffers) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *InsuranceOffers) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ InsuranceOffers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ InsuranceOffers) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ InsuranceOffers) HintSize(int)                     { panic("Unsupported operation") }
func (_ InsuranceOffers) Finalize()                        {}

func (_ InsuranceOffers) AvroCRC64Fingerprint() []byte {
	return []byte(InsuranceOffersAvroCRC64Fingerprint)
}

func (r InsuranceOffers) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["offer_id"], err = json.Marshal(r.Offer_id)
	if err != nil {
		return nil, err
	}
	output["offer_name"], err = json.Marshal(r.Offer_name)
	if err != nil {
		return nil, err
	}
	output["offer_url"], err = json.Marshal(r.Offer_url)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *InsuranceOffers) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["offer_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Offer_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for offer_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["offer_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Offer_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for offer_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["offer_url"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Offer_url); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for offer_url")
	}
	return nil
}
