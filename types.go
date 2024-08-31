package gobalt

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
	// VIdeoCodecVP9 same quality as av1, but file size is 2x larger. Supports 4k and HDR.
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
	FilenamePatternClassic FilenamePattern = "classic"
	FilenamePatternPretty  FilenamePattern = "pretty"
	FilenamePatternBasic   FilenamePattern = "basic"
	FilenamePatternNerdy   FilenamePattern = "nerdy"
)
