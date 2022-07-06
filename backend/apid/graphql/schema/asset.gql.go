// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	errors "errors"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

//
// AssetFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Asset' type.
type AssetFieldResolvers interface {
	// ID implements response to request for 'id' field.
	ID(p graphql.ResolveParams) (string, error)

	// Namespace implements response to request for 'namespace' field.
	Namespace(p graphql.ResolveParams) (string, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)

	// Metadata implements response to request for 'metadata' field.
	Metadata(p graphql.ResolveParams) (interface{}, error)

	// Url implements response to request for 'url' field.
	Url(p graphql.ResolveParams) (string, error)

	// Sha512 implements response to request for 'sha512' field.
	Sha512(p graphql.ResolveParams) (string, error)

	// Filters implements response to request for 'filters' field.
	Filters(p graphql.ResolveParams) ([]string, error)

	// ToJSON implements response to request for 'toJSON' field.
	ToJSON(p graphql.ResolveParams) (interface{}, error)

	// Builds implements response to request for 'builds' field.
	Builds(p graphql.ResolveParams) (interface{}, error)
}

// AssetAliases implements all methods on AssetFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type AssetAliases struct{}

// ID implements response to request for 'id' field.
func (_ AssetAliases) ID(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'id'")
	}
	return ret, err
}

// Namespace implements response to request for 'namespace' field.
func (_ AssetAliases) Namespace(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'namespace'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ AssetAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// Metadata implements response to request for 'metadata' field.
func (_ AssetAliases) Metadata(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Url implements response to request for 'url' field.
func (_ AssetAliases) Url(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'url'")
	}
	return ret, err
}

// Sha512 implements response to request for 'sha512' field.
func (_ AssetAliases) Sha512(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'sha512'")
	}
	return ret, err
}

// Filters implements response to request for 'filters' field.
func (_ AssetAliases) Filters(p graphql.ResolveParams) ([]string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.([]string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'filters'")
	}
	return ret, err
}

// ToJSON implements response to request for 'toJSON' field.
func (_ AssetAliases) ToJSON(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Builds implements response to request for 'builds' field.
func (_ AssetAliases) Builds(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// AssetType Asset defines an archive, an agent will install as a dependency for a check.
var AssetType = graphql.NewType("Asset", graphql.ObjectKind)

// RegisterAsset registers Asset object type with given service.
func RegisterAsset(svc *graphql.Service, impl AssetFieldResolvers) {
	svc.RegisterObject(_ObjectTypeAssetDesc, impl)
}
func _ObjTypeAssetIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		ID(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ID(frp)
	}
}

func _ObjTypeAssetNamespaceHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Namespace(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Namespace(frp)
	}
}

func _ObjTypeAssetNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Name(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjTypeAssetMetadataHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Metadata(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Metadata(frp)
	}
}

func _ObjTypeAssetUrlHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Url(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Url(frp)
	}
}

func _ObjTypeAssetSha512Handler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Sha512(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Sha512(frp)
	}
}

func _ObjTypeAssetFiltersHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Filters(p graphql.ResolveParams) ([]string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Filters(frp)
	}
}

func _ObjTypeAssetToJSONHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		ToJSON(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ToJSON(frp)
	}
}

func _ObjTypeAssetBuildsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Builds(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Builds(frp)
	}
}

func _ObjectTypeAssetConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Asset defines an archive, an agent will install as a dependency for a check.",
		Fields: graphql1.Fields{
			"builds": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Builds defines a collection of assets that this asset can install as a dependency for a check, handler, mutator, etc. .",
				Name:              "builds",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("CoreV2AssetBuild")))),
			},
			"filters": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Filters are a collection of sensu queries, used by the system to determine\nif the asset should be installed. If more than one filter is present the\nqueries are joined by the \"AND\" operator.",
				Name:              "filters",
				Type:              graphql1.NewList(graphql1.String),
			},
			"id": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "The globally unique identifier of the record",
				Name:              "id",
				Type:              graphql1.NewNonNull(graphql1.ID),
			},
			"metadata": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "metadata contains name, namespace, labels and annotations of the record",
				Name:              "metadata",
				Type:              graphql.OutputType("ObjectMeta"),
			},
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "use metadata",
				Description:       "Name is the unique identifier for an asset",
				Name:              "name",
				Type:              graphql1.String,
			},
			"namespace": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "use metadata",
				Description:       "Namespace in which this record resides",
				Name:              "namespace",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"sha512": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Sha512 is the SHA-512 checksum of the asset",
				Name:              "sha512",
				Type:              graphql1.String,
			},
			"toJSON": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "toJSON returns a REST API compatible representation of the resource. Handy for\nsharing snippets that can then be imported with `sensuctl create`.",
				Name:              "toJSON",
				Type:              graphql1.NewNonNull(graphql.OutputType("JSON")),
			},
			"url": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "URL is the location of the asset",
				Name:              "url",
				Type:              graphql1.String,
			},
		},
		Interfaces: []*graphql1.Interface{
			graphql.Interface("Node"),
			graphql.Interface("Namespaced"),
			graphql.Interface("Resource")},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see AssetFieldResolvers.")
		},
		Name: "Asset",
	}
}

// describe Asset's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeAssetDesc = graphql.ObjectDesc{
	Config: _ObjectTypeAssetConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"builds":    _ObjTypeAssetBuildsHandler,
		"filters":   _ObjTypeAssetFiltersHandler,
		"id":        _ObjTypeAssetIDHandler,
		"metadata":  _ObjTypeAssetMetadataHandler,
		"name":      _ObjTypeAssetNameHandler,
		"namespace": _ObjTypeAssetNamespaceHandler,
		"sha512":    _ObjTypeAssetSha512Handler,
		"toJSON":    _ObjTypeAssetToJSONHandler,
		"url":       _ObjTypeAssetUrlHandler,
	},
}
