package repr

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// boolPtr returns a *bool for the given value
func boolPtr(b bool) *bool { return &b }

// intPtr returns a *int for the given value
func intPtr(i int) *int { return &i }

// float32Ptr returns a *float for the given value
func float32Ptr(f float32) *float32 { return &f }

// float64Ptr returns a *float for the given value
func float64Ptr(f float64) *float64 { return &f }

// stringPtr returns a *string for the given value
func stringPtr(s string) *string { return &s }

// timePtr returns a *string for the given value
func timePtr(t time.Time) *time.Time { return &t }

type Project struct {
	Deprecated               *bool      `json:"deprecated,omitempty"`
	Description              *string    `json:"description,omitempty"`
	Forks                    *int       `json:"forks,omitempty"`
	Keywords                 []*string  `json:"keywords,omitempty"`
	LatestReleasePublishedAt *time.Time `json:"latest_release_published_at,omitempty"`
	LatestStableRelease      *Release   `json:"latest_stable_release,omitempty"`
	Versions                 []*Release `json:"versions,omitempty"`

	// This is not a pointer to cover this case as well in the tests
	Name string `json:"name,omitempty"`
}

type Release struct {
	Number      *string    `json:"number,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

func TestString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"helloworld", `"helloworld"`},
		{"G'Day Mate", `"G'Day Mate"`},
		{"1234", `"1234"`},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			w := &bytes.Buffer{}
			v := reflect.ValueOf(tc.s)

			String(w, v)

			if got := w.String(); got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	m := map[string]bool{"hello": true, "world": false}

	// Map keys come back in unspecified order
	// So we check that either want1 or want2 match
	want1 := `map["hello":true "world":false]`
	want2 := `map["world":false "hello":true]`

	w := &bytes.Buffer{}
	Map(w, reflect.ValueOf(m))

	if got := w.String(); got != want1 && got != want2 {
		t.Errorf("Map() = %v, want %v or %v", got, want1, want2)
	}
}

func TestSlice_bool(t *testing.T) {
	s := []bool{true, false, true}
	want := `[true false true]`

	w := &bytes.Buffer{}
	Slice(w, reflect.ValueOf(s))

	if got := w.String(); got != want {
		t.Errorf("Slice() = %v, want %v", got, want)
	}
}

func TestSlice_string(t *testing.T) {
	s := []string{"Hello", "World", "Foo", "Bar"}
	want := `["Hello" "World" "Foo" "Bar"]`

	w := &bytes.Buffer{}
	Slice(w, reflect.ValueOf(s))

	if got := w.String(); got != want {
		t.Errorf("Slice() = %v, want %v", got, want)
	}
}

func Test_toString(t *testing.T) {
	tests := []struct {
		name string
		obj  interface{}
		want string
	}{
		{"string", "helloworld", `"helloworld"`},
		{"stringPointer", stringPtr("helloworld"), `"helloworld"`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			v := reflect.ValueOf(tc.obj)

			toString(w, v)
			if got := w.String(); got != tc.want {
				t.Errorf("toString() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestTime(t *testing.T) {
	d := time.Date(2017, 01, 02, 15, 04, 05, 0, time.UTC)
	want := "{2017-01-02 15:04:05 +0000 UTC}"

	w := &bytes.Buffer{}
	Time(w, reflect.ValueOf(d))

	if got := w.String(); got != want {
		t.Errorf("Time() = %v, want %v", got, want)
	}
}

func TestStruct(t *testing.T) {
	r := &Release{
		Number:      stringPtr("3.0.6"),
		PublishedAt: timePtr(time.Date(2017, 01, 02, 15, 04, 05, 0, time.UTC)),
	}
	p := &Project{
		Deprecated:               boolPtr(true),
		Description:              stringPtr("Python testing framework"),
		Forks:                    intPtr(350),
		Keywords:                 []*string{stringPtr("Python"), stringPtr("testing"), stringPtr("pytest")},
		LatestReleasePublishedAt: timePtr(time.Date(2017, 01, 02, 15, 04, 05, 0, time.UTC)),
		LatestStableRelease:      r,
		Versions:                 []*Release{r},
		Name:                     "pytest",
	}
	want := `repr.Project{` +
		`Deprecated:true, ` +
		`Description:"Python testing framework", ` +
		`Forks:350, ` +
		`Keywords:["Python" "testing" "pytest"], ` +
		`LatestReleasePublishedAt:time.Time{2017-01-02 15:04:05 +0000 UTC}, ` +
		`LatestStableRelease:repr.Release{Number:"3.0.6", PublishedAt:time.Time{2017-01-02 15:04:05 +0000 UTC}}, ` +
		`Versions:[repr.Release{Number:"3.0.6", PublishedAt:time.Time{2017-01-02 15:04:05 +0000 UTC}}], ` +
		`Name:"pytest"}`

	w := &bytes.Buffer{}
	Struct(w, reflect.ValueOf(p).Elem())

	if got := w.String(); got != want {
		t.Errorf("Struct() = %v, want %v", got, want)
	}
}

func TestRepr(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want string
	}{
		{"string", "helloworld", `"helloworld"`},
		{"string pointer", stringPtr("G'Day Mate"), `"G'Day Mate"`},
		{"bool", false, `false`},
		{"bool pointer", boolPtr(true), `true`},
		{"int", 1234, `1234`},
		{"int pointer", intPtr(404), `404`},
		{"float32", 3.14, `3.14`},
		{"float32 pointer", float32Ptr(3.14159), `3.14159`},
		{"float64", 3.14, `3.14`},
		{"float64 pointer", float64Ptr(3.14159), `3.14159`},
		{"map", map[string]string{"hello": "world"}, `map["hello":"world"]`},
		{"map pointer", map[*string]*bool{stringPtr("a"): boolPtr(true)}, `map["a":true]`},
		{"slice", []string{"gopherize", "me"}, `["gopherize" "me"]`},
		{"slice pointer", []*int{intPtr(1), intPtr(23)}, `[1 23]`},
		{"nil", nil, `<nil>`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Repr(tc.args); got != tc.want {
				t.Errorf("Repr() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRepr_struct(t *testing.T) {
	type Repo struct {
		URL *string
	}
	type Maintainer struct {
		AuthorName  *string
		Email       *string
		SocialLinks map[string]string
	}
	type Release struct {
		Number       string
		PublishedAt  *time.Time
		Contributors []*string
	}
	type Project struct {
		Name                *string
		forks               *int
		Keywords            []*string
		LatestStableRelease *Release
		repo                *Repo
		*Maintainer
	}
	p := &Project{
		stringPtr("pytest"),
		intPtr(123),
		[]*string{stringPtr("testing"), stringPtr("test"), nil},
		&Release{
			"3.0.6",
			timePtr(time.Date(2017, 01, 02, 15, 04, 05, 0, time.UTC)),
			nil,
		},
		&Repo{},
		&Maintainer{
			AuthorName:  stringPtr("Brianna"),
			SocialLinks: map[string]string{"twitter": "hackebrot"},
			Email:       nil,
		},
	}
	want := `repr.Project{` +
		`Name:"pytest", ` +
		`Keywords:["testing" "test" <nil>], ` +
		`LatestStableRelease:repr.Release{Number:"3.0.6", PublishedAt:time.Time{2017-01-02 15:04:05 +0000 UTC}}, ` +
		`Maintainer:repr.Maintainer{AuthorName:"Brianna", ` +
		`SocialLinks:map["twitter":"hackebrot"]}}`

	if got := Repr(p); got != want {
		t.Errorf("\nRepr() %8v\nwant   %8v", got, want)
	}
}
