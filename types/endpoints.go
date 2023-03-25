package types

type Endpoint struct {
	Path      string
	ApiPath   string // Path from the OpenAPI spec
	Operation string
}

func (e *Endpoint) GetApiPath() string {
	if e.ApiPath == "" {
		return e.Path
	}
	return e.ApiPath
}

var Endpoints = map[string]Endpoint{
	"Scan": {
		Path:      "/v1/scan",
		Operation: "POST",
	},
	"ScanMultiple": {
		Path:      "/v1/multiscan",
		Operation: "POST",
	},
	"ScanQuotas": {
		Path:      "/v1/quotas",
		Operation: "GET",
	},
	"MembersList": {
		Path:      "/v1/members",
		Operation: "GET",
	},
	"ListTeamMembership": {
		Path:      "/v1/members/%d/team_memberships",
		ApiPath:   "/v1/members/{member_id}/team_memberships",
		Operation: "GET",
	},
	"MembersGet": {
		Path:      "/v1/members/%d",
		ApiPath:   "/v1/members/{member_id}",
		Operation: "GET",
	},
	"AuditLogsList": {
		Path:      "/v1/audit_logs",
		Operation: "GET",
	},
	"UpdateNote": {
		Path:      "/v1/incidents/secrets/%d/notes/%d",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/notes/{note_id}",
		Operation: "PATCH",
	},
	"GetSecretIncidents": {
		Path:      "/v1/incidents/secrets/%d",
		ApiPath:   "/v1/incidents/secrets/{incident_id}",
		Operation: "GET",
	},
	"UpdateSecretIncidents": {
		Path:      "/v1/incidents/secrets/%d",
		ApiPath:   "/v1/incidents/secrets/{incident_id}",
		Operation: "PATCH",
	},
	"UnshareSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/unshare",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/unshare",
		Operation: "POST",
	},
	"UnassignSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/unassign",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/unassign",
		Operation: "POST",
	},
	"ShareSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/share",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/share",
		Operation: "POST",
	},
	"RevokeAccessSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/revoke_access",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/revoke_access",
		Operation: "POST",
	},
	"ResolveSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/resolve",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/resolve",
		Operation: "POST",
	},
	"ReopenSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/reopen",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/reopen",
		Operation: "POST",
	},
	"ListSecretMembers": {
		Path:      "/v1/incidents/secrets/%d/members",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/members",
		Operation: "GET",
	},
	"ListOccurrences": {
		Path:      "/v1/occurrences/secrets",
		Operation: "GET",
	},
	"ListNotes": {
		Path:      "/v1/incidents/secrets/%d/notes",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/notes",
		Operation: "GET",
	},
	"IgnoreSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/ignore",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/ignore",
		Operation: "POST",
	},
	"GrantAccessSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/grant_access",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/grant_access",
		Operation: "POST",
	},
	"DeleteNote": {
		Path:      "/v1/incidents/secrets/%d/notes/%d",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/notes/{note_id}",
		Operation: "DELETE",
	},
	"CreateNote": {
		Path:      "/v1/incidents/secrets/%d/notes",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/notes",
		Operation: "POST",
	},
	"AssignSecretIncident": {
		Path:      "/v1/incidents/secrets/%d/assign",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/assign",
		Operation: "POST",
	},
	"InvitationsResend": {
		Path:      "/v1/invitations/%d/resend",
		ApiPath:   "/v1/invitations/{invitation_id}/resend",
		Operation: "POST",
	},
	"InvitationsList": {
		Path:      "/v1/invitations",
		Operation: "GET",
	},
	"InvitationsDelete": {
		Path:      "/v1/invitations/%d",
		ApiPath:   "/v1/invitations/{invitation_id}",
		Operation: "DELETE",
	},
	"InvitationsCreate": {
		Path:      "/v1/invitations",
		Operation: "POST",
	},
	"SourcesList": {
		Path:      "/v1/sources",
		Operation: "GET",
	},
	"SourcesGet": {
		Path:      "/v1/sources/%d",
		ApiPath:   "/v1/sources/{source_id}",
		Operation: "GET",
	},
	"Health": {
		Path:      "/v1/health",
		Operation: "GET",
	},
	"TeamsUpdatePerimeter": {
		Path:      "/v1/teams/%d/sources",
		ApiPath:   "/v1/teams/{team_id}/sources",
		Operation: "POST",
	},
	"TeamsUpdate": {
		Path:      "/v1/teams/%d",
		ApiPath:   "/v1/teams/{team_id}",
		Operation: "PATCH",
	},
	"TeamsListMemberships": {
		Path:      "/v1/teams/%d/team_memberships",
		ApiPath:   "/v1/teams/{team_id}/team_memberships",
		Operation: "GET",
	},
	"TeamsList": {
		Path:      "/v1/teams",
		Operation: "GET",
	},
	"TeamsGet": {
		Path:      "/v1/teams/%d",
		ApiPath:   "/v1/teams/{team_id}",
		Operation: "GET",
	},
	"TeamsDelete": {
		Path:      "/v1/teams/%s",
		ApiPath:   "/v1/teams/{team_id}",
		Operation: "DELETE",
	},
	"TeamsCreate": {
		Path:      "/v1/teams",
		Operation: "POST",
	},
	"TeamsAddMember": {
		Path:      "/v1/teams/%d/team_memberships",
		ApiPath:   "/v1/teams/{team_id}/team_memberships",
		Operation: "POST",
	},
	"ListSecretIncidents": {
		Path:      "/v1/incidents/secrets",
		Operation: "GET",
	},
	"ListSecretInvitations": {
		Path:      "/v1/incidents/secrets/%d/invitations",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/invitations",
		Operation: "GET",
	},
	"ListSecretTeams": {
		Path:      "/v1/incidents/secrets/%d/teams",
		ApiPath:   "/v1/incidents/secrets/{incident_id}/teams",
		Operation: "GET",
	},
}
