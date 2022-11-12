package github

const URL = "https://api.github.com/repos/ReposForTopics/"

func GetTopicURL(topic *Topic) string {
	return URL + topic.Name + "/topics"
}

type Topic struct {
	Name        string
	DisplayName string `json:display_name`
}
