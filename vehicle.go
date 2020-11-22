package tesla

import (
	"encoding/json"
	"fmt"
)

type Vehicle struct {
	ID                     int64       `json:"id"`
	UserID                 int64       `json:"user_id"`
	VehicleID              int64       `json:"vehicle_id"`
	Vin                    string      `json:"vin"`
	DisplayName            string      `json:"display_name"`
	OptionCodes            string      `json:"option_codes"`
	Color                  interface{} `json:"color"`
	AccessType             string      `json:"access_type"`
	Tokens                 []string    `json:"tokens"`
	State                  string      `json:"state"` // "online", "asleep", "offline", "waking", "unknown"
	InService              bool        `json:"in_service"`
	IDS                    string      `json:"id_s"`
	CalendarEnabled        bool        `json:"calendar_enabled"`
	APIVersion             int64       `json:"api_version"`
	BackseatToken          interface{} `json:"backseat_token"`
	BackseatTokenUpdatedAt interface{} `json:"backseat_token_updated_at"`
}

type ChargeState struct {
	BatteryHeaterOn             bool        `json:"battery_heater_on"`
	BatteryLevel                int         `json:"battery_level"`
	BatteryRange                float64     `json:"battery_range"`
	ChargeCurrentRequest        int         `json:"charge_current_request"`
	ChargeCurrentRequestMax     int         `json:"charge_current_request_max"`
	ChargeEnableRequest         bool        `json:"charge_enable_request"`
	ChargeEnergyAdded           float64     `json:"charge_energy_added"`
	ChargeLimitSoc              int         `json:"charge_limit_soc"`
	ChargeLimitSocMax           int         `json:"charge_limit_soc_max"`
	ChargeLimitSocMin           int         `json:"charge_limit_soc_min"`
	ChargeLimitSocStd           int         `json:"charge_limit_soc_std"`
	ChargeMilesAddedIdeal       float64     `json:"charge_miles_added_ideal"`
	ChargeMilesAddedRated       float64     `json:"charge_miles_added_rated"`
	ChargePortColdWeatherMode   interface{} `json:"charge_port_cold_weather_mode"`
	ChargePortDoorOpen          bool        `json:"charge_port_door_open"`
	ChargePortLatch             string      `json:"charge_port_latch"`
	ChargeRate                  float64     `json:"charge_rate"`
	ChargeToMaxRange            bool        `json:"charge_to_max_range"`
	ChargerActualCurrent        int         `json:"charger_actual_current"`
	ChargerPhases               interface{} `json:"charger_phases"`
	ChargerPilotCurrent         int         `json:"charger_pilot_current"`
	ChargerPower                int         `json:"charger_power"`
	ChargerVoltage              int         `json:"charger_voltage"`
	ChargingState               string      `json:"charging_state"` // Starting, Charging, Complete, Stopped
	ConnChargeCable             string      `json:"conn_charge_cable"`
	EstBatteryRange             float64     `json:"est_battery_range"`
	FastChargerBrand            string      `json:"fast_charger_brand"`
	FastChargerPresent          bool        `json:"fast_charger_present"`
	FastChargerType             string      `json:"fast_charger_type"`
	IdealBatteryRange           float64     `json:"ideal_battery_range"`
	ManagedChargingActive       bool        `json:"managed_charging_active"`
	ManagedChargingStartTime    interface{} `json:"managed_charging_start_time"`
	ManagedChargingUserCanceled bool        `json:"managed_charging_user_canceled"`
	MaxRangeChargeCounter       int         `json:"max_range_charge_counter"`
	MinutesToFullCharge         int         `json:"minutes_to_full_charge"`
	NotEnoughPowerToHeat        bool        `json:"not_enough_power_to_heat"`
	ScheduledChargingPending    bool        `json:"scheduled_charging_pending"`
	ScheduledChargingStartTime  interface{} `json:"scheduled_charging_start_time"`
	TimeToFullCharge            float64     `json:"time_to_full_charge"`
	Timestamp                   int64       `json:"timestamp"`
	TripCharging                bool        `json:"trip_charging"`
	UsableBatteryLevel          int         `json:"usable_battery_level"`
	UserChargeEnableRequest     interface{} `json:"user_charge_enable_request"`
}

type ClimateState struct {
	BatteryHeater              bool    `json:"battery_heater"`
	BatteryHeaterNoPower       bool    `json:"battery_heater_no_power"`
	ClimateKeeperMode          string  `json:"climate_keeper_mode"`
	DefrostMode                int     `json:"defrost_mode"`
	DriverTempSetting          float64 `json:"driver_temp_setting"`
	FanStatus                  int     `json:"fan_status"`
	InsideTemp                 float64 `json:"inside_temp"`
	IsAutoConditioningOn       bool    `json:"is_auto_conditioning_on"`
	IsClimateOn                bool    `json:"is_climate_on"`
	IsFrontDefrosterOn         bool    `json:"is_front_defroster_on"`
	IsPreconditioning          bool    `json:"is_preconditioning"`
	IsRearDefrosterOn          bool    `json:"is_rear_defroster_on"`
	LeftTempDirection          int     `json:"left_temp_direction"`
	MaxAvailTemp               float64 `json:"max_avail_temp"`
	MinAvailTemp               float64 `json:"min_avail_temp"`
	OutsideTemp                float64 `json:"outside_temp"`
	PassengerTempSetting       float64 `json:"passenger_temp_setting"`
	RemoteHeaterControlEnabled bool    `json:"remote_heater_control_enabled"`
	RightTempDirection         int     `json:"right_temp_direction"`
	SeatHeaterLeft             int64   `json:"seat_heater_left"`
	SeatHeaterRight            int64   `json:"seat_heater_right"`
	SideMirrorHeaters          bool    `json:"side_mirror_heaters"`
	Timestamp                  int64   `json:"timestamp"`
	WiperBladeHeater           bool    `json:"wiper_blade_heater"`
}

type DriveState struct {
	GpsAsOf                 int64       `json:"gps_as_of"`
	Heading                 int64       `json:"heading"`
	Latitude                float64     `json:"latitude"`
	Longitude               float64     `json:"longitude"`
	NativeLatitude          float64     `json:"native_latitude"`
	NativeLocationSupported int64       `json:"native_location_supported"`
	NativeLongitude         float64     `json:"native_longitude"`
	NativeType              string      `json:"native_type"`
	Power                   int64       `json:"power"`
	ShiftState              interface{} `json:"shift_state"`
	Speed                   interface{} `json:"speed"`
	Timestamp               int64       `json:"timestamp"`
}

type GuiSettings struct {
	Gui24HourTime       bool   `json:"gui_24_hour_time"`
	GuiChargeRateUnits  string `json:"gui_charge_rate_units"`
	GuiDistanceUnits    string `json:"gui_distance_units"`
	GuiRangeDisplay     string `json:"gui_range_display"`
	GuiTemperatureUnits string `json:"gui_temperature_units"`
	ShowRangeUnits      bool   `json:"show_range_units"`
	Timestamp           int64  `json:"timestamp"`
}

type VehicleConfig struct {
	CanAcceptNavigationRequests bool   `json:"can_accept_navigation_requests"`
	CanActuateTrunks            bool   `json:"can_actuate_trunks"`
	CarSpecialType              string `json:"car_special_type"`
	CarType                     string `json:"car_type"`
	ChargePortType              string `json:"charge_port_type"`
	EceRestrictions             bool   `json:"ece_restrictions"`
	EuVehicle                   bool   `json:"eu_vehicle"`
	ExteriorColor               string `json:"exterior_color"`
	HasAirSuspension            bool   `json:"has_air_suspension"`
	HasLudicrousMode            bool   `json:"has_ludicrous_mode"`
	MotorizedChargePort         bool   `json:"motorized_charge_port"`
	Plg                         bool   `json:"plg"`
	RearSeatHeaters             int    `json:"rear_seat_heaters"`
	RearSeatType                int    `json:"rear_seat_type"`
	Rhd                         bool   `json:"rhd"`
	RoofColor                   string `json:"roof_color"`
	SeatType                    int    `json:"seat_type"`
	SpoilerType                 string `json:"spoiler_type"`
	SunRoofInstalled            int    `json:"sun_roof_installed"`
	ThirdRowSeats               string `json:"third_row_seats"`
	Timestamp                   int64  `json:"timestamp"`
	TrimBadging                 string `json:"trim_badging"`
	UseRangeBadging             bool   `json:"use_range_badging"`
	WheelType                   string `json:"wheel_type"`
}

type MediaState struct {
	RemoteControlEnabled bool `json:"remote_control_enabled"`
}

type SoftwareUpdate struct {
	DownloadPerc        int    `json:"download_perc"`
	ExpectedDurationSec int    `json:"expected_duration_sec"`
	InstallPerc         int    `json:"install_perc"`
	Status              string `json:"status"`
	Version             string `json:"version"`
}

type SpeedLimitMode struct {
	Active          bool    `json:"active"`
	CurrentLimitMph float64 `json:"current_limit_mph"`
	MaxLimitMph     int     `json:"max_limit_mph"`
	MinLimitMph     int     `json:"min_limit_mph"`
	PinCodeSet      bool    `json:"pin_code_set"`
}

type VehicleState struct {
	APIVersion               int            `json:"api_version"`
	AutoparkStateV2          string         `json:"autopark_state_v2"`
	AutoparkStyle            string         `json:"autopark_style"`
	CalendarSupported        bool           `json:"calendar_supported"`
	CarVersion               string         `json:"car_version"`
	CenterDisplayState       int            `json:"center_display_state"`
	Df                       int            `json:"df"`
	Dr                       int            `json:"dr"`
	Ft                       int            `json:"ft"`
	HomelinkDeviceCount      int            `json:"homelink_device_count"`
	HomelinkNearby           bool           `json:"homelink_nearby"`
	IsUserPresent            bool           `json:"is_user_present"`
	LastAutoparkError        string         `json:"last_autopark_error"`
	Locked                   bool           `json:"locked"`
	MediaState               MediaState     `json:"media_state"`
	NotificationsSupported   bool           `json:"notifications_supported"`
	Odometer                 float64        `json:"odometer"`
	ParsedCalendarSupported  bool           `json:"parsed_calendar_supported"`
	Pf                       int            `json:"pf"`
	Pr                       int            `json:"pr"`
	RemoteStart              bool           `json:"remote_start"`
	RemoteStartEnabled       bool           `json:"remote_start_enabled"`
	RemoteStartSupported     bool           `json:"remote_start_supported"`
	Rt                       int            `json:"rt"`
	SentryMode               bool           `json:"sentry_mode"`
	SentryModeAvailable      bool           `json:"sentry_mode_available"`
	SmartSummonAvailable     bool           `json:"smart_summon_available"`
	SoftwareUpdate           SoftwareUpdate `json:"software_update"`
	SpeedLimitMode           SpeedLimitMode `json:"speed_limit_mode"`
	SummonStandbyModeEnabled bool           `json:"summon_standby_mode_enabled"`
	SunRoofPercentOpen       int            `json:"sun_roof_percent_open"`
	SunRoofState             string         `json:"sun_roof_state"`
	Timestamp                int64          `json:"timestamp"`
	ValetMode                bool           `json:"valet_mode"`
	ValetPinNeeded           bool           `json:"valet_pin_needed"`
	VehicleName              string         `json:"vehicle_name"`
}

type VehicleData struct {
	ID                     int64         `json:"id"`
	UserID                 int64         `json:"user_id"`
	VehicleID              int64         `json:"vehicle_id"`
	Vin                    string        `json:"vin"`
	DisplayName            string        `json:"display_name"`
	OptionCodes            string        `json:"option_codes"`
	Color                  interface{}   `json:"color"`
	AccessType             string        `json:"access_type"`
	Tokens                 []string      `json:"tokens"`
	State                  string        `json:"state"` // "online", "asleep", "offline", "waking", "unknown"
	InService              bool          `json:"in_service"`
	IDS                    string        `json:"id_s"`
	CalendarEnabled        bool          `json:"calendar_enabled"`
	APIVersion             int64         `json:"api_version"`
	BackseatToken          interface{}   `json:"backseat_token"`
	BackseatTokenUpdatedAt interface{}   `json:"backseat_token_updated_at"`
	ChargeState            ChargeState   `json:"charge_state"`
	ClimateState           ClimateState  `json:"climate_state"`
	DriveState             DriveState    `json:"drive_state"`
	GuiSettings            GuiSettings   `json:"gui_settings"`
	MobileEnabled          bool          `json:"mobile_enabled"`
	VehicleState           VehicleState  `json:"vehicle_state"`
	VehicleConfig          VehicleConfig `json:"vehicle_config"`
}

func (v VehicleData) ToVehicle() Vehicle {
	return Vehicle{
		ID:                     v.ID,
		UserID:                 v.UserID,
		VehicleID:              v.VehicleID,
		Vin:                    v.Vin,
		DisplayName:            v.DisplayName,
		OptionCodes:            v.OptionCodes,
		Color:                  v.Color,
		AccessType:             v.AccessType,
		Tokens:                 v.Tokens,
		State:                  v.State,
		InService:              v.InService,
		IDS:                    v.IDS,
		CalendarEnabled:        v.CalendarEnabled,
		APIVersion:             v.APIVersion,
		BackseatToken:          v.BackseatToken,
		BackseatTokenUpdatedAt: v.BackseatTokenUpdatedAt,
	}
}

func (v Vehicle) IsOnline() bool {
	return v.State == "online"
}

type VehiclesResponse struct {
	Response []Vehicle `json:"response"`
	Count    int64     `json:"count"`
}

type VehicleResponse struct {
	Response Vehicle `json:"response"`
}

type VehicleDataResponse struct {
	Response VehicleData `json:"response"`
}

type VehicleChargeStateResponse struct {
	Response ChargeState `json:"response"`
}

type VehicleClimateStateResponse struct {
	Response ClimateState `json:"response"`
}

type VehicleDriveStateResponse struct {
	Response DriveState `json:"response"`
}

type VehicleGuiSettingsResponse struct {
	Response GuiSettings `json:"response"`
}

type VehicleVehicleStateResponse struct {
	Response VehicleState `json:"response"`
}

type VehicleVehicleConfigResponse struct {
	Response VehicleConfig `json:"response"`
}

type BoolResponse struct {
	Response bool `json:"response"`
}

type BoolReasonResponse struct {
	Response bool   `json:"response"`
	Reason   string `json:"reason"`
}

type StringResponse struct {
	Response string `json:"response"`
}

func (c *Client) Vehicles() ([]Vehicle, error) {
	data, err := c.get(BaseURL + "/vehicles")
	if err != nil {
		return nil, err
	}

	var vehicles VehiclesResponse
	err = json.Unmarshal(data, &vehicles)
	if err != nil {
		return nil, err
	}

	return vehicles.Response, nil
}

func (c *Client) MobileEnabled(id int64) bool {
	data, err := c.get(fmt.Sprintf("/api/v1/vehicles/%d/mobile_enabled", id))
	if err != nil {
		return false
	}

	var br BoolResponse
	err = json.Unmarshal(data, &br)
	if err != nil {
		return false
	}

	return br.Response
}
