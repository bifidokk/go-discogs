package discogs

type Community struct {
	Contributors []Contributor `json:"contributors"`
	DataQuality  string        `json:"data_quality"`
	Have         int           `json:"have"`
	Rating       Rating        `json:"rating"`
	Status       string        `json:"status"`
	Submitter    Submitter     `json:"submitter"`
	Want         int           `json:"expected"`
}

type Company struct {
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
}

type Contributor struct {
	ResourceURL string `json:"resource_url"`
	Username    string `json:"username"`
}

type Format struct {
	Descriptions []string `json:"descriptions"`
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
}

type Identifier struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	Value       string `json:"value"`
}

type Image struct {
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	ResourceURL string `json:"resource_url"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	URI150      string `json:"uri150"`
}

type LabelSource struct {
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
}

type Rating struct {
	Average float32 `json:"average"`
	Count   int     `json:"count"`
}

type ReleaseArtist struct {
	Anv         string `json:"anv"`
	ID          int    `json:"id"`
	Join        string `json:"join"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
}

type Submitter struct {
	ResourceURL string `json:"resource_url"`
	Username    string `json:"username"`
}

type Track struct {
	Duration     string          `json:"duration"`
	Position     string          `json:"position"`
	Title        string          `json:"title"`
	Type         string          `json:"type_"`
	Extraartists []ReleaseArtist `json:"extraartists,omitempty"`
	Artists      []ReleaseArtist `json:"artists,omitempty"`
}

type Video struct {
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
	Title       string `json:"title"`
	URI         string `json:"uri"`
}

type Series struct {
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
	ThumbnailURL   string `json:"thumbnail_url,omitempty"`
}
