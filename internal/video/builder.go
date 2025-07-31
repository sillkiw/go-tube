package video

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// buildFFmpegCmd builds either a video+audio command or only video,
// either on thumbnail or audio only, depending on the flags in the task.
func (s *Service) buildFFmpegCmd(t task) *exec.Cmd {
	args := []string{"-i", t.videoPath, "-map_metadata", t.metadataFlag, "-threads", s.cfg.Video.FFmpeg.Threads}
	switch {
	case t.withAudio:
		// VP9 + audio → WebM
		args = append(args,
			"-c:v", "libvpx-vp9",
			"-b:v", t.bitrate,
			"-vf", "scale="+t.resolution,
			t.outputPath,
		)
	case t.createThumb:
		// thumbnail → JPEG
		args = append(args,
			"-ss", "00:00:01",
			"-vframes", "1",
			"-s", "640x480",
			"-f", "image2",
			t.outputPath,
		)
	case t.processAudio:
		// audio-only → AAC/mp4
		args = append(args,
			"-c:a", "aac",
			"-b:a", t.audioQuality,
			"-vn",
			"-f", "mp4",
			t.outputPath,
		)
	default:
		// video-only → H264 MP4 + dash fragment flag
		args = append(args,
			"-c:v", "libx264",
			"-level", "4.1",
			"-b:v", t.bitrate,
			"-g", "60",
			"-vf", "scale="+t.resolution,
			"-preset", s.cfg.Video.FFmpeg.Preset,
			"-keyint_min", "60",
			"-sc_threshold", "0",
			"-an",
			"-f", "mp4",
			"-dash", "1",
			t.outputPath,
		)
	}

	return exec.Command("/usr/bin/ffmpeg", args...)
}

// buildMPDCommand constructs an MP4Box command to generate a DASH manifest (MPD)
// from the H264 video segments and optional audio segment.
//
// It expects that video-only segments (“high_<base>.mp4”, “med_<base>.mp4”,
// “low_<base>.mp4”) already exist in the same output directory, and that an
// audio segment (“audio_<base>.mp4”) exists unless a no-audio flag file is present.
//
// The returned *exec.Cmd should be run via cmd.Run().
func buildMPDCommand(t task) *exec.Cmd {
	// Directory where all segments are stored
	dir := filepath.Dir(t.outputPath)
	base := t.baseName // filename without extension

	// Base MPDBox arguments: 2s segments/fragments, on-demand profile
	args := []string{
		"-dash", "2000",
		"-frag", "2000",
		"-rap",
		"-profile", "onDemand",
		"-out", t.outputPath, // this is the .mpd file
	}

	// Add video representations: high, med, low
	for _, quality := range []string{"high", "med", "low"} {
		seg := fmt.Sprintf("%s/%s_%s.mp4#video", dir, quality, base)
		args = append(args, seg)
	}

	// If no-audio flag file is absent, include the audio representation
	flagFile := filepath.Join(dir, base+"noaudio.txt")
	if _, err := os.Stat(flagFile); os.IsNotExist(err) {
		audioSeg := fmt.Sprintf("%s/audio_%s.mp4#audio", dir, base)
		args = append(args, audioSeg)
	}

	// Join into a shell command to support the “#video” syntax
	cmdLine := "MP4Box " + strings.Join(args, " ")
	return exec.Command("/bin/sh", "-c", cmdLine)
}
