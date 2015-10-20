// Code generated by protoc-gen-go.
// source: resume.proto
// DO NOT EDIT!

/*
Package goresume is a generated protocol buffer package.

It is generated from these files:
	resume.proto

It has these top-level messages:
	Date
	DateRange
	Resume
*/
package goresume

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Date_Month int32

const (
	Date_UNKNOWN   Date_Month = 0
	Date_JANUARY   Date_Month = 1
	Date_FEBRUARY  Date_Month = 2
	Date_MARCH     Date_Month = 3
	Date_APRIL     Date_Month = 4
	Date_MAY       Date_Month = 5
	Date_JUNE      Date_Month = 6
	Date_JULY      Date_Month = 7
	Date_AUGUST    Date_Month = 8
	Date_SEPTEMBER Date_Month = 9
	Date_OCTOBER   Date_Month = 10
	Date_NOVEMBER  Date_Month = 11
	Date_DECEMBER  Date_Month = 12
)

var Date_Month_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "JANUARY",
	2:  "FEBRUARY",
	3:  "MARCH",
	4:  "APRIL",
	5:  "MAY",
	6:  "JUNE",
	7:  "JULY",
	8:  "AUGUST",
	9:  "SEPTEMBER",
	10: "OCTOBER",
	11: "NOVEMBER",
	12: "DECEMBER",
}
var Date_Month_value = map[string]int32{
	"UNKNOWN":   0,
	"JANUARY":   1,
	"FEBRUARY":  2,
	"MARCH":     3,
	"APRIL":     4,
	"MAY":       5,
	"JUNE":      6,
	"JULY":      7,
	"AUGUST":    8,
	"SEPTEMBER": 9,
	"OCTOBER":   10,
	"NOVEMBER":  11,
	"DECEMBER":  12,
}

func (x Date_Month) String() string {
	return proto.EnumName(Date_Month_name, int32(x))
}

// The Date message representing a calendar date. A date is represented by an
// optional year, month, and day. The format in which this information is
// displayed is undefined. A date containing a day without a year is most likely
// meaningless.
type Date struct {
	Year  int32      `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	Month Date_Month `protobuf:"varint,2,opt,name=month,enum=goresume.Date_Month" json:"month,omitempty"`
	Day   uint32     `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}

// The DateRange message representing a range of calendar dates. The date range
// is represented by a beginning date and an ending date. If the ending date is
// omitted, the ending date is assumed to be the present. If only the ending
// date is present, the date range will be displayed as a single date.
type DateRange struct {
	Began *Date `protobuf:"bytes,1,opt,name=began" json:"began,omitempty"`
	Ended *Date `protobuf:"bytes,2,opt,name=ended" json:"ended,omitempty"`
}

func (m *DateRange) Reset()         { *m = DateRange{} }
func (m *DateRange) String() string { return proto.CompactTextString(m) }
func (*DateRange) ProtoMessage()    {}

func (m *DateRange) GetBegan() *Date {
	if m != nil {
		return m.Began
	}
	return nil
}

func (m *DateRange) GetEnded() *Date {
	if m != nil {
		return m.Ended
	}
	return nil
}

// The main Resume container for an individual. If a repeated section does not
// have any elements, the header will not be displayed. Any strings should
// be UTF-8 strings formatted as they are expected to be displayed in the
// resume. Rich text formatting is not supported, and should be handled by
// any translation layer that turns the resume buffer into formatted text.
type Resume struct {
	Name         string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Address      *Resume_Address       `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	EmailAddress string                `protobuf:"bytes,3,opt,name=email_address" json:"email_address,omitempty"`
	PhoneNumber  string                `protobuf:"bytes,4,opt,name=phone_number" json:"phone_number,omitempty"`
	Website      string                `protobuf:"bytes,5,opt,name=website" json:"website,omitempty"`
	Objective    *Resume_Objective     `protobuf:"bytes,6,opt,name=objective" json:"objective,omitempty"`
	Education    []*Resume_Education   `protobuf:"bytes,7,rep,name=education" json:"education,omitempty"`
	Employment   []*Resume_Employment  `protobuf:"bytes,8,rep,name=employment" json:"employment,omitempty"`
	Project      []*Resume_Project     `protobuf:"bytes,9,rep,name=project" json:"project,omitempty"`
	Publication  []*Resume_Publication `protobuf:"bytes,10,rep,name=publication" json:"publication,omitempty"`
	// The url to access a formatted version of this resume.
	FormattedUrl string `protobuf:"bytes,11,opt,name=formatted_url" json:"formatted_url,omitempty"`
}

func (m *Resume) Reset()         { *m = Resume{} }
func (m *Resume) String() string { return proto.CompactTextString(m) }
func (*Resume) ProtoMessage()    {}

func (m *Resume) GetAddress() *Resume_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Resume) GetObjective() *Resume_Objective {
	if m != nil {
		return m.Objective
	}
	return nil
}

func (m *Resume) GetEducation() []*Resume_Education {
	if m != nil {
		return m.Education
	}
	return nil
}

func (m *Resume) GetEmployment() []*Resume_Employment {
	if m != nil {
		return m.Employment
	}
	return nil
}

func (m *Resume) GetProject() []*Resume_Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *Resume) GetPublication() []*Resume_Publication {
	if m != nil {
		return m.Publication
	}
	return nil
}

// The Address message representing street address. The address is composed of
// an ordered list of strings. No assumptions about the content of the address
// are made. The address may be displayed with each string on a separate line,
// or it may be displayed on a single line with each string separated by a
// comma.
type Resume_Address struct {
	Line []string `protobuf:"bytes,1,rep,name=line" json:"line,omitempty"`
}

func (m *Resume_Address) Reset()         { *m = Resume_Address{} }
func (m *Resume_Address) String() string { return proto.CompactTextString(m) }
func (*Resume_Address) ProtoMessage()    {}

// An Education message representing time spent at an educational institute.
// Each note, if any, is presented as a bulleted list of accomplishments.
type Resume_Education struct {
	Institute string          `protobuf:"bytes,1,opt,name=institute" json:"institute,omitempty"`
	Address   *Resume_Address `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Degree    string          `protobuf:"bytes,3,opt,name=degree" json:"degree,omitempty"`
	DateRange *DateRange      `protobuf:"bytes,4,opt,name=date_range" json:"date_range,omitempty"`
	Note      []string        `protobuf:"bytes,5,rep,name=note" json:"note,omitempty"`
}

func (m *Resume_Education) Reset()         { *m = Resume_Education{} }
func (m *Resume_Education) String() string { return proto.CompactTextString(m) }
func (*Resume_Education) ProtoMessage()    {}

func (m *Resume_Education) GetAddress() *Resume_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Resume_Education) GetDateRange() *DateRange {
	if m != nil {
		return m.DateRange
	}
	return nil
}

// An Employment message representing time spent as an employee of a company.
// Each note, if any, is presented as a bulleted list of accomplishments.
type Resume_Employment struct {
	Company   string          `protobuf:"bytes,1,opt,name=company" json:"company,omitempty"`
	Title     string          `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Address   *Resume_Address `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Website   string          `protobuf:"bytes,4,opt,name=website" json:"website,omitempty"`
	DateRange *DateRange      `protobuf:"bytes,5,opt,name=date_range" json:"date_range,omitempty"`
	Note      []string        `protobuf:"bytes,6,rep,name=note" json:"note,omitempty"`
}

func (m *Resume_Employment) Reset()         { *m = Resume_Employment{} }
func (m *Resume_Employment) String() string { return proto.CompactTextString(m) }
func (*Resume_Employment) ProtoMessage()    {}

func (m *Resume_Employment) GetAddress() *Resume_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Resume_Employment) GetDateRange() *DateRange {
	if m != nil {
		return m.DateRange
	}
	return nil
}

// A Objective message representing a resume objective.
type Resume_Objective struct {
	Title    string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Contents string `protobuf:"bytes,2,opt,name=contents" json:"contents,omitempty"`
}

func (m *Resume_Objective) Reset()         { *m = Resume_Objective{} }
func (m *Resume_Objective) String() string { return proto.CompactTextString(m) }
func (*Resume_Objective) ProtoMessage()    {}

// A Project message representing a project worked on outside of employment.
// Each note, if any, is presented as a bulleted list of accomplishments.
type Resume_Project struct {
	Title     string     `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Role      string     `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	Website   string     `protobuf:"bytes,3,opt,name=website" json:"website,omitempty"`
	DateRange *DateRange `protobuf:"bytes,4,opt,name=date_range" json:"date_range,omitempty"`
	Note      []string   `protobuf:"bytes,5,rep,name=note" json:"note,omitempty"`
}

func (m *Resume_Project) Reset()         { *m = Resume_Project{} }
func (m *Resume_Project) String() string { return proto.CompactTextString(m) }
func (*Resume_Project) ProtoMessage()    {}

func (m *Resume_Project) GetDateRange() *DateRange {
	if m != nil {
		return m.DateRange
	}
	return nil
}

// A Publication message representing a single published work with multiple
// authors.
type Resume_Publication struct {
	Title   string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Date    *Date    `protobuf:"bytes,2,opt,name=date" json:"date,omitempty"`
	Url     string   `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Author  []string `protobuf:"bytes,4,rep,name=author" json:"author,omitempty"`
	Journal string   `protobuf:"bytes,5,opt,name=journal" json:"journal,omitempty"`
	Note    []string `protobuf:"bytes,6,rep,name=note" json:"note,omitempty"`
}

func (m *Resume_Publication) Reset()         { *m = Resume_Publication{} }
func (m *Resume_Publication) String() string { return proto.CompactTextString(m) }
func (*Resume_Publication) ProtoMessage()    {}

func (m *Resume_Publication) GetDate() *Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func init() {
	proto.RegisterEnum("goresume.Date_Month", Date_Month_name, Date_Month_value)
}
