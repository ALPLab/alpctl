// Code generated by protoc-gen-go. DO NOT EDIT.
// source: osi_sensordata.proto

package osi3

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// Data qualifier communicates the overall availability of the
// interface.
//
type DetectedEntityHeader_DataQualifier int32

const (
	// Unknown (must not be used in ground truth).
	//
	DetectedEntityHeader_DATA_QUALIFIER_UNKNOWN DetectedEntityHeader_DataQualifier = 0
	// Other (unspecified but known).
	//
	DetectedEntityHeader_DATA_QUALIFIER_OTHER DetectedEntityHeader_DataQualifier = 1
	// Data is available.
	//
	DetectedEntityHeader_DATA_QUALIFIER_AVAILABLE DetectedEntityHeader_DataQualifier = 2
	// Reduced data is available.
	//
	DetectedEntityHeader_DATA_QUALIFIER_AVAILABLE_REDUCED DetectedEntityHeader_DataQualifier = 3
	// Data is not available.
	//
	DetectedEntityHeader_DATA_QUALIFIER_NOT_AVAILABLE DetectedEntityHeader_DataQualifier = 4
	// Sensor is blind.
	//
	DetectedEntityHeader_DATA_QUALIFIER_BLINDNESS DetectedEntityHeader_DataQualifier = 5
	// Sensor temporary available.
	//
	DetectedEntityHeader_DATA_QUALIFIER_TEMPORARY_AVAILABLE DetectedEntityHeader_DataQualifier = 6
)

var DetectedEntityHeader_DataQualifier_name = map[int32]string{
	0: "DATA_QUALIFIER_UNKNOWN",
	1: "DATA_QUALIFIER_OTHER",
	2: "DATA_QUALIFIER_AVAILABLE",
	3: "DATA_QUALIFIER_AVAILABLE_REDUCED",
	4: "DATA_QUALIFIER_NOT_AVAILABLE",
	5: "DATA_QUALIFIER_BLINDNESS",
	6: "DATA_QUALIFIER_TEMPORARY_AVAILABLE",
}

var DetectedEntityHeader_DataQualifier_value = map[string]int32{
	"DATA_QUALIFIER_UNKNOWN":             0,
	"DATA_QUALIFIER_OTHER":               1,
	"DATA_QUALIFIER_AVAILABLE":           2,
	"DATA_QUALIFIER_AVAILABLE_REDUCED":   3,
	"DATA_QUALIFIER_NOT_AVAILABLE":       4,
	"DATA_QUALIFIER_BLINDNESS":           5,
	"DATA_QUALIFIER_TEMPORARY_AVAILABLE": 6,
}

func (x DetectedEntityHeader_DataQualifier) String() string {
	return proto.EnumName(DetectedEntityHeader_DataQualifier_name, int32(x))
}

func (DetectedEntityHeader_DataQualifier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fed9f41b516b5af3, []int{0, 0}
}

//
// \brief The header attributes of each detected entity.
//
type DetectedEntityHeader struct {
	// Time stamp at which the measurement was taken (not the time at which it
	// was processed or at which it is transmitted) in the global synchronized
	// time.
	//
	// \note See \c SensorData::timestamp and \c
	// SensorData::last_measurement_time for detailed discussions on the
	// semantics of time-related fields.
	//
	MeasurementTime *Timestamp `protobuf:"bytes,1,opt,name=measurement_time,json=measurementTime,proto3" json:"measurement_time,omitempty"`
	// Continuous up counter to identify the cycle.
	//
	CycleCounter uint64 `protobuf:"varint,2,opt,name=cycle_counter,json=cycleCounter,proto3" json:"cycle_counter,omitempty"`
	// Data Qualifier expresses to what extent the content of this event can be
	// relied on.
	//
	DataQualifier        DetectedEntityHeader_DataQualifier `protobuf:"varint,3,opt,name=data_qualifier,json=dataQualifier,proto3,enum=osi3.DetectedEntityHeader_DataQualifier" json:"data_qualifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *DetectedEntityHeader) Reset()         { *m = DetectedEntityHeader{} }
func (m *DetectedEntityHeader) String() string { return proto.CompactTextString(m) }
func (*DetectedEntityHeader) ProtoMessage()    {}
func (*DetectedEntityHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed9f41b516b5af3, []int{0}
}

func (m *DetectedEntityHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectedEntityHeader.Unmarshal(m, b)
}
func (m *DetectedEntityHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectedEntityHeader.Marshal(b, m, deterministic)
}
func (m *DetectedEntityHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectedEntityHeader.Merge(m, src)
}
func (m *DetectedEntityHeader) XXX_Size() int {
	return xxx_messageInfo_DetectedEntityHeader.Size(m)
}
func (m *DetectedEntityHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectedEntityHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DetectedEntityHeader proto.InternalMessageInfo

func (m *DetectedEntityHeader) GetMeasurementTime() *Timestamp {
	if m != nil {
		return m.MeasurementTime
	}
	return nil
}

func (m *DetectedEntityHeader) GetCycleCounter() uint64 {
	if m != nil {
		return m.CycleCounter
	}
	return 0
}

func (m *DetectedEntityHeader) GetDataQualifier() DetectedEntityHeader_DataQualifier {
	if m != nil {
		return m.DataQualifier
	}
	return DetectedEntityHeader_DATA_QUALIFIER_UNKNOWN
}

//
// \brief The sensor information derived from \c GroundTruth and processed by
// sensor-models.
//
// The sensor information is supposed to imitate the output of real sensors.
// All information regarding the environment is given with respect to
// the virtual sensor coordinate system specified in
// \c SensorData::mounting_position, except for feature data, which is given
// with respect to the physical sensor coordinate system specified in the
// corresponding physical sensor's coordinate system.
//
// When simulating multiple distinct sensors, each sensor can produce an
// individual copy of the \c SensorData interface. This allows an independent
// treatment of the sensors.
//
// Sensor fusion models can consolidate multiple \c SensorData interfaces into
// one consolidated \c SensorData interface.  This can happen either in
// seperate logical models, consuming and producing \c SensorData interfaces,
// or it can happen as part of a combined sensor/logical model, that consumes
// \c SensorView interfaces and directly produces one consolidated \c SensorData
// output.
//
type SensorData struct {
	// The interface version used by the sender.
	//
	Version *InterfaceVersion `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The timestamp of the sensor data. Zero time is arbitrary but must be
	// identical for all messages. Zero time does not need to coincide with
	// the unix epoch. Recommended is the starting time point of the
	// simulation.
	//
	// \note This is the point in time that the sensor data message becomes
	// available to the rest of the system (i.e. the driving functions), so
	// it corresponds with the sending time and thus takes the latency of
	// internal processing of the sensor into account. Latencies of bus
	// communications, etc., that occur after the sensor output have to be
	// applied on top of this, if needed.
	//
	// The time that the actual measurement was performed (which will usually
	// correspond with the timestamp of the \c GroundTruth the sensor model
	// processed to arrive at these results) can be found in the additional
	// field \c SensorData::last_measurement_time.
	//
	// For an ideal zero latency sensor the two timestamps would be the same
	// and would correspond with the timestamp from the current \c GroundTruth
	// message.
	//
	// For a sensor model that does not know its own internal latencies (e.g.
	// a dumb sensor with no internal time concept), the two timestamps might
	// also be identical, but delayed from the \c GroundTruth timestamp.
	//
	Timestamp *Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The sensors estimated location of the host vehicle
	//
	// \note This value is only set by sensors that are able to
	// provide an own estimation of the host vehicle location.
	// \note Note that dimension and base_polygon need not be set.
	// \note The parent frame of \c host_vehicle_location is the sensor frame.
	//
	HostVehicleLocation *BaseMoving `protobuf:"bytes,3,opt,name=host_vehicle_location,json=hostVehicleLocation,proto3" json:"host_vehicle_location,omitempty"`
	// The sensors estimated location error of the host vehicle
	//
	// \note This value is only set by sensors that are able to
	// provide an own estimation of the host vehicle location.
	// \note Note that dimension and base_polygon need not be set.
	// \note The parent frame of \c host_vehicle_location_rmse is the sensor frame.
	//
	HostVehicleLocationRmse *BaseMoving `protobuf:"bytes,4,opt,name=host_vehicle_location_rmse,json=hostVehicleLocationRmse,proto3" json:"host_vehicle_location_rmse,omitempty"`
	// The ID of the sensor at host vehicle's mounting_position.
	//
	// This is the ID of the virtual sensor, to be used in its detected
	// object output; it is distinct from the IDs of its physical detectors,
	// which are used in the detected features.
	//
	SensorId *Identifier `protobuf:"bytes,5,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	// The virtual mounting position of the sensor (origin and orientation
	// of the sensor coordinate system) given in vehicle coordinates [1].
	// The virtual position pertains to the sensor as a whole, regardless
	// of the actual position of individual physical detectors, and governs
	// the sensor-relative coordinates in detected objects of the sensor
	// as a whole.  Individual features detected by individual physical
	// detectors are governed by the actual physical mounting positions
	// of the detectors, as indicated in the technology-specific sub-views
	// and sub-view configurations.
	//
	// \arg \b x-direction of sensor coordinate system: sensor viewing direction
	// \arg \b z-direction of sensor coordinate system: sensor (up)
	// \arg \b y-direction of sensor coordinate system: perpendicular to x and z
	// right hand system
	//
	// \par References:
	// - [1] DIN ISO 8855:2013-11
	//
	// \note This field is usually static during the simulation.
	// \note The origin of vehicle's coordinate system in world frame is
	// ( \c MovingObject::base . \c BaseMoving::position +
	// Inverse_Rotation_yaw_pitch_roll( \c MovingObject::base . \c
	// BaseMoving::orientation) * \c
	// MovingObject::VehicleAttributes::bbcenter_to_rear) . The orientation of
	// the vehicle's coordinate system is equal to the orientation of the
	// vehicle's bounding box \c MovingObject::base . \c
	// BaseMoving::orientation.
	//
	MountingPosition *MountingPosition `protobuf:"bytes,6,opt,name=mounting_position,json=mountingPosition,proto3" json:"mounting_position,omitempty"`
	// The root mean squared error of the mounting position.
	//
	MountingPositionRmse *MountingPosition `protobuf:"bytes,7,opt,name=mounting_position_rmse,json=mountingPositionRmse,proto3" json:"mounting_position_rmse,omitempty"`
	// Sensor view w.r.t. the sensor coordinate system
	//
	// This provides a copy of the \c SensorView data received by the sensor
	// for reference purposes.  For complex sensors or logic models this
	// can be multiple copies.
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	SensorView []*SensorView `protobuf:"bytes,8,rep,name=sensor_view,json=sensorView,proto3" json:"sensor_view,omitempty"`
	// The timestamp of the last real-world measurement (e.g. GT input) that
	// this set of sensor data takes into account. This in effect is the last
	// time instance of reality the measurements correspond to. See field
	// \c SensorData::timestamp for a detailed discussion. This value is also
	// the upper bound to the \c DetectedEntityHeader::measurement_time and the
	// feature data \c SensorDetectionHeader::measurement_time fields.
	//
	LastMeasurementTime *Timestamp `protobuf:"bytes,9,opt,name=last_measurement_time,json=lastMeasurementTime,proto3" json:"last_measurement_time,omitempty"`
	// General information about the \c DetectedStationaryObject .
	//
	StationaryObjectHeader *DetectedEntityHeader `protobuf:"bytes,10,opt,name=stationary_object_header,json=stationaryObjectHeader,proto3" json:"stationary_object_header,omitempty"`
	// The list of stationary objects (e.g. landmarks) detected by the sensor.
	//
	StationaryObject []*DetectedStationaryObject `protobuf:"bytes,11,rep,name=stationary_object,json=stationaryObject,proto3" json:"stationary_object,omitempty"`
	// General information about the \c DetectedMovingObject .
	//
	MovingObjectHeader *DetectedEntityHeader `protobuf:"bytes,12,opt,name=moving_object_header,json=movingObjectHeader,proto3" json:"moving_object_header,omitempty"`
	// The list of moving objects detected by the sensor as perceived by
	// the sensor.
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	MovingObject []*DetectedMovingObject `protobuf:"bytes,13,rep,name=moving_object,json=movingObject,proto3" json:"moving_object,omitempty"`
	// General information about the \c DetectedTrafficSign .
	//
	TrafficSignHeader *DetectedEntityHeader `protobuf:"bytes,14,opt,name=traffic_sign_header,json=trafficSignHeader,proto3" json:"traffic_sign_header,omitempty"`
	// The list of traffic signs detected by the sensor.
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	TrafficSign []*DetectedTrafficSign `protobuf:"bytes,15,rep,name=traffic_sign,json=trafficSign,proto3" json:"traffic_sign,omitempty"`
	// General information about the \c DetectedTrafficLight .
	//
	TrafficLightHeader *DetectedEntityHeader `protobuf:"bytes,16,opt,name=traffic_light_header,json=trafficLightHeader,proto3" json:"traffic_light_header,omitempty"`
	// The list of traffic lights detected by the sensor.
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	TrafficLight []*DetectedTrafficLight `protobuf:"bytes,17,rep,name=traffic_light,json=trafficLight,proto3" json:"traffic_light,omitempty"`
	// General information about the \c DetectedRoadMarking .
	//
	RoadMarkingHeader *DetectedEntityHeader `protobuf:"bytes,18,opt,name=road_marking_header,json=roadMarkingHeader,proto3" json:"road_marking_header,omitempty"`
	// The list of road markings detected by the sensor.
	// This excludes lane boundary markings.
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	RoadMarking []*DetectedRoadMarking `protobuf:"bytes,19,rep,name=road_marking,json=roadMarking,proto3" json:"road_marking,omitempty"`
	// General information about the \c DetectedLaneBoundary .
	//
	LaneBoundaryHeader *DetectedEntityHeader `protobuf:"bytes,20,opt,name=lane_boundary_header,json=laneBoundaryHeader,proto3" json:"lane_boundary_header,omitempty"`
	// The list of lane boundary markings detected by the sensor.
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	LaneBoundary []*DetectedLaneBoundary `protobuf:"bytes,21,rep,name=lane_boundary,json=laneBoundary,proto3" json:"lane_boundary,omitempty"`
	// General information about the \c DetectedLane .
	//
	LaneHeader *DetectedEntityHeader `protobuf:"bytes,22,opt,name=lane_header,json=laneHeader,proto3" json:"lane_header,omitempty"`
	// The list of lanes detected by the sensor
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	Lane []*DetectedLane `protobuf:"bytes,23,rep,name=lane,proto3" json:"lane,omitempty"`
	// General information about the \c DetectedOccupant .
	//
	OccupantHeader *DetectedEntityHeader `protobuf:"bytes,24,opt,name=occupant_header,json=occupantHeader,proto3" json:"occupant_header,omitempty"`
	// The list of occupants of the host vehicle
	//
	// \note OSI uses singular instead of plural for repeated field names.
	//
	Occupant []*DetectedOccupant `protobuf:"bytes,25,rep,name=occupant,proto3" json:"occupant,omitempty"`
	// Low level feature data interface.
	//
	// Low Level feature data is optionally provided by sensor models that
	// model sensors giving access to this low level data, i.e. data prior to
	// object hypothesis and tracking.
	//
	FeatureData          *FeatureData `protobuf:"bytes,26,opt,name=feature_data,json=featureData,proto3" json:"feature_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SensorData) Reset()         { *m = SensorData{} }
func (m *SensorData) String() string { return proto.CompactTextString(m) }
func (*SensorData) ProtoMessage()    {}
func (*SensorData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fed9f41b516b5af3, []int{1}
}

func (m *SensorData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorData.Unmarshal(m, b)
}
func (m *SensorData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorData.Marshal(b, m, deterministic)
}
func (m *SensorData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorData.Merge(m, src)
}
func (m *SensorData) XXX_Size() int {
	return xxx_messageInfo_SensorData.Size(m)
}
func (m *SensorData) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorData.DiscardUnknown(m)
}

var xxx_messageInfo_SensorData proto.InternalMessageInfo

func (m *SensorData) GetVersion() *InterfaceVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *SensorData) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *SensorData) GetHostVehicleLocation() *BaseMoving {
	if m != nil {
		return m.HostVehicleLocation
	}
	return nil
}

func (m *SensorData) GetHostVehicleLocationRmse() *BaseMoving {
	if m != nil {
		return m.HostVehicleLocationRmse
	}
	return nil
}

func (m *SensorData) GetSensorId() *Identifier {
	if m != nil {
		return m.SensorId
	}
	return nil
}

func (m *SensorData) GetMountingPosition() *MountingPosition {
	if m != nil {
		return m.MountingPosition
	}
	return nil
}

func (m *SensorData) GetMountingPositionRmse() *MountingPosition {
	if m != nil {
		return m.MountingPositionRmse
	}
	return nil
}

func (m *SensorData) GetSensorView() []*SensorView {
	if m != nil {
		return m.SensorView
	}
	return nil
}

func (m *SensorData) GetLastMeasurementTime() *Timestamp {
	if m != nil {
		return m.LastMeasurementTime
	}
	return nil
}

func (m *SensorData) GetStationaryObjectHeader() *DetectedEntityHeader {
	if m != nil {
		return m.StationaryObjectHeader
	}
	return nil
}

func (m *SensorData) GetStationaryObject() []*DetectedStationaryObject {
	if m != nil {
		return m.StationaryObject
	}
	return nil
}

func (m *SensorData) GetMovingObjectHeader() *DetectedEntityHeader {
	if m != nil {
		return m.MovingObjectHeader
	}
	return nil
}

func (m *SensorData) GetMovingObject() []*DetectedMovingObject {
	if m != nil {
		return m.MovingObject
	}
	return nil
}

func (m *SensorData) GetTrafficSignHeader() *DetectedEntityHeader {
	if m != nil {
		return m.TrafficSignHeader
	}
	return nil
}

func (m *SensorData) GetTrafficSign() []*DetectedTrafficSign {
	if m != nil {
		return m.TrafficSign
	}
	return nil
}

func (m *SensorData) GetTrafficLightHeader() *DetectedEntityHeader {
	if m != nil {
		return m.TrafficLightHeader
	}
	return nil
}

func (m *SensorData) GetTrafficLight() []*DetectedTrafficLight {
	if m != nil {
		return m.TrafficLight
	}
	return nil
}

func (m *SensorData) GetRoadMarkingHeader() *DetectedEntityHeader {
	if m != nil {
		return m.RoadMarkingHeader
	}
	return nil
}

func (m *SensorData) GetRoadMarking() []*DetectedRoadMarking {
	if m != nil {
		return m.RoadMarking
	}
	return nil
}

func (m *SensorData) GetLaneBoundaryHeader() *DetectedEntityHeader {
	if m != nil {
		return m.LaneBoundaryHeader
	}
	return nil
}

func (m *SensorData) GetLaneBoundary() []*DetectedLaneBoundary {
	if m != nil {
		return m.LaneBoundary
	}
	return nil
}

func (m *SensorData) GetLaneHeader() *DetectedEntityHeader {
	if m != nil {
		return m.LaneHeader
	}
	return nil
}

func (m *SensorData) GetLane() []*DetectedLane {
	if m != nil {
		return m.Lane
	}
	return nil
}

func (m *SensorData) GetOccupantHeader() *DetectedEntityHeader {
	if m != nil {
		return m.OccupantHeader
	}
	return nil
}

func (m *SensorData) GetOccupant() []*DetectedOccupant {
	if m != nil {
		return m.Occupant
	}
	return nil
}

func (m *SensorData) GetFeatureData() *FeatureData {
	if m != nil {
		return m.FeatureData
	}
	return nil
}

func init() {
	proto.RegisterEnum("osi3.DetectedEntityHeader_DataQualifier", DetectedEntityHeader_DataQualifier_name, DetectedEntityHeader_DataQualifier_value)
	proto.RegisterType((*DetectedEntityHeader)(nil), "osi3.DetectedEntityHeader")
	proto.RegisterType((*SensorData)(nil), "osi3.SensorData")
}

func init() { proto.RegisterFile("osi_sensordata.proto", fileDescriptor_fed9f41b516b5af3) }

var fileDescriptor_fed9f41b516b5af3 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x9d, 0x13, 0x37, 0x4d, 0xaf, 0xed, 0x44, 0x66, 0x1c, 0x97, 0x35, 0xba, 0xc2, 0xc8, 0x86,
	0x22, 0x2f, 0x0b, 0xb6, 0x74, 0x4f, 0xdb, 0x80, 0xc1, 0x5f, 0x45, 0xbc, 0xfa, 0xa3, 0xa5, 0x9d,
	0x0c, 0x7b, 0x22, 0x18, 0x89, 0x76, 0xb8, 0x59, 0x52, 0x26, 0xd1, 0x29, 0xf2, 0xbf, 0xf6, 0xb8,
	0x3f, 0xb5, 0x7f, 0x30, 0x90, 0xa2, 0x22, 0x5a, 0x51, 0xea, 0x37, 0xfb, 0x9e, 0x7b, 0xce, 0xbd,
	0x3c, 0x97, 0x1f, 0x82, 0x46, 0x18, 0x0b, 0x1a, 0xf3, 0x20, 0x0e, 0x23, 0x8f, 0x49, 0x76, 0x76,
	0x1b, 0x85, 0x32, 0x44, 0xe5, 0x30, 0x16, 0xef, 0x5a, 0x75, 0x85, 0xdd, 0xf1, 0x28, 0x16, 0x61,
	0x90, 0x00, 0x2d, 0x47, 0x85, 0xdc, 0xd0, 0xf7, 0x1f, 0x22, 0x5f, 0xab, 0x88, 0xc7, 0x25, 0x77,
	0x25, 0xf7, 0x64, 0xc4, 0x16, 0x0b, 0xe1, 0xc6, 0x62, 0x99, 0xc2, 0x6f, 0x0a, 0xe0, 0x95, 0x58,
	0xde, 0xc8, 0x22, 0x7a, 0x14, 0x32, 0xcf, 0x67, 0xd1, 0x5f, 0x22, 0x58, 0x1a, 0xb8, 0x69, 0xc3,
	0x2b, 0x16, 0x70, 0x13, 0xc7, 0x76, 0x3c, 0xbc, 0xfe, 0x93, 0xbb, 0xa9, 0x60, 0x6b, 0x03, 0x71,
	0xdd, 0xf5, 0x2d, 0x0b, 0x52, 0xcc, 0x5a, 0xec, 0x9d, 0xe0, 0x9f, 0x4d, 0xf4, 0x58, 0x45, 0x17,
	0x9c, 0xc9, 0x75, 0xc4, 0x33, 0x0f, 0x4e, 0xfe, 0xdd, 0x85, 0x46, 0xdf, 0xe8, 0x0c, 0x02, 0x29,
	0xe4, 0xfd, 0x05, 0x67, 0x1e, 0x8f, 0xd0, 0x4f, 0xe0, 0xf8, 0x9c, 0xc5, 0xeb, 0x88, 0xfb, 0x3c,
	0x90, 0x54, 0x0a, 0x9f, 0xe3, 0x52, 0xbb, 0x74, 0x5a, 0x39, 0x3f, 0x3c, 0x53, 0xbe, 0x9d, 0xcd,
	0x85, 0xcf, 0x63, 0xc9, 0xfc, 0x5b, 0x72, 0x68, 0x25, 0xaa, 0x28, 0xfa, 0x06, 0x6a, 0xee, 0xbd,
	0xbb, 0xe2, 0xd4, 0x0d, 0xd7, 0x81, 0xe4, 0x11, 0xde, 0x69, 0x97, 0x4e, 0xcb, 0xa4, 0xaa, 0x83,
	0xbd, 0x24, 0x86, 0xa6, 0x70, 0xa0, 0xfa, 0xa0, 0x7f, 0xaf, 0xd9, 0x4a, 0x2c, 0x04, 0x8f, 0xf0,
	0x6e, 0xbb, 0x74, 0x7a, 0x70, 0x7e, 0x9a, 0xc8, 0x17, 0x35, 0x75, 0xd6, 0x67, 0x92, 0x7d, 0x4a,
	0xf3, 0x49, 0xcd, 0xb3, 0xff, 0x9e, 0xfc, 0x57, 0x82, 0xda, 0x46, 0x02, 0x6a, 0x41, 0xb3, 0xdf,
	0x99, 0x77, 0xe8, 0xa7, 0xcb, 0xce, 0x68, 0xf8, 0x7e, 0x38, 0x20, 0xf4, 0x72, 0xf2, 0x61, 0x32,
	0xfd, 0x7d, 0xe2, 0x7c, 0x85, 0x30, 0x34, 0x72, 0xd8, 0x74, 0x7e, 0x31, 0x20, 0x4e, 0x09, 0xbd,
	0x06, 0x9c, 0x43, 0x3a, 0x57, 0x9d, 0xe1, 0xa8, 0xd3, 0x1d, 0x0d, 0x9c, 0x1d, 0xf4, 0x2d, 0xb4,
	0x9f, 0x42, 0x29, 0x19, 0xf4, 0x2f, 0x7b, 0x83, 0xbe, 0xb3, 0x8b, 0xda, 0xf0, 0x3a, 0x97, 0x35,
	0x99, 0xce, 0x2d, 0x9d, 0x72, 0x41, 0x95, 0xee, 0x68, 0x38, 0xe9, 0x4f, 0x06, 0xb3, 0x99, 0xf3,
	0x0c, 0xbd, 0x85, 0x93, 0x1c, 0x3a, 0x1f, 0x8c, 0x3f, 0x4e, 0x49, 0x87, 0xfc, 0x61, 0xa9, 0xec,
	0x9d, 0xfc, 0x53, 0x03, 0x98, 0xe9, 0x51, 0xab, 0x95, 0xa3, 0xef, 0xe1, 0xb9, 0xd9, 0xc9, 0x66,
	0x56, 0xcd, 0xc4, 0xcc, 0xa1, 0x72, 0x7c, 0xc1, 0x5c, 0x7e, 0x95, 0xa0, 0x24, 0x4d, 0x43, 0xdf,
	0xc1, 0x0b, 0x99, 0x0e, 0x52, 0x8f, 0xa9, 0x60, 0xbe, 0x59, 0x06, 0xea, 0xc3, 0xf1, 0x4d, 0x18,
	0x4b, 0x7a, 0xc7, 0x6f, 0x84, 0x1a, 0xf0, 0x2a, 0x74, 0x99, 0x54, 0xe5, 0x76, 0x35, 0xd5, 0x49,
	0xa8, 0x5d, 0x16, 0xf3, 0x71, 0x78, 0x27, 0x82, 0x25, 0x39, 0x52, 0xe9, 0x57, 0x49, 0xf6, 0xc8,
	0x24, 0xa3, 0x31, 0xb4, 0x0a, 0x55, 0x68, 0xe4, 0xc7, 0x1c, 0x97, 0x9f, 0x90, 0x7a, 0x59, 0x20,
	0x45, 0xfc, 0x98, 0xab, 0x35, 0x24, 0xdb, 0x9d, 0x0a, 0x0f, 0x3f, 0xb3, 0xd9, 0x43, 0x8f, 0x07,
	0x32, 0xd9, 0x2c, 0xfb, 0x49, 0xca, 0xd0, 0x43, 0x3d, 0xa8, 0xfb, 0x6a, 0x0f, 0x8a, 0x60, 0x49,
	0x6f, 0xc3, 0x58, 0xe8, 0xfe, 0xf7, 0x6c, 0xbb, 0xc6, 0x06, 0xfe, 0x68, 0x50, 0xe2, 0xf8, 0xb9,
	0x08, 0x1a, 0x41, 0xf3, 0x91, 0x48, 0xd2, 0xfe, 0xf3, 0x2f, 0x2a, 0x35, 0xf2, 0x4a, 0x7a, 0x05,
	0x3f, 0x40, 0xc5, 0xac, 0x40, 0x9d, 0x58, 0xbc, 0xdf, 0xde, 0xcd, 0xd6, 0x90, 0x8c, 0xf7, 0x4a,
	0xf0, 0xcf, 0x04, 0xe2, 0x87, 0xdf, 0xa8, 0x07, 0xc7, 0x2b, 0x16, 0x4b, 0xfa, 0xe8, 0x90, 0xbe,
	0x28, 0x1e, 0xe2, 0x91, 0xca, 0x1e, 0xe7, 0x0e, 0xea, 0x1c, 0x70, 0x2c, 0xb5, 0x91, 0x2c, 0xba,
	0xa7, 0xc9, 0x0d, 0x43, 0x6f, 0xf4, 0x59, 0xc3, 0xa0, 0x75, 0x5a, 0x4f, 0x9f, 0x46, 0xd2, 0xcc,
	0xb8, 0x53, 0x4d, 0x35, 0x57, 0xc7, 0x07, 0xa8, 0x3f, 0x52, 0xc5, 0x15, 0xbd, 0xa6, 0x37, 0x9b,
	0x72, 0xb3, 0x9c, 0x00, 0x71, 0xf2, 0x92, 0x68, 0x04, 0x0d, 0x5f, 0xcf, 0x3f, 0xd7, 0x5e, 0x75,
	0x6b, 0x7b, 0x28, 0xe1, 0x6d, 0xb4, 0xf6, 0x2b, 0xd4, 0x36, 0xd4, 0x70, 0x4d, 0xb7, 0x95, 0x93,
	0x19, 0x5b, 0x44, 0x52, 0xb5, 0x65, 0xd0, 0x6f, 0x70, 0x64, 0xee, 0x77, 0xaa, 0xee, 0xff, 0xb4,
	0x9b, 0x83, 0xad, 0xdd, 0xd4, 0x0d, 0x6d, 0x26, 0x96, 0x81, 0x69, 0xe6, 0x17, 0xa8, 0xda, 0x5a,
	0xf8, 0x50, 0xf7, 0xf2, 0x6a, 0x53, 0x64, 0x9e, 0xd1, 0x48, 0xc5, 0xd2, 0x50, 0xc6, 0xa4, 0x6c,
	0xfd, 0xd4, 0xa4, 0xad, 0x38, 0xdb, 0x8d, 0x31, 0xbc, 0x91, 0xa2, 0x65, 0xc6, 0x6c, 0xa8, 0xe1,
	0x7a, 0x91, 0x31, 0x73, 0x8b, 0x48, 0xaa, 0xb6, 0x8c, 0x32, 0x46, 0x3d, 0x6c, 0xd4, 0xbc, 0x6c,
	0x69, 0x37, 0x68, 0xbb, 0x31, 0x8a, 0x36, 0x4e, 0x58, 0x99, 0x31, 0xb6, 0x16, 0x3e, 0x2a, 0x32,
	0x86, 0x64, 0x34, 0x52, 0xb1, 0x34, 0x94, 0x31, 0xea, 0x0d, 0xa5, 0xd7, 0xe1, 0x3a, 0xf0, 0xd4,
	0x0e, 0x34, 0xad, 0x34, 0xb6, 0x1b, 0xa3, 0x78, 0x5d, 0x43, 0xcb, 0x8c, 0xd9, 0x50, 0xc3, 0xc7,
	0x45, 0xc6, 0x8c, 0x2c, 0x22, 0xa9, 0xda, 0x32, 0xe8, 0x67, 0xa8, 0x68, 0x01, 0xd3, 0x45, 0x73,
	0x6b, 0x17, 0xa0, 0xd2, 0x4d, 0xf5, 0xb7, 0x50, 0x56, 0xff, 0xf0, 0x4b, 0x5d, 0x14, 0x3d, 0x2e,
	0x4a, 0x34, 0x8e, 0x7a, 0x70, 0x98, 0x7e, 0x05, 0xa4, 0x85, 0xf0, 0xd6, 0x42, 0x07, 0x29, 0xc5,
	0x14, 0x3b, 0x87, 0xfd, 0x34, 0x82, 0x5f, 0xe9, 0x82, 0xcd, 0x4d, 0xf6, 0xd4, 0xa0, 0xe4, 0x21,
	0x0f, 0xfd, 0x08, 0x55, 0xf3, 0x51, 0x41, 0xd5, 0x6b, 0x8c, 0x5b, 0xba, 0x6a, 0x3d, 0xe1, 0xbd,
	0x4f, 0x10, 0xf5, 0x34, 0x91, 0xca, 0x22, 0xfb, 0xd3, 0xdd, 0xb9, 0x28, 0x5d, 0xef, 0xe9, 0x0f,
	0x90, 0x77, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x56, 0xc0, 0x86, 0x7e, 0x9c, 0x09, 0x00, 0x00,
}
