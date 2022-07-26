package scopes

import "github.com/gobuffalo/pop/v6"

func UserScope(userId string) pop.ScopeFunc {
	return func(q *pop.Query) *pop.Query {
		return q.Where("user_id = ?", userId)
	}
}

func ClusterScope(clusterId string) pop.ScopeFunc {
	return func(q *pop.Query) *pop.Query {
		return q.Where("cluster_id = ?", clusterId)
	}
}
