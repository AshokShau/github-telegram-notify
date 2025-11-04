package src

import (
	"fmt"
	"github-webhook/src/config"
	"github-webhook/src/utils"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v77/github"
)

// GitHubWebhook processes GitHub webhooks
func GitHubWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	payload, err := github.ValidatePayload(r, nil)
	if err != nil {
		log.Printf("Error validating payload: %v\n", err)
		http.Error(w, "Invalid payload", http.StatusUnauthorized)
		return
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("Error parsing webhook: %v\n", err)
		http.Error(w, "Error parsing webhook", http.StatusInternalServerError)
		return
	}

	// Prioritize critical or frequent event types
	var message string
	var markup *utils.InlineKeyboardMarkup
	switch e := event.(type) {
	case *github.PushEvent:
		message, markup = utils.HandlePushEvent(e)
	case *github.PullRequestEvent:
		message, markup = utils.HandlePullRequestEvent(e)
	case *github.IssuesEvent:
		message, markup = utils.HandleIssuesEvent(e)
	case *github.PingEvent:
		message, markup = utils.HandlePingEvent(e)

	// Handle review-related events
	case *github.PullRequestReviewEvent:
		message, markup = utils.HandlePullRequestReviewEvent(e)
	case *github.PullRequestReviewCommentEvent:
		message, markup = utils.HandlePullRequestReviewCommentEvent(e)

	// Handle repository and organization events
	case *github.RepositoryEvent:
		message, markup = utils.HandleRepositoryEvent(e)
	case *github.RepositoryDispatchEvent:
		message, markup = utils.HandleRepositoryDispatchEvent(e)
	case *github.OrganizationEvent:
		message, markup = utils.HandleOrganizationEvent(e)
	case *github.OrgBlockEvent:
		message, markup = utils.HandleOrgBlockEvent(e)

	// Handle CI/CD and deployment-related events
	case *github.CheckRunEvent:
		message, markup = utils.HandleCheckRunEvent(e)
	case *github.CheckSuiteEvent:
		message, markup = utils.HandleCheckSuiteEvent(e)
	case *github.WorkflowRunEvent:
		message, markup = utils.HandleWorkflowRunEvent(e)
	case *github.WorkflowJobEvent:
		message, markup = utils.HandleWorkflowJobEvent(e)
	case *github.DeploymentEvent:
		message, markup = utils.HandleDeploymentEvent(e)
	case *github.DeploymentStatusEvent:
		message, markup = utils.HandleDeploymentStatusEvent(e)

	// Handle advisory and security-related events
	case *github.SecurityAdvisoryEvent:
		message, markup = utils.HandleSecurityAdvisoryEvent(e)
	case *github.MembershipEvent:
		message, markup = utils.HandleMembershipEvent(e)
	case *github.MilestoneEvent:
		message, markup = utils.HandleMilestoneEvent(e)

	// Handle less frequent or low-priority events
	case *github.CommitCommentEvent:
		message, markup = utils.HandleCommitCommentEvent(e)
	case *github.ForkEvent:
		message, markup = utils.HandleForkEvent(e)
	case *github.ReleaseEvent:
		message, markup = utils.HandleReleaseEvent(e)
	case *github.StarEvent:
		message, markup = utils.HandleStarEvent(e)
	case *github.WatchEvent:
		message, markup = utils.HandleWatchEvent(e)
	case *github.LabelEvent:
		message, markup = utils.HandleLabelEvent(e)
	case *github.MarketplacePurchaseEvent:
		message, markup = utils.HandleMarketplacePurchaseEvent(e)
	case *github.PageBuildEvent:
		message, markup = utils.HandlePageBuildEvent(e)
	case *github.DeployKeyEvent:
		message, markup = utils.HandleDeployKeyEvent(e)
	case *github.StarredRepository:
		message, markup = utils.HandleStarredEvent(e)
	case *github.CreateEvent:
		message, markup = utils.HandleCreateEvent(e)
	case *github.DeleteEvent:
		message, markup = utils.HandleDeleteEvent(e)
	case *github.IssueCommentEvent:
		message, markup = utils.HandleIssueCommentEvent(e)
	case *github.MemberEvent:
		message, markup = utils.HandleMemberEvent(e)
	case *github.PublicEvent:
		message, markup = utils.HandlePublicEvent(e)
	case *github.StatusEvent:
		message, markup = utils.HandleStatusEvent(e)
	case *github.WorkflowDispatchEvent:
		message, markup = utils.HandleWorkflowDispatchEvent(e)
	case *github.TeamAddEvent:
		message, markup = utils.HandleTeamAddEvent(e)
	case *github.TeamEvent:
		message, markup = utils.HandleTeamEvent(e)
	case *github.PackageEvent:
		message, markup = utils.HandlePackageEvent(e)
	case *github.GollumEvent:
		message, markup = utils.HandleGollumEvent(e)
	case *github.MetaEvent:
		message, markup = utils.HandleMetaEvent(e)
	case *github.InstallationEvent:
		message, markup = utils.HandleInstallationEvent(e)
	// Catch-all fallback for unhandled events
	default:
		log.Printf("Unhandled event type: %s\n", github.WebHookType(r))
		message = fmt.Sprintf("Unhandled event type: %s", github.WebHookType(r))
	}

	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		http.Error(w, "Missing chat_id query parameter", http.StatusBadRequest)
		return
	}

	if message == "" {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
		return
	}

	err = utils.SendToTelegram(chatID, message, markup)
	if err != nil {
		http.Error(w, strings.ReplaceAll(err.Error(), config.BotToken, "$Bot"), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(message))
}
