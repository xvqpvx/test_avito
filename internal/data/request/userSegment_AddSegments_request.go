package request

type AddSegmentsRequest struct {
	IdUser           int      `json:"id_user"`
	SegmentsToAdd    []string `json:"segments_to_add,omitempty"`
	SegmentsToDelete []string `json:"segments_to_delete,omitempty"`
	Ttl              string   `json:"ttl,omitempty"`
}
