package response

type UserSegmentResponse struct {
	IdUser         int      `json:"id_user"`
	ActiveSegments []string `json:"active_segments"`
}
