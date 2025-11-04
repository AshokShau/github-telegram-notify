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
	switch e := event.(type) {
	case *github.PushEvent:
		message = utils.HandlePushEvent(e)
	case *github.PullRequestEvent:
		message = utils.HandlePullRequestEvent(e)
	case *github.IssuesEvent:
		message = utils.HandleIssuesEvent(e)
	case *github.PingEvent:
		message = utils.HandlePingEvent(e)

	// Handle review-related events
	case *github.PullRequestReviewEvent:
		message = utils.HandlePullRequestReviewEvent(e)
	case *github.PullRequestReviewCommentEvent:
		message = utils.HandlePullRequestReviewCommentEvent(e)

	// Handle repository and organization events
	case *github.RepositoryEvent:
		message = utils.HandleRepositoryEvent(e)
	case *github.RepositoryDispatchEvent:
		message = utils.HandleRepositoryDispatchEvent(e)
	case *github.OrganizationEvent:
		message = utils.HandleOrganizationEvent(e)
	case *github.OrgBlockEvent:
		message = utils.HandleOrgBlockEvent(e)

	// Handle CI/CD and deployment-related events
	case *github.CheckRunEvent:
		message = utils.HandleCheckRunEvent(e)
	case *github.CheckSuiteEvent:
		message = utils.HandleCheckSuiteEvent(e)
	case *github.WorkflowRunEvent:
		message = utils.HandleWorkflowRunEvent(e)
	case *github.WorkflowJobEvent:
		message = utils.HandleWorkflowJobEvent(e)
	case *github.DeploymentEvent:
		message = utils.HandleDeploymentEvent(e)
	case *github.DeploymentStatusEvent:
		message = utils.HandleDeploymentStatusEvent(e)

	// Handle advisory and security-related events
	case *github.SecurityAdvisoryEvent:
		message = utils.HandleSecurityAdvisoryEvent(e)
	case *github.MembershipEvent:
		message = utils.HandleMembershipEvent(e)
	case *github.MilestoneEvent:
		message = utils.HandleMilestoneEvent(e)

	// Handle less frequent or low-priority events
	case *github.CommitCommentEvent:
		message = utils.HandleCommitCommentEvent(e)
	case *github.ForkEvent:
		message = utils.HandleForkEvent(e)
	case *github.ReleaseEvent:
		message = utils.HandleReleaseEvent(e)
	case *github.StarEvent:
		message = utils.HandleStarEvent(e)
	case *github.WatchEvent:
		message = utils.HandleWatchEvent(e)
	case *github.LabelEvent:
		message = utils.HandleLabelEvent(e)
	case *github.MarketplacePurchaseEvent:
		message = utils.HandleMarketplacePurchaseEvent(e)
	case *github.PageBuildEvent:
		message = utils.HandlePageBuildEvent(e)
	case *github.DeployKeyEvent:
		message = utils.HandleDeployKeyEvent(e)
	case *github.StarredRepository:
		message = utils.HandleStarredEvent(e)
	case *github.CreateEvent:
		message = utils.HandleCreateEvent(e)
	case *github.DeleteEvent:
		message = utils.HandleDeleteEvent(e)
	case *github.IssueCommentEvent:
		message = utils.HandleIssueCommentEvent(e)
	case *github.MemberEvent:
		message = utils.HandleMemberEvent(e)
	case *github.PublicEvent:
		message = utils.HandlePublicEvent(e)
	case *github.StatusEvent:
		message = utils.HandleStatusEvent(e)
	case *github.WorkflowDispatchEvent:
		message = utils.HandleWorkflowDispatchEvent(e)
	case *github.TeamAddEvent:
		message = utils.HandleTeamAddEvent(e)
	case *github.TeamEvent:
		message = utils.HandleTeamEvent(e)
	case *github.PackageEvent:
		message = utils.HandlePackageEvent(e)
	case *github.GollumEvent:
		message = utils.HandleGollumEvent(e)
	case *github.MetaEvent:
		message = utils.HandleMetaEvent(e)
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

	err = utils.SendToTelegram(chatID, message)
	if err != nil {
		http.Error(w, strings.ReplaceAll(err.Error(), config.BotToken, "$Bot"), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(message))
}
