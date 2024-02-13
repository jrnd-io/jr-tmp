// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaignfinance.avsc
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
 *     insurancecustomeractivity.avsc
 *     insurancecustomers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payrollbonus.avsc
 *     payrollemployee.avsc
 *     payrollemployeelocation.avsc
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

type ShoeOrder struct {
	Order_id int32 `json:"order_id"`

	Product_id string `json:"product_id"`

	Customer_id string `json:"customer_id"`

	Ts int64 `json:"ts"`
}

const ShoeOrderAvroCRC64Fingerprint = "\xf7\xa3\xdeP#v4\x1d"

func NewShoeOrder() ShoeOrder {
	r := ShoeOrder{}
	return r
}

func DeserializeShoeOrder(r io.Reader) (ShoeOrder, error) {
	t := NewShoeOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeShoeOrderFromSchema(r io.Reader, schema string) (ShoeOrder, error) {
	t := NewShoeOrder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeShoeOrder(r ShoeOrder, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Order_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Product_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Customer_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	return err
}

func (r ShoeOrder) Serialize(w io.Writer) error {
	return writeShoeOrder(r, w)
}

func (r ShoeOrder) Schema() string {
	return "{\"fields\":[{\"name\":\"order_id\",\"type\":\"int\"},{\"name\":\"product_id\",\"type\":\"string\"},{\"name\":\"customer_id\",\"type\":\"string\"},{\"name\":\"ts\",\"type\":\"long\"}],\"name\":\"shoes.ShoeOrder\",\"type\":\"record\"}"
}

func (r ShoeOrder) SchemaName() string {
	return "shoes.ShoeOrder"
}

func (_ ShoeOrder) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ShoeOrder) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ShoeOrder) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ShoeOrder) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ShoeOrder) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ShoeOrder) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ShoeOrder) SetString(v string)   { panic("Unsupported operation") }
func (_ ShoeOrder) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ShoeOrder) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Order_id}

		return w

	case 1:
		w := types.String{Target: &r.Product_id}

		return w

	case 2:
		w := types.String{Target: &r.Customer_id}

		return w

	case 3:
		w := types.Long{Target: &r.Ts}

		return w

	}
	panic("Unknown field index")
}

func (r *ShoeOrder) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ShoeOrder) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ShoeOrder) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ShoeOrder) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ShoeOrder) HintSize(int)                     { panic("Unsupported operation") }
func (_ ShoeOrder) Finalize()                        {}

func (_ ShoeOrder) AvroCRC64Fingerprint() []byte {
	return []byte(ShoeOrderAvroCRC64Fingerprint)
}

func (r ShoeOrder) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["order_id"], err = json.Marshal(r.Order_id)
	if err != nil {
		return nil, err
	}
	output["product_id"], err = json.Marshal(r.Product_id)
	if err != nil {
		return nil, err
	}
	output["customer_id"], err = json.Marshal(r.Customer_id)
	if err != nil {
		return nil, err
	}
	output["ts"], err = json.Marshal(r.Ts)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ShoeOrder) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["order_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Order_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for order_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["product_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Product_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for product_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["customer_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Customer_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for customer_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ts"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ts); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ts")
	}
	return nil
}
