package visitors

import (
	"fmt"
	"io"
	"mydemos/mycmd/pkg/visitors/unstructured"
	"mydemos/mycmd/pkg/visitors/yaml"
)

// Info contains temporary info to execute a REST call, or show the results
// of an already completed REST call.
type Info struct {
	// Client will only be present if this builder was not local
	//Client RESTClient
	// Mapping will only be present if this builder was not local
	//Mapping *meta.RESTMapping

	// Namespace will be set if the object is namespaced and has a specified value.
	Namespace string
	Name      string

	// Optional, Source is the filename or URL to template file (.json or .yaml),
	// or stdin to use to handle the resource
	Source string
	// Optional, this is the most recent value returned by the server if available. It will
	// typically be in unstructured or internal forms, depending on how the Builder was
	// defined. If retrieved from the server, the Builder expects the mapping client to
	// decide the final form. Use the AsVersioned, AsUnstructured, and AsInternal helpers
	// to alter the object versions.
	// If Subresource is specified, this will be the object for the subresource.
	//Object runtime.Object
	// Optional, this is the most recent resource version the server knows about for
	// this type of resource. It may not match the resource version of the object,
	// but if set it should be equal to or newer than the resource version of the
	// object (however the server defines resource version).
	ResourceVersion string
	// Optional, if specified, the object is the most recent value of the subresource
	// returned by the server if available.
	Subresource string
}

type VisitorFunc func(*Info, error) error

// Visitor lets clients walk a list of resources.
type Visitor interface {
	Visit(VisitorFunc) error
}

type StreamVisitor struct {
	io.Reader
	//*mapper

	Source string
	//Schema ContentValidator
}

// NewStreamVisitor is a helper function that is useful when we want to change the fields of the struct but keep calls the same.
func NewStreamVisitor(r io.Reader, source string) *StreamVisitor {
	return &StreamVisitor{
		Reader: r,
		//mapper: mapper,
		Source: source,
		//Schema: schema,
	}
}

//type RawExtension struct {
//	// Raw is the underlying serialization of this object.
//	//
//	// TODO: Determine how to detect ContentType and ContentEncoding of 'Raw' data.
//	Raw []byte `json:"-" protobuf:"bytes,1,opt,name=raw"`
//	// Object can hold a representation of this extension - useful for working with versioned
//	// structs.
//	Object Object `json:"-"`
//}

// Visit implements Visitor over a stream. StreamVisitor is able to distinct multiple resources in one stream.
func (v *StreamVisitor) Visit(fn VisitorFunc) error {
	d := yaml.NewYAMLOrJSONDecoder(v.Reader, 4096)
	for {
		ext := unstructured.Unstructured{}
		if err := d.Decode(&ext.Object); err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("error parsing %s: %v", v.Source, err)
		}
		// TODO: This needs to be able to handle object in other encodings and schemas.
		//ext.Raw = bytes.TrimSpace(ext.Raw)
		//if len(ext.Raw) == 0 || bytes.Equal(ext.Raw, []byte("null")) {
		//	continue
		//}
		fmt.Printf("object string: %#v\n\n", ext)
		//if err := ValidateSchema(ext.Raw, v.Schema); err != nil {
		//	return fmt.Errorf("error validating %q: %v", v.Source, err)
		//}
		//info, err := v.infoForData(ext.Raw, v.Source)
		//if err != nil {
		//	if fnErr := fn(info, err); fnErr != nil {
		//		return fnErr
		//	}
		//	continue
		//}
		//if err := fn(info, nil); err != nil {
		//	return err
		//}
	}
}
