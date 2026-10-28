package sdwebui

type txt2img struct {
	Prompt           string           `json:"prompt"`
	NegativePrompt   string           `json:"negative_prompt"`
	Seed             int              `json:"seed"`
	BatchSize        int              `json:"batch_size"`
	NIter            int              `json:"n_iter"`
	Steps            int              `json:"steps"`
	CfgScale         float64          `json:"cfg_scale"`
	Width            int              `json:"width"`
	Height           int              `json:"height"`
	RestoreFaces     bool             `json:"restore_faces"`
	Tiling           bool             `json:"tiling"`
	SamplerIndex     string           `json:"sampler_index"`
	SendImages       bool             `json:"send_images"`
	ClipSkip         int              `json:"clip_skip"`
	OverrideSettings overrideSettings `json:"override_settings"`
}

type info struct {
	Prompt                string      `json:"prompt"`
	AllPrompts            []string    `json:"all_prompts"`
	NegativePrompt        string      `json:"negative_prompt"`
	AllNegativePrompts    []string    `json:"all_negative_prompts"`
	Seed                  int         `json:"seed"`
	AllSeeds              []int       `json:"all_seeds"`
	Subseed               int64       `json:"subseed"`
	AllSubseeds           []int64     `json:"all_subseeds"`
	SubseedStrength       float64     `json:"subseed_strength"`
	Width                 int         `json:"width"`
	Height                int         `json:"height"`
	SamplerName           string      `json:"sampler_name"`
	CfgScale              float32     `json:"cfg_scale"`
	Steps                 int         `json:"steps"`
	BatchSize             int         `json:"batch_size"`
	RestoreFaces          bool        `json:"restore_faces"`
	FaceRestorationModel  string      `json:"face_restoration_model"`
	SdModelName           string      `json:"sd_model_name"`
	SdModelHash           interface{} `json:"sd_model_hash"`
	SdVaeName             interface{} `json:"sd_vae_name"`
	SdVaeHash             interface{} `json:"sd_vae_hash"`
	SeedResizeFromW       int         `json:"seed_resize_from_w"`
	SeedResizeFromH       int         `json:"seed_resize_from_h"`
	DenoisingStrength     float64     `json:"denoising_strength"`
	ExtraGenerationParams struct {
		Eta float64 `json:"Eta"`
	} `json:"extra_generation_params"`
	IndexOfFirstImage             int           `json:"index_of_first_image"`
	Infotexts                     []string      `json:"infotexts"`
	Styles                        []interface{} `json:"styles"`
	JobTimestamp                  string        `json:"job_timestamp"`
	ClipSkip                      int           `json:"clip_skip"`
	IsUsingInpaintingConditioning bool          `json:"is_using_inpainting_conditioning"`
	Version                       string        `json:"version"`
}

type imginfo struct {
	Images     []string `json:"images"`
	Parameters txt2img  `json:"parameters"`
	Info       string   `json:"info"`
}

type models []modelinfo

type modelinfo struct {
	Title     string      `json:"title"`
	ModelName string      `json:"model_name"`
	Hash      interface{} `json:"hash"`
	Sha256    interface{} `json:"sha256"`
	Filename  string      `json:"filename"`
	Config    interface{} `json:"config"`
}

type options struct {
	SDModelCheckpoint    string `json:"sd_model_checkpoint"`
	CLIPStopAtLastLayers int    `json:"CLIP_stop_at_last_layers"`
}

type overrideSettings struct {
	CLIPStopAtLastLayers int `json:"CLIP_stop_at_last_layers"`
}
