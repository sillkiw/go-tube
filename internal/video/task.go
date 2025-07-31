package video

// task represents a single conversion task.
type task struct {
	videoPath    string // full path to the source video file
	outputPath   string // full path to the output file (video, audio or thumbnail)
	bitrate      string // e.g. "500k", "2M"
	resolution   string // e.g. "640x360"
	metadataFlag string // e.g. "-2" to strip metadata

	withAudio    bool   // include audio in the video output
	processAudio bool   // only convert the audio track
	audioQuality string // e.g. "64k"

	createThumb bool // extract a thumbnail image
	createMPD   bool // generate a DASH manifest (MPD)

	baseName string // base filename (without extension) for templating
}
