// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaign_finance.avsc
 *     clickstream.avsc
 *     clickstream_codes.avsc
 *     clickstream_users.avsc
 *     creditcards.avsc
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
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siemlogs.avsc
 *     stock_trade.avsc
 *     stores.avsc
 *     syslog_logs.avsc
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

type Shoe struct {
	Id string `json:"id"`

	Sale_price string `json:"sale_price"`

	Brand string `json:"brand"`

	Name string `json:"name"`

	Rating float32 `json:"rating"`
}

const ShoeAvroCRC64Fingerprint = "x\x13y\xfe\xab\xa5\x82\xf9"

func NewShoe() Shoe {
	r := Shoe{}
	return r
}

func DeserializeShoe(r io.Reader) (Shoe, error) {
	t := NewShoe()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeShoeFromSchema(r io.Reader, schema string) (Shoe, error) {
	t := NewShoe()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeShoe(r Shoe, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Sale_price, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Brand, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Rating, w)
	if err != nil {
		return err
	}
	return err
}

func (r Shoe) Serialize(w io.Writer) error {
	return writeShoe(r, w)
}

func (r Shoe) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"sale_price\",\"type\":\"string\"},{\"name\":\"brand\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"rating\",\"type\":\"float\"}],\"name\":\"shoes.Shoe\",\"type\":\"record\"}"
}

func (r Shoe) SchemaName() string {
	return "shoes.Shoe"
}

func (_ Shoe) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Shoe) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Shoe) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Shoe) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Shoe) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Shoe) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Shoe) SetString(v string)   { panic("Unsupported operation") }
func (_ Shoe) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Shoe) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Sale_price}

		return w

	case 2:
		w := types.String{Target: &r.Brand}

		return w

	case 3:
		w := types.String{Target: &r.Name}

		return w

	case 4:
		w := types.Float{Target: &r.Rating}

		return w

	}
	panic("Unknown field index")
}

func (r *Shoe) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Shoe) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Shoe) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Shoe) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Shoe) HintSize(int)                     { panic("Unsupported operation") }
func (_ Shoe) Finalize()                        {}

func (_ Shoe) AvroCRC64Fingerprint() []byte {
	return []byte(ShoeAvroCRC64Fingerprint)
}

func (r Shoe) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["sale_price"], err = json.Marshal(r.Sale_price)
	if err != nil {
		return nil, err
	}
	output["brand"], err = json.Marshal(r.Brand)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["rating"], err = json.Marshal(r.Rating)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Shoe) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sale_price"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sale_price); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sale_price")
	}
	val = func() json.RawMessage {
		if v, ok := fields["brand"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Brand); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for brand")
	}
	val = func() json.RawMessage {
		if v, ok := fields["name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["rating"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Rating); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rating")
	}
	return nil
}
