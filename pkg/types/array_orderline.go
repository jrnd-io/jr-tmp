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
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayOrderline(r []Orderline, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeOrderline(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayOrderlineWrapper struct {
	Target *[]Orderline
}

func (_ ArrayOrderlineWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayOrderlineWrapper) Finalize()                        {}
func (_ ArrayOrderlineWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayOrderlineWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Orderline, 0, s)
	}
}
func (r ArrayOrderlineWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayOrderlineWrapper) AppendArray() types.Field {
	var v Orderline
	v = NewOrderline()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
