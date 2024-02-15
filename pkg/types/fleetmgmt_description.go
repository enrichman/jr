// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     csv_product.avsc
 *     csv_user.avsc
 *     finance_stock_trade.avsc
 *     fleetmgmt_description.avsc
 *     fleetmgmt_location.avsc
 *     fleetmgmt_sensor.avsc
 *     gaming_game.avsc
 *     gaming_player.avsc
 *     gaming_player_activity.avsc
 *     genericstore_purchase.avsc
 *     insurance_customer.avsc
 *     insurance_customer_activity.avsc
 *     insurance_offer.avsc
 *     inventorymgmt_inventory.avsc
 *     inventorymgmt_product.avsc
 *     iot_device_information.avsc
 *     map_dumb_schema.avsc
 *     marketing_campaign_finance.avsc
 *     net_device.avsc
 *     payment_credit_card.avsc
 *     payment_transaction.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizzastore_order.avsc
 *     pizzastore_order_cancelled.avsc
 *     pizzastore_order_completed.avsc
 *     shoestore_clickstream.avsc
 *     shoestore_customer.avsc
 *     shoestore_order.avsc
 *     shoestore_shoe.avsc
 *     shopping_order.avsc
 *     shopping_rating.avsc
 *     siem_log.avsc
 *     store.avsc
 *     syslog_log.avsc
 *     user.avsc
 *     users.avsc
 *     users_array_map.avsc
 *     webanalytics_clickstream.avsc
 *     webanalytics_code.avsc
 *     webanalytics_page_view.avsc
 *     webanalytics_user.avsc
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

type FleetmgmtDescription struct {
	Vehicle_id int32 `json:"vehicle_id"`

	Driver_name string `json:"driver_name"`

	License_plate string `json:"license_plate"`
}

const FleetmgmtDescriptionAvroCRC64Fingerprint = "\xf8\xbc\xda<\xa0\xd1}:"

func NewFleetmgmtDescription() FleetmgmtDescription {
	r := FleetmgmtDescription{}
	return r
}

func DeserializeFleetmgmtDescription(r io.Reader) (FleetmgmtDescription, error) {
	t := NewFleetmgmtDescription()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFleetmgmtDescriptionFromSchema(r io.Reader, schema string) (FleetmgmtDescription, error) {
	t := NewFleetmgmtDescription()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFleetmgmtDescription(r FleetmgmtDescription, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Vehicle_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Driver_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.License_plate, w)
	if err != nil {
		return err
	}
	return err
}

func (r FleetmgmtDescription) Serialize(w io.Writer) error {
	return writeFleetmgmtDescription(r, w)
}

func (r FleetmgmtDescription) Schema() string {
	return "{\"fields\":[{\"name\":\"vehicle_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":9999,\"min\":1000}},\"type\":\"int\"}},{\"name\":\"driver_name\",\"type\":{\"arg.properties\":{\"options\":[\"Frieda Baldi\",\"Cherrita Gallaccio\",\"Matt Cleugh\",\"Dulciana Murfill\",\"Germayne Streetley\",\"Brenna Woolfall\",\"Gerhardt Tenbrug\",\"Hayley Tuma\",\"Winny Cadigan\",\"Bonnibelle Macek\",\"Lionel Byneth\",\"Trev Roper\",\"Lena MacFadzean\",\"Benton Allcorn\",\"Avis Moyler\",\"Marchall Rochewell\",\"Adele Bohl\",\"Barnett Mcall\",\"Frieda Pirrone\",\"Pattin Eringey\",\"Kalila Fewings\",\"Giacobo Beuscher\",\"Rozalin Hair\",\"Egon Beagan\",\"Owen Strotton\",\"Fernando Rosensaft\",\"Carleton Gwyther\",\"Kata Coll\",\"Rossie Hobben\",\"Stephanie Gookey\",\"Robyn Milazzo\",\"Tilda O'Lunney\",\"Nolan Kidney\",\"Jori Ottiwill\",\"Benito Graveson\",\"Zechariah Wrate\",\"Chelsae Napton\",\"Jeremy Heffernon\",\"Derk McAviy\",\"Constantin Mears\",\"Fitz Ballin\",\"Essy Bettles\",\"Gene Klemt\",\"Nikolai Arnopp\",\"Gustave Westhofer\",\"Simona Mayhow\",\"Cort Bainbridge\",\"Sibyl Vockins\",\"Andriette Gaze\",\"Shaughn De Simoni\",\"Nathaniel Hallowell\",\"Charley Dudill\",\"Cirstoforo Joblin\",\"Hyacinthia Kinastan\",\"Dur Lasselle\",\"Gay Chadburn\",\"Livvie Hawyes\",\"Aldrich MacVay\",\"Riva Rossant\",\"Johanna Reichartz\",\"Trent Gantlett\",\"Aryn Haskell\",\"Byrann Barock\",\"Gerda Cleugher\",\"Sonnie Guildford\",\"Vergil Borge\",\"Lurline Rocco\",\"Geoff Eddy\",\"Zea Leighton\",\"Leif Baden\",\"Quint Bidgod\",\"Talbot Cashell\",\"Sheridan Foulsham\",\"Camile Shrimplin\",\"Marcel Nayshe\",\"Lea Murrish\",\"Lucais Midson\",\"Zeb Rylatt\",\"Nertie Zuker\",\"Babara Henderson\",\"Electra Ridgley\",\"Jere Standingford\",\"Cyril Yellowlea\",\"Isadora Peegrem\",\"Caria Smewings\",\"Karena Kauffman\",\"Haywood Snowball\",\"Winslow Starcks\",\"Alis Ponton\",\"Marietta Lezemere\",\"Emilee Broadbridge\",\"Faye Beaument\",\"Shannah Beatson\",\"West Doy\",\"Chryste Wren\",\"Trumann Labba\",\"Anatollo Beckwith\",\"Konstanze Dunsford\",\"Raychel Roset\",\"Heindrick Ravenscroft\"]},\"type\":\"string\"}},{\"name\":\"license_plate\",\"type\":{\"arg.properties\":{\"regex\":\"[A-Z]{2}\\\\d{6}[A-Z]\"},\"type\":\"string\"}}],\"name\":\"fleetmgmt.FleetmgmtDescription\",\"type\":\"record\"}"
}

func (r FleetmgmtDescription) SchemaName() string {
	return "fleetmgmt.FleetmgmtDescription"
}

func (_ FleetmgmtDescription) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FleetmgmtDescription) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FleetmgmtDescription) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FleetmgmtDescription) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FleetmgmtDescription) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FleetmgmtDescription) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FleetmgmtDescription) SetString(v string)   { panic("Unsupported operation") }
func (_ FleetmgmtDescription) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FleetmgmtDescription) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Vehicle_id}

		return w

	case 1:
		w := types.String{Target: &r.Driver_name}

		return w

	case 2:
		w := types.String{Target: &r.License_plate}

		return w

	}
	panic("Unknown field index")
}

func (r *FleetmgmtDescription) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *FleetmgmtDescription) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ FleetmgmtDescription) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FleetmgmtDescription) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FleetmgmtDescription) HintSize(int)                     { panic("Unsupported operation") }
func (_ FleetmgmtDescription) Finalize()                        {}

func (_ FleetmgmtDescription) AvroCRC64Fingerprint() []byte {
	return []byte(FleetmgmtDescriptionAvroCRC64Fingerprint)
}

func (r FleetmgmtDescription) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["vehicle_id"], err = json.Marshal(r.Vehicle_id)
	if err != nil {
		return nil, err
	}
	output["driver_name"], err = json.Marshal(r.Driver_name)
	if err != nil {
		return nil, err
	}
	output["license_plate"], err = json.Marshal(r.License_plate)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FleetmgmtDescription) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["driver_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Driver_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for driver_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["license_plate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.License_plate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for license_plate")
	}
	return nil
}