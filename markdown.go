package goresume

import (
  "strconv"
	"strings"

  pb "proto"
)

const (
	nl = "\n"
	nlnl = "\n\n"
	br = "<br>\n"
)

func resumeToMarkdown(r *pb.Resume) (out string) {
	// Name as title:
	if r.Name != "" {
		out += title("=", r.Name)
	}

	// Header info:
	if r.PhoneNumber != "" {
		out += r.PhoneNumber + br
	}
	if r.EmailAddress != "" {
		out += mailto(r.EmailAddress) + br
	}
	if r.Website != "" {
		out += web(r.Website) + br
	}
	if r.Address != nil {
		out += strings.Join(r.Address.Line, ", ")
	}
	out += br + nl

	if r.Objective != nil {
		o := r.Objective
		oTitle := o.Title
		if oTitle == "" {
			oTitle = "Objective"
		}
		out += title("-", oTitle)
		out += o.Contents + nlnl
	}

	if len(r.Education) > 0 {
		out += title("-", "Education")
		for _, e := range r.Education {
			out += bold(e.Institute) + br
			out += e.Degree + br
			out += dateRange(e.DateRange) + br
			out += bulletList(e.Note) + nl
		}
	}

	if len(r.Employment) > 0 {
		out += title("-", "Employment")
		for _, e := range r.Employment {
			out += bold(e.Company) + br
			out += e.Title + br
			out += dateRange(e.DateRange) + br
			out += bulletList(e.Note) + nl
		}
	}

	if len(r.Project) > 0 {
		out += title("-", "Projects")
		for _, p := range r.Project {
			out += bold(p.Title) + br
			out += p.Role + br
			out += dateRange(p.DateRange) + br
			if p.Website != "" {
				out += web(p.Website) + br
			}
			out += bulletList(p.Note) + nl
		}
	}

	if len(r.Publication) > 0 {
		out += title("-", "Publications")
		for _, p := range r.Publication {
			out += bold(p.Title) + br
			out += andList(p.Author) + br
			if p.Date != nil {
				out += date(p.Date) + br
			}
			if p.Url != "" {
				out += web(p.Url) + br
			}
			out += nl
			out += bulletList(p.Note)
		}
	}
	return
}

func title(sep, t string) (out string) {
	out += t + nl
	out += strings.Repeat(sep, len(t)) + nlnl
	return
}

func bulletList(l []string) (out string) {
	out += nl
	for _, e := range l {
		out += "*   " + e + nl
	}
	return
}

func andList(l []string) string {
	if len(l) == 0 {
		return ""
	} else if len(l) == 1 {
		return l[0]
	}
	return strings.Join(l[:len(l)-1], ", ") + " and " + l[len(l)-1]
}

func date(d *pb.Date) string {
	if d == nil || (d.Year == 0 && d.Month == pb.Date_UNKNOWN) {
		return "Present"
	}
	var s string
	if d.Month != pb.Date_UNKNOWN {
		s += strings.Title(strings.ToLower(d.Month.String()))
		if d.Day != 0 {
			s += " " + strconv.FormatUint(uint64(d.Day), 10)
		}
	}
	if d.Year != 0 {
		if len(s) > 0 {
			s += " "
		}
		s += strconv.FormatUint(uint64(d.Year), 10)
	}
	return s
}

func dateRange(dr *pb.DateRange) string {
	return date(dr.Began) + " \u2013 " + date(dr.Ended)
}

func code(s string) string {
	return "`" + s + "`"
}

func bold(s string) string {
	return "__" + s + "__"
}

func mailto(e string) string {
	return "[" + code(e) + "](mailto:" + e + ")" 
}
	
func web(u string) string {
	return "[" + code(u) + "](" + u + ")"
}
