package repos

type RepoRecord struct {
}

type CLI_BorgInfo struct {
	Cache struct {
		Path  string `json:"path"`
		Stats struct {
			TotalChunks       int   `json:"total_chunks"`
			TotalCsize        int64 `json:"total_csize"`
			TotalSize         int64 `json:"total_size"`
			TotalUniqueChunks int   `json:"total_unique_chunks"`
			UniqueCsize       int64 `json:"unique_csize"`
			UniqueSize        int64 `json:"unique_size"`
		} `json:"stats"`
	} `json:"cache"`
	Encryption struct {
		Mode string `json:"mode"`
	} `json:"encryption"`
	Repository struct {
		ID           string `json:"id"`
		LastModified string `json:"last_modified"`
		Location     string `json:"location"`
	} `json:"repository"`
	SecurityDir string `json:"security_dir"`
}

type CLI_BorgList struct {
	Archives []struct {
		Archive  string `json:"archive"`
		Barchive string `json:"barchive"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		Start    string `json:"start"`
		Time     string `json:"time"`
	} `json:"archives"`
	Encryption struct {
		Mode string `json:"mode"`
	} `json:"encryption"`
	Repository struct {
		ID           string `json:"id"`
		LastModified string `json:"last_modified"`
		Location     string `json:"location"`
	} `json:"repository"`
}
