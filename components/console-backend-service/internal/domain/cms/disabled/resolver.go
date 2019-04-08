// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import context "context"
import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

// Resolver is an autogenerated failing mock type for the Resolver type
type Resolver struct {
	err error
}

// NewResolver creates a new Resolver type instance
func NewResolver(err error) *Resolver {
	return &Resolver{err: err}
}

// ClusterDocsTopicAssetsField provides a failing mock function with given fields: ctx, obj, types
func (_m *Resolver) ClusterDocsTopicAssetsField(ctx context.Context, obj *gqlschema.ClusterDocsTopic, types []string) ([]gqlschema.ClusterAsset, error) {
	var r0 []gqlschema.ClusterAsset
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ClusterDocsTopicEventSubscription provides a failing mock function with given fields: ctx
func (_m *Resolver) ClusterDocsTopicEventSubscription(ctx context.Context) (<-chan gqlschema.ClusterDocsTopicEvent, error) {
	var r0 <-chan gqlschema.ClusterDocsTopicEvent
	var r1 error
	r1 = _m.err

	return r0, r1
}

// ClusterDocsTopicsQuery provides a failing mock function with given fields: ctx, viewContext, groupName
func (_m *Resolver) ClusterDocsTopicsQuery(ctx context.Context, viewContext *string, groupName *string) ([]gqlschema.ClusterDocsTopic, error) {
	var r0 []gqlschema.ClusterDocsTopic
	var r1 error
	r1 = _m.err

	return r0, r1
}

// DocsTopicAssetsField provides a failing mock function with given fields: ctx, obj, types
func (_m *Resolver) DocsTopicAssetsField(ctx context.Context, obj *gqlschema.DocsTopic, types []string) ([]gqlschema.Asset, error) {
	var r0 []gqlschema.Asset
	var r1 error
	r1 = _m.err

	return r0, r1
}

// DocsTopicEventSubscription provides a failing mock function with given fields: ctx, namespace
func (_m *Resolver) DocsTopicEventSubscription(ctx context.Context, namespace string) (<-chan gqlschema.DocsTopicEvent, error) {
	var r0 <-chan gqlschema.DocsTopicEvent
	var r1 error
	r1 = _m.err

	return r0, r1
}