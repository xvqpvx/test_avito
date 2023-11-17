package request

type SegmentUpdateRequest struct {
	IdSegment int    `json:"id_segment,omitempty"`
	Name      string `json:"name"`
}
