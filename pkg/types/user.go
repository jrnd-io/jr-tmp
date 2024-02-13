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

type User struct {
	Guid string `json:"guid"`

	IsActive bool `json:"isActive"`

	Balance string `json:"balance"`

	Age int32 `json:"age"`

	EyeColor string `json:"eyeColor"`

	Name string `json:"name"`

	Company string `json:"company"`

	Work_email string `json:"work_email"`

	Email string `json:"email"`

	About string `json:"about"`

	Latitude float32 `json:"latitude"`

	Longitude float32 `json:"longitude"`
}

const UserAvroCRC64Fingerprint = "\xc14\xf27I\xfa\x83\x11"

func NewUser() User {
	r := User{}
	return r
}

func DeserializeUser(r io.Reader) (User, error) {
	t := NewUser()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeUserFromSchema(r io.Reader, schema string) (User, error) {
	t := NewUser()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeUser(r User, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Guid, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.IsActive, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Balance, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Age, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EyeColor, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Company, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Work_email, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Email, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.About, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Latitude, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.Longitude, w)
	if err != nil {
		return err
	}
	return err
}

func (r User) Serialize(w io.Writer) error {
	return writeUser(r, w)
}

func (r User) Schema() string {
	return "{\"fields\":[{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"isActive\",\"type\":\"boolean\"},{\"name\":\"balance\",\"type\":\"string\"},{\"name\":\"age\",\"type\":\"int\"},{\"name\":\"eyeColor\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"company\",\"type\":\"string\"},{\"name\":\"work_email\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"about\",\"type\":\"string\"},{\"name\":\"latitude\",\"type\":\"float\"},{\"name\":\"longitude\",\"type\":\"float\"}],\"name\":\"jr.User\",\"type\":\"record\"}"
}

func (r User) SchemaName() string {
	return "jr.User"
}

func (_ User) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ User) SetInt(v int32)       { panic("Unsupported operation") }
func (_ User) SetLong(v int64)      { panic("Unsupported operation") }
func (_ User) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ User) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ User) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ User) SetString(v string)   { panic("Unsupported operation") }
func (_ User) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *User) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Guid}

		return w

	case 1:
		w := types.Boolean{Target: &r.IsActive}

		return w

	case 2:
		w := types.String{Target: &r.Balance}

		return w

	case 3:
		w := types.Int{Target: &r.Age}

		return w

	case 4:
		w := types.String{Target: &r.EyeColor}

		return w

	case 5:
		w := types.String{Target: &r.Name}

		return w

	case 6:
		w := types.String{Target: &r.Company}

		return w

	case 7:
		w := types.String{Target: &r.Work_email}

		return w

	case 8:
		w := types.String{Target: &r.Email}

		return w

	case 9:
		w := types.String{Target: &r.About}

		return w

	case 10:
		w := types.Float{Target: &r.Latitude}

		return w

	case 11:
		w := types.Float{Target: &r.Longitude}

		return w

	}
	panic("Unknown field index")
}

func (r *User) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *User) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ User) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ User) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ User) HintSize(int)                     { panic("Unsupported operation") }
func (_ User) Finalize()                        {}

func (_ User) AvroCRC64Fingerprint() []byte {
	return []byte(UserAvroCRC64Fingerprint)
}

func (r User) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["guid"], err = json.Marshal(r.Guid)
	if err != nil {
		return nil, err
	}
	output["isActive"], err = json.Marshal(r.IsActive)
	if err != nil {
		return nil, err
	}
	output["balance"], err = json.Marshal(r.Balance)
	if err != nil {
		return nil, err
	}
	output["age"], err = json.Marshal(r.Age)
	if err != nil {
		return nil, err
	}
	output["eyeColor"], err = json.Marshal(r.EyeColor)
	if err != nil {
		return nil, err
	}
	output["name"], err = json.Marshal(r.Name)
	if err != nil {
		return nil, err
	}
	output["company"], err = json.Marshal(r.Company)
	if err != nil {
		return nil, err
	}
	output["work_email"], err = json.Marshal(r.Work_email)
	if err != nil {
		return nil, err
	}
	output["email"], err = json.Marshal(r.Email)
	if err != nil {
		return nil, err
	}
	output["about"], err = json.Marshal(r.About)
	if err != nil {
		return nil, err
	}
	output["latitude"], err = json.Marshal(r.Latitude)
	if err != nil {
		return nil, err
	}
	output["longitude"], err = json.Marshal(r.Longitude)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *User) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["guid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Guid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for guid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["isActive"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IsActive); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for isActive")
	}
	val = func() json.RawMessage {
		if v, ok := fields["balance"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Balance); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for balance")
	}
	val = func() json.RawMessage {
		if v, ok := fields["age"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Age); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for age")
	}
	val = func() json.RawMessage {
		if v, ok := fields["eyeColor"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EyeColor); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for eyeColor")
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
		if v, ok := fields["company"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Company); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for company")
	}
	val = func() json.RawMessage {
		if v, ok := fields["work_email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Work_email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for work_email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["email"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Email); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for email")
	}
	val = func() json.RawMessage {
		if v, ok := fields["about"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.About); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for about")
	}
	val = func() json.RawMessage {
		if v, ok := fields["latitude"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Latitude); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for latitude")
	}
	val = func() json.RawMessage {
		if v, ok := fields["longitude"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Longitude); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for longitude")
	}
	return nil
}
