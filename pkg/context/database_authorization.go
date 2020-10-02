package context

import containersv1alpha1 "github.com/Kafei59/bol/api/v1alpha1"

var dbAuthContext = "_context_database_authorization"

// DatabaseAuthorization returns a kubernetes db auth resource from the context
func DatabaseAuthorization(ctx Context) *containersv1alpha1.DatabaseAuthorization {
	return ctx.Value(&dbAuthContext).(*containersv1alpha1.DatabaseAuthorization)
}

// WithDatabaseAuthorization inserts a kubernetes db auth resource into the context
func WithDatabaseAuthorization(ctx Context, auth *containersv1alpha1.DatabaseAuthorization) Context {
	return WithValue(ctx, &dbAuthContext, auth)
}
