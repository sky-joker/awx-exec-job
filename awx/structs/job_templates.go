package structs

type JobTemplates struct {
	Count    int32  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Id      int32  `json:"id"`
		Type    string `json:"type"`
		Url     string `json:"url"`
		Related struct {
			CreateBy                     string `json:"created_by"`
			ModifiedBy                   string `json:"modified_by"`
			Labels                       string `json:"labels"`
			Inventory                    string `json:"inventory"`
			Project                      string `json:"project"`
			NotificationTemplatesError   string `json:"notification_templates_error"`
			NotificationTemplatesSuccess string `json:"notification_templates_success"`
			Jobs                         string `json:"jobs"`
			ObjectRoles                  string `json:"object_roles"`
			NotificationTemplatesAny     string `json:"notification_templates_any"`
			AccessList                   string `json:"access_list"`
			Launch                       string `json:"launch"`
			InstanceGroups               string `json:"instance_groups"`
			Schedules                    string `json:"schedules"`
			Copy                         string `json:"copy"`
			ActivityStream               string `json:"activity_stream"`
			SurveySpec                   string `json:"survey_spec"`
		} `json:"related"`
		SummaryFields struct {
			Inventory struct {
				Id                           int32  `json:"id"`
				Name                         string `json:"name"`
				Description                  string `json:"description"`
				HasActiveFailures            bool   `json:"has_active_failures"`
				TotalHosts                   int32  `json:"total_hosts"`
				HostsWithActiveFailures      int32  `json:"hosts_with_active_failures"`
				TotalGroups                  int32  `json:"total_groups"`
				GroupsWithActiveFailures     int32  `json:"groups_with_active_failures"`
				HasInventorySources          bool   `json:"has_inventory_sources"`
				TotalInventorySources        int32  `json:"total_inventory_sources"`
				InventorySourcesWithFailures int32  `json:"inventory_sources_with_failures"`
				OrganizationId               int32  `json:"organization_id"`
			} `json:"inventory"`
			Project struct {
				Id          int32  `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Status      string `json:"status"`
				ScmType     string `json:"scm_type"`
			} `json:"project"`
			CreatedBy struct {
				Id        int32  `json:"id"`
				Username  string `json:"username"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			} `json:"created_by"`
			ModifiedBy struct {
				Id        int32  `json:"id"`
				Username  string `json:"username"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			} `json:"modified_by"`
			ObjectRoles struct {
				AdminRole struct {
					Id          int32  `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"admin_role"`
				ExecuteRole struct {
					Id          int32  `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"execute_role"`
				ReadRole struct {
					Id          int32  `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
				} `json:"read_role"`
			} `json:"object_roles"`
			UserCapabilities struct {
				Edit     bool `json:"edit"`
				Start    bool `json:"start"`
				Copy     bool `json:"copy"`
				Schedule bool `json:"schedule"`
				Delete   bool `json:"delete"`
			} `json:"user_capabilities"`
			Labels struct {
				Count   int32    `json:"count"`
				Results []string `json:"results"`
			} `json:"labels"`
			RecentJobs []struct {
				Status   string `json:"status"`
				Finished string `json:"finished"`
				Id       int32  `json:id`
			} `json:"recent_jobs"`
		} `json:"summary_fields"`
		Created               string `json:"created"`
		Modified              string `json:"modified"`
		Name                  string `json:"name"`
		Description           string `json:"description"`
		JobType               string `json:"job_type"`
		Inventory             int32  `json:"inventory"`
		Project               int32  `json:"project"`
		Playbook              string `json:"playbook"`
		Forks                 int32  `json:"forks"`
		Limit                 string `json:"limit"`
		Verbosity             int32  `json:"verbosity"`
		ExtraVars             string `json:"extra_vars"`
		JobTags               string `json:"job_tags"`
		ForceHandlers         bool   `json:"force_handlers"`
		SkipTags              string `json:"skip_tags"`
		StartAtTask           string `json:"start_at_task"`
		Timeout               int32  `json:"timeout"`
		UseFactCache          bool   `json:"use_fact_cache"`
		LastJobRun            string `json:"last_job_run"`
		LastJobFailed         bool   `json:"last_job_failed"`
		NextJobRun            string `json:"next_job_run"`
		Status                string `json:"status"`
		HostConfigKey         string `json:"host_config_key"`
		AskDiffModeOnLaunch   bool   `json:"ask_diff_mode_on_launch"`
		AskVariablesOnLaunch  bool   `json:"ask_variables_on_launch"`
		AskLimitOnLaunch      bool   `json:"ask_limit_on_launch"`
		AskTagsOnLaunch       bool   `json:"ask_tags_on_launch"`
		AskSkipTagsOnLaunch   bool   `json:"ask_skip_tags_on_launch"`
		AskJobTypeOnLaunch    bool   `json:"ask_job_type_on_launch"`
		AskVerbosityOnLaunch  bool   `json:"ask_verbosity_on_launch"`
		AskInventoryOnLaunch  bool   `json:"ask_inventory_on_launch"`
		AskCredentialOnLaunch bool   `json:"ask_credential_on_launch"`
		SurveyEnabled         bool   `json:"survey_enabled"`
		BecomeEnabled         bool   `json:"become_enabled"`
		DiffMode              bool   `json:"diff_mode"`
		AllowSimultaneous     bool   `json:"allow_simultaneous"`
		CustomVirtualenv      string `json:"custom_virtualenv"`
		CloudCredential       string `json:"cloud_credential"`
		NetworkCredential     string `json:"network_credential"`
		Credential            int32  `json:"credential"`
		Vaultredential        string `json:"vault_credential"`
	} `json:"results"`
}
