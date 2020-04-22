package exporter

//SocialMedia interface with which implement feed method
type SocialMedia interface {
	Feed() Timeline
}

//Timeline as my data
type Timeline map[string]string

//Timeland for xml
type Timeland struct {
	x, y, z string
}

//Facebook type which implement feed method and fame method
type Facebook struct {
	URL       string
	Username  string
	Followers int
}

//Feed method for our SocialMedia interface
func (f *Facebook) Feed() Timeline {
	return map[string]string{"1": "This is start.NG internship,", "2": "I know you know it,", "3": "but I have to tell you,"}

}

//Linkedin type  which implement feed method
type Linkedin struct {
	URL       string
	Username  string
	Followers int
}

//Feed method for our SocialMedia interface
func (i *Linkedin) Feed() Timeline {
	return map[string]string{"1": "I know you know I need the money,", "2": "Time for business,", "3": "I love my Job"}

}

//Twitter type  which implement feed method
type Twitter struct {
	URL       string
	Username  string
	Followers int
}

//Feed method for our SocialMedia interface
func (t *Twitter) Feed() Timeline {
	return map[string]string{"1": "I know you know I need the money,", "2": "Thanks to the Cyberelf himself", "3": "You are the best and I am happy being here"}
}
