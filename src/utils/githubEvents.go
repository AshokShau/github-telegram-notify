package utils

import (
	"encoding/json"
	"fmt"
	"html"
	"strings"

	"github.com/google/go-github/v71/github"
)

func HandleIssuesEvent(event *github.IssuesEvent) string {
	repo := event.GetRepo().GetFullName()
	action := event.GetAction()
	sender := event.GetSender().GetLogin()
	issue := event.GetIssue()
	title := issue.GetTitle()
	url := issue.GetHTMLURL()
	body := issue.GetBody()

	baseMessage := fmt.Sprintf(
		"📂 <b>Repository:</b> %s\n"+
			"⚡ <b>Action:</b> %s\n"+
			"👤 <b>Sender:</b> %s\n"+
			"📝 <b>Issue:</b> %s\n"+
			"🔗 <b>URL:</b> <a href='%s'>%s</a>\n",
		repo,
		action,
		sender,
		title,
		url, url,
	)

	switch action {
	case "opened":
		return baseMessage + fmt.Sprintf("✍️ <b>Description:</b>\n%s\n", body)

	case "edited":
		return baseMessage + fmt.Sprintf("✏️ <i>Issue was edited.</i>\n✍️ <b>Description:</b>\n%s\n", body)

	case "deleted":
		return baseMessage + "🗑️ <i>The issue was deleted.</i>\n"

	case "transferred":
		return baseMessage + "🔄 <i>The issue was transferred to a different repository.</i>\n"

	case "pinned":
		return baseMessage + "📌 <i>The issue was pinned.</i>\n"

	case "unpinned":
		return baseMessage + "📌❌ <i>The issue was unpinned.</i>\n"

	case "closed":
		closedBy := issue.GetClosedBy()
		closer := ""
		if closedBy != nil {
			closer = fmt.Sprintf("🔒 <b>Closed by:</b> %s\n", closedBy.GetLogin())
		}
		return baseMessage + closer + "🔒 <i>The issue is now closed.</i>\n"

	case "reopened":
		return baseMessage + "🔓 <i>The issue was reopened.</i>\n"

	case "assigned":
		assignees := issue.Assignees
		var assigneeNames []string
		for _, assignee := range assignees {
			assigneeNames = append(assigneeNames, assignee.GetLogin())
		}
		return baseMessage + fmt.Sprintf("🙋 <b>Assignees:</b> %s\n", strings.Join(assigneeNames, ", "))

	case "unassigned":
		return baseMessage + "🙅 <i>An assignee was removed from the issue.</i>\n"

	case "labeled":
		labels := issue.Labels
		var labelNames []string
		for _, label := range labels {
			labelNames = append(labelNames, label.GetName())
		}
		return baseMessage + fmt.Sprintf("🏷️ <b>Labels:</b> %s\n", strings.Join(labelNames, ", "))

	case "unlabeled":
		return baseMessage + "🏷️❌ <i>A label was removed from the issue.</i>\n"

	case "locked":
		return baseMessage + "🔒 <i>The issue was locked.</i>\n"

	case "unlocked":
		return baseMessage + "🔓 <i>The issue was unlocked.</i>\n"

	case "milestoned":
		milestone := issue.GetMilestone()
		if milestone != nil {
			return baseMessage + fmt.Sprintf("🏁 <b>Milestone:</b> %s\n", milestone.GetTitle())
		}
		return baseMessage + "🏁 <i>The issue was added to a milestone.</i>\n"

	case "demilestoned":
		return baseMessage + "🏁❌ <i>The issue was removed from a milestone.</i>\n"

	default:
		return baseMessage + "❓ <i>Unhandled action.</i>\n"
	}
}

func HandlePullRequestEvent(event *github.PullRequestEvent) string {
	repo := event.GetRepo().GetFullName()
	action := event.GetAction()
	sender := event.GetSender().GetLogin()
	pr := event.GetPullRequest()
	title := pr.GetTitle()
	url := pr.GetHTMLURL()
	body := pr.GetBody()
	state := pr.GetState()

	baseMessage := fmt.Sprintf(
		"📂 <b>Repository:</b> <a href='https://github.com/%s'>%s</a>\n"+
			"⚡ <b>Action:</b> %s\n"+
			"👤 <b>Sender:</b> <a href='https://github.com/%s'>%s</a>\n"+
			"🔗 <b>Pull Request:</b> <a href='%s'>%s</a>\n"+
			"📌 <b>State:</b> %s\n",
		repo, repo,
		action,
		sender, sender,
		url, title,
		state,
	)

	switch action {
	case "opened":
		return baseMessage + fmt.Sprintf("✍️ <b>Description:</b>\n%s\n", body)

	case "closed":
		if pr.GetMerged() {
			return baseMessage + "✅ <i>The pull request was successfully merged.</i>\n"
		}
		return baseMessage + "❌ <i>The pull request was closed without merging.</i>\n"

	case "reopened":
		return baseMessage + "🔄 <i>The pull request was reopened.</i>\n"

	case "edited":
		return baseMessage + fmt.Sprintf("✏️ <i>The pull request was edited.</i>\n✍️ <b>Description:</b>\n%s\n", body)

	case "assigned":
		assignees := pr.Assignees
		var assigneeLinks []string
		for _, assignee := range assignees {
			assigneeLinks = append(assigneeLinks, fmt.Sprintf("<a href='https://github.com/%s'>%s</a>", assignee.GetLogin(), assignee.GetLogin()))
		}
		return baseMessage + fmt.Sprintf("🙋 <b>Assigned to:</b> %s\n", strings.Join(assigneeLinks, ", "))

	case "unassigned":
		return baseMessage + "🙅 <i>An assignee was removed from the pull request.</i>\n"

	case "review_requested":
		reviewers := pr.RequestedReviewers
		var reviewerLinks []string
		for _, reviewer := range reviewers {
			reviewerLinks = append(reviewerLinks, fmt.Sprintf("<a href='https://github.com/%s'>%s</a>", reviewer.GetLogin(), reviewer.GetLogin()))
		}
		return baseMessage + fmt.Sprintf("📝 <b>Review requested from:</b> %s\n", strings.Join(reviewerLinks, ", "))

	case "review_request_removed":
		return baseMessage + "❌ <i>A review request was removed from the pull request.</i>\n"

	case "labeled":
		labels := pr.Labels
		var labelNames []string
		for _, label := range labels {
			labelNames = append(labelNames, label.GetName())
		}
		return baseMessage + fmt.Sprintf("🏷️ <b>Labels:</b> %s\n", strings.Join(labelNames, ", "))

	case "unlabeled":
		return baseMessage + "🏷️❌ <i>A label was removed from the pull request.</i>\n"

	case "locked":
		return baseMessage + "🔒 <i>The pull request was locked.</i>\n"

	case "unlocked":
		return baseMessage + "🔓 <i>The pull request was unlocked.</i>\n"

	case "synchronize":
		return baseMessage + "🔄 <i>The pull request was synchronized (updated with new commits).</i>\n"

	default:
		return baseMessage + "❓ <i>Unhandled action.</i>\n"
	}
}

func HandleStarredEvent(event *github.StarredRepository) string {
	repo := event.Repository.GetFullName()
	repoURL := event.Repository.GetHTMLURL()
	sender := event.Repository.Owner.GetLogin()
	stars := event.Repository.GetStargazersCount()
	forks := event.Repository.GetForksCount()

	return fmt.Sprintf(
		"🌟 <b>Repository Starred Event</b>\n\n"+
			"📂 <b>Repository:</b> <a href='%s'>%s</a>\n"+
			"👤 <b>Starred by:</b> <a href='https://github.com/%s'>%s</a>\n\n"+
			"✨ <b>Total Stars:</b> %d | 🍴 <b>Total Forks:</b> %d\n\n"+
			"💡 <i>Keep up the great work!</i>",
		repoURL,
		repo,
		sender, sender,
		stars,
		forks,
	)
}
func HandlePushEvent(event *github.PushEvent) string {
	repo := event.Repo.GetFullName()
	repoURL := event.Repo.GetHTMLURL()
	sender := event.Sender.GetLogin()
	ref := event.GetRef()
	compareURL := event.GetCompare()
	commitCount := len(event.Commits)
	distinctCommits := event.GetDistinctSize()
	created := event.GetCreated()
	deleted := event.GetDeleted()
	forced := event.GetForced()

	// Push summary
	message := fmt.Sprintf(
		"<b>🚀 Push Event in Repository:</b> <a href='%s'>%s</a>\n"+
			"<b>📂 Branch/Ref:</b> <i>%s</i>\n"+
			"<b>👤 Pushed by:</b> %s\n\n",
		repoURL,
		repo,
		ref,
		sender,
	)

	// Include branch creation/deletion/force-push details
	if created {
		message += "<i>🌱 A new branch was created.</i>\n\n"
	} else if deleted {
		message += "<i>🗑️ The branch was deleted.</i>\n\n"
	} else if forced {
		message += "<i>⚠️ This was a force-push.</i>\n\n"
	}

	// Commit summary
	message += fmt.Sprintf(
		"<b>📊 Total Commits:</b> %d <i>(Distinct: %d)</i>\n"+
			"<b>🔗 Compare Changes:</b> <a href='%s'>View Comparison</a>\n\n",
		commitCount,
		distinctCommits,
		compareURL,
	)

	// List individual commits
	if commitCount > 0 {
		message += "<b>🔨 Commit Details:</b>\n\n"
		for _, commit := range event.Commits {
			commitMessage := commit.GetMessage()
			author := commit.Author.GetName()
			url := commit.GetURL()

			message += fmt.Sprintf(
				"• <i>%s</i> by <b>%s</b> (<a href='%s'>View Commit</a>)\n",
				commitMessage,
				author,
				url,
			)
			/*
				// Add details about files changed in each commit
				added := commit.Added
				removed := commit.Removed
				modified := commit.Modified

				if len(added) > 0 || len(removed) > 0 || len(modified) > 0 {
					message += "<i>📁 Changed Files:</i>\n"
					if len(added) > 0 {
						message += fmt.Sprintf("<b>➕ Added:</b> %v\n", added)
					}
					if len(removed) > 0 {
						message += fmt.Sprintf("<b>❌ Removed:</b> %v\n", removed)
					}
					if len(modified) > 0 {
						message += fmt.Sprintf("<b>✏️ Modified:</b> %v\n", modified)
					}
				}
			*/
			message += "\n"
		}
	} else if event.HeadCommit != nil {
		// Add details for the head commit if no commit list is provided
		headCommit := event.HeadCommit
		message += fmt.Sprintf(
			"<b>📍 Head Commit:</b> <i>%s</i> by <b>%s</b> (<a href='%s'>View Commit</a>)\n",
			headCommit.GetMessage(),
			headCommit.Author.GetName(),
			headCommit.GetURL(),
		)
		/*
			// Add files changed in the head commit
			added := headCommit.Added
			removed := headCommit.Removed
			modified := headCommit.Modified

			if len(added) > 0 || len(removed) > 0 || len(modified) > 0 {
				message += "<i>📁 Changed Files:</i>\n"
				if len(added) > 0 {
					message += fmt.Sprintf("<b>➕ Added:</b> %v\n", added)
				}
				if len(removed) > 0 {
					message += fmt.Sprintf("<b>❌ Removed:</b> %v\n", removed)
				}
				if len(modified) > 0 {
					message += fmt.Sprintf("<b>✏️ Modified:</b> %v\n", modified)
				}
			}
		*/
	}

	return message
}
func HandleCreateEvent(event *github.CreateEvent) string {
	repo := event.Repo.GetFullName()
	repoURL := event.Repo.GetHTMLURL()
	sender := event.Sender.GetLogin()
	refType := event.GetRefType() // "branch", "tag", etc.
	ref := event.GetRef()         // Name of the branch, tag, or reference
	description := event.GetDescription()
	masterBranch := event.GetMasterBranch()
	pusherType := event.GetPusherType() // "user" or "bot"

	// Base message
	message := fmt.Sprintf(
		"<b>🆕 %s created a new %s <i>%s</i> in <a href='%s'>%s</a></b>\n",
		sender,
		refType,
		ref,
		repoURL,
		repo,
	)

	// Add repository description if available
	if description != "" {
		message += fmt.Sprintf("<i>📖 Repository Description:</i> %s\n\n", description)
	}

	// Add master branch information if the reference type is a branch
	if refType == "branch" && masterBranch != "" {
		message += fmt.Sprintf("<i>🌟 Default Branch:</i> <b>%s</b>\n\n", masterBranch)
	}

	// Add pusher type if available
	if pusherType != "" {
		message += fmt.Sprintf("<i>👤 Pusher Type:</i> <b>%s</b>\n\n", pusherType)
	}

	return message
}
func HandleDeleteEvent(event *github.DeleteEvent) string {
	repo := event.Repo.GetFullName()
	repoURL := event.Repo.GetHTMLURL()
	sender := event.Sender.GetLogin()
	refType := event.GetRefType() // "branch" or "tag"
	ref := event.GetRef()         // Name of the deleted branch or tag

	// Format message based on the type of deletion
	var message string
	switch refType {
	case "branch":
		message = fmt.Sprintf(
			"<b>🗑️ %s deleted the branch <i>%s</i> in <a href='%s'>%s</a></b>.\n",
			sender,
			ref,
			repoURL,
			repo,
		)
	case "tag":
		message = fmt.Sprintf(
			"<b>🏷️ %s deleted the tag <i>%s</i> in <a href='%s'>%s</a></b>.\n",
			sender,
			ref,
			repoURL,
			repo,
		)
	default:
		message = fmt.Sprintf(
			"<b>❌ %s deleted a %s <i>%s</i> in <a href='%s'>%s</a></b>.\n",
			sender,
			refType,
			ref,
			repoURL,
			repo,
		)
	}

	return message
}
func HandleForkEvent(event *github.ForkEvent) string {
	originalRepo := event.Repo.GetFullName() // The original repository's full name
	forkedRepo := event.Forkee.GetFullName() // The forked repository's full name
	sender := event.Sender.GetLogin()        // The user who created the fork

	originalForkCount := event.Repo.GetForksCount()
	originalStarCount := event.Repo.GetStargazersCount()

	// Enhanced message with clickable repository links and better formatting
	message := fmt.Sprintf(
		"<b>🍴 %s forked the repository <a href='https://github.com/%s'>%s</a> to create <a href='https://github.com/%s'>%s</a></b>\n"+
			"🌟 The original repository has <b>%d stars</b> and <b>%d forks</b>.\n",
		sender,
		originalRepo,
		originalRepo,
		forkedRepo,
		forkedRepo,
		originalStarCount,
		originalForkCount,
	)

	return message
}
func HandleCommitCommentEvent(event *github.CommitCommentEvent) string {
	comment := event.Comment.GetBody()       // The body of the commit comment
	commitSHA := event.Comment.GetCommitID() // The commit ID (SHA)
	repo := event.Repo.GetFullName()         // The repository's full name
	sender := event.Sender.GetLogin()        // The user who made the comment
	action := event.GetAction()              // The action (created, edited, deleted)

	// Common base URL for GitHub
	baseURL := "https://github.com/" + repo

	// Prepare message based on action
	switch action {
	case "created":
		return fmt.Sprintf(
			"💬 <b>%s</b> commented on commit <b>%s</b> in <a href='%s'>%s</a>.\n"+
				"📝 <i>Comment:</i> %s\n",
			sender,
			commitSHA[:7], // First 7 characters of commit SHA for brevity
			baseURL,
			repo,
			comment,
		)
	case "edited":
		return fmt.Sprintf(
			"✏️ <b>%s</b> edited their comment on commit <b>%s</b> in <a href='%s'>%s</a>.\n"+
				"📝 <i>Comment:</i> %s\n",
			sender,
			commitSHA[:7],
			baseURL,
			repo,
			comment,
		)
	case "deleted":
		return fmt.Sprintf(
			"❌ <b>%s</b> deleted their comment on commit <b>%s</b> in <a href='%s'>%s</a>.\n",
			sender,
			commitSHA[:7],
			baseURL,
			repo,
		)
	default:
		return fmt.Sprintf(
			"⚠️ <b>%s</b> performed an unknown action on their comment on commit <b>%s</b> in <a href='%s'>%s</a>.\n",
			sender,
			commitSHA[:7],
			baseURL,
			repo,
		)
	}
}

func HandlePublicEvent(event *github.PublicEvent) string {
	repoName := event.Repo.GetFullName() // Full name of the repository
	repoURL := event.Repo.GetHTMLURL()   // URL of the repository
	sender := event.Sender.GetLogin()    // User who made the repository public

	// Build the message
	message := fmt.Sprintf(
		"🔓 The repository <b>%s</b> is now public!\n"+
			"🌐 Repository URL: <a href=\"%s\">%s</a>\n"+
			"👤 Made public by: <b>%s</b>",
		repoName,
		repoURL,
		repoURL,
		sender,
	)

	return message
}

func HandleIssueCommentEvent(event *github.IssueCommentEvent) string {
	action := event.GetAction()            // The action performed (created, edited, deleted)
	issueTitle := event.Issue.GetTitle()   // The title of the issue
	issueURL := event.Issue.GetHTMLURL()   // The URL of the issue
	commentBody := event.Comment.GetBody() // The body of the comment
	issueNumber := event.Issue.GetNumber() // The issue number
	repoName := event.Repo.GetFullName()   // The full name of the repository
	sender := event.Sender.GetLogin()      // The user who performed the action

	// Format the message based on the action taken
	switch action {
	case "created":
		return fmt.Sprintf(
			"💬 <b>%s</b> commented on <a href='%s'>issue #%d</a> in <b>%s</b>.\n"+
				"📝 <i>Comment:</i> %s\n"+
				"📌 <i>Issue Title:</i> %s\n",
			sender,
			issueURL,
			issueNumber,
			repoName,
			commentBody,
			issueTitle,
		)
	case "edited":
		return fmt.Sprintf(
			"✏️ <b>%s</b> edited their comment on <a href='%s'>issue #%d</a> in <b>%s</b>.\n"+
				"📝 <i>Comment:</i> %s\n"+
				"📌 <i>Issue Title:</i> %s\n",
			sender,
			issueURL,
			issueNumber,
			repoName,
			commentBody,
			issueTitle,
		)
	case "deleted":
		return fmt.Sprintf(
			"❌ <b>%s</b> deleted their comment on <a href='%s'>issue #%d</a> in <b>%s</b>.\n",
			sender,
			issueURL,
			issueNumber,
			repoName,
		)
	default:
		return fmt.Sprintf(
			"⚠️ <b>%s</b> performed an unknown action on their comment on <a href='%s'>issue #%d</a> in <b>%s</b>.\n",
			sender,
			issueURL,
			issueNumber,
			repoName,
		)
	}
}
func HandleMemberEvent(event *github.MemberEvent) string {
	action := event.GetAction()          // The action performed (added, removed, etc.)
	member := event.Member.GetLogin()    // The login of the member
	repoName := event.Repo.GetFullName() // The full name of the repository
	orgName := event.Org.GetLogin()      // The organization name (if applicable)
	sender := event.Sender.GetLogin()    // The user who performed the action

	var message string

	// Format the message based on the action performed
	switch action {
	case "added":
		message = fmt.Sprintf(
			"🔹 <b>%s</b> was added as a member of <b>%s</b> in the organization <b>%s</b>.\n"+
				"👤 Added by: <b>%s</b>",
			member,
			repoName,
			orgName,
			sender,
		)
	case "removed":
		message = fmt.Sprintf(
			"🔸 <b>%s</b> was removed from the repository <b>%s</b> in the organization <b>%s</b>.\n"+
				"👤 Removed by: <b>%s</b>",
			member,
			repoName,
			orgName,
			sender,
		)
	case "edited":
		// Check for changes and include them if available
		if event.Changes != nil {
			message = fmt.Sprintf(
				"✏️ <b>%s</b>'s role was updated in the repository <b>%s</b> of the organization <b>%s</b>.\n"+
					"🔄 Changes: <pre>%v</pre>\n👤 Updated by: <b>%s</b>",
				member,
				repoName,
				orgName,
				event.Changes,
				sender,
			)
		} else {
			message = fmt.Sprintf(
				"✏️ <b>%s</b>'s role was edited in the repository <b>%s</b> of the organization <b>%s</b>.\n"+
					"👤 Edited by: <b>%s</b>",
				member,
				repoName,
				orgName,
				sender,
			)
		}
	default:
		message = fmt.Sprintf(
			"⚠️ <b>%s</b> performed an unknown action on <b>%s</b> in the organization <b>%s</b>.\n",
			sender,
			repoName,
			orgName,
		)
	}

	return message
}
func HandleRepositoryEvent(event *github.RepositoryEvent) string {
	action := event.GetAction()          // The action performed (e.g., created, renamed, archived)
	repoName := event.Repo.GetFullName() // Full name of the repository
	repoURL := event.Repo.GetHTMLURL()   // Repository's HTML URL
	sender := event.Sender.GetLogin()    // User who performed the action

	var message string

	switch action {
	case "created":
		message = fmt.Sprintf(
			"🎉 Repository <b>%s</b> has been created!\n"+
				"🌐 Repository URL: <a href='%s'>%s</a>\n"+
				"👤 Created by: <b>%s</b>",
			repoName,
			repoURL,
			repoURL,
			sender,
		)
	case "renamed":
		newName := event.Repo.GetName() // New name of the repository
		message = fmt.Sprintf(
			"🔄 Repository <b>%s</b> has been renamed to <b>%s</b>!\n"+
				"👤 Renamed by: <b>%s</b>",
			repoName,
			newName,
			sender,
		)
	case "archived":
		message = fmt.Sprintf(
			"🔒 Repository <b>%s</b> has been archived.\n"+
				"👤 Archived by: <b>%s</b>",
			repoName,
			sender,
		)
	case "unarchived":
		message = fmt.Sprintf(
			"🔓 Repository <b>%s</b> has been unarchived.\n"+
				"👤 Unarchived by: <b>%s</b>",
			repoName,
			sender,
		)
	default:
		message = fmt.Sprintf(
			"⚠️ <b>%s</b> performed an unknown action (<i>%s</i>) on repository <b>%s</b>.\n"+
				"🌐 Repository URL: <a href='%s'>%s</a>",
			sender,
			action,
			repoName,
			repoURL,
			repoURL,
		)
	}

	return message
}
func HandleReleaseEvent(event *github.ReleaseEvent) string {
	action := event.GetAction()               // Action performed (e.g., created, published, deleted, edited)
	release := event.GetRelease()             // Release details
	repoName := event.GetRepo().GetFullName() // Full name of the repository
	sender := event.GetSender().GetLogin()    // User who performed the action

	releaseName := release.GetName()        // Name of the release
	releaseTag := release.GetTagName()      // Tag name of the release
	releaseDescription := release.GetBody() // Description/body of the release
	releaseURL := release.GetHTMLURL()      // HTML URL of the release

	// Fallback for empty descriptions
	if releaseDescription == "" {
		releaseDescription = "No description provided."
	}

	var message string

	// Format the message based on the action performed
	switch action {
	case "created":
		message = fmt.Sprintf(
			"🎉 A new release has been created in <b>%s</b>!\n"+
				"📦 Release Name: <b>%s</b>\n"+
				"🏷️ Tag: <b>%s</b>\n"+
				"📝 Description: %s\n"+
				"🌐 <a href='%s'>View Release</a>\n"+
				"👤 Created by: <b>%s</b>",
			repoName,
			releaseName,
			releaseTag,
			releaseDescription,
			releaseURL,
			sender,
		)
	case "published":
		message = fmt.Sprintf(
			"🚀 The release <b>%s</b> has been published in <b>%s</b>!\n"+
				"🏷️ Tag: <b>%s</b>\n"+
				"🌐 <a href='%s'>View Release</a>\n"+
				"👤 Published by: <b>%s</b>",
			releaseName,
			repoName,
			releaseTag,
			releaseURL,
			sender,
		)
	case "deleted":
		message = fmt.Sprintf(
			"🗑️ The release <b>%s</b> (tag: <b>%s</b>) has been deleted from <b>%s</b>.\n"+
				"👤 Deleted by: <b>%s</b>",
			releaseName,
			releaseTag,
			repoName,
			sender,
		)
	case "edited":
		message = fmt.Sprintf(
			"📝 The release <b>%s</b> (tag: <b>%s</b>) in <b>%s</b> has been edited.\n"+
				"🌐 <a href='%s'>View Release</a>\n"+
				"👤 Edited by: <b>%s</b>",
			releaseName,
			releaseTag,
			repoName,
			releaseURL,
			sender,
		)
	default:
		message = fmt.Sprintf(
			"⚠️ An unknown action (<i>%s</i>) was performed on a release in <b>%s</b>.\n"+
				"👤 Performed by: <b>%s</b>",
			action,
			repoName,
			sender,
		)
	}

	return message
}
func HandleWatchEvent(event *github.WatchEvent) string {
	action := event.GetAction()                 // The action performed (always 'started')
	repoName := event.GetRepo().GetFullName()   // The full name of the repository (owner/repo-name)
	repoURL := event.GetRepo().GetHTMLURL()     // The HTML URL of the repository
	sender := event.GetSender().GetLogin()      // The user who performed the action
	senderURL := event.GetSender().GetHTMLURL() // The HTML URL of the user

	var message string

	// Format the message based on the action performed
	switch action {
	case "started":
		message = fmt.Sprintf(
			"⭐ <a href='%s'>%s</a> has starred the repository <a href='%s'>%s</a>.",
			senderURL,
			sender,
			repoURL,
			repoName,
		)
	default:
		message = fmt.Sprintf(
			"⚠️ <a href='%s'>%s</a> performed an unknown action (<i>%s</i>) on the repository <a href='%s'>%s</a>.",
			senderURL,
			sender,
			action,
			repoURL,
			repoName,
		)
	}

	return message
}
func HandleStatusEvent(event *github.StatusEvent) string {
	state := event.GetState()                   // The state of the status (success, error, pending)
	description := event.GetDescription()       // The description of the status
	commitSHA := event.GetCommit().GetSHA()     // The commit SHA associated with the status
	commitURL := event.GetCommit().GetHTMLURL() // The URL of the commit
	repoName := event.GetRepo().GetFullName()   // The full name of the repository (owner/repo-name)
	repoURL := event.GetRepo().GetHTMLURL()     // The URL of the repository
	sender := event.GetSender().GetLogin()      // The user who triggered the event
	senderURL := event.GetSender().GetHTMLURL() // The URL of the sender

	var message string

	// Format the message based on the status
	switch state {
	case "success":
		message = fmt.Sprintf(
			"✅ The status for commit <a href='%s'>%s</a> in repository <a href='%s'>%s</a> is <b>SUCCESS</b>.\n<i>%s</i> (by <a href='%s'>%s</a>)",
			commitURL,
			commitSHA[:7], // First 7 characters of the commit SHA
			repoURL,
			repoName,
			description,
			senderURL,
			sender,
		)
	case "error":
		message = fmt.Sprintf(
			"❌ The status for commit <a href='%s'>%s</a> in repository <a href='%s'>%s</a> is <b>ERROR</b>.\n<i>%s</i> (by <a href='%s'>%s</a>)",
			commitURL,
			commitSHA[:7],
			repoURL,
			repoName,
			description,
			senderURL,
			sender,
		)
	case "pending":
		message = fmt.Sprintf(
			"⏳ The status for commit <a href='%s'>%s</a> in repository <a href='%s'>%s</a> is <b>PENDING</b>.\n<i>%s</i> (by <a href='%s'>%s</a>)",
			commitURL,
			commitSHA[:7],
			repoURL,
			repoName,
			description,
			senderURL,
			sender,
		)
	default:
		message = fmt.Sprintf(
			"⚠️ The status for commit <a href='%s'>%s</a> in repository <a href='%s'>%s</a> has an <b>unknown state</b>.\n<i>%s</i> (by <a href='%s'>%s</a>)",
			commitURL,
			commitSHA[:7],
			repoURL,
			repoName,
			description,
			senderURL,
			sender,
		)
	}

	return message
}
func HandleWorkflowRunEvent(e *github.WorkflowRunEvent) string {
	workflowName := e.GetWorkflow().GetName()        // The name of the workflow
	runID := e.GetWorkflowRun().GetID()              // The ID of the workflow run
	status := e.GetWorkflowRun().GetStatus()         // The status of the workflow run (queued, in_progress, completed)
	conclusion := e.GetWorkflowRun().GetConclusion() // The conclusion of the workflow run (success, failure, etc.)
	runURL := e.GetWorkflowRun().GetHTMLURL()        // The URL for the workflow run details
	repoName := e.GetRepo().GetFullName()            // The full name of the repository (owner/repo-name)
	repoURL := e.GetRepo().GetHTMLURL()              // The URL of the repository
	sender := e.GetSender().GetLogin()               // The username of the sender who triggered the event
	senderURL := e.GetSender().GetHTMLURL()          // The URL of the sender's GitHub profile

	var message string

	// Build message based on workflow run status and conclusion
	switch status {
	case "queued":
		message = fmt.Sprintf(
			"🔄 Workflow run '<b>%s</b>' is <b>QUEUED</b> in repository <a href='%s'>%s</a>.\nRun ID: <code>%d</code> (by <a href='%s'>%s</a>)\nDetails: <a href='%s'>View Run</a>",
			workflowName,
			repoURL,
			repoName,
			runID,
			senderURL,
			sender,
			runURL,
		)
	case "in_progress":
		message = fmt.Sprintf(
			"⏳ Workflow run '<b>%s</b>' is <b>IN PROGRESS</b> in repository <a href='%s'>%s</a>.\nRun ID: <code>%d</code> (by <a href='%s'>%s</a>)\nDetails: <a href='%s'>View Run</a>",
			workflowName,
			repoURL,
			repoName,
			runID,
			senderURL,
			sender,
			runURL,
		)
	case "completed":
		switch conclusion {
		case "success":
			message = fmt.Sprintf(
				"✅ Workflow run '<b>%s</b>' COMPLETED successfully in repository <a href='%s'>%s</a>.\nRun ID: <code>%d</code> (by <a href='%s'>%s</a>)\nDetails: <a href='%s'>View Run</a>",
				workflowName,
				repoURL,
				repoName,
				runID,
				senderURL,
				sender,
				runURL,
			)
		case "failure":
			message = fmt.Sprintf(
				"❌ Workflow run '<b>%s</b>' FAILED in repository <a href='%s'>%s</a>.\nRun ID: <code>%d</code> (by <a href='%s'>%s</a>)\nDetails: <a href='%s'>View Run</a>",
				workflowName,
				repoURL,
				repoName,
				runID,
				senderURL,
				sender,
				runURL,
			)
		case "neutral":
			message = fmt.Sprintf(
				"⚖️ Workflow run '<b>%s</b>' ended with <b>NEUTRAL</b> conclusion in repository <a href='%s'>%s</a>.\nRun ID: <code>%d</code> (by <a href='%s'>%s</a>)\nDetails: <a href='%s'>View Run</a>",
				workflowName,
				repoURL,
				repoName,
				runID,
				senderURL,
				sender,
				runURL,
			)
		case "cancelled":
			message = fmt.Sprintf(
				"⛔ Workflow run '<b>%s</b>' was CANCELLED in repository <a href='%s'>%s</a>.\nRun ID: <code>%d</code> (by <a href='%s'>%s</a>)\nDetails: <a href='%s'>View Run</a>",
				workflowName,
				repoURL,
				repoName,
				runID,
				senderURL,
				sender,
				runURL,
			)
		default:
			message = fmt.Sprintf(
				"⚠️ Workflow run '<b>%s</b>' COMPLETED with an unknown conclusion in repository <a href='%s'>%s</a>.\nRun ID: <code>%d</code> (by <a href='%s'>%s</a>)\nDetails: <a href='%s'>View Run</a>",
				workflowName,
				repoURL,
				repoName,
				runID,
				senderURL,
				sender,
				runURL,
			)
		}
	default:
		message = fmt.Sprintf(
			"⚠️ Workflow run '<b>%s</b>' has an UNKNOWN status in repository <a href='%s'>%s</a>.\nRun ID: <code>%d</code> (by <a href='%s'>%s</a>)\nDetails: <a href='%s'>View Run</a>",
			workflowName,
			repoURL,
			repoName,
			runID,
			senderURL,
			sender,
			runURL,
		)
	}

	return message
}
func HandleWorkflowJobEvent(e *github.WorkflowJobEvent) string {
	jobName := e.GetWorkflowJob().GetName()
	jobID := e.GetWorkflowJob().GetID()
	runID := e.GetWorkflowJob().GetRunID()
	status := e.GetWorkflowJob().GetStatus()
	conclusion := e.GetWorkflowJob().GetConclusion()
	jobURL := fmt.Sprintf("<a href='%s'>%s</a>", e.GetWorkflowJob().GetHTMLURL(), "Job Details")
	repoName := e.GetRepo().GetFullName()
	sender := e.GetSender().GetLogin()

	var message string

	// Build message based on workflow job status and conclusion
	switch status {
	case "queued":
		message = fmt.Sprintf(
			"🔄 Job '%s' is QUEUED in workflow run ID %d in repository *%s*. Job ID: %d (by %s)\nDetails: %s",
			jobName, runID, repoName, jobID, sender, jobURL,
		)
	case "in_progress":
		message = fmt.Sprintf(
			"⏳ Job '%s' is IN PROGRESS in workflow run ID %d in repository *%s*. Job ID: %d (by %s)\nDetails: %s",
			jobName, runID, repoName, jobID, sender, jobURL,
		)
	case "completed":
		switch conclusion {
		case "success":
			message = fmt.Sprintf(
				"✅ Job '%s' COMPLETED successfully in workflow run ID %d in repository *%s*. Job ID: %d (by %s)\nDetails: %s",
				jobName, runID, repoName, jobID, sender, jobURL,
			)
		case "failure":
			message = fmt.Sprintf(
				"❌ Job '%s' FAILED in workflow run ID %d in repository *%s*. Job ID: %d (by %s)\nDetails: %s",
				jobName, runID, repoName, jobID, sender, jobURL,
			)
		case "neutral":
			message = fmt.Sprintf(
				"⚖️ Job '%s' ended with NEUTRAL conclusion in workflow run ID %d in repository *%s*. Job ID: %d (by %s)\nDetails: %s",
				jobName, runID, repoName, jobID, sender, jobURL,
			)
		case "cancelled":
			message = fmt.Sprintf(
				"⛔ Job '%s' was CANCELLED in workflow run ID %d in repository *%s*. Job ID: %d (by %s)\nDetails: %s",
				jobName, runID, repoName, jobID, sender, jobURL,
			)
		default:
			message = fmt.Sprintf(
				"⚠️ Job '%s' COMPLETED with an unknown conclusion in workflow run ID %d in repository *%s*. Job ID: %d (by %s)\nDetails: %s",
				jobName, runID, repoName, jobID, sender, jobURL,
			)
		}
	default:
		message = fmt.Sprintf(
			"⚠️ Job '%s' has an UNKNOWN status ('%s') in workflow run ID %d in repository *%s*. Job ID: %d (by %s)\nDetails: %s",
			jobName, status, runID, repoName, jobID, sender, jobURL,
		)
	}

	return message
}
func HandleWorkflowDispatchEvent(e *github.WorkflowDispatchEvent) string {
	// Extract workflow dispatch event details
	repoName := e.GetRepo().GetFullName() // Full repository name (owner/repo-name)
	sender := e.GetSender().GetLogin()    // Username of the sender who triggered the event
	workflowName := e.GetWorkflow()       // Get workflow name safely
	eventType := "workflow_dispatch"      // Event type for manual workflow dispatch
	ref := e.GetRef()                     // Get ref safely (branch or tag)

	// Provide default if no workflow name is available
	if workflowName == "" {
		workflowName = "(Unnamed workflow)"
	}

	// Parse inputs if provided
	var inputs string
	if e.Inputs != nil {
		var inputsMap map[string]interface{}
		if err := json.Unmarshal(e.Inputs, &inputsMap); err == nil {
			// Convert inputs map to a formatted string
			var formattedInputs []string
			for key, value := range inputsMap {
				formattedInputs = append(formattedInputs, fmt.Sprintf("<b>%s</b>: <i>%s</i>", key, html.EscapeString(fmt.Sprintf("%v", value))))
			}
			inputs = strings.Join(formattedInputs, ", ")
		} else {
			inputs = "<i>(Invalid JSON inputs)</i>"
		}
	} else {
		inputs = "<i>No inputs provided</i>"
	}

	// Build the message using HTML
	message := fmt.Sprintf(
		"🔧 <b>Workflow</b> '%s' has been manually triggered in repository <b>%s</b> by <b>%s</b>.\n"+
			"<b>Event:</b> %s\n"+
			"<b>Ref:</b> %s\n"+
			"<b>Inputs:</b> %s",
		workflowName,
		repoName,
		sender,
		eventType,
		ref,
		inputs,
	)

	return message
}
func HandleTeamAddEvent(e *github.TeamAddEvent) string {
	// Extract team add event details
	teamName := e.GetTeam().GetName()     // Team name
	repoName := e.GetRepo().GetFullName() // Repository full name (owner/repo-name)
	orgName := e.GetOrg().GetLogin()      // Organization name
	sender := e.GetSender().GetLogin()    // Username of the sender who triggered the event

	// Handle missing values
	if teamName == "" {
		teamName = "(Unnamed team)"
	}
	if repoName == "" {
		repoName = "(Unknown repository)"
	}
	if orgName == "" {
		orgName = "(Unknown organization)"
	}
	if sender == "" {
		sender = "(Unknown sender)"
	}

	// Sanitize values to avoid potential HTML issues
	teamName = html.EscapeString(teamName)
	repoName = html.EscapeString(repoName)
	orgName = html.EscapeString(orgName)
	sender = html.EscapeString(sender)

	// Build the message using HTML formatting
	message := fmt.Sprintf(
		"👥 <b>Team</b> '%s' has been added to repository <b>%s</b> in the organization <b>%s</b> by <b>%s</b>.",
		teamName,
		repoName,
		orgName,
		sender,
	)

	return message
}
func HandleTeamEvent(e *github.TeamEvent) string {
	// Extract team event details
	action := e.GetAction()            // Action like "created", "edited", "deleted"
	teamName := e.GetTeam().GetName()  // Team name
	orgName := e.GetOrg().GetLogin()   // Organization name
	sender := e.GetSender().GetLogin() // Username of the sender who triggered the event

	// Handle missing values
	if teamName == "" {
		teamName = "(Unnamed team)"
	}
	if orgName == "" {
		orgName = "(Unknown organization)"
	}
	if sender == "" {
		sender = "(Unknown sender)"
	}

	// Sanitize values to avoid potential HTML issues
	teamName = html.EscapeString(teamName)
	orgName = html.EscapeString(orgName)
	sender = html.EscapeString(sender)

	// Build the message using HTML formatting
	var message string
	switch action {
	case "created":
		message = fmt.Sprintf("🎉 <b>Team</b> '%s' has been created in the organization <b>%s</b> by <b>%s</b>.", teamName, orgName, sender)
	case "edited":
		message = fmt.Sprintf("✏️ <b>Team</b> '%s' has been edited in the organization <b>%s</b> by <b>%s</b>.", teamName, orgName, sender)
	case "deleted":
		message = fmt.Sprintf("❌ <b>Team</b> '%s' has been deleted from the organization <b>%s</b> by <b>%s</b>.", teamName, orgName, sender)
	default:
		message = fmt.Sprintf("⚙️ <b>Team</b> '%s' has undergone an event in the organization <b>%s</b> by <b>%s</b>: %s.", teamName, orgName, sender, action)
	}

	return message
}

func HandleStarEvent(e *github.StarEvent) string {
	// Extract star event details
	repoName := e.GetRepo().GetFullName() // Repository full name (owner/repo-name)
	userName := e.GetSender().GetLogin()  // Username of the person who starred the repo
	repoName = html.EscapeString(repoName)
	userName = html.EscapeString(userName)

	message := fmt.Sprintf("⭐ <b>%s</b> has starred the repository <b>%s</b>.", userName, repoName)
	return message
}

func HandleRepositoryDispatchEvent(e *github.RepositoryDispatchEvent) string {
	// Extract event details
	action := e.GetAction()               // Action performed
	branch := e.Branch                    // Branch where the event occurred
	clientPayload := e.ClientPayload      // Custom payload data
	repoName := e.GetRepo().GetFullName() // Repository name (owner/repo-name)
	organization := e.GetOrg().GetLogin() // Organization name (if applicable)
	sender := e.GetSender().GetLogin()    // User responsible for the event

	// Decode the client payload if needed
	var payloadMap map[string]interface{}
	if clientPayload != nil {
		if err := json.Unmarshal(clientPayload, &payloadMap); err != nil {
			payloadMap = map[string]interface{}{
				"error": "Unable to parse client payload",
			}
		}
	}

	// Convert payloadMap to a readable string (JSON formatted)
	var payloadStr string
	if payloadBytes, err := json.MarshalIndent(payloadMap, "", "  "); err == nil {
		payloadStr = string(payloadBytes)
	} else {
		payloadStr = "Invalid Payload Data"
	}

	// Build the response message
	message := fmt.Sprintf(
		"📦 Repository Dispatch Event triggered for repository *%s* by %s.\n"+
			"🔧 Action: %s\n"+
			"🌿 Branch: %s\n"+
			"🏢 Organization: %s\n"+
			"📋 Client Payload:\n%s",
		repoName,
		sender,
		action,
		branchOrDefault(branch),
		organization,
		payloadStr, // Include nicely formatted payload
	)

	return message
}

// Helper function to handle branch field
func branchOrDefault(branch *string) string {
	if branch != nil {
		return *branch
	}
	return "default branch"
}
func HandlePullRequestReviewCommentEvent(e *github.PullRequestReviewCommentEvent) string {
	// Extract details
	action := e.GetAction()                    // Action performed on the comment
	repoName := e.GetRepo().GetFullName()      // Repository full name (owner/repo-name)
	sender := e.GetSender()                    // User responsible for the event
	comment := e.GetComment().GetBody()        // Comment body
	commentURL := e.GetComment().GetHTMLURL()  // Link to the comment
	prNumber := e.GetPullRequest().GetNumber() // Pull request number
	prTitle := e.GetPullRequest().GetTitle()   // Pull request title
	prURL := e.GetPullRequest().GetHTMLURL()   // Pull request link
	orgName := e.GetOrg().GetLogin()           // Organization name (if applicable)

	// Determine sender or organization context
	var actor string
	if sender != nil {
		actor = sender.GetLogin()
	} else if orgName != "" {
		actor = fmt.Sprintf("Organization: %s", orgName)
	} else {
		actor = "Unknown Actor"
	}

	// Build response message with HTML tags for easy access
	message := fmt.Sprintf(
		"💬 A Pull Request Review Comment event occurred in repository <b>%s</b>.\n"+
			"👤 Actor: <b>%s</b>\n"+
			"🔧 Action: <b>%s</b>\n"+
			"🔗 Pull Request #%d: <a href='%s'>%s</a>\n"+
			"💡 Comment: <i>%s</i>\n"+
			"🌐 <a href='%s'>View Comment</a>",
		repoName,
		actor,
		action,
		prNumber,
		prURL,
		prTitle,
		truncateComment(comment, 169), // Truncate the comment for easy readability
		commentURL,
	)

	return message
}

// Helper function to truncate comments to fit within a reasonable size
func truncateComment(comment string, maxLength int) string {
	if len(comment) > maxLength {
		return comment[:maxLength] + "..." // Truncate and append "..." for brevity
	}
	return comment
}

func HandlePullRequestReviewEvent(e *github.PullRequestReviewEvent) string {
	// Extract event details
	action := e.GetAction()                    // Action performed on the review (e.g., submitted, edited, dismissed)
	repoName := e.GetRepo().GetFullName()      // Repository full name (owner/repo-name)
	sender := e.GetSender()                    // User responsible for the event
	review := e.GetReview()                    // Pull request review object
	reviewState := review.GetState()           // State of the review (e.g., approved, changes_requested)
	reviewBody := review.GetBody()             // Review body text
	reviewURL := review.GetHTMLURL()           // Link to the review
	prNumber := e.GetPullRequest().GetNumber() // Pull request number
	prTitle := e.GetPullRequest().GetTitle()   // Pull request title
	prURL := e.GetPullRequest().GetHTMLURL()   // Pull request link
	orgName := e.Organization.GetLogin()       // Organization name (if applicable)

	// Determine sender or organization context
	var actor string
	if sender != nil {
		actor = sender.GetLogin()
	} else if orgName != "" {
		actor = fmt.Sprintf("Organization: %s", orgName)
	} else {
		actor = "Unknown Actor"
	}

	// Build response message with HTML tags for easier access
	message := fmt.Sprintf(
		"🔍 A Pull Request Review event occurred in repository <b>%s</b>.\n"+
			"👤 Actor: <b>%s</b>\n"+
			"🔧 Action: <b>%s</b>\n"+
			"🌟 Review State: <b>%s</b>\n"+
			"🔗 Pull Request #%d: <a href='%s'>%s</a>\n"+
			"💡 Review Comment: <i>%s</i>\n"+
			"🌐 <a href='%s'>View Review</a>",
		repoName,
		actor,
		action,
		reviewState,
		prNumber,
		prURL,
		prTitle,
		truncateComment(reviewBody, 150), // Helper to truncate long review comments
		reviewURL,
	)

	return message
}

func HandlePingEvent(e *github.PingEvent) string {
	var responseMessage string

	// Zen Message
	if e.Zen != nil {
		responseMessage += fmt.Sprintf("<b>Zen message:</b> %s\n", *e.Zen)
	}

	// Repository Details
	if e.Repo != nil {
		responseMessage += fmt.Sprintf("<b>Repository:</b> <a href='https://github.com/%s'>%s</a>\n", *e.Repo.FullName, *e.Repo.Name)
		responseMessage += fmt.Sprintf("<i>Description:</i> %s\n", *e.Repo.Description)
	}

	// Sender Information
	if e.Sender != nil {
		responseMessage += fmt.Sprintf("<b>Sender:</b> %s\n", *e.Sender.Login)
	}

	// Organization Information
	if e.Org != nil {
		responseMessage += fmt.Sprintf("<b>Organization:</b> %s\n", *e.Org.Login)
	}

	// Final Confirmation
	responseMessage += "<b>Webhook ping received successfully!</b>"

	return responseMessage
}

func HandlePageBuildEvent(e *github.PageBuildEvent) string {
	var responseMessage string

	// Page Build Details
	if e.Build != nil {
		responseMessage += fmt.Sprintf("<b>Page Build ID:</b> %d\n", *e.ID)
		if e.Build.Status != nil {
			responseMessage += fmt.Sprintf("<b>Build Status:</b> %s\n", *e.Build.Status)
		}
		if e.Build.Error != nil {
			responseMessage += fmt.Sprintf("<b>Build Error:</b> %v\n", *e.Build.Error)
		}
	}

	// Repository Details
	if e.Repo != nil {
		responseMessage += fmt.Sprintf("<b>Repository:</b> <a href='https://github.com/%s'>%s</a>\n", *e.Repo.FullName, *e.Repo.Name)
		responseMessage += fmt.Sprintf("<i>Description:</i> %s\n", *e.Repo.Description)
	}

	// Sender Information
	if e.Sender != nil {
		responseMessage += fmt.Sprintf("<b>Sender:</b> %s\n", *e.Sender.Login)
	}

	// Organization Information
	if e.Org != nil {
		responseMessage += fmt.Sprintf("<b>Organization:</b> %s\n", *e.Org.Login)
	}

	// Final Confirmation
	responseMessage += "<b>Page build event handled successfully!</b>"

	return responseMessage
}
func HandlePackageEvent(e *github.PackageEvent) string {
	var responseMessage string

	// Package details
	if e.Package != nil && e.Package.Name != nil {
		responseMessage += fmt.Sprintf("<b>Package Name:</b> %s\n", *e.Package.Name)
	}

	// Repository details
	if e.Repo != nil && e.Repo.Name != nil {
		responseMessage += fmt.Sprintf("<b>Repository:</b> <a href='https://github.com/%s'>%s</a>\n", *e.Repo.FullName, *e.Repo.Name)
	}

	// Sender details
	if e.Sender != nil && e.Sender.Login != nil {
		responseMessage += fmt.Sprintf("<b>Sender:</b> %s\n", *e.Sender.Login)
	}

	// Organization details
	if e.Org != nil && e.Org.Login != nil {
		responseMessage += fmt.Sprintf("<b>Organization:</b> %s\n", *e.Org.Login)
	}

	// Fallback message
	if responseMessage == "" {
		responseMessage = "<b>No details available for the package event.</b>"
	} else {
		responseMessage += "<b>Package event handled successfully!</b>"
	}

	return responseMessage
}
func HandleOrgBlockEvent(e *github.OrgBlockEvent) string {
	var responseMessage string

	// Blocked User details
	if e.BlockedUser != nil && e.BlockedUser.Login != nil {
		responseMessage += fmt.Sprintf("<b>Blocked User:</b> %s\n", *e.BlockedUser.Login)
	}

	// Sender details
	if e.Sender != nil && e.Sender.Login != nil {
		responseMessage += fmt.Sprintf("<b>Sender:</b> %s\n", *e.Sender.Login)
	}

	// Fallback message
	if responseMessage == "" {
		responseMessage = "<b>No details available for the organization block event.</b>"
	} else {
		responseMessage += "<b>Organization block event handled successfully!</b>"
	}

	return responseMessage
}
func HandleOrganizationEvent(e *github.OrganizationEvent) string {
	var responseMessage string

	// Action details
	if e.Action != nil {
		responseMessage += fmt.Sprintf("<b>Action:</b> %s\n", *e.Action)
	}

	// Sender details
	if e.Sender != nil && e.Sender.Login != nil {
		responseMessage += fmt.Sprintf("<b>Sender:</b> %s\n", *e.Sender.Login)
	}

	// Fallback message
	if responseMessage == "" {
		responseMessage = "<b>No details available for the organization event.</b>"
	} else {
		responseMessage += "<b>Organization event handled successfully!</b>"
	}

	return responseMessage
}

func HandleMilestoneEvent(e *github.MilestoneEvent) string {
	var responseMessage string

	// Milestone details
	if e.Milestone != nil {
		if e.Milestone.Title != nil {
			responseMessage += fmt.Sprintf("<b>Milestone Title:</b> %s\n", *e.Milestone.Title)
		}
		if e.Milestone.Description != nil {
			responseMessage += fmt.Sprintf("<b>Description:</b> %s\n", *e.Milestone.Description)
		}
	}

	// Action details
	if e.Action != nil {
		responseMessage += fmt.Sprintf("<b>Action:</b> %s\n", *e.Action)
	}

	// Sender details
	if e.Sender != nil && e.Sender.Login != nil {
		responseMessage += fmt.Sprintf("<b>Sender:</b> %s\n", *e.Sender.Login)
	}

	// Fallback message
	if responseMessage == "" {
		responseMessage = "<b>No details available for the milestone event.</b>"
	} else {
		responseMessage += "<b>Milestone event handled successfully!</b>"
	}

	return responseMessage
}

func HandleMetaEvent(e *github.MetaEvent) string {
	var responseMessage string
	if e.HookID != nil {
		responseMessage += fmt.Sprintf("Hook ID: %d\n", *e.HookID)
	}

	if e.Repo != nil && e.Repo.Name != nil {
		responseMessage += fmt.Sprintf("Repository: %s\n", *e.Repo.Name)
	}

	if e.Sender != nil && e.Sender.Login != nil {
		responseMessage += fmt.Sprintf("Sender: %s\n", *e.Sender.Login)
	}

	if e.Org != nil && e.Org.Login != nil {
		responseMessage += fmt.Sprintf("Organization: %s\n", *e.Org.Login)
	}

	if e.Installation != nil && e.Installation.ID != nil {
		responseMessage += fmt.Sprintf("Installation ID: %d\n", *e.Installation.ID)
	}

	if responseMessage == "" {
		responseMessage = "No details available for the meta event."
	} else {
		responseMessage += "Delete event handled successfully!"
	}

	return responseMessage
}

func HandleMembershipEvent(e *github.MembershipEvent) string {
	if e == nil {
		return "<b>No membership event data available.</b>"
	}

	var response strings.Builder

	// Action
	if e.Action != nil {
		response.WriteString(fmt.Sprintf("<b>Action:</b> %s\n", *e.Action))
	}

	// Scope (team or repository collaborators)
	if e.Scope != nil {
		response.WriteString(fmt.Sprintf("<b>Scope:</b> %s\n", *e.Scope))
	}

	// Member Information
	if e.Member != nil && e.Member.Login != nil {
		response.WriteString(fmt.Sprintf("<b>Member:</b> %s\n", *e.Member.Login))
	}

	// Team Information (if available)
	if e.Team != nil {
		if e.Team.Name != nil {
			response.WriteString(fmt.Sprintf("<b>Team:</b> %s\n", *e.Team.Name))
		}
		if e.Team.ID != nil {
			response.WriteString(fmt.Sprintf("<b>Team ID:</b> %d\n", *e.Team.ID))
		}
		if e.Team.Description != nil {
			response.WriteString(fmt.Sprintf("<b>Team Description:</b> %s\n", *e.Team.Description))
		}
	}

	// Sender Information
	if e.Sender != nil && e.Sender.Login != nil {
		response.WriteString(fmt.Sprintf("<b>Action by:</b> %s\n", *e.Sender.Login))
	}

	return response.String()
}

func HandleDeploymentEvent(e *github.DeploymentEvent) string {
	var responseMessage string

	if e.Deployment != nil {
		responseMessage += fmt.Sprintf("Deployment ID: %d\n", *e.Deployment.ID)
		if e.Deployment.Description != nil {
			responseMessage += fmt.Sprintf("Description: %s\n", *e.Deployment.Description)
		}
	}

	if e.Repo != nil && e.Repo.Name != nil {
		responseMessage += fmt.Sprintf("Repository: %s\n", *e.Repo.Name)
	}

	if e.Sender != nil && e.Sender.Login != nil {
		responseMessage += fmt.Sprintf("Sender: %s\n", *e.Sender.Login)
	}

	if e.Installation != nil && e.Installation.ID != nil {
		responseMessage += fmt.Sprintf("Installation ID: %d\n", *e.Installation.ID)
	}

	if responseMessage == "" {
		responseMessage = "No details available for the deployment event."
	}

	return responseMessage
}

func HandleLabelEvent(e *github.LabelEvent) string {
	if e == nil {
		return "No label event data available."
	}

	var response strings.Builder

	// Action
	if e.Action != nil {
		response.WriteString(fmt.Sprintf("Action: %s\n", *e.Action))
	}

	// Label Details
	if e.Label != nil {
		if e.Label.Name != nil {
			response.WriteString(fmt.Sprintf("Label Name: %s\n", *e.Label.Name))
		}
		if e.Label.Color != nil {
			response.WriteString(fmt.Sprintf("Label Color: #%s\n", *e.Label.Color))
		}
		if e.Label.Description != nil {
			response.WriteString(fmt.Sprintf("Label Description: %s\n", *e.Label.Description))
		}
	}

	// Changes
	if e.Changes != nil {
		response.WriteString("Changes:\n")
		if e.Changes.Title != nil && e.Changes.Title.From != nil {
			response.WriteString(fmt.Sprintf("  Previous Title: %s\n", *e.Changes.Title.From))
		}
		if e.Changes.Body != nil && e.Changes.Body.From != nil {
			response.WriteString(fmt.Sprintf("  Previous Body: %s\n", *e.Changes.Body.From))
		}
	}
	return response.String()
}

func HandleMarketplacePurchaseEvent(e *github.MarketplacePurchaseEvent) string {
	if e == nil {
		return "<b>No marketplace purchase event data available.</b>"
	}

	var response strings.Builder

	// Action
	if e.Action != nil {
		response.WriteString(fmt.Sprintf("<b>Action:</b> %s\n", *e.Action))
	}

	// Marketplace Plan Info
	if e.MarketplacePurchase != nil {
		if e.MarketplacePurchase.Plan != nil && e.MarketplacePurchase.Plan.Name != nil {
			response.WriteString(fmt.Sprintf("<b>Plan Name:</b> %s\n", *e.MarketplacePurchase.Plan.Name))
		}
		if e.MarketplacePurchase.BillingCycle != nil {
			response.WriteString(fmt.Sprintf("<b>Billing Cycle:</b> %s\n", *e.MarketplacePurchase.BillingCycle))
		}
		if e.MarketplacePurchase.UnitCount != nil {
			response.WriteString(fmt.Sprintf("<b>Unit Count:</b> %d\n", *e.MarketplacePurchase.UnitCount))
		}
		if e.MarketplacePurchase.NextBillingDate != nil {
			response.WriteString(fmt.Sprintf("<b>Next Billing Date:</b> %s\n", e.MarketplacePurchase.NextBillingDate.String()))
		}
	}

	// Account Info
	if e.MarketplacePurchase.Account != nil {
		if e.MarketplacePurchase.Account.Login != nil {
			response.WriteString(fmt.Sprintf("<b>Account Login:</b> %s\n", *e.MarketplacePurchase.Account.Login))
		}
		if e.MarketplacePurchase.Account.Type != nil {
			response.WriteString(fmt.Sprintf("<b>Account Type:</b> %s\n", *e.MarketplacePurchase.Account.Type))
		}
	}

	// Sender Info
	if e.Sender != nil && e.Sender.Login != nil {
		response.WriteString(fmt.Sprintf("<b>Sender:</b> %s\n", *e.Sender.Login))
	}

	return response.String()
}

func HandleGollumEvent(e *github.GollumEvent) string {
	if e == nil {
		return "No Gollum event data available."
	}

	var response strings.Builder

	// Repository Info
	if e.Repo != nil && e.Repo.Name != nil {
		response.WriteString(fmt.Sprintf("Repository: %s\n", *e.Repo.Name))
	}

	// Sender Info
	if e.Sender != nil && e.Sender.Login != nil {
		response.WriteString(fmt.Sprintf("Sender: %s\n", *e.Sender.Login))
	}

	// Pages Info
	if e.Pages != nil && len(e.Pages) > 0 {
		response.WriteString("Wiki Pages:\n")
		for _, page := range e.Pages {
			if page.Title != nil {
				response.WriteString(fmt.Sprintf("- Title: %s\n", *page.Title))
			}
			if page.Action != nil {
				response.WriteString(fmt.Sprintf("  Action: %s\n", *page.Action))
			}
			if page.HTMLURL != nil {
				response.WriteString(fmt.Sprintf("  URL: %s\n", *page.HTMLURL))
			}
		}
	}

	return response.String()
}

func HandleDeployKeyEvent(e *github.DeployKeyEvent) string {
	if e == nil {
		return "No deploy key event data available."
	}

	var response strings.Builder

	// Action
	if e.Action != nil {
		response.WriteString(fmt.Sprintf("Action: %s\n", *e.Action))
	}

	// Deploy Key Info
	if e.Key != nil {
		if e.Key.Title != nil {
			response.WriteString(fmt.Sprintf("Deploy Key Title: %s\n", *e.Key.Title))
		}
		if e.Key.URL != nil {
			response.WriteString(fmt.Sprintf("Deploy Key URL: %s\n", *e.Key.URL))
		}
	}

	// Repository Info
	if e.Repo != nil && e.Repo.Name != nil {
		response.WriteString(fmt.Sprintf("Repository: %s\n", *e.Repo.Name))
	}

	// Sender Info
	if e.Sender != nil && e.Sender.Login != nil {
		response.WriteString(fmt.Sprintf("Sender: %s\n", *e.Sender.Login))
	}

	return response.String()
}

func HandleCheckSuiteEvent(e *github.CheckSuiteEvent) string {
	if e == nil {
		return "No check suite event data available."
	}

	var response strings.Builder

	// Action
	if e.Action != nil {
		response.WriteString(fmt.Sprintf("Action: %s\n", *e.Action))
	}

	// Check Suite Info
	if e.CheckSuite != nil {
		if e.CheckSuite.Status != nil {
			response.WriteString(fmt.Sprintf("Status: %s\n", *e.CheckSuite.Status))
		}
		if e.CheckSuite.Conclusion != nil {
			response.WriteString(fmt.Sprintf("Conclusion: %s\n", *e.CheckSuite.Conclusion))
		}
		if e.CheckSuite.URL != nil {
			response.WriteString(fmt.Sprintf("Details URL: %s\n", *e.CheckSuite.URL))
		}
	}

	// Repository Info
	if e.Repo != nil && e.Repo.Name != nil {
		response.WriteString(fmt.Sprintf("Repository: %s\n", *e.Repo.Name))
	}

	// Sender Info
	if e.Sender != nil && e.Sender.Login != nil {
		response.WriteString(fmt.Sprintf("Sender: %s\n", *e.Sender.Login))
	}

	return response.String()
}

func HandleCheckRunEvent(e *github.CheckRunEvent) string {
	if e == nil {
		return "No check run event data available."
	}

	var response strings.Builder

	// Action
	if e.Action != nil {
		response.WriteString(fmt.Sprintf("Action: %s\n", *e.Action))
	}

	// Check Run Details
	if e.CheckRun != nil {
		if e.CheckRun.Name != nil {
			response.WriteString(fmt.Sprintf("Check Run Name: %s\n", *e.CheckRun.Name))
		}
		if e.CheckRun.Status != nil {
			response.WriteString(fmt.Sprintf("Status: %s\n", *e.CheckRun.Status))
		}
		if e.CheckRun.Conclusion != nil {
			response.WriteString(fmt.Sprintf("Conclusion: %s\n", *e.CheckRun.Conclusion))
		}
		if e.CheckRun.StartedAt != nil {
			response.WriteString(fmt.Sprintf("Started At: %s\n", e.CheckRun.StartedAt.String()))
		}
		if e.CheckRun.CompletedAt != nil {
			response.WriteString(fmt.Sprintf("Completed At: %s\n", e.CheckRun.CompletedAt.String()))
		}
	}

	// Repository Information
	if e.Repo != nil && e.Repo.Name != nil {
		response.WriteString(fmt.Sprintf("Repository: %s\n", *e.Repo.Name))
	}

	// Sender Information
	if e.Sender != nil && e.Sender.Login != nil {
		response.WriteString(fmt.Sprintf("Sender: %s\n", *e.Sender.Login))
	}

	return response.String()
}

func HandleDeploymentStatusEvent(e *github.DeploymentStatusEvent) string {
	var responseMessage string

	// Include deployment status details
	if e.DeploymentStatus != nil {
		responseMessage += fmt.Sprintf("Deployment Status: %s\n", *e.DeploymentStatus.State)
		if e.DeploymentStatus.Description != nil {
			responseMessage += fmt.Sprintf("Description: %s\n", *e.DeploymentStatus.Description)
		}
	}

	// Include repository details
	if e.Repo != nil && e.Repo.Name != nil {
		responseMessage += fmt.Sprintf("Repository: %s\n", *e.Repo.Name)
	}

	// Include the sender information
	if e.Sender != nil && e.Sender.Login != nil {
		responseMessage += fmt.Sprintf("Sender: %s\n", *e.Sender.Login)
	}

	// Add installation details if available
	if e.Installation != nil && e.Installation.ID != nil {
		responseMessage += fmt.Sprintf("Installation ID: %d\n", *e.Installation.ID)
	}

	// Default message if no other details are available
	if responseMessage == "" {
		responseMessage = "No details available for the deployment status event."
	}

	return responseMessage
}

func HandleSecurityAdvisoryEvent(e *github.SecurityAdvisoryEvent) string {
	if e == nil {
		return "No security advisory event data available."
	}

	var response strings.Builder

	// Action
	if e.Action != nil {
		response.WriteString(fmt.Sprintf("Action: %s\n", *e.Action))
	}

	// Security Advisory Details
	if e.SecurityAdvisory != nil {
		advisory := e.SecurityAdvisory

		if advisory.Summary != nil {
			response.WriteString(fmt.Sprintf("Summary: %s\n", *advisory.Summary))
		}

		if advisory.Description != nil {
			response.WriteString(fmt.Sprintf("Description: %s\n", *advisory.Description))
		}

		if advisory.Severity != nil {
			response.WriteString(fmt.Sprintf("Severity: %s\n", *advisory.Severity))
		}

		if advisory.CVEID != nil {
			response.WriteString(fmt.Sprintf("CVE ID: %s\n", *advisory.CVEID))
		}

		if advisory.URL != nil {
			response.WriteString(fmt.Sprintf("Advisory URL: %s\n", *advisory.URL))
		}

		if advisory.PublishedAt != nil {
			response.WriteString(fmt.Sprintf("Published At: %s\n", advisory.PublishedAt.String()))
		}

		if advisory.WithdrawnAt != nil {
			response.WriteString(fmt.Sprintf("Withdrawn At: %s\n", advisory.WithdrawnAt.String()))
		}

		if advisory.Author != nil && advisory.Author.Login != nil {
			response.WriteString(fmt.Sprintf("Reported By: %s\n", *advisory.Author.Login))
		}
	}

	// Repository
	if e.Repository != nil && e.Repository.FullName != nil {
		response.WriteString(fmt.Sprintf("Repository: %s\n", *e.Repository.FullName))
	}

	// Organization
	if e.Organization != nil && e.Organization.Login != nil {
		response.WriteString(fmt.Sprintf("Organization: %s\n", *e.Organization.Login))
	}

	// Sender
	if e.Sender != nil && e.Sender.Login != nil {
		response.WriteString(fmt.Sprintf("Sender: %s\n", *e.Sender.Login))
	}

	return response.String()
}
