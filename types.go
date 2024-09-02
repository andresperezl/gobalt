package gobalt

import "net/http"

type Request struct {
	// URL must be included in every request
	URL string `json:"url"`

	// VideoCodec applies only for Youtube downloads.
	// Pick VideoCodecH264 if you want best compatibility. Pick VideoCodecAV1 if you want best quality and efficiency.
	// Default VideoQualityH264
	VideoCodec VideoCodec `json:"vCodec,omitempty"`

	// VideoQuality if the selected quality isn't available, closest one is used instead.
	// Default VideoQuality720p
	VideoQuality VideoQuality `json:"vQuality,omitempty"`

	// AudioFormat Format to re-encode audio into. If AudioFormatBest is selected, you get the audio the way it is on service's side.
	// Default AudioFormatMP3
	AudioFormat AudioFormat `json:"aFormat,omitempty"`

	// FilenamePattern changes the way files are named.
	// Some services don't support rich file names and always use the classic style.
	// Default FilenamePatternClassic
	FilenamePattern FilenamePattern `json:"filenamePattern,omitempty"`

	// AudioOnly download only the audio.
	// Default false
	AudioOnly bool `json:"isAudioOnly,omitempty"`

	// TiktokFullAudio download original sound used in a tiktok video.
	// Default false
	OriginalTiktokAudio bool `json:"isTTFullAudio,omitempty"`

	// MuteAudio disable audio track in video downloads.
	// Default false
	MuteAudio bool `json:"isAudioMuted,omitempty"`

	// BrowserDubLanguage use your browser's default language for youtube dubbed audio tracks.
	// Default false
	BrowserLanguage bool `json:"dubLang,omitempty"`

	// DisableMetadata disables file metadata when set to true.
	// Default false
	DisableMetadata bool `json:"disableMetadata,omitempty"`

	// TwitterGif encode twitter gif and proper .gif (not mp4)
	// Default false
	TwitterGif bool `json:"twitterGif,omitempty"`

	// TiktokH265 wheter to download Tiktok 1080p H264 video
	// Default false
	TiktokH265 bool `json:"tiktokH265,omitempty"`
}

// VideoCodec is the codec used to download the video. Only applicable to Youtube.
type VideoCodec string

const (
	// VideoCodecH264 best support across apps/platforms, average detail level. Max VideoQuality is VideoQuality1080p
	VideoCodecH264 VideoCodec = "h264"
	// VideoCodecAV1 best quality, small file size, most detail. Supports 8k and HDR.
	VideoCodecAV1 VideoCodec = "av1"
	// VideoCodecVP9 same quality as av1, but file size is 2x larger. Supports 4k and HDR.
	VideoCodecVP9 VideoCodec = "vp9"
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
	VideoQualityMax   VideoQuality = "max"

	// VideoQuality4k is an alias for VideoQuality2160p
	VideoQuality4K VideoQuality = VideoQuality2160p
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

// FilenamePattern is the way resulting files are named.
type FilenamePattern string

const (
	// FilenamePatternClassic default cobalt file name pattern.
	//
	// Video: youtube_dQw4w9WgXcQ_2560x1440_h264.mp4
	//
	// Audio: youtube_dQw4w9WgXcQ_audio.mp3
	FilenamePatternClassic FilenamePattern = "classic"

	// FilenamePatternPretty title and info in brackets.
	//
	// Video: Video Title (1440p, h264, youtube).mp4
	//
	// Audio: Audio Title - Audio Author (soundcloud).mp3
	FilenamePatternPretty FilenamePattern = "pretty"

	// FilenamePatternBasic title and basic info in brackets.
	//
	// Video: Video Title (1440p, h264).mp4
	//
	// Audio: Audio Title - Audio Author.mp3
	FilenamePatternBasic FilenamePattern = "basic"

	// FilenamePatternNerdy title and full info in brackets.
	//
	// Video: Video Title (1440p, h264, youtube, dQw4w9WgXcQ).mp4
	//
	// Audio: Audio Title - Audio Author (soundcloud, 1242868615).mp3
	FilenamePatternNerdy FilenamePattern = "nerdy"
)

type Media struct {
	client   *http.Client
	filename string

	// Status is the type of response from the API
	Status ResponseStatus `json:"status"`
	// Text is mostly used for errors
	Text string `json:"text,omitempty"`
	// URL direct link to a file or a link to cobalt's live render
	URL string `json:"url,omitempty"`
	// PickerType used when ResponseStatusPicker
	PickerType PickerType `json:"pickerType,omitempty"`
	// Picker array of PickerItem
	Picker []PickerItem `json:"picker,omitempty"`
	// Audio direct link to a file or a link to cobalt's live render
	Audio string `json:"audio,omitempty"`
}

type ResponseStatus string

const (
	ResponseStatusError     ResponseStatus = "error"
	ResponseStatusRedirect  ResponseStatus = "redirect"
	ResponseStatusStream    ResponseStatus = "stream"
	ResponseStatusSuccess   ResponseStatus = "success"
	ResponseStatusRateLimit ResponseStatus = "rate-limit"
	ResponseStatusPicker    ResponseStatus = "picker"
)

type PickerType string

const (
	PickerTypeVarious PickerType = "various"
	PickerTypeImages  PickerType = "images"
)

type PickerItem struct {
	// Type used only when parent PickerType is PickerTypeVarious
	Type PickerItemType `json:"type,omitempty"`
	// URL direct link to a file or a link to cobalt's live render
	URL string `json:"url"`
	// Thumb item thumbnail that is displayed in the picker
	Thumb string `json:"thumb"`
}

type PickerItemType string

const (
	PickerItemTypeVideo PickerItemType = "video"
	PickerItemTypePhoto PickerItemType = "photo"
	PickerItemTypeGIF   PickerItemType = "gif"
)

type ServerInfo struct {
	// Version Cobalt version
	Version string `json:"version"`
	// Commit Git commit
	Commit string `json:"commit"`
	// Branch Git Branch
	Branch string `json:"branch"`
	// Name server name
	Name string `json:"name"`
	// URL server URL
	URL string `json:"url"`
	// CORS status of CORS
	CORS string `json:"cors"`
	// StartTime server start time
	StartTime string `json:"startTime"`
}
