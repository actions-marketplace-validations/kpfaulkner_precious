# Precious
Github action to react to gollum events (wiki updates)

There seems to be plenty of Github actions that react to events and sends notifications to Slack. But very few seem to be filtering.
Precious purely reacts to Github wiki updates, but filters the wiki titles. This way it can be configured to only message slack
when "highly important" pages get updated and not every one (which can be very "spammy")

Simply configure the action with the yaml:

on: gollum
name: Precious
jobs:
  slackNotification:
    name: Slack Notification
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Slack Notification
      uses: kpfaulkner/precious@0.2.0
      env:
        SLACK_MESSAGE: "Wiki updated!!!"
        WIKI_TITLES_TO_ALERT: "home,test page"
        SLACK_WEBHOOK: ${{ secrets.SLACK_WEBHOOK }}

secrets.SLACK_WEBHOOK is a regular Slack webhook ( https://slack.com/help/articles/115005265063-Incoming-Webhooks-for-Slack )

WIKI_TITLES_TO_ALERT is a comma separated list of wiki titles that you want to alert on. If you happen to have a comma in your title, then 
I might need to change things.
