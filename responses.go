// QCLauncher by syncore <syncore@syncore.org> 2017
// https://github.com/syncore/qclauncher

package qclauncher

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type remoteResponseType int

const (
	rrPreSave remoteResponseType = iota
	rrAuth
	rrBuildInfo
	rrBranchInfo
	rrLaunchArgs
	rrGameCode
	rrServerStatus
	rrUpdateQC
	rrUpdateLauncher
	rrEntitlementInfo
	rrEntitlementCheckAPI
)

type remoteResponse interface {
	parse(j json.RawMessage) error
	validate() error
}

type remoteResponseData struct {
	ResponseType remoteResponseType
	Data         json.RawMessage
}

type Project struct {
	CheckFilter bool   `json:"check_filter"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

type Branch struct {
	ID         int    `json:"id"`
	Project    int    `json:"project"`
	BranchType int    `json:"branch_type"`
	BuildID    int    `json:"build_id"`
	Name       string `json:"name"`
}

type Depot struct {
	Depot252298 DepotItem `json:"252298"`
}

type DepotItem struct {
	ID              int    `json:"id"`
	Platform        int    `json:"platform"`
	Region          int    `json:"region"`
	CompressionType int    `json:"compression_type"`
	DepotType       int    `json:"depot_type"`
	DeploymentOrder int    `json:"deployment_order"`
	DefaultRegion   bool   `json:"default_region"`
	EncryptionType  int    `json:"encryption_type"`
	Language        int    `json:"language"`
	SizeOnDisk      int64  `json:"size_on_disk"`
	Name            string `json:"name"`
	DefaultLanguage bool   `json:"default_language"`
	Build           int    `json:"build"`
	DownloadSize    int64  `json:"download_size"`
	Architecture    int    `json:"architecture"`
	BytesPerChunk   int    `json:"bytes_per_chunk"`
	PropertiesID    int    `json:"properties_id"`
}

type LaunchInfo struct {
	Default         LaunchInfoItem `json:"8"`  // NOTE: Default is the only relevant item; fragile (?) may change in the future
	BetaTemp        LaunchInfoItem `json:"9"`  // NOTE: fragile (?) may change in the future
	TestMaxFPS      LaunchInfoItem `json:"10"` // NOTE: fragile (?) may change in the future
	PTSArenaBackend LaunchInfoItem `json:"14"` // NOTE: fragile (?) may change in the future
}

type LaunchInfoItem struct {
	Architecture int    `json:"architecture"`
	Description  string `json:"description"`
	ExePath      string `json:"exe_path"`
	LaunchArgs   string `json:"launch_args"`
	Name         string `json:"name"`
	Platform     int    `json:"platform"`
	Registry     string `json:"registry"`
	WorkingDir   string `json:"working_dir"`
}

type Dependency struct {
	Architecture  int    `json:"architecture"`
	CmdlineArgs   string `json:"cmdline_args"`
	ID            int    `json:"id"`
	InstallerLink string `json:"installer_link"`
	Name          string `json:"name"`
	Platform      int    `json:"platform"`
}

type AuthResponse struct {
	OAuthToken            interface{} `json:"oauth_token"`
	BeamClientAPIKey      string      `json:"beam_client_api_key"`
	Token                 string      `json:"token"`
	SessionID             string      `json:"session_id"`
	BeamToken             []string    `json:"beam_token"`
	EntitlementIDs        []int       `json:"entitlement_ids"`
	isPreSaveVerification bool        `json:"-"` // custom flag for internal launcher use
}

type BuildInfoResponse struct {
	Projects []Project `json:"projects"`
	Branches []Branch  `json:"branches"`
}

type FileDiffContainer struct {
	ID        int `json:"id"`
	FromBuild int `json:"from_build"`
	ToBuild   int `json:"to_build"`
}

type BranchBuildHistory struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type BranchInfoResponse struct {
	StorageURL         string               `json:"storage_url"`
	LaunchinfoList     []int                `json:"launchinfo_list"`
	FileDiffBuildList  []int                `json:"file_diff_build_list"`
	FileDiffContainers []FileDiffContainer  `json:"filediffcontainers"`
	BuildHistory       []BranchBuildHistory `json:"build_history"`
	Preload            bool                 `json:"preload"`
	PreloadOndeck      bool                 `json:"preload_ondeck"`
	Available          bool                 `json:"available"`
	BranchType         int                  `json:"branch_type"`
	DiffType           int                  `json:"diff_type"`
	Project            int                  `json:"project"`
	Name               string               `json:"name"`
	OnDeckBuild        interface{}          `json:"on_deck_build"`
	DepotList          Depot                `json:"depot_list"`
	Build              int                  `json:"build"`
	PreloadLiveTime    interface{}          `json:"preload_live_time"`
}

type LaunchArgsResponse struct {
	CheckFilter      bool          `json:"check_filter"`
	DefaultBranch    int           `json:"default_branch"`
	DependencyList   []Dependency  `json:"dependency_list"`
	EulaLink         string        `json:"eula_link"`
	FirewallLabel    string        `json:"firewall_label"`
	FirewallPath     string        `json:"firewall_path"`
	HasOauthClientID bool          `json:"has_oauth_client_id"`
	IconLink         string        `json:"icon_link"`
	InstallFolder    string        `json:"install_folder"`
	InstallRegistry  string        `json:"install_registry"`
	LaunchinfoSet    LaunchInfo    `json:"launchinfo_set"`
	Name             string        `json:"name"`
	NewChunkFormat   bool          `json:"new_chunk_format"`
	NewChunkDownload bool          `json:"new_chunk_download"`
	RequireLatest    bool          `json:"require_latest"`
	State            int           `json:"state"`
	StorageList      []interface{} `json:"storage_list"`
	SupportLink      string        `json:"support_link"`
}

type GameCodeResponse struct {
	Gamecode string `json:"gamecode"`
	Project  int    `json:"project"`
}

type ServerStatusResponse struct {
	Platform struct {
		Code     int    `json:"code"`
		Message  string `json:"message"`
		Response struct {
			Quake string `json:"Quake"`
		} `json:"response"`
	} `json:"platform"`
}

type FileHash struct {
	File string `json:"file"`
	Hash string `json:"hash"`
}

type UpdateQCResponse struct {
	ID     int        `json:"id"`
	Date   time.Time  `json:"date"`
	Hashes []FileHash `json:"hashes"`
	BVer   string     `json:"bver"`
}

type UpdateLauncherResponse struct {
	LatestVersion float32   `json:"latest"`
	Date          time.Time `json:"date"`
	URL           string    `json:"url"`
}

type EntitlementCheckAPIResponse struct {
	UseEntitlementAPI bool `json:"useEntitlementAPI"`
}

type EntitlementInfoResponse struct {
	Blacklist EntitlementBlacklist `json:"blacklist"`
	Branches  []EntitlementBranch  `json:"branches"`
	Projects  []EntitlementProject `json:"projects"`
}

type EntitlementBlacklist struct {
	Branches []EntitlementBlacklistBranch  `json:"branches"`
	Country  string                        `json:"country"`
	IP       string                        `json:"ip"`
	Projects []EntitlementBlacklistProject `json:"projects"`
}

type EntitlementBlacklistBranch struct {
	Available  bool   `json:"available"`
	BranchType int    `json:"branch_type"`
	Build      int    `json:"build"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Preload    bool   `json:"preload"`
	Project    int    `json:"project"`
}

type EntitlementBlacklistProject struct {
	ID int `json:"id"`
}

type EntitlementBranch struct {
	Available  bool   `json:"available"`
	BranchType int    `json:"branch_type"`
	Build      int    `json:"build"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Preload    bool   `json:"preload"`
	Project    int    `json:"project"`
}

type EntitlementProject struct {
	BeamClientKey    bool   `json:"beam_client_key"`
	Buildinfo        bool   `json:"buildinfo"`
	DefaultBranch    int    `json:"default_branch"`
	ID               int    `json:"id"`
	Name             string `json:"name"`
	NewChunkDownload bool   `json:"new_chunk_download"`
	NewChunkFormat   bool   `json:"new_chunk_format"`
}

func (response *AuthResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw auth response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := updateAuthToken(response.isPreSaveVerification, response.Token); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error updating auth token from response", GetCaller()), "error", err, "data", response.Token)
		return err
	}
	return nil
}

func (response *BuildInfoResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw build info response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: build info response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

func (response *BranchInfoResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw branch info response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: branch info response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

func (response *LaunchArgsResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw launch args response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: launch args response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

func (response *GameCodeResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw game code response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: game code response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

func (response *ServerStatusResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw server status response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: server status response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

func (response *UpdateQCResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw QC update response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: QC update response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

func (response *UpdateLauncherResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw launcher update response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: launcher update response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

func (response *EntitlementInfoResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw entitlement info response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: entitlement info response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

func (response *EntitlementCheckAPIResponse) parse(j json.RawMessage) error {
	if err := json.Unmarshal(j, response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: error parsing raw entitlement check API response message", GetCaller()), "error", err, "data", string(j))
		return err
	}
	if err := validateResponse(response); err != nil {
		logger.Errorw(fmt.Sprintf("%s: entitlement check API response failed validation", GetCaller()), "error", err, "data", response)
		return err
	}
	return nil
}

// yeah, cyclomatic complexity and stuff...
func parseRemoteResponseData(rd *remoteResponseData) (interface{}, error) {
	switch rd.ResponseType {
	case rrAuth, rrPreSave:
		var r AuthResponse
		if rd.ResponseType == rrPreSave {
			r.isPreSaveVerification = true
		}
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrBuildInfo:
		var r BuildInfoResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrEntitlementInfo:
		var r EntitlementInfoResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrBranchInfo:
		var r BranchInfoResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrLaunchArgs:
		var r LaunchArgsResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrGameCode:
		var r GameCodeResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrServerStatus:
		var r ServerStatusResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrUpdateQC:
		var r UpdateQCResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrUpdateLauncher:
		var r UpdateLauncherResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	case rrEntitlementCheckAPI:
		var r EntitlementCheckAPIResponse
		err := r.parse(rd.Data)
		if err != nil {
			return nil, err
		}
		return r, nil
	default:
		logger.Error(fmt.Sprintf("%s: received unknown response type: %v", GetCaller(), rd.ResponseType))
		return nil, errors.New("Unknown response type")
	}
}
