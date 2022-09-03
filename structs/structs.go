package structs

import (
	"time"
)

type YamlStruct struct {
	Override              *[]OverrideStruct `yaml:"Override,omitempty"`
	OrganizationStandards *[]OrgStandards   `yaml:"OrganizationStandards,omitempty"`
	Organizations         []string          `yaml:"Organizations,omitempty"`
}

type ChangeLogStruct struct {
	Type string
	Path string
	From string
	To   string
}

type TeamClassifier struct {
	Slug        string           `yaml:"Slug,omitempty"`
	Permissions *RepoPermissions `yaml:"Permissions,omitempty"`
}

type RepoPermissions struct {
	Admin    *bool `yaml:"admin,omitempty"`
	Maintain *bool `yaml:"maintain,omitempty"`
	Pull     *bool `yaml:"pull,omitempty"`
	Push     *bool `yaml:"push,omitempty"`
	Triage   *bool `yaml:"triage,omitempty"`
}

type OrgStandards struct {
	Organization string               `yaml:"Organization,omitempty"`
	Repository   *Repository          `yaml:"Repository,omitempty"`
	Branches     []BranchesClassifier `yaml:"Branches,omitempty"`
	Team         []TeamClassifier     `yaml:"Team,omitempty"`
}

type OverrideStruct struct {
	Organization string               `yaml:"Organization,omitempty"`
	RepoName     string               `yaml:"repoName,omitempty"`
	Repository   *Repository          `yaml:"Repository,omitempty"`
	Protection   *Protection          `yaml:"Protection,omitempty"`
	Team         []TeamClassifier     `yaml:"Team,omitempty"`
	Branches     []BranchesClassifier `yaml:"Branches,omitempty"`
}

type Convert struct {
	RequiredStatusChecks           *RequiredStatusChecks           `yaml:"requiredstatuschecks,omitempty"`
	RequiredPullRequestReviews     *PullRequestReviewsEnforcement  `yaml:"requiredpullrequestreviews,omitempty"`
	EnforceAdmins                  *AdminEnforcement               `yaml:"enforceadmins,omitempty"`
	Restrictions                   *BranchRestrictions             `yaml:"restrictions,omitempty"`
	RequireLinearHistory           *RequireLinearHistory           `yaml:"requiredlinearhistory,omitempty"`
	AllowForcePushes               *AllowForcePushes               `yaml:"allowforcepushes,omitempty"`
	AllowDeletions                 *AllowDeletions                 `yaml:"allowdeletions,omitempty"`
	RequiredConversationResolution *RequiredConversationResolution `yaml:"requiredconversationresolution,omitempty"`
}

type BranchesClassifier struct {
	BranchName string      `yaml:"branchName,omitempty"`
	Protection *Protection `yaml:"Protection,omitempty"`
}

type EmptyProtection struct {
	RequiredStatusChecks           *RequiredStatusChecks           `yaml:"required_status_checks"`
	RequiredPullRequestReviews     *PullRequestReviewsEnforcement  `yaml:"required_pull_request_reviews"`
	EnforceAdmins                  *AdminEnforcement               `yaml:"enforce_admins"`
	Restrictions                   *BranchRestrictions             `yaml:"restrictions"`
	RequireLinearHistory           *RequireLinearHistory           `yaml:"required_linear_history"`
	AllowForcePushes               *AllowForcePushes               `yaml:"allow_force_pushes"`
	AllowDeletions                 *AllowDeletions                 `yaml:"allow_deletions"`
	RequiredConversationResolution *RequiredConversationResolution `yaml:"required_conversation_resolution"`
}

type Protection struct {
	RequiredStatusChecks           *RequiredStatusChecks           `yaml:"required_status_checks,omitempty"`
	RequiredPullRequestReviews     *PullRequestReviewsEnforcement  `yaml:"required_pull_request_reviews,omitempty"`
	EnforceAdmins                  *AdminEnforcement               `yaml:"enforce_admins,omitempty"`
	Restrictions                   *BranchRestrictions             `yaml:"restrictions,omitempty"`
	RequireLinearHistory           *RequireLinearHistory           `yaml:"required_linear_history,omitempty"`
	AllowForcePushes               *AllowForcePushes               `yaml:"allow_force_pushes,omitempty"`
	AllowDeletions                 *AllowDeletions                 `yaml:"allow_deletions,omitempty"`
	RequiredConversationResolution *RequiredConversationResolution `yaml:"required_conversation_resolution,omitempty"`
}

type Hook struct {
	CreatedAt    *time.Time             `yaml:"created_at,omitempty"`
	UpdatedAt    *time.Time             `yaml:"updated_at,omitempty"`
	URL          *string                `yaml:"url,omitempty"`
	ID           *int64                 `yaml:"id,omitempty"`
	Type         *string                `yaml:"type,omitempty"`
	Name         *string                `yaml:"name,omitempty"`
	TestURL      *string                `yaml:"test_url,omitempty"`
	PingURL      *string                `yaml:"ping_url,omitempty"`
	LastResponse map[string]interface{} `yaml:"last_response,omitempty"`

	// Only the following fields are used when creating a hook.
	// Config is required.
	Config map[string]interface{} `yaml:"config,omitempty"`
	Events []string               `yaml:"events,omitempty"`
	Active *bool                  `yaml:"active"`
}

type AdminEnforcement struct {
	URL     *string `yaml:"url,omitempty"`
	Enabled *bool   `yaml:"enabled,omitempty"`
}

type RequiredStatusChecks struct {
	// Require branches to be up to date before merging. (Required.)
	Strict *bool `yaml:"strict,omitempty"`
	// The list of status checks to require in order to merge into this
	// branch. (Deprecated. Note: only one of Contexts/Checks can be populated,
	// but at least one must be populated).
	Contexts []string `yaml:"contexts,omitempty"`
	// The list of status checks to require in order to merge into this
	// branch.
	Checks []*RequiredStatusCheck `yaml:"checks,omitempty"`
}

type RequiredStatusCheck struct {
	// The name of the required check.
	Context string `yaml:"context,omitempty"`
	// The ID of the GitHub App that must provide this check.
	// Omit this field to automatically select the GitHub App
	// that has recently provided this check,
	// or any app if it was not set by a GitHub App.
	// Pass -1 to explicitly allow any app to set the status.
	AppID *int64 `yaml:"app_id,omitempty"`
}

type PullRequestReviewsEnforcement struct {
	// Specifies which users and teams can dismiss pull request reviews.
	DismissalRestrictions *DismissalRestrictions `yaml:"dismissal_restrictions,omitempty"`
	// Specifies if approved reviews are dismissed automatically, when a new commit is pushed.
	DismissStaleReviews *bool `yaml:"dismiss_stale_reviews,omitempty"`
	// RequireCodeOwnerReviews specifies if an approved review is required in pull requests including files with a designated code owner.
	RequireCodeOwnerReviews *bool `yaml:"require_code_owner_reviews,omitempty"`
	// RequiredApprovingReviewCount specifies the number of approvals required before the pull request can be merged.
	// Valid values are 1-6.
	RequiredApprovingReviewCount int `yaml:"required_approving_review_count,omitempty"`
}

type DismissalRestrictions struct {
	// The list of users who can dimiss pull request reviews.
	Users []*User `yaml:"users,omitempty"`
	// The list of teams which can dismiss pull request reviews.
	Teams []*Team `yaml:"teams,omitempty"`
}

type User struct {
	Login                   *string    `yaml:"login,omitempty"`
	ID                      *int64     `yaml:"id,omitempty"`
	NodeID                  *string    `yaml:"node_id,omitempty"`
	AvatarURL               *string    `yaml:"avatar_url,omitempty"`
	HTMLURL                 *string    `yaml:"html_url,omitempty"`
	GravatarID              *string    `yaml:"gravatar_id,omitempty"`
	Name                    *string    `yaml:"name,omitempty"`
	Company                 *string    `yaml:"company,omitempty"`
	Blog                    *string    `yaml:"blog,omitempty"`
	Location                *string    `yaml:"location,omitempty"`
	Email                   *string    `yaml:"email,omitempty"`
	Hireable                *bool      `yaml:"hireable,omitempty"`
	Bio                     *string    `yaml:"bio,omitempty"`
	TwitterUsername         *string    `yaml:"twitter_username,omitempty"`
	PublicRepos             *int       `yaml:"public_repos,omitempty"`
	PublicGists             *int       `yaml:"public_gists,omitempty"`
	Followers               *int       `yaml:"followers,omitempty"`
	Following               *int       `yaml:"following,omitempty"`
	CreatedAt               *Timestamp `yaml:"created_at,omitempty"`
	UpdatedAt               *Timestamp `yaml:"updated_at,omitempty"`
	SuspendedAt             *Timestamp `yaml:"suspended_at,omitempty"`
	Type                    *string    `yaml:"type,omitempty"`
	SiteAdmin               *bool      `yaml:"site_admin,omitempty"`
	TotalPrivateRepos       *int       `yaml:"total_private_repos,omitempty"`
	OwnedPrivateRepos       *int       `yaml:"owned_private_repos,omitempty"`
	PrivateGists            *int       `yaml:"private_gists,omitempty"`
	DiskUsage               *int       `yaml:"disk_usage,omitempty"`
	Collaborators           *int       `yaml:"collaborators,omitempty"`
	TwoFactorAuthentication *bool      `yaml:"two_factor_authentication,omitempty"`
	Plan                    *Plan      `yaml:"plan,omitempty"`
	LdapDn                  *string    `yaml:"ldap_dn,omitempty"`

	// API URLs
	URL               *string `yaml:"url,omitempty"`
	EventsURL         *string `yaml:"events_url,omitempty"`
	FollowingURL      *string `yaml:"following_url,omitempty"`
	FollowersURL      *string `yaml:"followers_url,omitempty"`
	GistsURL          *string `yaml:"gists_url,omitempty"`
	OrganizationsURL  *string `yaml:"organizations_url,omitempty"`
	ReceivedEventsURL *string `yaml:"received_events_url,omitempty"`
	ReposURL          *string `yaml:"repos_url,omitempty"`
	StarredURL        *string `yaml:"starred_url,omitempty"`
	SubscriptionsURL  *string `yaml:"subscriptions_url,omitempty"`

	// Permissions and RoleName identify the permissions and role that a user has on a given
	// repository. These are only populated when calling Repositories.ListCollaborators.
	Permissions map[string]bool `yaml:"permissions"`
	RoleName    *string         `yaml:"role_name,omitempty"`
}

type Team struct {
	ID          *int64  `yaml:"id,omitempty"`
	NodeID      *string `yaml:"node_id,omitempty"`
	Name        *string `yaml:"name,omitempty"`
	Description *string `yaml:"description,omitempty"`
	URL         *string `yaml:"url,omitempty"`
	Slug        *string `yaml:"Slug,omitempty"`

	// Permission specifies the default permission for repositories owned by the team.
	Permission *string `yaml:"permission,omitempty"`

	// Permissions identifies the permissions that a team has on a given
	// repository. This is only populated when calling Repositories.ListTeams.
	Permissions map[string]bool `yaml:"permissions,omitempty"`

	// Privacy identifies the level of privacy this team should have.
	// Possible values are:
	//     secret - only visible to organization owners and members of this team
	//     closed - visible to all members of this organization
	// Default is "secret".
	Privacy *string `yaml:"privacy,omitempty"`

	MembersCount    *int          `yaml:"members_count,omitempty"`
	ReposCount      *int          `yaml:"repos_count,omitempty"`
	Organization    *Organization `yaml:"organization,omitempty"`
	HTMLURL         *string       `yaml:"html_url,omitempty"`
	MembersURL      *string       `yaml:"members_url,omitempty"`
	RepositoriesURL *string       `yaml:"repositories_url,omitempty"`
	Parent          *Team         `yaml:"parent,omitempty"`

	// LDAPDN is only available in GitHub Enterprise and when the team
	// membership is synchronized with LDAP.
	LDAPDN *string `yaml:"ldap_dn,omitempty"`
}

type Timestamp struct {
	time.Time
}

type Plan struct {
	Name          *string `yaml:"name,omitempty"`
	Space         *int    `yaml:"space,omitempty"`
	Collaborators *int    `yaml:"collaborators,omitempty"`
	PrivateRepos  *int    `yaml:"private_repos,omitempty"`
	FilledSeats   *int    `yaml:"filled_seats,omitempty"`
	Seats         *int    `yaml:"seats,omitempty"`
}

type Organization struct {
	Login                       *string    `yaml:"login,omitempty"`
	ID                          *int64     `yaml:"id,omitempty"`
	NodeID                      *string    `yaml:"node_id,omitempty"`
	AvatarURL                   *string    `yaml:"avatar_url,omitempty"`
	HTMLURL                     *string    `yaml:"html_url,omitempty"`
	Name                        *string    `yaml:"name,omitempty"`
	Company                     *string    `yaml:"company,omitempty"`
	Blog                        *string    `yaml:"blog,omitempty"`
	Location                    *string    `yaml:"location,omitempty"`
	Email                       *string    `yaml:"email,omitempty"`
	TwitterUsername             *string    `yaml:"twitter_username,omitempty"`
	Description                 *string    `yaml:"description,omitempty"`
	PublicRepos                 *int       `yaml:"public_repos,omitempty"`
	PublicGists                 *int       `yaml:"public_gists,omitempty"`
	Followers                   *int       `yaml:"followers,omitempty"`
	Following                   *int       `yaml:"following,omitempty"`
	CreatedAt                   *time.Time `yaml:"created_at,omitempty"`
	UpdatedAt                   *time.Time `yaml:"updated_at,omitempty"`
	TotalPrivateRepos           *int       `yaml:"total_private_repos,omitempty"`
	OwnedPrivateRepos           *int       `yaml:"owned_private_repos,omitempty"`
	PrivateGists                *int       `yaml:"private_gists,omitempty"`
	DiskUsage                   *int       `yaml:"disk_usage,omitempty"`
	Collaborators               *int       `yaml:"collaborators,omitempty"`
	BillingEmail                *string    `yaml:"billing_email,omitempty"`
	Type                        *string    `yaml:"type,omitempty"`
	Plan                        *Plan      `yaml:"plan,omitempty"`
	TwoFactorRequirementEnabled *bool      `yaml:"two_factor_requirement_enabled,omitempty"`
	IsVerified                  *bool      `yaml:"is_verified,omitempty"`
	HasOrganizationProjects     *bool      `yaml:"has_organization_projects,omitempty"`
	HasRepositoryProjects       *bool      `yaml:"has_repository_projects,omitempty"`

	// DefaultRepoPermission can be one of: "read", "write", "admin", or "none". (Default: "read").
	// It is only used in OrganizationsService.Edit.
	DefaultRepoPermission *string `yaml:"default_repository_permission,omitempty"`
	// DefaultRepoSettings can be one of: "read", "write", "admin", or "none". (Default: "read").
	// It is only used in OrganizationsService.Get.
	DefaultRepoSettings *string `yaml:"default_repository_settings,omitempty"`

	// MembersCanCreateRepos default value is true and is only used in Organizations.Edit.
	MembersCanCreateRepos *bool `yaml:"members_can_create_repositories,omitempty"`

	// https://developer.github.com/changes/2019-12-03-internal-visibility-changes/#rest-v3-api
	MembersCanCreatePublicRepos   *bool `yaml:"members_can_create_public_repositories,omitempty"`
	MembersCanCreatePrivateRepos  *bool `yaml:"members_can_create_private_repositories,omitempty"`
	MembersCanCreateInternalRepos *bool `yaml:"members_can_create_internal_repositories,omitempty"`

	// MembersCanForkPrivateRepos toggles whether organization members can fork private organization repositories.
	MembersCanForkPrivateRepos *bool `yaml:"members_can_fork_private_repositories,omitempty"`

	// MembersAllowedRepositoryCreationType denotes if organization members can create repositories
	// and the type of repositories they can create. Possible values are: "all", "private", or "none".
	//
	// Deprecated: Use MembersCanCreatePublicRepos, MembersCanCreatePrivateRepos, MembersCanCreateInternalRepos
	// instead. The new fields overrides the existing MembersAllowedRepositoryCreationType during 'edit'
	// operation and does not consider 'internal' repositories during 'get' operation
	MembersAllowedRepositoryCreationType *string `yaml:"members_allowed_repository_creation_type,omitempty"`

	// MembersCanCreatePages toggles whether organization members can create GitHub Pages sites.
	MembersCanCreatePages *bool `yaml:"members_can_create_pages,omitempty"`
	// MembersCanCreatePublicPages toggles whether organization members can create public GitHub Pages sites.
	MembersCanCreatePublicPages *bool `yaml:"members_can_create_public_pages,omitempty"`
	// MembersCanCreatePrivatePages toggles whether organization members can create private GitHub Pages sites.
	MembersCanCreatePrivatePages *bool `yaml:"members_can_create_private_pages,omitempty"`

	// API URLs
	URL              *string `yaml:"url,omitempty"`
	EventsURL        *string `yaml:"events_url,omitempty"`
	HooksURL         *string `yaml:"hooks_url,omitempty"`
	IssuesURL        *string `yaml:"issues_url,omitempty"`
	MembersURL       *string `yaml:"members_url,omitempty"`
	PublicMembersURL *string `yaml:"public_members_url,omitempty"`
	ReposURL         *string `yaml:"repos_url,omitempty"`
}

type BranchRestrictions struct {
	// The list of user logins with push access.
	Users []*User `yaml:"users,omitempty"`
	// The list of team slugs with push access.
	Teams []*Team `yaml:"teams,omitempty"`
	// The list of app slugs with push access.
	Apps []*App `yaml:"apps,omitempty"`
}
type App struct {
	ID          *int64                   `yaml:"id,omitempty"`
	Slug        *string                  `yaml:"slug,omitempty"`
	NodeID      *string                  `yaml:"node_id,omitempty"`
	Owner       *User                    `yaml:"owner,omitempty"`
	Name        *string                  `yaml:"name,omitempty"`
	Description *string                  `yaml:"description,omitempty"`
	ExternalURL *string                  `yaml:"external_url,omitempty"`
	HTMLURL     *string                  `yaml:"html_url,omitempty"`
	CreatedAt   *Timestamp               `yaml:"created_at,omitempty"`
	UpdatedAt   *Timestamp               `yaml:"updated_at,omitempty"`
	Permissions *InstallationPermissions `yaml:"permissions,omitempty"`
	Events      []string                 `yaml:"events,omitempty"`
}
type InstallationPermissions struct {
	Actions                       *string `yaml:"actions,omitempty"`
	Administration                *string `yaml:"administration,omitempty"`
	Blocking                      *string `yaml:"blocking,omitempty"`
	Checks                        *string `yaml:"checks,omitempty"`
	Contents                      *string `yaml:"contents,omitempty"`
	ContentReferences             *string `yaml:"content_references,omitempty"`
	Deployments                   *string `yaml:"deployments,omitempty"`
	Emails                        *string `yaml:"emails,omitempty"`
	Environments                  *string `yaml:"environments,omitempty"`
	Followers                     *string `yaml:"followers,omitempty"`
	Issues                        *string `yaml:"issues,omitempty"`
	Metadata                      *string `yaml:"metadata,omitempty"`
	Members                       *string `yaml:"members,omitempty"`
	OrganizationAdministration    *string `yaml:"organization_administration,omitempty"`
	OrganizationHooks             *string `yaml:"organization_hooks,omitempty"`
	OrganizationPlan              *string `yaml:"organization_plan,omitempty"`
	OrganizationPreReceiveHooks   *string `yaml:"organization_pre_receive_hooks,omitempty"`
	OrganizationProjects          *string `yaml:"organization_projects,omitempty"`
	OrganizationSecrets           *string `yaml:"organization_secrets,omitempty"`
	OrganizationSelfHostedRunners *string `yaml:"organization_self_hosted_runners,omitempty"`
	OrganizationUserBlocking      *string `yaml:"organization_user_blocking,omitempty"`
	Packages                      *string `yaml:"packages,omitempty"`
	Pages                         *string `yaml:"pages,omitempty"`
	PullRequests                  *string `yaml:"pull_requests,omitempty"`
	RepositoryHooks               *string `yaml:"repository_hooks,omitempty"`
	RepositoryProjects            *string `yaml:"repository_projects,omitempty"`
	RepositoryPreReceiveHooks     *string `yaml:"repository_pre_receive_hooks,omitempty"`
	Secrets                       *string `yaml:"secrets,omitempty"`
	SecretScanningAlerts          *string `yaml:"secret_scanning_alerts,omitempty"`
	SecurityEvents                *string `yaml:"security_events,omitempty"`
	SingleFile                    *string `yaml:"single_file,omitempty"`
	Statuses                      *string `yaml:"statuses,omitempty"`
	TeamDiscussions               *string `yaml:"team_discussions,omitempty"`
	VulnerabilityAlerts           *string `yaml:"vulnerability_alerts,omitempty"`
	Workflows                     *string `yaml:"workflows,omitempty"`
}

type RequireLinearHistory struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type AllowForcePushes struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type AllowDeletions struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type RequiredConversationResolution struct {
	Enabled *bool `yaml:"enabled,omitempty"`
}

type Repository struct {
	ID                        *int64          `yaml:"id,omitempty"`
	NodeID                    *string         `yaml:"node_id,omitempty"`
	Owner                     *User           `yaml:"owner,omitempty"`
	Name                      *string         `yaml:"name,omitempty"`
	FullName                  *string         `yaml:"full_name,omitempty"`
	Description               *string         `yaml:"description,omitempty"`
	Homepage                  *string         `yaml:"homepage,omitempty"`
	CodeOfConduct             *CodeOfConduct  `yaml:"code_of_conduct,omitempty"`
	DefaultBranch             *string         `yaml:"default_branch,omitempty"`
	MasterBranch              *string         `yaml:"master_branch,omitempty"`
	CreatedAt                 *Timestamp      `yaml:"created_at,omitempty"`
	PushedAt                  *Timestamp      `yaml:"pushed_at,omitempty"`
	UpdatedAt                 *Timestamp      `yaml:"updated_at,omitempty"`
	HTMLURL                   *string         `yaml:"html_url,omitempty"`
	CloneURL                  *string         `yaml:"clone_url,omitempty"`
	GitURL                    *string         `yaml:"git_url,omitempty"`
	MirrorURL                 *string         `yaml:"mirror_url,omitempty"`
	SSHURL                    *string         `yaml:"ssh_url,omitempty"`
	SVNURL                    *string         `yaml:"svn_url,omitempty"`
	Language                  *string         `yaml:"language,omitempty"`
	Fork                      *bool           `yaml:"fork,omitempty"`
	ForksCount                *int            `yaml:"forks_count,omitempty"`
	NetworkCount              *int            `yaml:"network_count,omitempty"`
	OpenIssuesCount           *int            `yaml:"open_issues_count,omitempty"`
	OpenIssues                *int            `yaml:"open_issues,omitempty"` // Deprecated: Replaced by OpenIssuesCount. For backward compatibility OpenIssues is still populated.
	StargazersCount           *int            `yaml:"stargazers_count,omitempty"`
	SubscribersCount          *int            `yaml:"subscribers_count,omitempty"`
	WatchersCount             *int            `yaml:"watchers_count,omitempty"` // Deprecated: Replaced by StargazersCount. For backward compatibility WatchersCount is still populated.
	Watchers                  *int            `yaml:"watchers,omitempty"`       // Deprecated: Replaced by StargazersCount. For backward compatibility Watchers is still populated.
	Size                      *int            `yaml:"size,omitempty"`
	AutoInit                  *bool           `yaml:"auto_init,omitempty"`
	Parent                    *Repository     `yaml:"parent,omitempty"`
	Source                    *Repository     `yaml:"source,omitempty"`
	TemplateRepository        *Repository     `yaml:"template_repository,omitempty"`
	Organization              *Organization   `yaml:"organization,omitempty"`
	Permissions               map[string]bool `yaml:"permissions,omitempty"`
	AllowRebaseMerge          *bool           `yaml:"allow_rebase_merge,omitempty"`
	AllowUpdateBranch         *bool           `yaml:"allow_update_branch,omitempty"`
	AllowSquashMerge          *bool           `yaml:"allow_squash_merge,omitempty"`
	AllowMergeCommit          *bool           `yaml:"allow_merge_commit,omitempty"`
	AllowAutoMerge            *bool           `yaml:"allow_auto_merge,omitempty"`
	AllowForking              *bool           `yaml:"allow_forking,omitempty"`
	DeleteBranchOnMerge       *bool           `yaml:"delete_branch_on_merge,omitempty"`
	UseSquashPRTitleAsDefault *bool           `yaml:"use_squash_pr_title_as_default,omitempty"`
	Topics                    []string        `yaml:"topics,omitempty"`
	Archived                  *bool           `yaml:"archived,omitempty"`
	Disabled                  *bool           `yaml:"disabled,omitempty"`

	// Only provided when using RepositoriesService.Get while in preview
	License *License `yaml:"license,omitempty"`

	// Additional mutable fields when creating and editing a repository
	Private           *bool   `yaml:"private,omitempty"`
	HasIssues         *bool   `yaml:"has_issues,omitempty"`
	HasWiki           *bool   `yaml:"has_wiki,omitempty"`
	HasPages          *bool   `yaml:"has_pages,omitempty"`
	HasProjects       *bool   `yaml:"has_projects,omitempty"`
	HasDownloads      *bool   `yaml:"has_downloads,omitempty"`
	IsTemplate        *bool   `yaml:"is_template,omitempty"`
	LicenseTemplate   *string `yaml:"license_template,omitempty"`
	GitignoreTemplate *string `yaml:"gitignore_template,omitempty"`

	// Options for configuring Advanced Security and Secret Scanning
	SecurityAndAnalysis *SecurityAndAnalysis `yaml:"security_and_analysis,omitempty"`

	// Creating an organization repository. Required for non-owners.
	TeamID *int64 `yaml:"team_id,omitempty"`

	// API URLs
	URL              *string `yaml:"url,omitempty"`
	ArchiveURL       *string `yaml:"archive_url,omitempty"`
	AssigneesURL     *string `yaml:"assignees_url,omitempty"`
	BlobsURL         *string `yaml:"blobs_url,omitempty"`
	BranchesURL      *string `yaml:"branches_url,omitempty"`
	CollaboratorsURL *string `yaml:"collaborators_url,omitempty"`
	CommentsURL      *string `yaml:"comments_url,omitempty"`
	CommitsURL       *string `yaml:"commits_url,omitempty"`
	CompareURL       *string `yaml:"compare_url,omitempty"`
	ContentsURL      *string `yaml:"contents_url,omitempty"`
	ContributorsURL  *string `yaml:"contributors_url,omitempty"`
	DeploymentsURL   *string `yaml:"deployments_url,omitempty"`
	DownloadsURL     *string `yaml:"downloads_url,omitempty"`
	EventsURL        *string `yaml:"events_url,omitempty"`
	ForksURL         *string `yaml:"forks_url,omitempty"`
	GitCommitsURL    *string `yaml:"git_commits_url,omitempty"`
	GitRefsURL       *string `yaml:"git_refs_url,omitempty"`
	GitTagsURL       *string `yaml:"git_tags_url,omitempty"`
	HooksURL         *string `yaml:"hooks_url,omitempty"`
	IssueCommentURL  *string `yaml:"issue_comment_url,omitempty"`
	IssueEventsURL   *string `yaml:"issue_events_url,omitempty"`
	IssuesURL        *string `yaml:"issues_url,omitempty"`
	KeysURL          *string `yaml:"keys_url,omitempty"`
	LabelsURL        *string `yaml:"labels_url,omitempty"`
	LanguagesURL     *string `yaml:"languages_url,omitempty"`
	MergesURL        *string `yaml:"merges_url,omitempty"`
	MilestonesURL    *string `yaml:"milestones_url,omitempty"`
	NotificationsURL *string `yaml:"notifications_url,omitempty"`
	PullsURL         *string `yaml:"pulls_url,omitempty"`
	ReleasesURL      *string `yaml:"releases_url,omitempty"`
	StargazersURL    *string `yaml:"stargazers_url,omitempty"`
	StatusesURL      *string `yaml:"statuses_url,omitempty"`
	SubscribersURL   *string `yaml:"subscribers_url,omitempty"`
	SubscriptionURL  *string `yaml:"subscription_url,omitempty"`
	TagsURL          *string `yaml:"tags_url,omitempty"`
	TreesURL         *string `yaml:"trees_url,omitempty"`
	TeamsURL         *string `yaml:"teams_url,omitempty"`

	// Visibility is only used for Create and Edit endpoints. The visibility field
	// overrides the field parameter when both are used.
	// Can be one of public, private or internal.
	Visibility *string `yaml:"visibility,omitempty"`

	// RoleName is only returned by the API 'check team permissions for a repository'.
	// See: teams.go (IsTeamRepoByID) https://docs.github.com/en/rest/teams/teams#check-team-permissions-for-a-repository
	RoleName *string `yaml:"role_name,omitempty"`
}

type SecurityAndAnalysis struct {
	AdvancedSecurity *AdvancedSecurity `yaml:"advanced_security,omitempty"`
	SecretScanning   *SecretScanning   `yaml:"secret_scanning,omitempty"`
}

type AdvancedSecurity struct {
	Status *string `yaml:"status,omitempty"`
}

type SecretScanning struct {
	Status *string `yaml:"status,omitempty"`
}

type CodeOfConduct struct {
	Name *string `yaml:"name,omitempty"`
	Key  *string `yaml:"key,omitempty"`
	URL  *string `yaml:"url,omitempty"`
	Body *string `yaml:"body,omitempty"`
}

type License struct {
	Key  *string `yaml:"key,omitempty"`
	Name *string `yaml:"name,omitempty"`
	URL  *string `yaml:"url,omitempty"`

	SPDXID         *string   `yaml:"spdx_id,omitempty"`
	HTMLURL        *string   `yaml:"html_url,omitempty"`
	Featured       *bool     `yaml:"featured,omitempty"`
	Description    *string   `yaml:"description,omitempty"`
	Implementation *string   `yaml:"implementation,omitempty"`
	Permissions    *[]string `yaml:"permissions,omitempty"`
	Conditions     *[]string `yaml:"conditions,omitempty"`
	Limitations    *[]string `yaml:"limitations,omitempty"`
	Body           *string   `yaml:"body,omitempty"`
}
