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

type CreditCards struct {
	Card_id int32 `json:"card_id"`

	Card_number int64 `json:"card_number"`

	Cvv int32 `json:"cvv"`

	Expiration_date string `json:"expiration_date"`
}

const CreditCardsAvroCRC64Fingerprint = "\xc9\xf2i\xa1(\"=\xca"

func NewCreditCards() CreditCards {
	r := CreditCards{}
	return r
}

func DeserializeCreditCards(r io.Reader) (CreditCards, error) {
	t := NewCreditCards()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCreditCardsFromSchema(r io.Reader, schema string) (CreditCards, error) {
	t := NewCreditCards()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCreditCards(r CreditCards, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Card_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Card_number, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Cvv, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Expiration_date, w)
	if err != nil {
		return err
	}
	return err
}

func (r CreditCards) Serialize(w io.Writer) error {
	return writeCreditCards(r, w)
}

func (r CreditCards) Schema() string {
	return "{\"arg.properties\":{\"options\":[{\"card_id\":1,\"card_number\":6011601160116611,\"cvv\":832,\"expiration_date\":\"12/27\"},{\"card_id\":2,\"card_number\":6011111111111117,\"cvv\":833,\"expiration_date\":\"12/26\"},{\"card_id\":3,\"card_number\":6011000990139424,\"cvv\":348,\"expiration_date\":\"12/27\"},{\"card_id\":4,\"card_number\":3530111333300000,\"cvv\":288,\"expiration_date\":\"11/27\"},{\"card_id\":5,\"card_number\":3566002020360505,\"cvv\":558,\"expiration_date\":\"11/27\"},{\"card_id\":6,\"card_number\":5431111111111111,\"cvv\":383,\"expiration_date\":\"11/26\"},{\"card_id\":7,\"card_number\":5555555555554444,\"cvv\":273,\"expiration_date\":\"10/27\"},{\"card_id\":8,\"card_number\":5105105105105100,\"cvv\":663,\"expiration_date\":\"10/27\"},{\"card_id\":9,\"card_number\":4111111111111111,\"cvv\":822,\"expiration_date\":\"10/26\"},{\"card_id\":10,\"card_number\":4242424242424242,\"cvv\":190,\"expiration_date\":\"09/27\"},{\"card_id\":11,\"card_number\":4000056655665556,\"cvv\":729,\"expiration_date\":\"09/26\"},{\"card_id\":12,\"card_number\":5555555555554444,\"cvv\":289,\"expiration_date\":\"08/27\"},{\"card_id\":13,\"card_number\":2223003122003222,\"cvv\":115,\"expiration_date\":\"08/26\"},{\"card_id\":14,\"card_number\":5200828282828210,\"cvv\":927,\"expiration_date\":\"08/27\"},{\"card_id\":15,\"card_number\":5105105105105100,\"cvv\":329,\"expiration_date\":\"07/27\"},{\"card_id\":16,\"card_number\":6200000000000005,\"cvv\":727,\"expiration_date\":\"06/27\"},{\"card_id\":17,\"card_number\":3056930009020004,\"cvv\":563,\"expiration_date\":\"05/27\"},{\"card_id\":18,\"card_number\":4000000760000002,\"cvv\":922,\"expiration_date\":\"04/27\"},{\"card_id\":19,\"card_number\":4000001240000000,\"cvv\":293,\"expiration_date\":\"03/27\"},{\"card_id\":20,\"card_number\":4000004840008001,\"cvv\":779,\"expiration_date\":\"02/27\"},{\"card_id\":21,\"card_number\":4000002500003155,\"cvv\":443,\"expiration_date\":\"02/26\"},{\"card_id\":22,\"card_number\":4000002760003184,\"cvv\":487,\"expiration_date\":\"06/27\"},{\"card_id\":23,\"card_number\":4000008260003178,\"cvv\":663,\"expiration_date\":\"01/26\"},{\"card_id\":24,\"card_number\":4000008260003179,\"cvv\":723,\"expiration_date\":\"01/26\"}]},\"fields\":[{\"name\":\"card_id\",\"type\":\"int\"},{\"name\":\"card_number\",\"type\":\"long\"},{\"name\":\"cvv\",\"type\":\"int\"},{\"name\":\"expiration_date\",\"type\":\"string\"}],\"name\":\"datagen.example.CreditCards\",\"type\":\"record\"}"
}

func (r CreditCards) SchemaName() string {
	return "datagen.example.CreditCards"
}

func (_ CreditCards) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CreditCards) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CreditCards) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CreditCards) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CreditCards) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CreditCards) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CreditCards) SetString(v string)   { panic("Unsupported operation") }
func (_ CreditCards) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CreditCards) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Card_id}

		return w

	case 1:
		w := types.Long{Target: &r.Card_number}

		return w

	case 2:
		w := types.Int{Target: &r.Cvv}

		return w

	case 3:
		w := types.String{Target: &r.Expiration_date}

		return w

	}
	panic("Unknown field index")
}

func (r *CreditCards) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CreditCards) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CreditCards) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CreditCards) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CreditCards) HintSize(int)                     { panic("Unsupported operation") }
func (_ CreditCards) Finalize()                        {}

func (_ CreditCards) AvroCRC64Fingerprint() []byte {
	return []byte(CreditCardsAvroCRC64Fingerprint)
}

func (r CreditCards) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["card_id"], err = json.Marshal(r.Card_id)
	if err != nil {
		return nil, err
	}
	output["card_number"], err = json.Marshal(r.Card_number)
	if err != nil {
		return nil, err
	}
	output["cvv"], err = json.Marshal(r.Cvv)
	if err != nil {
		return nil, err
	}
	output["expiration_date"], err = json.Marshal(r.Expiration_date)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CreditCards) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["card_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Card_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for card_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["card_number"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Card_number); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for card_number")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cvv"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cvv); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cvv")
	}
	val = func() json.RawMessage {
		if v, ok := fields["expiration_date"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Expiration_date); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for expiration_date")
	}
	return nil
}
