package gobalt

import "net/http"

type PostRequest struct {
	// URL must be included in every request
	URL string `json:"url"`

	// VideoQuality if the selected quality isn't available, closest one is used instead.
	// Default VideoQuality1080p
	VideoQuality VideoQuality `json:"vQuality,omitempty"`

	// AudioFormat Format to re-encode audio into. If AudioFormatBest is selected, you get the audio the way it is on service's side.
	// Default AudioFormatMP3
	AudioFormat AudioFormat `json:"audioFormat,omitempty"`

	// AudioBitrate Specifies the birate to use for the audio. Applies only to audio conversion.
	// Default AudioBitrate128
	AudioBitrate AudioBitrate `json:"audioBitrate,omitempty"`

	// FilenameStyle changes the way files are named.
	// Some services don't support rich file names and always use the classic style.
	// Default FilenamePatternClassic
	FilenameStyle FilenameStyle `json:"filenameStyle,omitempty"`

	// DownloadMode selects if to download only the audio, or mute the audio in video tracks
	// Default auto
	DownloadMode DownloadMode `json:"downloadMode,omitempty"`

	// YoutubeVideoCodec applies only for Youtube downloads.
	// Pick YoutubeVideoCodecH264 if you want best compatibility. Pick YoutubeVideoCodecAV1 if you want best quality and efficiency.
	// Default YoutubeVideoCodecH264
	YoutubeVideoCodec YoutubeVideoCodec `json:"youtubeVideoCodec,omitempty"`

	// YoutubeDubLang Specifies the language of audio to download when a youtube video is dubbed.
	// this must be a valid locale tag (BCP 47 tag). Examples: en, ru, cs, ja, es-US.
	// Default is empty (the video original audio track)
	YoutubeDubLang string `json:"youtubeDubLang,omitempty"`

	// AlwaysProxy tunnels all downloads through the processing server, even when not necessary.
	// Default false
	AlwaysProxy bool `json:"alwaysProxy,omitempty"`

	// DisableMetadata disables file metadata when set to true.
	// Default false
	DisableMetadata bool `json:"disableMetadata,omitempty"`

	// TiktokFullAudio download original sound used in a tiktok video.
	// Default false
	TiktokFullAudio bool `json:"tiktokFullAudio,omitempty"`

	// TiktokH265 wheter to download Tiktok 1080p H264 video
	// Default false
	TiktokH265 bool `json:"tiktokH265,omitempty"`

	// TwitterGif encode twitter gif and proper .gif (not mp4)
	// Default false
	TwitterGif bool `json:"twitterGif,omitempty"`

	// YoutubeHLS specified whether to use HLS for downloading video or audio from Youtube.
	YoutubeHLS bool `json:"youtubeHLS,omitempty"`
}

// YoutubeVideoCodec is the codec used to download the video. Only applicable to Youtube.
type YoutubeVideoCodec string

const (
	// YoutubeVideoCodecH264 best support across apps/platforms, average detail level. Max VideoQuality is VideoQuality1080p
	YoutubeVideoCodecH264 YoutubeVideoCodec = "h264"
	// YoutubeVideoCodecAV1 best quality, small file size, most detail. Supports 8k and HDR.
	YoutubeVideoCodecAV1 YoutubeVideoCodec = "av1"
	// YoutubeVideoCodecVP9 same quality as av1, but file size is 2x larger. Supports 4k and HDR.
	YoutubeVideoCodecVP9 YoutubeVideoCodec = "vp9"
)

// VideoQuality is the video resolution to use
type VideoQuality string

const (
	VideoQuality144p  VideoQuality = "144"
	VideoQuality240p  VideoQuality = "240"
	VideoQuality360p  VideoQuality = "360"
	VideoQuality480p  VideoQuality = "480"
	VideoQuality720p  VideoQuality = "720"
	VideoQuality1080p VideoQuality = "1080"
	VideoQuality2160p VideoQuality = "2160"
	VideoQuality4320p VideoQuality = "4320"
	VideoQualityMax   VideoQuality = "max"

	// VideoQuality4k is an alias for VideoQuality2160p
	VideoQuality4K VideoQuality = VideoQuality2160p

	// VideoQuality8k is an alias for VideoQuality4320p
	VideoQuality8K VideoQuality = VideoQuality4320p
)

// AudioFormat is the encoding to use for audio.
type AudioFormat string

const (
	AudioFormatBest AudioFormat = "best"
	AudioFormatMP3  AudioFormat = "mp3"
	AudioFormatOGG  AudioFormat = "ogg"
	AudioFormatWAW  AudioFormat = "wav"
	AudioFormatOpus AudioFormat = "opus"
)

// AudioBitrate bitrate for audio conversion
type AudioBitrate string

const (
	AudioBitrate320 AudioBitrate = "320"
	AudioBitrate256 AudioBitrate = "256"
	AudioBitrate128 AudioBitrate = "128"
	AudioBitrate96  AudioBitrate = "96"
	AudioBitrate64  AudioBitrate = "64"
	AudioBitrate8   AudioBitrate = "8"
)

// FilenameStyle is the way resulting files are named.
type FilenameStyle string

const (
	// FilenameStyleClassic default cobalt file name pattern.
	//
	// Video: youtube_dQw4w9WgXcQ_2560x1440_h264.mp4
	//
	// Audio: youtube_dQw4w9WgXcQ_audio.mp3
	FilenameStyleClassic FilenameStyle = "classic"

	// FilenameStylePretty title and info in brackets.
	//
	// Video: Video Title (1440p, h264, youtube).mp4
	//
	// Audio: Audio Title - Audio Author (soundcloud).mp3
	FilenameStylePretty FilenameStyle = "pretty"

	// FilenameStyleBasic title and basic info in brackets.
	//
	// Video: Video Title (1440p, h264).mp4
	//
	// Audio: Audio Title - Audio Author.mp3
	FilenameStyleBasic FilenameStyle = "basic"

	// FilenameStyleNerdy title and full info in brackets.
	//
	// Video: Video Title (1440p, h264, youtube, dQw4w9WgXcQ).mp4
	//
	// Audio: Audio Title - Audio Author (soundcloud, 1242868615).mp3
	FilenameStyleNerdy FilenameStyle = "nerdy"
)

// DownloadMode is used to include, skip, or only download the audio track
type DownloadMode string

const (
	// DownloadModeAuto keeps all the original media
	DownloadModeAuto DownloadMode = "auto"

	// DownloadModeAudio downloads only the audio
	DownloadModeAudio DownloadMode = "audio"

	// DownloadModeMute skips the audio track in videos
	DownloadModeMute DownloadMode = "mute"
)

type ResponseStatus string

const (
	// ResponseStatusError something went wrong
	ResponseStatusError ResponseStatus = "error"
	// ResponseStatusError you are being redirected to the direct service URL
	ResponseStatusRedirect ResponseStatus = "redirect"
	// ResponseStatusTunnel cobalt is proxying the download for you
	ResponseStatusTunnel ResponseStatus = "tunnel"
	// ResponseStatusPicker we have multiple items to choose from
	ResponseStatusPicker ResponseStatus = "picker"
)

type PostResponse struct {
	client *http.Client

	// Status is the type of response from the API
	Status ResponseStatus `json:"status"`

	// ResponseStatusTunnel and ResponseStatusRedirect fields

	// URL for the cobalt tunnel, or redirect to an external link
	URL string `json:"url,omitempty"`
	// Filename cobalt-generated filename for the file being downloaded
	Filename string `json:"filename,omitempty"`

	// ResponseStatusPicker fields

	// Audio returned when an image slideshow (such as tiktok) has a general background audio
	Audio string `json:"audio,omitempty"`
	// AudioFilename cobalt-generated filename, returned if Audio exists.
	AudioFilename string `json:"audioFilename,omitempty"`
	// Picker array of objects containing the individual media
	Picker []PickerItem `json:"picker,omitempty"`

	// ResponseStatusError fields

	// ErrorInfo contains more context about the error
	ErrorInfo ResponseErrorInfo `json:"error,omitempty"`
}

type PickerItem struct {
	Type  PickerItemType `json:"type"`
	URL   string         `json:"url"`
	Thumb string         `json:"thumb"`
}

type PickerItemType string

const (
	PickerItemTypeVideo PickerItemType = "video"
	PickerItemTypePhoto PickerItemType = "photo"
	PickerItemTypeGIF   PickerItemType = "gif"
)

type ResponseErrorInfo struct {
	// Code machine-readable error code explaining the failure reason
	Code string `json:"code"`
	// Context container for providing more context
	Context ResponseErrorContext `json:"context,omitempty"`
}

type ResponseErrorContext struct {
	// Service states which service was being downloaded from
	Service string `json:"service,omitempty"`
	// Limit number providing the ratelimit maximum number of requests, or maximum downloadable video duration
	Limit int `json:"limit,omitempty"`
}

type GetResponse struct {
	Cobalt ServerInfo `json:"cobalt"`
	Git    GitInfo    `json:"git"`
}

type ServerInfo struct {
	// Version current version
	Version string `json:"version"`
	// URL server URL
	URL string `json:"url"`
	// StartTime server start time in unix miliseconds
	StartTime string `json:"startTime"`
	// DurationLimit maximum downloadable video lenght in seconds
	DurationLimit int `json:"limit"`
	// Services array of services which the instance supports
	Services []string `json:"services"`
}

type GitInfo struct {
	// Commit hash
	Commit string `json:"commit"`
	// Branch git Branch
	Branch string `json:"branch"`
	// Remote git remote
	Remote string `json:"remote"`
}

type SessionResponse struct {
	// Token a bearer token used for later request authentication
	Token string `json:"token,omitempty"`
	// Exp number is seconds indicating the token lifetime
	Exp int `json:"exp,omitempty"`

	// Fields only returned when errors occurs
	Status    ResponseStatus    `json:"status,omitempty"`
	ErrorInfo ResponseErrorInfo `json:"error,omitempty"`
}
