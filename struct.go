package utils

import "time"

/*codice generato da https://mholt.github.io/json-to-go/*/

type ScannerR struct {
	ScanResultHistoryLength int    `json:"scan_result_history_length"`
	FileID                  string `json:"file_id"`
	DataID                  string `json:"data_id"`
	ScanResults             struct {
		ScanDetails struct {
			Lavasoft struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Lavasoft"`
			STOPzilla struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"STOPzilla"`
			Zillya struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Zillya!"`
			VirusBlokAda struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"VirusBlokAda"`
			TrendMicro struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"TrendMicro"`
			SUPERAntiSpyware struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"SUPERAntiSpyware"`
			NProtect struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"nProtect"`
			Nanoav struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"NANOAV"`
			FSecure struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"F-secure"`
			Eset struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"ESET"`
			BitDefender struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"BitDefender"`
			Baidu struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Baidu"`
			Ahnlab struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Ahnlab"`
			AegisLab struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"AegisLab"`
			Zoner struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Zoner"`
			ThreatTrack struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"ThreatTrack"`
			Sophos struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Sophos"`
			Preventon struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Preventon"`
			McAfee struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"McAfee"`
			K7 struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"K7"`
			Jiangmin struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Jiangmin"`
			Hauri struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Hauri"`
			FProt struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"F-prot"`
			Fortinet struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Fortinet"`
			Filseclab struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Filseclab"`
			Emsisoft struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Emsisoft"`
			ClamAV struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"ClamAV"`
			ByteHero struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"ByteHero"`
			Avira struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Avira"`
			Avg struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"AVG"`
			Agnitum struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Agnitum"`
			Ikarus struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Ikarus"`
			Cyren struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Cyren"`
			MicrosoftSecurityEssentials struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Microsoft Security Essentials"`
			QuickHeal struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Quick Heal"`
			TotalDefense struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Total Defense"`
			TrendMicroHouseCall struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"TrendMicro House Call"`
			XvirusPersonalGuard struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Xvirus Personal Guard"`
			DrWebGateway struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Dr.Web Gateway"`
			VirITEXplorer struct {
				ScanTime    int       `json:"scan_time"`
				DefTime     time.Time `json:"def_time"`
				ScanResultI int       `json:"scan_result_i"`
				ThreatFound string    `json:"threat_found"`
			} `json:"Vir.IT eXplorer"`
		} `json:"scan_details"`
		RescanAvailable    bool      `json:"rescan_available"`
		ScanAllResultI     int       `json:"scan_all_result_i"`
		StartTime          time.Time `json:"start_time"`
		TotalTime          int       `json:"total_time"`
		TotalAvs           int       `json:"total_avs"`
		TotalDetectedAvs   int       `json:"total_detected_avs"`
		ProgressPercentage int       `json:"progress_percentage"`
		ScanAllResultA     string    `json:"scan_all_result_a"`
	} `json:"scan_results"`
	FileInfo struct {
		FileSize            int       `json:"file_size"`
		UploadTimestamp     time.Time `json:"upload_timestamp"`
		Md5                 string    `json:"md5"`
		Sha1                string    `json:"sha1"`
		Sha256              string    `json:"sha256"`
		FileTypeCategory    string    `json:"file_type_category"`
		FileTypeDescription string    `json:"file_type_description"`
		FileTypeExtension   string    `json:"file_type_extension"`
		DisplayName         string    `json:"display_name"`
	} `json:"file_info"`
	ShareFile      int           `json:"share_file"`
	RestVersion    string        `json:"rest_version"`
	AdditionalInfo []interface{} `json:"additional_info"`
	Votes          struct {
		Up   int `json:"up"`
		Down int `json:"down"`
	} `json:"votes"`
}
