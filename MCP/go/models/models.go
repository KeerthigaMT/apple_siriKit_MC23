package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UpdateActivityRequest represents the UpdateActivityRequest schema from the OpenAPI specification
type UpdateActivityRequest struct {
	Constraints Constraints `json:"constraints,omitempty"`
	Nowplaying PlayerContext `json:"nowPlaying,omitempty"`
	Previouslyplaying PlayerContext `json:"previouslyPlaying,omitempty"`
	Report string `json:"report"`
	Timestamp string `json:"timestamp"`
	Useractivity UserActivity `json:"userActivity"`
	Version string `json:"version"`
}

// MediaAffinityTypeResolutionResult represents the MediaAffinityTypeResolutionResult schema from the OpenAPI specification
type MediaAffinityTypeResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// PlayMediaIntentResponse represents the PlayMediaIntentResponse schema from the OpenAPI specification
type PlayMediaIntentResponse struct {
	Class string `json:"class"`
	Useractivity UserActivity `json:"userActivity"`
}

// UpdateMediaAffinityIntentHandlingHandleInvocationResponse represents the UpdateMediaAffinityIntentHandlingHandleInvocationResponse schema from the OpenAPI specification
type UpdateMediaAffinityIntentHandlingHandleInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// PlayMediaIntentHandlingResolvePlayShuffledInvocationResponse represents the PlayMediaIntentHandlingResolvePlayShuffledInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolvePlayShuffledInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// DateComponentsRange represents the DateComponentsRange schema from the OpenAPI specification
type DateComponentsRange struct {
	Enddatecomponents interface{} `json:"endDateComponents,omitempty"`
	Startdatecomponents interface{} `json:"startDateComponents,omitempty"`
}

// MediaDestination represents the MediaDestination schema from the OpenAPI specification
type MediaDestination struct {
	Mediadestinationtype string `json:"mediaDestinationType"`
}

// ExplicitDateComponents represents the ExplicitDateComponents schema from the OpenAPI specification
type ExplicitDateComponents struct {
	Timezone string `json:"timeZone,omitempty"`
	Minute int `json:"minute,omitempty"`
	Month int `json:"month,omitempty"`
	Year int `json:"year,omitempty"`
	Calendaridentifier string `json:"calendarIdentifier,omitempty"`
	Hour int `json:"hour,omitempty"`
	Nanosecond int `json:"nanosecond,omitempty"`
	Era int `json:"era,omitempty"`
	Day int `json:"day,omitempty"`
	Second int `json:"second,omitempty"`
}

// UpdateMediaAffinityIntentHandlingResolveAffinityTypeInvocationResponse represents the UpdateMediaAffinityIntentHandlingResolveAffinityTypeInvocationResponse schema from the OpenAPI specification
type UpdateMediaAffinityIntentHandlingResolveAffinityTypeInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// PlayMediaIntentHandlingResolveMediaItemsInvocationResponse represents the PlayMediaIntentHandlingResolveMediaItemsInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolveMediaItemsInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// ExtensionEndpointConfig represents the ExtensionEndpointConfig schema from the OpenAPI specification
type ExtensionEndpointConfig struct {
	Hdr map[string]interface{} `json:"hdr,omitempty"`
	Url string `json:"url,omitempty"`
}

// PlayMediaRequest represents the PlayMediaRequest schema from the OpenAPI specification
type PlayMediaRequest struct {
	Constraints Constraints `json:"constraints"`
	Useractivity UserActivity `json:"userActivity"`
	Version string `json:"version"`
}

// MediaDestinationLibrary represents the MediaDestinationLibrary schema from the OpenAPI specification
type MediaDestinationLibrary struct {
	Mediadestinationtype string `json:"mediaDestinationType"`
}

// PlayerContext represents the PlayerContext schema from the OpenAPI specification
type PlayerContext struct {
	Activityidentifier string `json:"activityIdentifier,omitempty"`
	Contentidentifier string `json:"contentIdentifier,omitempty"`
	Offsetinmillis int64 `json:"offsetInMillis,omitempty"`
	Playbackspeed float64 `json:"playbackSpeed,omitempty"`
	Queueidentifier string `json:"queueIdentifier,omitempty"`
}

// PlayMediaControlCommandSet represents the PlayMediaControlCommandSet schema from the OpenAPI specification
type PlayMediaControlCommandSet struct {
	Nexttrack bool `json:"nextTrack,omitempty"`
	Seektoplaybackposition bool `json:"seekToPlaybackPosition,omitempty"`
	Skipbackward bool `json:"skipBackward,omitempty"`
	Preferskipforward bool `json:"preferSkipForward,omitempty"`
	Skipforward bool `json:"skipForward,omitempty"`
	Disliketrack bool `json:"dislikeTrack,omitempty"`
	Liketrack bool `json:"likeTrack,omitempty"`
	Preferskipbackward bool `json:"preferSkipBackward,omitempty"`
	Previoustrack bool `json:"previousTrack,omitempty"`
	Bookmarktrack bool `json:"bookmarkTrack,omitempty"`
}

// InvocationResponse represents the InvocationResponse schema from the OpenAPI specification
type InvocationResponse struct {
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
}

// MediaDestinationPlaylist represents the MediaDestinationPlaylist schema from the OpenAPI specification
type MediaDestinationPlaylist struct {
	Mediadestinationtype string `json:"mediaDestinationType"`
}

// AddMediaMediaItemResolutionResult represents the AddMediaMediaItemResolutionResult schema from the OpenAPI specification
type AddMediaMediaItemResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// Invocation represents the Invocation schema from the OpenAPI specification
type Invocation struct {
	Params map[string]interface{} `json:"params"`
	Session Session `json:"session,omitempty"`
	Method string `json:"method"`
}

// UpdateMediaAffinityIntentHandlingInvocation represents the UpdateMediaAffinityIntentHandlingInvocation schema from the OpenAPI specification
type UpdateMediaAffinityIntentHandlingInvocation struct {
	Method string `json:"method"`
	Params map[string]interface{} `json:"params"`
	Session Session `json:"session,omitempty"`
}

// MediaItem represents the MediaItem schema from the OpenAPI specification
type MediaItem struct {
	Artist string `json:"artist,omitempty"`
	Identifier string `json:"identifier"`
	Title string `json:"title,omitempty"`
	TypeField string `json:"type"`
}

// PlayMediaIntentHandlingResolvePlaybackRepeatModeInvocationResponse represents the PlayMediaIntentHandlingResolvePlaybackRepeatModeInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolvePlaybackRepeatModeInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// AddMediaIntentHandlingResolveMediaItemsInvocationResponse represents the AddMediaIntentHandlingResolveMediaItemsInvocationResponse schema from the OpenAPI specification
type AddMediaIntentHandlingResolveMediaItemsInvocationResponse struct {
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
}

// PlayMediaMediaItemResolutionResult represents the PlayMediaMediaItemResolutionResult schema from the OpenAPI specification
type PlayMediaMediaItemResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// PlayMediaIntentHandlingHandleInvocationResponse represents the PlayMediaIntentHandlingHandleInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingHandleInvocationResponse struct {
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
}

// Queue represents the Queue schema from the OpenAPI specification
type Queue struct {
	Version string `json:"version"`
	Prerollseconds float64 `json:"prerollSeconds,omitempty"`
	Skipsremaining int `json:"skipsRemaining,omitempty"`
	Nextcontenturl string `json:"nextContentUrl,omitempty"`
	Playpointer QueuePlayPointer `json:"playPointer,omitempty"`
	Previouscontenturl string `json:"previousContentUrl,omitempty"`
	Content []Content `json:"content"`
	Contentitemscount int `json:"contentItemsCount,omitempty"`
	Insertpointer QueueInsertPointer `json:"insertPointer,omitempty"`
	Controls QueueControlMapping `json:"controls,omitempty"`
	Identifier string `json:"identifier"`
}

// PlayMediaControlActivity represents the PlayMediaControlActivity schema from the OpenAPI specification
type PlayMediaControlActivity struct {
	Playelapsed int `json:"playElapsed,omitempty"`
	Playelapsedinterval int `json:"playElapsedInterval,omitempty"`
	Playpaused int `json:"playPaused,omitempty"`
}

// AddMediaIntentHandlingInvocation represents the AddMediaIntentHandlingInvocation schema from the OpenAPI specification
type AddMediaIntentHandlingInvocation struct {
	Session Session `json:"session,omitempty"`
	Method string `json:"method"`
	Params map[string]interface{} `json:"params"`
}

// PlaybackRepeatModeResolutionResult represents the PlaybackRepeatModeResolutionResult schema from the OpenAPI specification
type PlaybackRepeatModeResolutionResult struct {
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
}

// AddMediaIntentResponse represents the AddMediaIntentResponse schema from the OpenAPI specification
type AddMediaIntentResponse struct {
	Useractivity UserActivity `json:"userActivity"`
	Class string `json:"class"`
}

// UpdateMediaAffinityMediaItemResolutionResult represents the UpdateMediaAffinityMediaItemResolutionResult schema from the OpenAPI specification
type UpdateMediaAffinityMediaItemResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// Session represents the Session schema from the OpenAPI specification
type Session struct {
	Version string `json:"version"`
	Constraints Constraints `json:"constraints"`
	Deadline string `json:"deadline"`
	Identifier string `json:"identifier"`
	Playercontext PlayerContext `json:"playerContext,omitempty"`
	Requested string `json:"requested"`
}

// UserActivity represents the UserActivity schema from the OpenAPI specification
type UserActivity struct {
	Title string `json:"title,omitempty"`
	Userinfo map[string]interface{} `json:"userInfo,omitempty"`
	Version string `json:"version"`
	Activitytype string `json:"activityType"`
	Persistentidentifier string `json:"persistentIdentifier,omitempty"`
}

// BooleanResolutionResult represents the BooleanResolutionResult schema from the OpenAPI specification
type BooleanResolutionResult struct {
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
}

// Constraints represents the Constraints schema from the OpenAPI specification
type Constraints struct {
	Maximumqueuesegmentitemcount int `json:"maximumQueueSegmentItemCount,omitempty"`
	Updateusertasteprofile bool `json:"updateUserTasteProfile,omitempty"`
	Allowexplicitcontent bool `json:"allowExplicitContent,omitempty"`
}

// PlayMediaIntentHandlingResolveResumePlaybackInvocationResponse represents the PlayMediaIntentHandlingResolveResumePlaybackInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolveResumePlaybackInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// ContentAttributes represents the ContentAttributes schema from the OpenAPI specification
type ContentAttributes struct {
	Composername string `json:"composerName,omitempty"`
	Durationinmillis int `json:"durationInMillis,omitempty"`
	Genrenames []string `json:"genreNames,omitempty"`
	Name string `json:"name,omitempty"`
	Tracknumber int `json:"trackNumber,omitempty"`
	Albumname string `json:"albumName,omitempty"`
	Artistname string `json:"artistName,omitempty"`
	Artwork map[string]interface{} `json:"artwork,omitempty"`
}

// ExtensionConfig represents the ExtensionConfig schema from the OpenAPI specification
type ExtensionConfig struct {
	Media map[string]interface{} `json:"media"`
	Url string `json:"url,omitempty"`
	Version string `json:"version"`
	Hdr map[string]interface{} `json:"hdr,omitempty"`
	Intent map[string]interface{} `json:"intent"`
}

// IntentResponse represents the IntentResponse schema from the OpenAPI specification
type IntentResponse struct {
	Class string `json:"class"`
	Useractivity UserActivity `json:"userActivity"`
}

// UpdateActivityResponse represents the UpdateActivityResponse schema from the OpenAPI specification
type UpdateActivityResponse struct {
	Queue Queue `json:"queue,omitempty"`
	Useractivity UserActivity `json:"userActivity,omitempty"`
}

// UpdateMediaAffinityIntentHandlingResolveMediaItemsInvocationResponse represents the UpdateMediaAffinityIntentHandlingResolveMediaItemsInvocationResponse schema from the OpenAPI specification
type UpdateMediaAffinityIntentHandlingResolveMediaItemsInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// UpdateMediaAffinityIntent represents the UpdateMediaAffinityIntent schema from the OpenAPI specification
type UpdateMediaAffinityIntent struct {
	Identifier string `json:"identifier"`
	Class string `json:"class"`
}

// AddMediaIntentHandlingConfirmInvocationResponse represents the AddMediaIntentHandlingConfirmInvocationResponse schema from the OpenAPI specification
type AddMediaIntentHandlingConfirmInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// AddMediaIntentHandlingResolveMediaDestinationInvocationResponse represents the AddMediaIntentHandlingResolveMediaDestinationInvocationResponse schema from the OpenAPI specification
type AddMediaIntentHandlingResolveMediaDestinationInvocationResponse struct {
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
}

// AddMediaIntentHandlingHandleInvocationResponse represents the AddMediaIntentHandlingHandleInvocationResponse schema from the OpenAPI specification
type AddMediaIntentHandlingHandleInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// PlayMediaIntentHandlingResolvePlaybackQueueLocationInvocationResponse represents the PlayMediaIntentHandlingResolvePlaybackQueueLocationInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolvePlaybackQueueLocationInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// Intent represents the Intent schema from the OpenAPI specification
type Intent struct {
	Identifier string `json:"identifier"`
	Class string `json:"class"`
}

// UpdateMediaAffinityIntentResponse represents the UpdateMediaAffinityIntentResponse schema from the OpenAPI specification
type UpdateMediaAffinityIntentResponse struct {
	Useractivity UserActivity `json:"userActivity"`
	Class string `json:"class"`
}

// ProtocolExceptionInvocationResponse represents the ProtocolExceptionInvocationResponse schema from the OpenAPI specification
type ProtocolExceptionInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// Content represents the Content schema from the OpenAPI specification
type Content struct {
	Playindex int `json:"playIndex,omitempty"`
	Url string `json:"url,omitempty"`
	Attributes ContentAttributes `json:"attributes,omitempty"`
	Control string `json:"control,omitempty"`
	Identifier string `json:"identifier"`
	Islive bool `json:"isLive,omitempty"`
}

// ProtocolException represents the ProtocolException schema from the OpenAPI specification
type ProtocolException struct {
	Trace []string `json:"trace,omitempty"`
	Code int64 `json:"code,omitempty"`
	Methodindex int `json:"methodIndex,omitempty"`
	Methodname string `json:"methodName,omitempty"`
	Reason string `json:"reason"`
	Retrywithdelay float32 `json:"retryWithDelay,omitempty"`
}

// ExecutionMetrics represents the ExecutionMetrics schema from the OpenAPI specification
type ExecutionMetrics struct {
	Duration float32 `json:"duration,omitempty"`
	Received string `json:"received,omitempty"`
	Completed string `json:"completed,omitempty"`
}

// QueuePlayPointer represents the QueuePlayPointer schema from the OpenAPI specification
type QueuePlayPointer struct {
	Contentidentifier string `json:"contentIdentifier,omitempty"`
	Offsetinmillis int64 `json:"offsetInMillis,omitempty"`
}

// IntentResolutionResult represents the IntentResolutionResult schema from the OpenAPI specification
type IntentResolutionResult struct {
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
}

// AddMediaIntent represents the AddMediaIntent schema from the OpenAPI specification
type AddMediaIntent struct {
	Class string `json:"class"`
	Identifier string `json:"identifier"`
}

// AddMediaMediaDestinationResolutionResult represents the AddMediaMediaDestinationResolutionResult schema from the OpenAPI specification
type AddMediaMediaDestinationResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// QueueInsertPointer represents the QueueInsertPointer schema from the OpenAPI specification
type QueueInsertPointer struct {
	Afteridentifier string `json:"afterIdentifier,omitempty"`
	Replace bool `json:"replace,omitempty"`
}

// PlayMediaControl represents the PlayMediaControl schema from the OpenAPI specification
type PlayMediaControl struct {
	Scheme string `json:"scheme"`
	Activity PlayMediaControlActivity `json:"activity,omitempty"`
	Commands PlayMediaControlCommandSet `json:"commands,omitempty"`
}

// MediaSearch represents the MediaSearch schema from the OpenAPI specification
type MediaSearch struct {
	Mediaidentifier string `json:"mediaIdentifier,omitempty"`
	Mediatype string `json:"mediaType,omitempty"`
	Reference string `json:"reference,omitempty"`
	Artistname string `json:"artistName,omitempty"`
	Genrenames []string `json:"genreNames,omitempty"`
	Moodnames []string `json:"moodNames,omitempty"`
	Sortorder string `json:"sortOrder,omitempty"`
	Albumname string `json:"albumName,omitempty"`
	Medianame string `json:"mediaName,omitempty"`
	Releasedate DateComponentsRange `json:"releaseDate,omitempty"`
}

// PlayMediaIntentHandlingInvocation represents the PlayMediaIntentHandlingInvocation schema from the OpenAPI specification
type PlayMediaIntentHandlingInvocation struct {
	Params map[string]interface{} `json:"params"`
	Session Session `json:"session,omitempty"`
	Method string `json:"method"`
}

// PlayMediaIntent represents the PlayMediaIntent schema from the OpenAPI specification
type PlayMediaIntent struct {
	Identifier string `json:"identifier"`
	Class string `json:"class"`
}

// PlaybackQueueLocationResolutionResult represents the PlaybackQueueLocationResolutionResult schema from the OpenAPI specification
type PlaybackQueueLocationResolutionResult struct {
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
}

// QueueControlMapping represents the QueueControlMapping schema from the OpenAPI specification
type QueueControlMapping struct {
	DefaultField PlayMediaControl `json:"default"`
}
