// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaignfinance.avsc
 *     clickstream.avsc
 *     clickstreamcodes.avsc
 *     clickstreamusers.avsc
 *     creditcards.avsc
 *     deviceinformation.avsc
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
 *     pageviews.avsc
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
 *     stockTrades.avsc
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

type Fleetmgmtsensors struct {
	Vehicle_id int32 `json:"vehicle_id"`

	Engine_temperature int32 `json:"engine_temperature"`

	Average_rpm int32 `json:"average_rpm"`
}

const FleetmgmtsensorsAvroCRC64Fingerprint = "\x9a\x01A\xa3'\x92\x1f\x8f"

func NewFleetmgmtsensors() Fleetmgmtsensors {
	r := Fleetmgmtsensors{}
	return r
}

func DeserializeFleetmgmtsensors(r io.Reader) (Fleetmgmtsensors, error) {
	t := NewFleetmgmtsensors()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFleetmgmtsensorsFromSchema(r io.Reader, schema string) (Fleetmgmtsensors, error) {
	t := NewFleetmgmtsensors()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFleetmgmtsensors(r Fleetmgmtsensors, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Vehicle_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Engine_temperature, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Average_rpm, w)
	if err != nil {
		return err
	}
	return err
}

func (r Fleetmgmtsensors) Serialize(w io.Writer) error {
	return writeFleetmgmtsensors(r, w)
}

func (r Fleetmgmtsensors) Schema() string {
	return "{\"fields\":[{\"name\":\"vehicle_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":9999,\"min\":1000}},\"type\":\"int\"}},{\"name\":\"engine_temperature\",\"type\":{\"arg.properties\":{\"range\":{\"max\":250,\"min\":150}},\"type\":\"int\"}},{\"name\":\"average_rpm\",\"type\":{\"arg.properties\":{\"range\":{\"max\":5000,\"min\":1800}},\"type\":\"int\"}}],\"name\":\"fleet_mgmt.fleetmgmtsensors\",\"type\":\"record\"}"
}

func (r Fleetmgmtsensors) SchemaName() string {
	return "fleet_mgmt.fleetmgmtsensors"
}

func (_ Fleetmgmtsensors) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) SetString(v string)   { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Fleetmgmtsensors) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Vehicle_id}

		return w

	case 1:
		w := types.Int{Target: &r.Engine_temperature}

		return w

	case 2:
		w := types.Int{Target: &r.Average_rpm}

		return w

	}
	panic("Unknown field index")
}

func (r *Fleetmgmtsensors) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Fleetmgmtsensors) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Fleetmgmtsensors) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) HintSize(int)                     { panic("Unsupported operation") }
func (_ Fleetmgmtsensors) Finalize()                        {}

func (_ Fleetmgmtsensors) AvroCRC64Fingerprint() []byte {
	return []byte(FleetmgmtsensorsAvroCRC64Fingerprint)
}

func (r Fleetmgmtsensors) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["vehicle_id"], err = json.Marshal(r.Vehicle_id)
	if err != nil {
		return nil, err
	}
	output["engine_temperature"], err = json.Marshal(r.Engine_temperature)
	if err != nil {
		return nil, err
	}
	output["average_rpm"], err = json.Marshal(r.Average_rpm)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Fleetmgmtsensors) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["vehicle_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Vehicle_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vehicle_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["engine_temperature"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Engine_temperature); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for engine_temperature")
	}
	val = func() json.RawMessage {
		if v, ok := fields["average_rpm"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Average_rpm); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for average_rpm")
	}
	return nil
}
