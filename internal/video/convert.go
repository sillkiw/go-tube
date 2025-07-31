package video

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

// StartConvertVideo creates the output directory, builds conversion tasks, and enqueues them.
func (s *Service) StartConvertVideo(filePath, baseName string) {
	// Create the output directory
	outDir := filepath.Join(s.cfg.Video.ConvertPath, baseName)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		s.logger.Error("convert: failed to create output directory",
			slog.String("path", outDir),
			slog.String("error", err.Error()),
		)
		return
	}

	// Build all conversion tasks
	tasks := s.buildConversionTasks(filePath, baseName, outDir)

	// Atomically increment the queue length by the number of tasks
	s.addToQueue(len(tasks))

	// Enqueue each task
	for _, task := range tasks {
		s.taskCh <- task
	}

	// Optionally remove the original file after upload
	if s.cfg.Video.DeleteOriginal {
		if err := os.Remove(filePath); err != nil {
			s.logger.Error("convert: failed to remove original file",
				slog.String("file", filePath),
				slog.String("error", err.Error()),
			)
		}
	}
}

// buildConversionTasks returns a slice of conversion tasks for the given video.
// It creates tasks for multiple profiles: WebM+audio, thumbnails, MP4 at various
// qualities, audio-only, and the final DASH manifest.
func (s *Service) buildConversionTasks(filePath, baseName, outDir string) []task {
	// helper for quickly assembling a task
	mk := func(filename string, withAudio, processAudio, createThumb, createMPD bool, bitrate, resolution string) task {
		return task{
			videoPath:    filePath,
			outputPath:   filepath.Join(outDir, filename),
			bitrate:      bitrate,
			resolution:   resolution,
			metadataFlag: "-2",
			withAudio:    withAudio,
			processAudio: processAudio,
			audioQuality: "64k",
			createThumb:  createThumb,
			createMPD:    createMPD,
			baseName:     baseName,
		}
	}

	return []task{
		// WebM with audio (low resolution)
		mk(
			fmt.Sprintf("low_%s_audio.webm", baseName),
			true,  /* withAudio */
			false, /* processAudio */
			false, /* createThumb */
			false, /* createMPD */
			s.cfg.Video.Bitrates.Low,
			s.cfg.Video.Resolutions.Low,
		),
		// Thumbnail extraction
		mk(
			fmt.Sprintf("thumb_%s.jpeg", baseName),
			false, /* withAudio */
			false, /* processAudio */
			true,  /* createThumb */
			false, /* createMPD */
			"",    /* bitrate unused */
			"",    /* resolution unused */
		),
		// MP4 low quality (video only)
		mk(
			fmt.Sprintf("low_%s.mp4", baseName),
			false,
			false,
			false,
			false,
			s.cfg.Video.Bitrates.Low,
			s.cfg.Video.Resolutions.Low,
		),
		// MP4 medium quality (video only)
		mk(
			fmt.Sprintf("med_%s.mp4", baseName),
			false,
			false,
			false,
			false,
			s.cfg.Video.Bitrates.Med,
			s.cfg.Video.Resolutions.Med,
		),
		// MP4 high quality (video only)
		mk(
			fmt.Sprintf("high_%s.mp4", baseName),
			false,
			false,
			false,
			false,
			s.cfg.Video.Bitrates.High,
			s.cfg.Video.Resolutions.High,
		),
		// Audio-only MP4
		mk(
			fmt.Sprintf("audio_%s.mp4", baseName),
			false,
			true,
			false,
			false,
			s.cfg.Video.Bitrates.High,
			"", /* resolution unused */
		),
		// Final DASH manifest
		mk(
			"output.mpd",
			false,
			false,
			false,
			true,
			"", /* bitrate unused */
			"", /* resolution unused */
		),
	}
}

// convertVideo reads conversion tasks from s.taskCh, executes the appropriate
// FFmpeg (or MP4Box) command for each, and decrements the queue counter when done.
func (s *Service) convertVideo() {
	for task := range s.taskCh {
		// build and run ffmpeg or MP4Box based on t’s flags
		var cmd *exec.Cmd
		if task.createMPD {
			cmd = buildMPDCommand(task)
		} else {
			cmd = s.buildFFmpegCmd(task)
		}

		if err := cmd.Run(); err != nil {
			if task.processAudio {
				s.logger.Info("audio conversion failed, creating noaudio flag",
					slog.String("video", task.videoPath),
					slog.String("err", err.Error()),
				)
				s.create_noaudio_file(task)
			} else {
				s.logger.Error("convert: conversion failed",
					slog.String("video", task.videoPath),
					slog.String("err", err.Error()),
				)
			}
		} else {
			s.logger.Info("convert: conversion succeeded",
				slog.String("output", task.outputPath),
			)
		}
		s.DecreaseQueue(1)
	}
}

func (s *Service) create_noaudio_file(task task) {
	// Create a flag file so that the MPD generator knows about the absence of audio
	flagPath := filepath.Join(task.outputPath, task.baseName+"noaudio.txt")
	if ferr := os.WriteFile(flagPath, nil, 0644); ferr != nil {
		s.logger.Error("failed to write noaudio flag",
			slog.String("flagPath", flagPath),
			slog.String("error", ferr.Error()),
		)
	}
}
